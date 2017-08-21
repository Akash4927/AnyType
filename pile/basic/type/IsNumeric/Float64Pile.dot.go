// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: originally inspired by parts of "cmd/doc/dirs.go"

// Float64Pile is a structure for
// a lazily populated sequence (= slice)
// of items (of type `float64`)
// which are cached in a growing-only list.
// Next() traverses the Float64Pile.
// Reset() allows a new transversal from the beginning.
//
// You may either
// traverse the Float64Pile lazily -following its (buffered) growth that is-
// or
// await the signal from Wait() before starting traversal.
//
// Note: Pile() may be used concurrently,
// Next() (and Reset) should be confined to a single go routine (thread),
// as the iteration is not intended to by concurrency safe.
type Float64Pile struct {
	pile   chan float64 // channel to receive further items
	list   []float64    // list of known items
	offset int          // index for Next()
}

// MakeFloat64Pile returns a (pointer to a) fresh pile
// of items (of type `float64`)
// with size as initial capacity
// and
// with buff non-blocking Add's before respective Next's
func MakeFloat64Pile(size, buff int) *Float64Pile {
	pile := new(Float64Pile)
	pile.list = make([]float64, 0, size)
	pile.pile = make(chan float64, buff)
	return pile
}

// Reset puts the pile iterator `Next()` back at the beginning.
func (d *Float64Pile) Reset() {
	d.offset = 0
}

// Next returns the next item,
// or false iff the pile is exhausted.
// Next may block, awaiting another Pile(),
// iff the pile is not Closed().
func (d *Float64Pile) Next() (item float64, ok bool) {
	if d.offset < len(d.list) {
		ok = true
		item = d.list[d.offset]
		d.offset++
	} else if item, ok = <-d.pile; ok {
		d.list = append(d.list, item)
		d.offset++
	}
	return item, ok
}

// Pile adds
// an item (of type `float64`)
// to the Float64Pile.
//
// Note: Pile() may block, iff buff is exceeded and no corresponding Next()'s were called.
func (d *Float64Pile) Pile(item float64) {
	d.pile <- item
}

// Close - call once after everything has been piled.
//
// Note: After Close(),
// any Close(...) will panic
// and
// any Pile(...) will panic
// and
// any Next() will return immediately: no eventual blocking, that is.
func (d *Float64Pile) Close() {
	close(d.pile)
}

// Wait returns a done channel which emits the size (=length) of the pile once it's been closed.
//
// Users of Wait() *must not* iterate (via Next() or Reset()) before the done-channel is closed!
//
// Wait is a convenience - useful iff You do not like/need to start any traversal before the pile is fully populated.
// (Or iff You just like to know the size to be traversed, i.e. in order to allocate some traversal-related structure.
// Once the pile is closed, Wait() will return in constant time.)
//
// Note: Upon close of the done-channel, the pile is Reset() so You may traverse it (via Next) right away.
func (d *Float64Pile) Wait() (done <-chan int) {
	cha := make(chan int)
	go func(cha chan<- int, d *Float64Pile) {
		defer close(cha)
		d.Reset()
		if len(d.list) > 0 { // skip what's already known
			d.offset = len(d.list) - 1
		}
		defer d.Reset()
		for {
			_, ok := d.Next() // keep draining
			if !ok {
				break
			}
		}
		cha <- len(d.list) // signal the length, and terminate
	}(cha, d)
	return cha
}
