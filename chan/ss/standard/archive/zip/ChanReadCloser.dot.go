// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"archive/zip"
)

// MakeReadCloserChan returns a new open channel
// (simply a 'chan zip.ReadCloser' that is).
//
// Note: No 'ReadCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadCloserPipelineStartsHere := MakeReadCloserChan()
//	// ... lot's of code to design and build Your favourite "myReadCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadCloserPipelineStartsHere <- drop
//	}
//	close(myReadCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadCloserBuffer) the channel is unbuffered.
//
func MakeReadCloserChan() (out chan zip.ReadCloser) {
	return make(chan zip.ReadCloser)
}

func sendReadCloser(out chan<- zip.ReadCloser, inp ...zip.ReadCloser) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanReadCloser returns a channel to receive all inputs before close.
func ChanReadCloser(inp ...zip.ReadCloser) (out <-chan zip.ReadCloser) {
	cha := make(chan zip.ReadCloser)
	go sendReadCloser(cha, inp...)
	return cha
}

func sendReadCloserSlice(out chan<- zip.ReadCloser, inp ...[]zip.ReadCloser) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanReadCloserSlice returns a channel to receive all inputs before close.
func ChanReadCloserSlice(inp ...[]zip.ReadCloser) (out <-chan zip.ReadCloser) {
	cha := make(chan zip.ReadCloser)
	go sendReadCloserSlice(cha, inp...)
	return cha
}

func joinReadCloser(done chan<- struct{}, out chan<- zip.ReadCloser, inp ...zip.ReadCloser) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadCloser
func JoinReadCloser(out chan<- zip.ReadCloser, inp ...zip.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadCloser(cha, out, inp...)
	return cha
}

func joinReadCloserSlice(done chan<- struct{}, out chan<- zip.ReadCloser, inp ...[]zip.ReadCloser) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinReadCloserSlice
func JoinReadCloserSlice(out chan<- zip.ReadCloser, inp ...[]zip.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadCloserSlice(cha, out, inp...)
	return cha
}

func joinReadCloserChan(done chan<- struct{}, out chan<- zip.ReadCloser, inp <-chan zip.ReadCloser) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadCloserChan
func JoinReadCloserChan(out chan<- zip.ReadCloser, inp <-chan zip.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadCloserChan(cha, out, inp)
	return cha
}

func doitReadCloser(done chan<- struct{}, inp <-chan zip.ReadCloser) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReadCloser returns a channel to receive one signal before close after inp has been drained.
func DoneReadCloser(inp <-chan zip.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReadCloser(cha, inp)
	return cha
}

func doitReadCloserSlice(done chan<- ([]zip.ReadCloser), inp <-chan zip.ReadCloser) {
	defer close(done)
	ReadCloserS := []zip.ReadCloser{}
	for i := range inp {
		ReadCloserS = append(ReadCloserS, i)
	}
	done <- ReadCloserS
}

// DoneReadCloserSlice returns a channel which will receive a slice
// of all the ReadClosers received on inp channel before close.
// Unlike DoneReadCloser, a full slice is sent once, not just an event.
func DoneReadCloserSlice(inp <-chan zip.ReadCloser) (done <-chan ([]zip.ReadCloser)) {
	cha := make(chan ([]zip.ReadCloser))
	go doitReadCloserSlice(cha, inp)
	return cha
}

func doitReadCloserFunc(done chan<- struct{}, inp <-chan zip.ReadCloser, act func(a zip.ReadCloser)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReadCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadCloserFunc(inp <-chan zip.ReadCloser, act func(a zip.ReadCloser)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a zip.ReadCloser) { return }
	}
	go doitReadCloserFunc(cha, inp, act)
	return cha
}

func pipeReadCloserBuffer(out chan<- zip.ReadCloser, inp <-chan zip.ReadCloser) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReadCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadCloserBuffer(inp <-chan zip.ReadCloser, cap int) (out <-chan zip.ReadCloser) {
	cha := make(chan zip.ReadCloser, cap)
	go pipeReadCloserBuffer(cha, inp)
	return cha
}

func pipeReadCloserFunc(out chan<- zip.ReadCloser, inp <-chan zip.ReadCloser, act func(a zip.ReadCloser) zip.ReadCloser) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReadCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadCloserFunc(inp <-chan zip.ReadCloser, act func(a zip.ReadCloser) zip.ReadCloser) (out <-chan zip.ReadCloser) {
	cha := make(chan zip.ReadCloser)
	if act == nil {
		act = func(a zip.ReadCloser) zip.ReadCloser { return a }
	}
	go pipeReadCloserFunc(cha, inp, act)
	return cha
}

func pipeReadCloserFork(out1, out2 chan<- zip.ReadCloser, inp <-chan zip.ReadCloser) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReadCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadCloserFork(inp <-chan zip.ReadCloser) (out1, out2 <-chan zip.ReadCloser) {
	cha1 := make(chan zip.ReadCloser)
	cha2 := make(chan zip.ReadCloser)
	go pipeReadCloserFork(cha1, cha2, inp)
	return cha1, cha2
}