// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type WriterAtChan interface { // bidirectional channel
	WriterAtROnlyChan // aka "<-chan" - receive only
	WriterAtSOnlyChan // aka "chan<-" - send only
}

type WriterAtROnlyChan interface { // receive-only channel
	RequestWriterAt() (dat io.WriterAt)        // the receive function - aka "some-new-WriterAt-var := <-MyKind"
	TryWriterAt() (dat io.WriterAt, open bool) // the multi-valued comma-ok receive function - aka "some-new-WriterAt-var, ok := <-MyKind"
}

type WriterAtSOnlyChan interface { // send-only channel
	ProvideWriterAt(dat io.WriterAt) // the send function - aka "MyKind <- some WriterAt"
}

type SChWriterAt struct { // supply channel
	dat chan io.WriterAt
	// req chan struct{}
}

func MakeSupplyWriterAtChan() *SChWriterAt {
	d := new(SChWriterAt)
	d.dat = make(chan io.WriterAt)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyWriterAtBuff(cap int) *SChWriterAt {
	d := new(SChWriterAt)
	d.dat = make(chan io.WriterAt, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriterAt is the send function - aka "MyKind <- some WriterAt"
func (c *SChWriterAt) ProvideWriterAt(dat io.WriterAt) {
	// .req
	c.dat <- dat
}

// RequestWriterAt is the receive function - aka "some WriterAt <- MyKind"
func (c *SChWriterAt) RequestWriterAt() (dat io.WriterAt) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriterAt is the comma-ok multi-valued form of RequestWriterAt and
// reports whether a received value was sent before the WriterAt channel was closed.
func (c *SChWriterAt) TryWriterAt() (dat io.WriterAt, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len