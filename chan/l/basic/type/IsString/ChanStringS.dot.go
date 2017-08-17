// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type StringSChan interface { // bidirectional channel
	StringSROnlyChan // aka "<-chan" - receive only
	StringSSOnlyChan // aka "chan<-" - send only
}

type StringSROnlyChan interface { // receive-only channel
	RequestStringS() (dat []string)        // the receive function - aka "some-new-StringS-var := <-MyKind"
	TryStringS() (dat []string, open bool) // the multi-valued comma-ok receive function - aka "some-new-StringS-var, ok := <-MyKind"
}

type StringSSOnlyChan interface { // send-only channel
	ProvideStringS(dat []string) // the send function - aka "MyKind <- some StringS"
}

type DChStringS struct { // demand channel
	dat chan []string
	req chan struct{}
}

func MakeDemandStringSChan() *DChStringS {
	d := new(DChStringS)
	d.dat = make(chan []string)
	d.req = make(chan struct{})
	return d
}

func MakeDemandStringSBuff(cap int) *DChStringS {
	d := new(DChStringS)
	d.dat = make(chan []string, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideStringS is the send function - aka "MyKind <- some StringS"
func (c *DChStringS) ProvideStringS(dat []string) {
	<-c.req
	c.dat <- dat
}

// RequestStringS is the receive function - aka "some StringS <- MyKind"
func (c *DChStringS) RequestStringS() (dat []string) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryStringS is the comma-ok multi-valued form of RequestStringS and
// reports whether a received value was sent before the StringS channel was closed.
func (c *DChStringS) TryStringS() (dat []string, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
