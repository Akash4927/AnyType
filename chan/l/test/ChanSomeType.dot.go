// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// SomeTypeChan represents a
// bidirectional
// channel
type SomeTypeChan interface {
	SomeTypeROnlyChan // aka "<-chan" - receive only
	SomeTypeSOnlyChan // aka "chan<-" - send only
}

// SomeTypeROnlyChan represents a
// receive-only
// channel
type SomeTypeROnlyChan interface {
	RequestSomeType() (dat SomeType)        // the receive function - aka "MySomeType := <-MySomeTypeROnlyChan"
	TrySomeType() (dat SomeType, open bool) // the multi-valued comma-ok receive function - aka "MySomeType, ok := <-MySomeTypeROnlyChan"
}

// SomeTypeSOnlyChan represents a
// send-only
// channel
type SomeTypeSOnlyChan interface {
	ProvideSomeType(dat SomeType) // the send function - aka "MyKind <- some SomeType"
}

// DChSomeType is a demand channel
type DChSomeType struct {
	dat chan SomeType
	req chan struct{}
}

// MakeDemandSomeTypeChan returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandSomeTypeChan() *DChSomeType {
	d := new(DChSomeType)
	d.dat = make(chan SomeType)
	d.req = make(chan struct{})
	return d
}

// MakeDemandSomeTypeBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandSomeTypeBuff(cap int) *DChSomeType {
	d := new(DChSomeType)
	d.dat = make(chan SomeType, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideSomeType is the send function - aka "MyKind <- some SomeType"
func (c *DChSomeType) ProvideSomeType(dat SomeType) {
	<-c.req
	c.dat <- dat
}

// RequestSomeType is the receive function - aka "some SomeType <- MyKind"
func (c *DChSomeType) RequestSomeType() (dat SomeType) {
	c.req <- struct{}{}
	return <-c.dat
}

// TrySomeType is the comma-ok multi-valued form of RequestSomeType and
// reports whether a received value was sent before the SomeType channel was closed.
func (c *DChSomeType) TrySomeType() (dat SomeType, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
