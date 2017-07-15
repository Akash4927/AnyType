// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

type ElementChan interface { // bidirectional channel
	ElementROnlyChan // aka "<-chan" - receive only
	ElementSOnlyChan // aka "chan<-" - send only
}

type ElementROnlyChan interface { // receive-only channel
	RequestElement() (dat list.Element)        // the receive function - aka "some-new-Element-var := <-MyKind"
	TryElement() (dat list.Element, open bool) // the multi-valued comma-ok receive function - aka "some-new-Element-var, ok := <-MyKind"
}

type ElementSOnlyChan interface { // send-only channel
	ProvideElement(dat list.Element) // the send function - aka "MyKind <- some Element"
}

type DChElement struct { // demand channel
	dat chan list.Element
	req chan struct{}
}

func MakeDemandElementChan() *DChElement {
	d := new(DChElement)
	d.dat = make(chan list.Element)
	d.req = make(chan struct{})
	return d
}

func MakeDemandElementBuff(cap int) *DChElement {
	d := new(DChElement)
	d.dat = make(chan list.Element, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideElement is the send function - aka "MyKind <- some Element"
func (c *DChElement) ProvideElement(dat list.Element) {
	<-c.req
	c.dat <- dat
}

// RequestElement is the receive function - aka "some Element <- MyKind"
func (c *DChElement) RequestElement() (dat list.Element) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryElement is the comma-ok multi-valued form of RequestElement and
// reports whether a received value was sent before the Element channel was closed.
func (c *DChElement) TryElement() (dat list.Element, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len