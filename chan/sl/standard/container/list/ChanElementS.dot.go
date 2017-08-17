// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"container/list"
)

type ElementSChan interface { // bidirectional channel
	ElementSROnlyChan // aka "<-chan" - receive only
	ElementSSOnlyChan // aka "chan<-" - send only
}

type ElementSROnlyChan interface { // receive-only channel
	RequestElementS() (dat []list.Element)        // the receive function - aka "some-new-ElementS-var := <-MyKind"
	TryElementS() (dat []list.Element, open bool) // the multi-valued comma-ok receive function - aka "some-new-ElementS-var, ok := <-MyKind"
}

type ElementSSOnlyChan interface { // send-only channel
	ProvideElementS(dat []list.Element) // the send function - aka "MyKind <- some ElementS"
}
