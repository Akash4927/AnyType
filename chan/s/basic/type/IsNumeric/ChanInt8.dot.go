// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int8Chan interface { // bidirectional channel
	Int8ROnlyChan // aka "<-chan" - receive only
	Int8SOnlyChan // aka "chan<-" - send only
}

type Int8ROnlyChan interface { // receive-only channel
	RequestInt8() (dat int8)        // the receive function - aka "some-new-Int8-var := <-MyKind"
	TryInt8() (dat int8, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int8-var, ok := <-MyKind"
}

type Int8SOnlyChan interface { // send-only channel
	ProvideInt8(dat int8) // the send function - aka "MyKind <- some Int8"
}

type SChInt8 struct { // supply channel
	dat chan int8
	// req chan struct{}
}

func MakeSupplyInt8Chan() *SChInt8 {
	d := new(SChInt8)
	d.dat = make(chan int8)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyInt8Buff(cap int) *SChInt8 {
	d := new(SChInt8)
	d.dat = make(chan int8, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt8 is the send function - aka "MyKind <- some Int8"
func (c *SChInt8) ProvideInt8(dat int8) {
	// .req
	c.dat <- dat
}

// RequestInt8 is the receive function - aka "some Int8 <- MyKind"
func (c *SChInt8) RequestInt8() (dat int8) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt8 is the comma-ok multi-valued form of RequestInt8 and
// reports whether a received value was sent before the Int8 channel was closed.
func (c *SChInt8) TryInt8() (dat int8, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len