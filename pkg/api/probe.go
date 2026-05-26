// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

// ProbeStatus is the status of a probe.
// ProbeStatus is treated as a sync/atomic int32.
type ProbeStatus int32

// get returns the value of the ProbeStatus.
func (ps *ProbeStatus) get() ProbeStatus { _ = "STUB: not implemented"; return *new(ProbeStatus) }

// set updates the value of the ProbeStatus.
func (ps *ProbeStatus) set(v ProbeStatus) { _ = "STUB: not implemented"; return }

// String implements the fmt.Stringer interface.
func (ps ProbeStatus) String() string { _ = "STUB: not implemented"; return "" }

const (
	// ProbeStatusOK indicates positive ProbeStatus status.
	ProbeStatusOK ProbeStatus = 1

	// ProbeStatusNOK indicates negative ProbeStatus status.
	ProbeStatusNOK ProbeStatus = 0
)

// Probe structure holds flags which indicate node healthiness (sometimes referred also as liveness) and readiness.
type Probe struct {
	// Healthy probe indicates if node, due to any reason, needs to restarted.
	healthy ProbeStatus
	// Ready probe indicates that node is ready to start accepting traffic.
	ready ProbeStatus
}

// NewProbe returns new Probe.
func NewProbe() *Probe {
	_ = "STUB: not implemented"

	// Healthy returns the value of the healthy status.
	return nil
}

func (p *Probe) Healthy() ProbeStatus { _ = "STUB: not implemented"; return *new(ProbeStatus) }

// SetHealthy updates the value of the healthy status.
func (p *Probe) SetHealthy(ps ProbeStatus) {
	_ = "STUB: not implemented"

	// Ready returns the value of the ready status.
	return
}

func (p *Probe) Ready() ProbeStatus { _ = "STUB: not implemented"; return *new(ProbeStatus) }

// SetReady updates the value of the ready status.
func (p *Probe) SetReady(ps ProbeStatus) { _ = "STUB: not implemented"; return }
