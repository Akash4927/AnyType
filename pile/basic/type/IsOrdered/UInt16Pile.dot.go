// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsOrdered

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: originally inspired by parts of "cmd/doc/dirs.go"

// UInt16Pile is a structure for
// a lazily populated sequence (= slice)
// of items (of type `uint16`)
// which are cached in a growing-only list.
// Next() traverses the UInt16Pile.
// Reset() allows a new transversal from the beginning.
//
// You may either
// traverse the UInt16Pile lazily -following its (buffered) growth that is-
// or
// await the signal from Wait() before starting traversal.
//
// Note: Pile() may be used concurrently,
// Next() (and Reset) should be confined to a single go routine (thread),
// as the iteration is not intended to by concurrency safe.
type UInt16Pile struct {
	pile   chan uint16 // channel to receive further items
	list   []uint16    // list of known items
	offset int         // index for Next()
}

// MakeUInt16Pile returns a (pointer to a) fresh pile
// of items (of type `uint16`)
// with size as initial capacity
// and
// with buff non-blocking Add's before respective Next's
func MakeUInt16Pile(size, buff int) *UInt16Pile {
	pile := new(UInt16Pile)
	pile.list = make([]uint16, 0, size)
	pile.pile = make(chan uint16, buff)
	return pile
}

// Reset puts the pile iterator `Next()` back at the beginning.
func (d *UInt16Pile) Reset() {
	d.offset = 0
}

// Next returns the next item,
// or false iff the pile is exhausted.
// Next may block, awaiting another Pile(),
// iff the pile is not Closed().
func (d *UInt16Pile) Next() (item uint16, ok bool) {
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
// an item (of type `uint16`)
// to the UInt16Pile.
//
// Note: Pile() may block, iff buff is exceeded and no corresponding Next()'s were called.
func (d *UInt16Pile) Pile(item uint16) {
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
func (d *UInt16Pile) Close() {
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
func (d *UInt16Pile) Wait() (done <-chan int) {
	cha := make(chan int)
	go func(cha chan<- int, d *UInt16Pile) {
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
