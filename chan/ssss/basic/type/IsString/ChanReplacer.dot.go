// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"strings"
)

// MakeReplacerChan returns a new open channel
// (simply a 'chan *strings.Replacer' that is).
//
// Note: No 'Replacer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReplacerPipelineStartsHere := MakeReplacerChan()
//	// ... lot's of code to design and build Your favourite "myReplacerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReplacerPipelineStartsHere <- drop
//	}
//	close(myReplacerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReplacerBuffer) the channel is unbuffered.
//
func MakeReplacerChan() chan *strings.Replacer {
	return make(chan *strings.Replacer)
}

// ChanReplacer returns a channel to receive all inputs before close.
func ChanReplacer(inp ...*strings.Replacer) chan *strings.Replacer {
	out := make(chan *strings.Replacer)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanReplacerSlice returns a channel to receive all inputs before close.
func ChanReplacerSlice(inp ...[]*strings.Replacer) chan *strings.Replacer {
	out := make(chan *strings.Replacer)
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

// JoinReplacer
func JoinReplacer(out chan<- *strings.Replacer, inp ...*strings.Replacer) chan struct{} {
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

// JoinReplacerSlice
func JoinReplacerSlice(out chan<- *strings.Replacer, inp ...[]*strings.Replacer) chan struct{} {
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

// JoinReplacerChan
func JoinReplacerChan(out chan<- *strings.Replacer, inp <-chan *strings.Replacer) chan struct{} {
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

// DoneReplacer returns a channel to receive one signal before close after inp has been drained.
func DoneReplacer(inp <-chan *strings.Replacer) chan struct{} {
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

// DoneReplacerSlice returns a channel which will receive a slice
// of all the Replacers received on inp channel before close.
// Unlike DoneReplacer, a full slice is sent once, not just an event.
func DoneReplacerSlice(inp <-chan *strings.Replacer) chan []*strings.Replacer {
	done := make(chan []*strings.Replacer)
	go func() {
		defer close(done)
		ReplacerS := []*strings.Replacer{}
		for i := range inp {
			ReplacerS = append(ReplacerS, i)
		}
		done <- ReplacerS
	}()
	return done
}

// DoneReplacerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReplacerFunc(inp <-chan *strings.Replacer, act func(a *strings.Replacer)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *strings.Replacer) { return }
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

// PipeReplacerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReplacerBuffer(inp <-chan *strings.Replacer, cap int) chan *strings.Replacer {
	out := make(chan *strings.Replacer, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReplacerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReplacerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReplacerFunc(inp <-chan *strings.Replacer, act func(a *strings.Replacer) *strings.Replacer) chan *strings.Replacer {
	out := make(chan *strings.Replacer)
	if act == nil {
		act = func(a *strings.Replacer) *strings.Replacer { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReplacerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReplacerFork(inp <-chan *strings.Replacer) (chan *strings.Replacer, chan *strings.Replacer) {
	out1 := make(chan *strings.Replacer)
	out2 := make(chan *strings.Replacer)
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