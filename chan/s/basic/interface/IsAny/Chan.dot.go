// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsAny

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Chan interface { // bidirectional channel
	ROnlyChan // aka "<-chan" - receive only
	SOnlyChan // aka "chan<-" - send only
}

type ROnlyChan interface { // receive-only channel
	Request() (dat interface{})        // the receive function - aka "some-new--var := <-MyKind"
	Try() (dat interface{}, open bool) // the multi-valued comma-ok receive function - aka "some-new--var, ok := <-MyKind"
}

type SOnlyChan interface { // send-only channel
	Provide(dat interface{}) // the send function - aka "MyKind <- some "
}

type SCh struct { // supply channel
	dat chan interface{}
	// req chan struct{}
}

func MakeSupplyChan() *SCh {
	d := new(SCh)
	d.dat = make(chan interface{})
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyBuff(cap int) *SCh {
	d := new(SCh)
	d.dat = make(chan interface{}, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// Provide is the send function - aka "MyKind <- some "
func (c *SCh) Provide(dat interface{}) {
	// .req
	c.dat <- dat
}

// Request is the receive function - aka "some  <- MyKind"
func (c *SCh) Request() (dat interface{}) {
	// eq <- struct{}{}
	return <-c.dat
}

// Try is the comma-ok multi-valued form of Request and
// reports whether a received value was sent before the  channel was closed.
func (c *SCh) Try() (dat interface{}, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len