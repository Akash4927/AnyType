// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type PipeReaderChan interface { // bidirectional channel
	PipeReaderROnlyChan // aka "<-chan" - receive only
	PipeReaderSOnlyChan // aka "chan<-" - send only
}

type PipeReaderROnlyChan interface { // receive-only channel
	RequestPipeReader() (dat *io.PipeReader)        // the receive function - aka "some-new-PipeReader-var := <-MyKind"
	TryPipeReader() (dat *io.PipeReader, open bool) // the multi-valued comma-ok receive function - aka "some-new-PipeReader-var, ok := <-MyKind"
}

type PipeReaderSOnlyChan interface { // send-only channel
	ProvidePipeReader(dat *io.PipeReader) // the send function - aka "MyKind <- some PipeReader"
}

type DChPipeReader struct { // demand channel
	dat chan *io.PipeReader
	req chan struct{}
}

func MakeDemandPipeReaderChan() *DChPipeReader {
	d := new(DChPipeReader)
	d.dat = make(chan *io.PipeReader)
	d.req = make(chan struct{})
	return d
}

func MakeDemandPipeReaderBuff(cap int) *DChPipeReader {
	d := new(DChPipeReader)
	d.dat = make(chan *io.PipeReader, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvidePipeReader is the send function - aka "MyKind <- some PipeReader"
func (c *DChPipeReader) ProvidePipeReader(dat *io.PipeReader) {
	<-c.req
	c.dat <- dat
}

// RequestPipeReader is the receive function - aka "some PipeReader <- MyKind"
func (c *DChPipeReader) RequestPipeReader() (dat *io.PipeReader) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryPipeReader is the comma-ok multi-valued form of RequestPipeReader and
// reports whether a received value was sent before the PipeReader channel was closed.
func (c *DChPipeReader) TryPipeReader() (dat *io.PipeReader, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
