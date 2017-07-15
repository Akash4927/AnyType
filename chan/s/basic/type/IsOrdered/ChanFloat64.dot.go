// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Float64Chan interface { // bidirectional channel
	Float64ROnlyChan // aka "<-chan" - receive only
	Float64SOnlyChan // aka "chan<-" - send only
}

type Float64ROnlyChan interface { // receive-only channel
	RequestFloat64() (dat float64)        // the receive function - aka "some-new-Float64-var := <-MyKind"
	TryFloat64() (dat float64, open bool) // the multi-valued comma-ok receive function - aka "some-new-Float64-var, ok := <-MyKind"
}

type Float64SOnlyChan interface { // send-only channel
	ProvideFloat64(dat float64) // the send function - aka "MyKind <- some Float64"
}

type SChFloat64 struct { // supply channel
	dat chan float64
	// req chan struct{}
}

func MakeSupplyFloat64Chan() *SChFloat64 {
	d := new(SChFloat64)
	d.dat = make(chan float64)
	// d.req = make(chan struct{})
	return d
}

func MakeSupplyFloat64Buff(cap int) *SChFloat64 {
	d := new(SChFloat64)
	d.dat = make(chan float64, cap)
	// eq = make(chan struct{}, cap)
	return d
}

// ProvideFloat64 is the send function - aka "MyKind <- some Float64"
func (c *SChFloat64) ProvideFloat64(dat float64) {
	// .req
	c.dat <- dat
}

// RequestFloat64 is the receive function - aka "some Float64 <- MyKind"
func (c *SChFloat64) RequestFloat64() (dat float64) {
	// eq <- struct{}{}
	return <-c.dat
}

// TryFloat64 is the comma-ok multi-valued form of RequestFloat64 and
// reports whether a received value was sent before the Float64 channel was closed.
func (c *SChFloat64) TryFloat64() (dat float64, open bool) {
	// eq <- struct{}{}
	dat, open = <-c.dat
	return dat, open
}

// TODO(apa): close, cap & len
// MergeFloat642 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func MergeFloat642(i1, i2 <-chan float64) (out <-chan float64) {
	cha := make(chan float64)
	go func(out chan<- float64, i1, i2 <-chan float64) {
		defer close(out)
		var (
			clos1, clos2 bool    // we found the chan closed
			buff1, buff2 bool    // we've read 'from', but not sent (yet)
			ok           bool    // did we read sucessfully?
			from1, from2 float64 // what we've read
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