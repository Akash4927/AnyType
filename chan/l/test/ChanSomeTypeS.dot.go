// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type SomeTypeSChan interface { // bidirectional channel
	SomeTypeSROnlyChan // aka "<-chan" - receive only
	SomeTypeSSOnlyChan // aka "chan<-" - send only
}

type SomeTypeSROnlyChan interface { // receive-only channel
	RequestSomeTypeS() (dat []SomeType)        // the receive function - aka "some-new-SomeTypeS-var := <-MyKind"
	TrySomeTypeS() (dat []SomeType, open bool) // the multi-valued comma-ok receive function - aka "some-new-SomeTypeS-var, ok := <-MyKind"
}

type SomeTypeSSOnlyChan interface { // send-only channel
	ProvideSomeTypeS(dat []SomeType) // the send function - aka "MyKind <- some SomeTypeS"
}

type DChSomeTypeS struct { // demand channel
	dat chan []SomeType
	req chan struct{}
}

func MakeDemandSomeTypeSChan() *DChSomeTypeS {
	d := new(DChSomeTypeS)
	d.dat = make(chan []SomeType)
	d.req = make(chan struct{})
	return d
}

func MakeDemandSomeTypeSBuff(cap int) *DChSomeTypeS {
	d := new(DChSomeTypeS)
	d.dat = make(chan []SomeType, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideSomeTypeS is the send function - aka "MyKind <- some SomeTypeS"
func (c *DChSomeTypeS) ProvideSomeTypeS(dat []SomeType) {
	<-c.req
	c.dat <- dat
}

// RequestSomeTypeS is the receive function - aka "some SomeTypeS <- MyKind"
func (c *DChSomeTypeS) RequestSomeTypeS() (dat []SomeType) {
	c.req <- struct{}{}
	return <-c.dat
}

// TrySomeTypeS is the comma-ok multi-valued form of RequestSomeTypeS and
// reports whether a received value was sent before the SomeTypeS channel was closed.
func (c *DChSomeTypeS) TrySomeTypeS() (dat []SomeType, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len