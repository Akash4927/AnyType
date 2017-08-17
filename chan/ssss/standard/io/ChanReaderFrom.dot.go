// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReaderFromChan returns a new open channel
// (simply a 'chan io.ReaderFrom' that is).
//
// Note: No 'ReaderFrom-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReaderFromPipelineStartsHere := MakeReaderFromChan()
//	// ... lot's of code to design and build Your favourite "myReaderFromWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReaderFromPipelineStartsHere <- drop
//	}
//	close(myReaderFromPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReaderFromBuffer) the channel is unbuffered.
//
func MakeReaderFromChan() chan io.ReaderFrom {
	return make(chan io.ReaderFrom)
}

// ChanReaderFrom returns a channel to receive all inputs before close.
func ChanReaderFrom(inp ...io.ReaderFrom) chan io.ReaderFrom {
	out := make(chan io.ReaderFrom)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanReaderFromSlice returns a channel to receive all inputs before close.
func ChanReaderFromSlice(inp ...[]io.ReaderFrom) chan io.ReaderFrom {
	out := make(chan io.ReaderFrom)
	go func() {
		defer close(out)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
	}()
	return out
}

// JoinReaderFrom
func JoinReaderFrom(out chan<- io.ReaderFrom, inp ...io.ReaderFrom) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// JoinReaderFromSlice
func JoinReaderFromSlice(out chan<- io.ReaderFrom, inp ...[]io.ReaderFrom) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, in := range inp {
			for _, i := range in {
				out <- i
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinReaderFromChan
func JoinReaderFromChan(out chan<- io.ReaderFrom, inp <-chan io.ReaderFrom) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneReaderFrom returns a channel to receive one signal before close after inp has been drained.
func DoneReaderFrom(inp <-chan io.ReaderFrom) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneReaderFromSlice returns a channel which will receive a slice
// of all the ReaderFroms received on inp channel before close.
// Unlike DoneReaderFrom, a full slice is sent once, not just an event.
func DoneReaderFromSlice(inp <-chan io.ReaderFrom) chan []io.ReaderFrom {
	done := make(chan []io.ReaderFrom)
	go func() {
		defer close(done)
		ReaderFromS := []io.ReaderFrom{}
		for i := range inp {
			ReaderFromS = append(ReaderFromS, i)
		}
		done <- ReaderFromS
	}()
	return done
}

// DoneReaderFromFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReaderFromFunc(inp <-chan io.ReaderFrom, act func(a io.ReaderFrom)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ReaderFrom) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeReaderFromBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReaderFromBuffer(inp <-chan io.ReaderFrom, cap int) chan io.ReaderFrom {
	out := make(chan io.ReaderFrom, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReaderFromFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReaderFromMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReaderFromFunc(inp <-chan io.ReaderFrom, act func(a io.ReaderFrom) io.ReaderFrom) chan io.ReaderFrom {
	out := make(chan io.ReaderFrom)
	if act == nil {
		act = func(a io.ReaderFrom) io.ReaderFrom { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReaderFromFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReaderFromFork(inp <-chan io.ReaderFrom) (chan io.ReaderFrom, chan io.ReaderFrom) {
	out1 := make(chan io.ReaderFrom)
	out2 := make(chan io.ReaderFrom)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}
