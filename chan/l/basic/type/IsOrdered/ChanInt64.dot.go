// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int64Chan interface { // bidirectional channel
	Int64ROnlyChan // aka "<-chan" - receive only
	Int64SOnlyChan // aka "chan<-" - send only
}

type Int64ROnlyChan interface { // receive-only channel
	RequestInt64() (dat int64)        // the receive function - aka "some-new-Int64-var := <-MyKind"
	TryInt64() (dat int64, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int64-var, ok := <-MyKind"
}

type Int64SOnlyChan interface { // send-only channel
	ProvideInt64(dat int64) // the send function - aka "MyKind <- some Int64"
}

type DChInt64 struct { // demand channel
	dat chan int64
	req chan struct{}
}

func MakeDemandInt64Chan() *DChInt64 {
	d := new(DChInt64)
	d.dat = make(chan int64)
	d.req = make(chan struct{})
	return d
}

func MakeDemandInt64Buff(cap int) *DChInt64 {
	d := new(DChInt64)
	d.dat = make(chan int64, cap)
	d.req = make(chan struct{}, cap)
	return d
}

// ProvideInt64 is the send function - aka "MyKind <- some Int64"
func (c *DChInt64) ProvideInt64(dat int64) {
	<-c.req
	c.dat <- dat
}

// RequestInt64 is the receive function - aka "some Int64 <- MyKind"
func (c *DChInt64) RequestInt64() (dat int64) {
	c.req <- struct{}{}
	return <-c.dat
}

// TryInt64 is the comma-ok multi-valued form of RequestInt64 and
// reports whether a received value was sent before the Int64 channel was closed.
func (c *DChInt64) TryInt64() (dat int64, open bool) {
	c.req <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
// MergeInt642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeInt642(i1, i2 <-chan int64) (out <-chan int64) {
	cha := make(chan int64)
	go func(out chan<- int64, i1, i2 <-chan int64) {
		defer close(out)
		var (
			clos1, clos2 bool  // we found the chan closed
			buff1, buff2 bool  // we've read 'from', but not sent (yet)
			ok           bool  // did we read sucessfully?
			from1, from2 int64 // what we've read
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