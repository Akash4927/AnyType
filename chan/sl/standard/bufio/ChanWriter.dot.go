// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
)

type WriterChan interface { // bidirectional channel
	WriterROnlyChan // aka "<-chan" - receive only
	WriterSOnlyChan // aka "chan<-" - send only
}

type WriterROnlyChan interface { // receive-only channel
	RequestWriter() (dat *bufio.Writer)        // the receive function - aka "some-new-Writer-var := <-MyKind"
	TryWriter() (dat *bufio.Writer, open bool) // the multi-valued comma-ok receive function - aka "some-new-Writer-var, ok := <-MyKind"
}

type WriterSOnlyChan interface { // send-only channel
	ProvideWriter(dat *bufio.Writer) // the send function - aka "MyKind <- some Writer"
}