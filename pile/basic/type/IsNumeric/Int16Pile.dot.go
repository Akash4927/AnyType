// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// Note: originally inspired by parts of "cmd/doc/dirs.go"

// Int16Pile is a structure for
// a lazily populated sequence (= slice)
// of items (of type `int16`)
// which are cached in a growing-only list.
// Next() traverses the Int16Pile.
// Reset() allows a new transversal from the beginning.
//
// Yoe may either
// traverse the Int16Pile lazily -following its (buffered) growth that is-
// or
// await the signal from Wait() before starting traversal.
//
// Note: Pile() may be used concurrently,
// Next() (and Reset) should be confinded to a single routine (thread),
// as the iteration is not concurrency safe.
type Int16Pile struct {
	pile   chan int16 // channel to receive further items
	list   []int16    // list of known items
	offset int        // index for Next()
}

// NewS returns a (pointer to a) fresh Int16Pile
// of items (of type `int16`)
// with size as initial capacity
// and
// with buff non-blocking Add's before respective Next's
func Int16New(size, buff int) *Int16Pile {
	pile := new(Int16Pile)
	pile.list = make([]int16, 0, size)
	pile.pile = make(chan int16, buff)
	return pile
}

// Reset puts the pile iterator `Next()` back at the beginning.
func (d *Int16Pile) Reset() {
	d.offset = 0
}

// Next returns the next item,
// or false iff the pile is exhausted.
// Next may block, awaiting another Pile(),
// iff the Int16Pile is not Closed().
func (d *Int16Pile) Next() (item int16, ok bool) {
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
// an item (of type `int16`)
// to the Int16Pile.
//
// Note: Pile() may block, iff buff is exceeded and no corresponding Next()'s were called.
func (d *Int16Pile) Pile(item int16) {
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
func (d *Int16Pile) Close() {
	close(d.pile)
}

// Wait returns a done channel which emits the size (=length) of the Int16Pile once it's been closed.
// Users of Wait() must not iterate (via Next() and Reset()) before the returned done-channel is closed!
//
// Wait is a convenience - useful iff You do not like/need to start any traversal before the Int16Pile is fully populated.
// (Or iff You just like to know the size to be traversed, i.e. in order to allocate some traversal-related structure.)
//
// Note: Upon close of the done-channel, the Int16Pile is Reset() so You may start traversing it (via Next) right away.
func (d *Int16Pile) Wait() (done <-chan int) {
	cha := make(chan int)
	go func(cha chan<- int, d *Int16Pile) {
		defer close(cha)
		d.Reset()
		defer d.Reset()
		for {
			_, ok := d.Next() // keep draining
			if !ok {
				break
			}
		}
		cha <- len(d.list) // signal the lenght, and terminate
	}(cha, d)
	return cha
}