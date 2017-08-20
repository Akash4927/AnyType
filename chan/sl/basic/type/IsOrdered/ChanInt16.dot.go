// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Int16Chan represents a
// bidirectional
// channel
type Int16Chan interface {
	Int16ROnlyChan // aka "<-chan" - receive only
	Int16SOnlyChan // aka "chan<-" - send only
}

// Int16ROnlyChan represents a
// receive-only
// channel
type Int16ROnlyChan interface {
	RequestInt16() (dat int16)        // the receive function - aka "MyInt16 := <-MyInt16ROnlyChan"
	TryInt16() (dat int16, open bool) // the multi-valued comma-ok receive function - aka "MyInt16, ok := <-MyInt16ROnlyChan"
}

// Int16SOnlyChan represents a
// send-only
// channel
type Int16SOnlyChan interface {
	ProvideInt16(dat int16) // the send function - aka "MyKind <- some Int16"
}

// DChInt16 is a demand channel
type DChInt16 struct {
	dat chan int16
	req chan struct{}
}

// MakeDemandInt16Chan() returns
// a (pointer to a) fresh
// unbuffered
// demand channel
func MakeDemandInt16Chan() *DChInt16 {
	d := new(DChInt16)
	d.dat = make(chan int16)
	d.req = make(chan struct{})
	return d
}

// MakeDemandInt16Buff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// demand channel
func MakeDemandInt16Buff(cap int) *DChInt16 {
	d := new(DChInt16)
	d.dat = make(chan int16, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt16 is the send function - aka "MyKind <- some Int16"
func (c *DChInt16) ProvideInt16(dat int16) {
	<-c.req
	c.dat <- dat
}

// RequestInt16 is the receive function - aka "some Int16 <- MyKind"
func (c *DChInt16) RequestInt16() (dat int16) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt16 is the comma-ok multi-valued form of RequestInt16 and
// reports whether a received value was sent before the Int16 channel was closed.
func (c *DChInt16) TryInt16() (dat int16, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// DChInt16 is a supply channel
type SChInt16 struct {
	dat chan int16
	// req chan struct{}
}

// MakeSupplyInt16Chan() returns
// a (pointer to a) fresh
// unbuffered
// supply channel
func MakeSupplyInt16Chan() *SChInt16 {
	d := new(SChInt16)
	d.dat = make(chan int16)
	// d.req = make(chan struct{})
	return d
}

// MakeSupplyInt16Buff() returns
// a (pointer to a) fresh
// buffered (with capacity cap)
// supply channel
func MakeSupplyInt16Buff(cap int) *SChInt16 {
	d := new(SChInt16)
	d.dat = make(chan int16, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideInt16 is the send function - aka "MyKind <- some Int16"
func (c *SChInt16) ProvideInt16(dat int16) {
	// .req
	c.dat <- dat
}

// RequestInt16 is the receive function - aka "some Int16 <- MyKind"
func (c *SChInt16) RequestInt16() (dat int16) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryInt16 is the comma-ok multi-valued form of RequestInt16 and
// reports whether a received value was sent before the Int16 channel was closed.
func (c *SChInt16) TryInt16() (dat int16, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len

// MergeInt16 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Int16channel is returned.
func MergeInt16(inps ...<-chan int16) (out <-chan int16) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan int16)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeInt162(inps[0], MergeInt16(inps[1:]...))
	}
}

// mergeInt162 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeInt162(i1, i2 <-chan int16) (out <-chan int16) {
	cha := make(chan int16)
	go func(out chan<- int16, i1, i2 <-chan int16) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 int16 // what we've read
		)

		for !clos1 || !clos2 {

			if !clos1 && !buff1 {
				if from1, ok = <-i1; ok {
					buff1 = true
				} else {
					clos1 = true
				}
			}

			if !clos2 && !buff2 {
				if from2, ok = <-i2; ok {
					buff2 = true
				} else {
					clos2 = true
				}
			}

			if clos1 && !buff1 {
				from1 = from2
			}
			if clos2 && !buff2 {
				from2 = from1
			}

			if from1 < from2 {
				out <- from1
				buff1 = false
			} else if from2 < from1 {
				out <- from2
				buff2 = false
			} else {
				out <- from1 // == from2
				buff1 = false
				buff2 = false
			}
		}
	}(cha, i1, i2)
	return cha
}

// Note: merge2 is not my own. Just: I forgot where found it - please accept my apologies.
// I'd love to learn about it's origin/author, so I can give credit. Any hint is highly appreciated!
