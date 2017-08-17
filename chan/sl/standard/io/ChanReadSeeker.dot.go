// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

type ReadSeekerChan interface { // bidirectional channel
	ReadSeekerROnlyChan // aka "<-chan" - receive only
	ReadSeekerSOnlyChan // aka "chan<-" - send only
}

type ReadSeekerROnlyChan interface { // receive-only channel
	RequestReadSeeker() (dat io.ReadSeeker)        // the receive function - aka "some-new-ReadSeeker-var := <-MyKind"
	TryReadSeeker() (dat io.ReadSeeker, open bool) // the multi-valued comma-ok receive function - aka "some-new-ReadSeeker-var, ok := <-MyKind"
}

type ReadSeekerSOnlyChan interface { // send-only channel
	ProvideReadSeeker(dat io.ReadSeeker) // the send function - aka "MyKind <- some ReadSeeker"
}
