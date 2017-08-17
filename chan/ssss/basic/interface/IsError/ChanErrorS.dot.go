// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsError

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

// MakeErrorSChan returns a new open channel
// (simply a 'chan []error' that is).
//
// Note: No 'ErrorS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myErrorSPipelineStartsHere := MakeErrorSChan()
//	// ... lot's of code to design and build Your favourite "myErrorSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myErrorSPipelineStartsHere <- drop
//	}
//	close(myErrorSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeErrorSBuffer) the channel is unbuffered.
//
func MakeErrorSChan() chan []error {
	return make(chan []error)
}

// ChanErrorS returns a channel to receive all inputs before close.
func ChanErrorS(inp ...[]error) chan []error {
	out := make(chan []error)
	go func() {
		defer close(out)
		for _, i := range inp {
			out <- i
		}
	}()
	return out
}

// ChanErrorSSlice returns a channel to receive all inputs before close.
func ChanErrorSSlice(inp ...[][]error) chan []error {
	out := make(chan []error)
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

// JoinErrorS
func JoinErrorS(out chan<- []error, inp ...[]error) chan struct{} {
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

// JoinErrorSSlice
func JoinErrorSSlice(out chan<- []error, inp ...[][]error) chan struct{} {
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

// JoinErrorSChan
func JoinErrorSChan(out chan<- []error, inp <-chan []error) chan struct{} {
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

// DoneErrorS returns a channel to receive one signal before close after inp has been drained.
func DoneErrorS(inp <-chan []error) chan struct{} {
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

// DoneErrorSSlice returns a channel which will receive a slice
// of all the ErrorSs received on inp channel before close.
// Unlike DoneErrorS, a full slice is sent once, not just an event.
func DoneErrorSSlice(inp <-chan []error) chan [][]error {
	done := make(chan [][]error)
	go func() {
		defer close(done)
		ErrorSS := [][]error{}
		for i := range inp {
			ErrorSS = append(ErrorSS, i)
		}
		done <- ErrorSS
	}()
	return done
}

// DoneErrorSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneErrorSFunc(inp <-chan []error, act func(a []error)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a []error) { return }
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

// PipeErrorSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeErrorSBuffer(inp <-chan []error, cap int) chan []error {
	out := make(chan []error, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeErrorSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeErrorSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeErrorSFunc(inp <-chan []error, act func(a []error) []error) chan []error {
	out := make(chan []error)
	if act == nil {
		act = func(a []error) []error { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeErrorSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeErrorSFork(inp <-chan []error) (chan []error, chan []error) {
	out1 := make(chan []error)
	out2 := make(chan []error)
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
