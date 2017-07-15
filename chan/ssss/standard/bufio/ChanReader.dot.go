// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

// MakeReaderChan returns a new open channel
// (simply a 'chan *bufio.Reader' that is).
//
// Note: No 'Reader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReaderPipelineStartsHere := MakeReaderChan()
//	// ... lot's of code to design and build Your favourite "myReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReaderPipelineStartsHere <- drop
//	}
//	close(myReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReaderBuffer) the channel is unbuffered.
//
func MakeReaderChan() chan *bufio.Reader {
	return make(chan *bufio.Reader)
}

// ChanReader returns a channel to receive all inputs before close.
func ChanReader(inp ...*bufio.Reader) chan *bufio.Reader {
	out := make(chan *bufio.Reader)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanReaderSlice returns a channel to receive all inputs before close.
func ChanReaderSlice(inp ...[]*bufio.Reader) chan *bufio.Reader {
	out := make(chan *bufio.Reader)
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

// JoinReader
func JoinReader(out chan<- *bufio.Reader, inp ...*bufio.Reader) chan struct{} {
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

// JoinReaderSlice
func JoinReaderSlice(out chan<- *bufio.Reader, inp ...[]*bufio.Reader) chan struct{} {
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

// JoinReaderChan
func JoinReaderChan(out chan<- *bufio.Reader, inp <-chan *bufio.Reader) chan struct{} {
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

// DoneReader returns a channel to receive one signal before close after inp has been drained.
func DoneReader(inp <-chan *bufio.Reader) chan struct{} {
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

// DoneReaderSlice returns a channel which will receive a slice
// of all the Readers received on inp channel before close.
// Unlike DoneReader, a full slice is sent once, not just an event.
func DoneReaderSlice(inp <-chan *bufio.Reader) chan []*bufio.Reader {
	done := make(chan []*bufio.Reader)
	go func() {
		defer close(done)
		ReaderS := []*bufio.Reader{}
		for i := range inp {
			ReaderS = append(ReaderS, i)
		}
		done <- ReaderS
	}()
	return done
}

// DoneReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReaderFunc(inp <-chan *bufio.Reader, act func(a *bufio.Reader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *bufio.Reader) { return }
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

// PipeReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReaderBuffer(inp <-chan *bufio.Reader, cap int) chan *bufio.Reader {
	out := make(chan *bufio.Reader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReaderFunc(inp <-chan *bufio.Reader, act func(a *bufio.Reader) *bufio.Reader) chan *bufio.Reader {
	out := make(chan *bufio.Reader)
	if act == nil {
		act = func(a *bufio.Reader) *bufio.Reader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReaderFork(inp <-chan *bufio.Reader) (chan *bufio.Reader, chan *bufio.Reader) {
	out1 := make(chan *bufio.Reader)
	out2 := make(chan *bufio.Reader)
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