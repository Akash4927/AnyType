// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// ReadWriterChan represents a
// bidirectional
// channel
type ReadWriterChan interface {
	ReadWriterROnlyChan // aka "<-chan" - receive only
	ReadWriterSOnlyChan // aka "chan<-" - send only
}

// ReadWriterROnlyChan represents a
// receive-only
// channel
type ReadWriterROnlyChan interface {
	RequestReadWriter() (dat *bufio.ReadWriter)        // the receive function - aka "MyReadWriter := <-MyReadWriterROnlyChan"
	TryReadWriter() (dat *bufio.ReadWriter, open bool) // the multi-valued comma-ok receive function - aka "MyReadWriter, ok := <-MyReadWriterROnlyChan"
}

// ReadWriterSOnlyChan represents a
// send-only
// channel
type ReadWriterSOnlyChan interface {
	ProvideReadWriter(dat *bufio.ReadWriter) // the send function - aka "MyKind <- some ReadWriter"
}

// DChReadWriter is a demand channel
type DChReadWriter struct {
	dat chan *bufio.ReadWriter
	req chan struct{}
}

// MakeDemandReadWriterChan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandReadWriterChan() *DChReadWriter {
	d := new(DChReadWriter)
	d.dat = make(chan *bufio.ReadWriter)
	d.req = make(chan struct{})
	return d
}

// MakeDemandReadWriterBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandReadWriterBuff(cap int) *DChReadWriter {
	d := new(DChReadWriter)
	d.dat = make(chan *bufio.ReadWriter, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideReadWriter is the send function - aka "MyKind <- some ReadWriter"
func (c *DChReadWriter) ProvideReadWriter(dat *bufio.ReadWriter) {
	<-c.req
	c.dat <- dat
}

// RequestReadWriter is the receive function - aka "some ReadWriter <- MyKind"
func (c *DChReadWriter) RequestReadWriter() (dat *bufio.ReadWriter) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryReadWriter is the comma-ok multi-valued form of RequestReadWriter and
// reports whether a received value was sent before the ReadWriter channel was closed.
func (c *DChReadWriter) TryReadWriter() (dat *bufio.ReadWriter, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// DChReadWriter is a supply channel
type SChReadWriter struct {
	dat chan *bufio.ReadWriter
	// req chan struct{}
}

// MakeSupplyReadWriterChan() returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyReadWriterChan() *SChReadWriter {
	d := new(SChReadWriter)
	d.dat = make(chan *bufio.ReadWriter)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyReadWriterBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyReadWriterBuff(cap int) *SChReadWriter {
	d := new(SChReadWriter)
	d.dat = make(chan *bufio.ReadWriter, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideReadWriter is the send function - aka "MyKind <- some ReadWriter"
func (c *SChReadWriter) ProvideReadWriter(dat *bufio.ReadWriter) {
	// .req
	c.dat <- dat
}

// RequestReadWriter is the receive function - aka "some ReadWriter <- MyKind"
func (c *SChReadWriter) RequestReadWriter() (dat *bufio.ReadWriter) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryReadWriter is the comma-ok multi-valued form of RequestReadWriter and
// reports whether a received value was sent before the ReadWriter channel was closed.
func (c *SChReadWriter) TryReadWriter() (dat *bufio.ReadWriter, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
