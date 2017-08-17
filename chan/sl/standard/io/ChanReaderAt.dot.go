// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReaderAtChan interface { // bidirectional channel
	ReaderAtROnlyChan // aka "<-chan" - receive only
	ReaderAtSOnlyChan // aka "chan<-" - send only
}

type ReaderAtROnlyChan interface { // receive-only channel
	RequestReaderAt() (dat io.ReaderAt)        // the receive function - aka "some-new-ReaderAt-var := <-MyKind"
	TryReaderAt() (dat io.ReaderAt, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReaderAt-var, ok := <-MyKind"
}

type ReaderAtSOnlyChan interface { // send-only channel
	ProvideReaderAt(dat io.ReaderAt) // the send function - aka "MyKind <- some ReaderAt"
}
