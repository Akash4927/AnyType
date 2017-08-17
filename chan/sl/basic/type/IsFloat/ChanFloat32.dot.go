// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsFloat

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Float32Chan interface { // bidirectional channel
	Float32ROnlyChan // aka "<-chan" - receive only
	Float32SOnlyChan // aka "chan<-" - send only
}

type Float32ROnlyChan interface { // receive-only channel
	RequestFloat32() (dat float32)        // the receive function - aka "some-new-Float32-var := <-MyKind"
	TryFloat32() (dat float32, open bool) // the multi-valued comma-ok receive function - aka "some-new-Float32-var, ok := <-MyKind"
}

type Float32SOnlyChan interface { // send-only channel
	ProvideFloat32(dat float32) // the send function - aka "MyKind <- some Float32"
}

// MergeFloat32 returns a channel to receive all inputs sorted and free of duplicates.
// Each input channel needs to be ascending; sorted and free of duplicates.
//  Note: If no inputs are given, a closed Float32channel is returned.
func MergeFloat32(inps ...<-chan float32) (out <-chan float32) {

	if len(inps) < 1 { // none: return a closed channel
		cha := make(chan float32)
		defer close(cha)
		return cha
	} else if len(inps) < 2 { // just one: return it
		return inps[0]
	} else { // tail recurse
		return mergeFloat322(inps[0], MergeFloat32(inps[1:]...))
	}
}

// mergeFloat322 takes two (eager) channels of comparable types,
// each of which needs to be sorted and free of duplicates,
// and merges them into a returned channel, which will be sorted and free of duplicates
func mergeFloat322(i1, i2 <-chan float32) (out <-chan float32) {
	cha := make(chan float32)
	go func(out chan<- float32, i1, i2 <-chan float32) {
		defer close(out)
		var (
			clos1, clos2 bool    // we found the chan closed
			buff1, buff2 bool    // we've read 'from', but not sent (yet)
			ok           bool    // did we read sucessfully?
			from1, from2 float32 // what we've read
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
