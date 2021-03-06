// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pat

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// PatSChan represents a
// bidirectional
// channel
type PatSChan interface {
	PatSROnlyChan // aka "<-chan" - receive only
	PatSSOnlyChan // aka "chan<-" - send only
}

// PatSROnlyChan represents a
// receive-only
// channel
type PatSROnlyChan interface {
	RequestPatS() (dat []struct{})        // the receive function - aka "MyPatS := <-MyPatSROnlyChan"
	TryPatS() (dat []struct{}, open bool) // the multi-valued comma-ok receive function - aka "MyPatS, ok := <-MyPatSROnlyChan"
}

// PatSSOnlyChan represents a
// send-only
// channel
type PatSSOnlyChan interface {
	ProvidePatS(dat []struct{}) // the send function - aka "MyKind <- some PatS"
}

// SChPatS is a supply channel
type SChPatS struct {
	dat chan []struct{}
	// req chan struct{}
}

// MakeSupplyPatSChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyPatSChan() *SChPatS {
	d := new(SChPatS)
	d.dat = make(chan []struct{})
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyPatSBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyPatSBuff(cap int) *SChPatS {
	d := new(SChPatS)
	d.dat = make(chan []struct{}, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvidePatS is the send function - aka "MyKind <- some PatS"
func (c *SChPatS) ProvidePatS(dat []struct{}) {
	// .req
	c.dat <- dat
}

// RequestPatS is the receive function - aka "some PatS <- MyKind"
func (c *SChPatS) RequestPatS() (dat []struct{}) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryPatS is the comma-ok multi-valued form of RequestPatS and
// reports whether a received value was sent before the PatS channel was closed.
func (c *SChPatS) TryPatS() (dat []struct{}, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
