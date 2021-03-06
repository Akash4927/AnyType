// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/tag"
)

// TagChan represents a
// bidirectional
// channel
type TagChan interface {
	TagROnlyChan // aka "<-chan" - receive only
	TagSOnlyChan // aka "chan<-" - send only
}

// TagROnlyChan represents a
// receive-only
// channel
type TagROnlyChan interface {
	RequestTag() (dat *tag.TagAny)        // the receive function - aka "MyTag := <-MyTagROnlyChan"
	TryTag() (dat *tag.TagAny, open bool) // the multi-valued comma-ok receive function - aka "MyTag, ok := <-MyTagROnlyChan"
}

// TagSOnlyChan represents a
// send-only
// channel
type TagSOnlyChan interface {
	ProvideTag(dat *tag.TagAny) // the send function - aka "MyKind <- some Tag"
}

// SChTag is a supply channel
type SChTag struct {
	dat chan *tag.TagAny
	// req chan struct{}
}

// MakeSupplyTagChan returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyTagChan() *SChTag {
	d := new(SChTag)
	d.dat = make(chan *tag.TagAny)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyTagBuff returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyTagBuff(cap int) *SChTag {
	d := new(SChTag)
	d.dat = make(chan *tag.TagAny, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideTag is the send function - aka "MyKind <- some Tag"
func (c *SChTag) ProvideTag(dat *tag.TagAny) {
	// .req
	c.dat <- dat
}

// RequestTag is the receive function - aka "some Tag <- MyKind"
func (c *SChTag) RequestTag() (dat *tag.TagAny) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryTag is the comma-ok multi-valued form of RequestTag and
// reports whether a received value was sent before the Tag channel was closed.
func (c *SChTag) TryTag() (dat *tag.TagAny, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
