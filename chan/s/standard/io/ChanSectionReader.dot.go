// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// SectionReaderChan represents a
// bidirectional
// channel
type SectionReaderChan interface {
	SectionReaderROnlyChan // aka "<-chan" - receive only
	SectionReaderSOnlyChan // aka "chan<-" - send only
}

// SectionReaderROnlyChan represents a
// receive-only
// channel
type SectionReaderROnlyChan interface {
	RequestSectionReader() (dat *io.SectionReader)        // the receive function - aka "MySectionReader := <-MySectionReaderROnlyChan"
	TrySectionReader() (dat *io.SectionReader, open bool) // the multi-valued comma-ok receive function - aka "MySectionReader, ok := <-MySectionReaderROnlyChan"
}

// SectionReaderSOnlyChan represents a
// send-only
// channel
type SectionReaderSOnlyChan interface {
	ProvideSectionReader(dat *io.SectionReader) // the send function - aka "MyKind <- some SectionReader"
}

// SChSectionReader is a supply channel
type SChSectionReader struct {
	dat chan *io.SectionReader
	// req chan struct{}
}

// MakeSupplySectionReaderChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplySectionReaderChan() *SChSectionReader {
	d := new(SChSectionReader)
	d.dat = make(chan *io.SectionReader)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplySectionReaderBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplySectionReaderBuff(cap int) *SChSectionReader {
	d := new(SChSectionReader)
	d.dat = make(chan *io.SectionReader, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideSectionReader is the send function - aka "MyKind <- some SectionReader"
func (c *SChSectionReader) ProvideSectionReader(dat *io.SectionReader) {
	// .req
	c.dat <- dat
}

// RequestSectionReader is the receive function - aka "some SectionReader <- MyKind"
func (c *SChSectionReader) RequestSectionReader() (dat *io.SectionReader) {
	// eq <- struct{}{}
	return <-c.dat
}

// TrySectionReader is the comma-ok multi-valued form of RequestSectionReader and
// reports whether a received value was sent before the SectionReader channel was closed.
func (c *SChSectionReader) TrySectionReader() (dat *io.SectionReader, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
