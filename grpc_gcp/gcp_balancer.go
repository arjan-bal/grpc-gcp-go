/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc_gcp

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	// "google.golang.org/grpc/metadata"
	// "runtime/debug"
)

// Name is the name of grpc_gcp balancer.
const Name = "grpc_gcp"

// Default settings for max pool size and max concurrent  streams.
const defaultMaxConn = 10
const defaultMaxStream = 100

type gcpBalancerBuilder struct{
	name          string
	pickerBuilder base.PickerBuilder
	config        base.Config
}

func (bb *gcpBalancerBuilder) Build(cc balancer.ClientConn, opt balancer.BuildOptions) balancer.Balancer {
	return &gcpBalancer{
		cc:            cc,
		pickerBuilder: bb.pickerBuilder,

		subConns: make(map[resolver.Address]balancer.SubConn),
		scStates: make(map[balancer.SubConn]connectivity.State),
		csEvltr:  &connectivityStateEvaluator{},
		// Initialize picker to a picker that always return
		// ErrNoSubConnAvailable, because when state of a SubConn changes, we
		// may call UpdateBalancerState with this picker.
		picker: base.NewErrPicker(balancer.ErrNoSubConnAvailable),
		config: bb.config,
	}
}

func (*gcpBalancerBuilder) Name() string {
	return Name
}

// newBuilder creates a new grpc_gcp balancer builder.
func newBuilder() balancer.Builder {
	return &gcpBalancerBuilder{
		name: Name,
		pickerBuilder: &gcpPickerBuilder{},
		config: base.Config{HealthCheck: true},
	}
}

// connectivityStateEvaluator gets updated by addrConns when their
// states transition, based on which it evaluates the state of
// ClientConn.
type connectivityStateEvaluator struct {
	numReady            uint64 // Number of addrConns in ready state.
	numConnecting       uint64 // Number of addrConns in connecting state.
	numTransientFailure uint64 // Number of addrConns in transientFailure.
}

// recordTransition records state change happening in every subConn and based on
// that it evaluates what aggregated state should be.
// It can only transition between Ready, Connecting and TransientFailure. Other states,
// Idle and Shutdown are transitioned into by ClientConn; in the beginning of the connection
// before any subConn is created ClientConn is in idle state. In the end when ClientConn
// closes it is in Shutdown state.
//
// recordTransition should only be called synchronously from the same goroutine.
func (cse *connectivityStateEvaluator) recordTransition(oldState, newState connectivity.State) connectivity.State {
	// Update counters.
	for idx, state := range []connectivity.State{oldState, newState} {
		updateVal := 2*uint64(idx) - 1 // -1 for oldState and +1 for new.
		switch state {
		case connectivity.Ready:
			cse.numReady += updateVal
		case connectivity.Connecting:
			cse.numConnecting += updateVal
		case connectivity.TransientFailure:
			cse.numTransientFailure += updateVal
		}
	}

	// Evaluate.
	if cse.numReady > 0 {
		return connectivity.Ready
	}
	if cse.numConnecting > 0 {
		return connectivity.Connecting
	}
	return connectivity.TransientFailure
}

type gcpBalancer struct {
	cc            balancer.ClientConn
	pickerBuilder base.PickerBuilder

	csEvltr *connectivityStateEvaluator
	state   connectivity.State

	subConns map[resolver.Address]balancer.SubConn
	scStates map[balancer.SubConn]connectivity.State
	picker   balancer.Picker
	config   base.Config
}

func (gb *gcpBalancer) HandleResolvedAddrs(addrs []resolver.Address, err error) {
	if err != nil {
		grpclog.Infof("base.baseBalancer: HandleResolvedAddrs called with error %v", err)
		return
	}
	grpclog.Infoln("base.baseBalancer: got new resolved addresses: ", addrs)
	// addrsSet is the set converted from addrs, it's used for quick lookup of an address.
	addrsSet := make(map[resolver.Address]struct{})
	for _, a := range addrs {
		addrsSet[a] = struct{}{}
		if _, ok := gb.subConns[a]; !ok {
			// a is a new address (not existing in b.subConns).
			sc, err := gb.cc.NewSubConn([]resolver.Address{a}, balancer.NewSubConnOptions{HealthCheckEnabled: gb.config.HealthCheck})
			if err != nil {
				grpclog.Warningf("base.baseBalancer: failed to create new SubConn: %v", err)
				continue
			}
			gb.subConns[a] = sc
			gb.scStates[sc] = connectivity.Idle
			sc.Connect()
		}
	}
	for a, sc := range gb.subConns {
		// a was removed by resolver.
		if _, ok := addrsSet[a]; !ok {
			gb.cc.RemoveSubConn(sc)
			delete(gb.subConns, a)
			// Keep the state of this sc in b.scStates until sc's state becomes Shutdown.
			// The entry will be deleted in HandleSubConnStateChange.
		}
	}
}

// regeneratePicker takes a snapshot of the balancer, and generates a picker
// from it. The picker is
//  - errPicker with ErrTransientFailure if the balancer is in TransientFailure,
//  - built by the pickerBuilder with all READY SubConns otherwise.
func (gb *gcpBalancer) regeneratePicker() {
	if gb.state == connectivity.TransientFailure {
		gb.picker = base.NewErrPicker(balancer.ErrTransientFailure)
		return
	}
	readySCs := make(map[resolver.Address]balancer.SubConn)

	// Filter out all ready SCs from full subConn map.
	for addr, sc := range gb.subConns {
		if st, ok := gb.scStates[sc]; ok && st == connectivity.Ready {
			readySCs[addr] = sc
		}
	}
	gb.picker = gb.pickerBuilder.Build(readySCs)
}

func (gb *gcpBalancer) HandleSubConnStateChange(sc balancer.SubConn, s connectivity.State) {
	grpclog.Infof("base.baseBalancer: handle SubConn state change: %p, %v", sc, s)
	oldS, ok := gb.scStates[sc]
	if !ok {
		grpclog.Infof("base.baseBalancer: got state changes for an unknown SubConn: %p, %v", sc, s)
		return
	}
	gb.scStates[sc] = s
	switch s {
	case connectivity.Idle:
		sc.Connect()
	case connectivity.Shutdown:
		// When an address was removed by resolver, b called RemoveSubConn but
		// kept the sc's state in scStates. Remove state for this sc here.
		delete(gb.scStates, sc)
	}

	oldAggrState := gb.state
	gb.state = gb.csEvltr.recordTransition(oldS, s)

	// Regenerate picker when one of the following happens:
	//  - this sc became ready from not-ready
	//  - this sc became not-ready from ready
	//  - the aggregated state of balancer became TransientFailure from non-TransientFailure
	//  - the aggregated state of balancer became non-TransientFailure from TransientFailure
	if (s == connectivity.Ready) != (oldS == connectivity.Ready) ||
		(gb.state == connectivity.TransientFailure) != (oldAggrState == connectivity.TransientFailure) {
		gb.regeneratePicker()
	}

	gb.cc.UpdateBalancerState(gb.state, gb.picker)
}

func (gb *gcpBalancer) Close() {
}

func init() {
	balancer.Register(newBuilder())
}