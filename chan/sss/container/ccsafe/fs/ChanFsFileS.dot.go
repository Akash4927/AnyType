// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsFileSChan returns a new open channel
// (simply a 'chan fs.FsFileS' that is).
//
// Note: No 'FsFileS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsFileSPipelineStartsHere := MakeFsFileSChan()
//	// ... lot's of code to design and build Your favourite "myFsFileSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsFileSPipelineStartsHere <- drop
//	}
//	close(myFsFileSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsFileSBuffer) the channel is unbuffered.
//
func MakeFsFileSChan() (out chan fs.FsFileS) {
	return make(chan fs.FsFileS)
}

// ChanFsFileS returns a channel to receive all inputs before close.
func ChanFsFileS(inp ...fs.FsFileS) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go func(out chan<- fs.FsFileS, inp ...fs.FsFileS) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanFsFileSSlice returns a channel to receive all inputs before close.
func ChanFsFileSSlice(inp ...[]fs.FsFileS) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go func(out chan<- fs.FsFileS, inp ...[]fs.FsFileS) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanFsFileSFuncNil returns a channel to receive all results of act until nil before close.
func ChanFsFileSFuncNil(act func() fs.FsFileS) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go func(out chan<- fs.FsFileS, act func() fs.FsFileS) {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanFsFileSFuncNok returns a channel to receive all results of act until nok before close.
func ChanFsFileSFuncNok(act func() (fs.FsFileS, bool)) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go func(out chan<- fs.FsFileS, act func() (fs.FsFileS, bool)) {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanFsFileSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFsFileSFuncErr(act func() (fs.FsFileS, error)) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go func(out chan<- fs.FsFileS, act func() (fs.FsFileS, error)) {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// JoinFsFileS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFileS(out chan<- fs.FsFileS, inp ...fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.FsFileS, inp ...fs.FsFileS) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFsFileSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFileSSlice(out chan<- fs.FsFileS, inp ...[]fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.FsFileS, inp ...[]fs.FsFileS) {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFsFileSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFsFileSChan(out chan<- fs.FsFileS, inp <-chan fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- fs.FsFileS, inp <-chan fs.FsFileS) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneFsFileS returns a channel to receive one signal before close after inp has been drained.
func DoneFsFileS(inp <-chan fs.FsFileS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan fs.FsFileS) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneFsFileSSlice returns a channel which will receive a slice
// of all the FsFileSs received on inp channel before close.
// Unlike DoneFsFileS, a full slice is sent once, not just an event.
func DoneFsFileSSlice(inp <-chan fs.FsFileS) (done <-chan []fs.FsFileS) {
	cha := make(chan []fs.FsFileS)
	go func(inp <-chan fs.FsFileS, done chan<- []fs.FsFileS) {
		defer close(done)
		FsFileSS := []fs.FsFileS{}
		for i := range inp {
			FsFileSS = append(FsFileSS, i)
		}
		done <- FsFileSS
	}(inp, cha)
	return cha
}

// DoneFsFileSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFileSFunc(inp <-chan fs.FsFileS, act func(a fs.FsFileS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsFileS) { return }
	}
	go func(done chan<- struct{}, inp <-chan fs.FsFileS, act func(a fs.FsFileS)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeFsFileSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFileSBuffer(inp <-chan fs.FsFileS, cap int) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS, cap)
	go func(out chan<- fs.FsFileS, inp <-chan fs.FsFileS) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFsFileSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFileSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFileSFunc(inp <-chan fs.FsFileS, act func(a fs.FsFileS) fs.FsFileS) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	if act == nil {
		act = func(a fs.FsFileS) fs.FsFileS { return a }
	}
	go func(out chan<- fs.FsFileS, inp <-chan fs.FsFileS, act func(a fs.FsFileS) fs.FsFileS) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFsFileSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFileSFork(inp <-chan fs.FsFileS) (out1, out2 <-chan fs.FsFileS) {
	cha1 := make(chan fs.FsFileS)
	cha2 := make(chan fs.FsFileS)
	go func(out1, out2 chan<- fs.FsFileS, inp <-chan fs.FsFileS) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// FsFileSTube is the signature for a pipe function.
type FsFileSTube func(inp <-chan fs.FsFileS, out <-chan fs.FsFileS)

// FsFileSDaisy returns a channel to receive all inp after having passed thru tube.
func FsFileSDaisy(inp <-chan fs.FsFileS, tube FsFileSTube) (out <-chan fs.FsFileS) {
	cha := make(chan fs.FsFileS)
	go tube(inp, cha)
	return cha
}

// FsFileSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FsFileSDaisyChain(inp <-chan fs.FsFileS, tubes ...FsFileSTube) (out <-chan fs.FsFileS) {
	cha := inp
	for i := range tubes {
		cha = FsFileSDaisy(cha, tubes[i])
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