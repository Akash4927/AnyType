// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeCloserChan returns a new open channel
// (simply a 'chan io.Closer' that is).
//
// Note: No 'Closer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myCloserPipelineStartsHere := MakeCloserChan()
//	// ... lot's of code to design and build Your favourite "myCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myCloserPipelineStartsHere <- drop
//	}
//	close(myCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeCloserBuffer) the channel is unbuffered.
//
func MakeCloserChan() chan io.Closer {
	return make(chan io.Closer)
}

// ChanCloser returns a channel to receive all inputs before close.
func ChanCloser(inp ...io.Closer) chan io.Closer {
	out := make(chan io.Closer)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanCloserSlice returns a channel to receive all inputs before close.
func ChanCloserSlice(inp ...[]io.Closer) chan io.Closer {
	out := make(chan io.Closer)
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

// JoinCloser
func JoinCloser(out chan<- io.Closer, inp ...io.Closer) chan struct{} {
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

// JoinCloserSlice
func JoinCloserSlice(out chan<- io.Closer, inp ...[]io.Closer) chan struct{} {
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

// JoinCloserChan
func JoinCloserChan(out chan<- io.Closer, inp <-chan io.Closer) chan struct{} {
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

// DoneCloser returns a channel to receive one signal before close after inp has been drained.
func DoneCloser(inp <-chan io.Closer) chan struct{} {
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

// DoneCloserSlice returns a channel which will receive a slice
// of all the Closers received on inp channel before close.
// Unlike DoneCloser, a full slice is sent once, not just an event.
func DoneCloserSlice(inp <-chan io.Closer) chan []io.Closer {
	done := make(chan []io.Closer)
	go func() {
		defer close(done)
		CloserS := []io.Closer{}
		for i := range inp {
			CloserS = append(CloserS, i)
		}
		done <- CloserS
	}()
	return done
}

// DoneCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneCloserFunc(inp <-chan io.Closer, act func(a io.Closer)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.Closer) { return }
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

// PipeCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeCloserBuffer(inp <-chan io.Closer, cap int) chan io.Closer {
	out := make(chan io.Closer, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeCloserFunc(inp <-chan io.Closer, act func(a io.Closer) io.Closer) chan io.Closer {
	out := make(chan io.Closer)
	if act == nil {
		act = func(a io.Closer) io.Closer { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeCloserFork(inp <-chan io.Closer) (chan io.Closer, chan io.Closer) {
	out1 := make(chan io.Closer)
	out2 := make(chan io.Closer)
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