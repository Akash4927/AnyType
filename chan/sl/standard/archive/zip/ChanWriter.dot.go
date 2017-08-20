// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

// WriterChan represents a
// bidirectional
// channel
type WriterChan interface {
	WriterROnlyChan // aka "<-chan" - receive only
	WriterSOnlyChan // aka "chan<-" - send only
}

// WriterROnlyChan represents a
// receive-only
// channel
type WriterROnlyChan interface {
	RequestWriter() (dat zip.Writer)        // the receive function - aka "MyWriter := <-MyWriterROnlyChan"
	TryWriter() (dat zip.Writer, open bool) // the multi-valued comma-ok receive function - aka "MyWriter, ok := <-MyWriterROnlyChan"
}

// WriterSOnlyChan represents a
// send-only
// channel
type WriterSOnlyChan interface {
	ProvideWriter(dat zip.Writer) // the send function - aka "MyKind <- some Writer"
}

// DChWriter is a demand channel
type DChWriter struct {
	dat chan zip.Writer
	req chan struct{}
}

// MakeDemandWriterChan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandWriterChan() *DChWriter {
	d := new(DChWriter)
	d.dat = make(chan zip.Writer)
	d.req = make(chan struct{})
	return d
}

// MakeDemandWriterBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandWriterBuff(cap int) *DChWriter {
	d := new(DChWriter)
	d.dat = make(chan zip.Writer, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideWriter is the send function - aka "MyKind <- some Writer"
func (c *DChWriter) ProvideWriter(dat zip.Writer) {
	<-c.req
	c.dat <- dat
}

// RequestWriter is the receive function - aka "some Writer <- MyKind"
func (c *DChWriter) RequestWriter() (dat zip.Writer) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryWriter is the comma-ok multi-valued form of RequestWriter and
// reports whether a received value was sent before the Writer channel was closed.
func (c *DChWriter) TryWriter() (dat zip.Writer, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// DChWriter is a supply channel
type SChWriter struct {
	dat chan zip.Writer
	// req chan struct{}
}

// MakeSupplyWriterChan() returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyWriterChan() *SChWriter {
	d := new(SChWriter)
	d.dat = make(chan zip.Writer)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyWriterBuff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyWriterBuff(cap int) *SChWriter {
	d := new(SChWriter)
	d.dat = make(chan zip.Writer, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideWriter is the send function - aka "MyKind <- some Writer"
func (c *SChWriter) ProvideWriter(dat zip.Writer) {
	// .req
	c.dat <- dat
}

// RequestWriter is the receive function - aka "some Writer <- MyKind"
func (c *SChWriter) RequestWriter() (dat zip.Writer) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryWriter is the comma-ok multi-valued form of RequestWriter and
// reports whether a received value was sent before the Writer channel was closed.
func (c *SChWriter) TryWriter() (dat zip.Writer, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
