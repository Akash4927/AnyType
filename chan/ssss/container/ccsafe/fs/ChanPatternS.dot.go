// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakePatternSChan returns a new open channel
// (simply a 'chan fs.PatternS' that is).
//
// Note: No 'PatternS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPatternSPipelineStartsHere := MakePatternSChan()
//	// ... lot's of code to design and build Your favourite "myPatternSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPatternSPipelineStartsHere <- drop
//	}
//	close(myPatternSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePatternSBuffer) the channel is unbuffered.
//
func MakePatternSChan() chan fs.PatternS {
	return make(chan fs.PatternS)
}

// ChanPatternS returns a channel to receive all inputs before close.
func ChanPatternS(inp ...fs.PatternS) chan fs.PatternS {
	out := make(chan fs.PatternS)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanPatternSSlice returns a channel to receive all inputs before close.
func ChanPatternSSlice(inp ...[]fs.PatternS) chan fs.PatternS {
	out := make(chan fs.PatternS)
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanPatternSFuncNil returns a channel to receive all results of act until nil before close.
func ChanPatternSFuncNil(act func() fs.PatternS) <-chan fs.PatternS {
	out := make(chan fs.PatternS)
	go func() {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanPatternSFuncNok returns a channel to receive all results of act until nok before close.
func ChanPatternSFuncNok(act func() (fs.PatternS, bool)) <-chan fs.PatternS {
	out := make(chan fs.PatternS)
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanPatternSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPatternSFuncErr(act func() (fs.PatternS, error)) <-chan fs.PatternS {
	out := make(chan fs.PatternS)
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// JoinPatternS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatternS(out chan<- fs.PatternS, inp ...fs.PatternS) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// JoinPatternSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatternSSlice(out chan<- fs.PatternS, inp ...[]fs.PatternS) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinPatternSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatternSChan(out chan<- fs.PatternS, inp <-chan fs.PatternS) chan struct{} {
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

// DonePatternS returns a channel to receive one signal before close after inp has been drained.
func DonePatternS(inp <-chan fs.PatternS) chan struct{} {
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

// DonePatternSSlice returns a channel which will receive a slice
// of all the PatternSs received on inp channel before close.
// Unlike DonePatternS, a full slice is sent once, not just an event.
func DonePatternSSlice(inp <-chan fs.PatternS) chan []fs.PatternS {
	done := make(chan []fs.PatternS)
	go func() {
		defer close(done)
		PatternSS := []fs.PatternS{}
		for i := range inp {
			PatternSS = append(PatternSS, i)
		}
		done <- PatternSS
	}()
	return done
}

// DonePatternSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a fs.PatternS) { return }
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

// PipePatternSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatternSBuffer(inp <-chan fs.PatternS, cap int) chan fs.PatternS {
	out := make(chan fs.PatternS, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipePatternSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatternSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatternSFunc(inp <-chan fs.PatternS, act func(a fs.PatternS) fs.PatternS) chan fs.PatternS {
	out := make(chan fs.PatternS)
	if act == nil {
		act = func(a fs.PatternS) fs.PatternS { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipePatternSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatternSFork(inp <-chan fs.PatternS) (chan fs.PatternS, chan fs.PatternS) {
	out1 := make(chan fs.PatternS)
	out2 := make(chan fs.PatternS)
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

// PatternSTube is the signature for a pipe function.
type PatternSTube func(inp <-chan fs.PatternS, out <-chan fs.PatternS)

// PatternSDaisy returns a channel to receive all inp after having passed thru tube.
func PatternSDaisy(inp <-chan fs.PatternS, tube PatternSTube) (out <-chan fs.PatternS) {
	cha := make(chan fs.PatternS)
	go tube(inp, cha)
	return cha
}

// PatternSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PatternSDaisyChain(inp <-chan fs.PatternS, tubes ...PatternSTube) (out <-chan fs.PatternS) {
	cha := inp
	for i := range tubes {
		cha = PatternSDaisy(cha, tubes[i])
	}
	return cha
}

/*
func sendOneInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwoInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

var fun = func(left chan<- int, right <-chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	right := daisyChain(leftmost, fun, 10000) // the chain - right to left!
	go sendTwoInto(right)
	fmt.Println(<-leftmost)
}
*/