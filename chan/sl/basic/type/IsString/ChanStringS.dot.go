// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsString

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type StringSChan interface { // bidirectional channel
	StringSROnlyChan // aka "<-chan" - receive only
	StringSSOnlyChan // aka "chan<-" - send only
}

type StringSROnlyChan interface { // receive-only channel
	RequestStringS() (dat []string)        // the receive function - aka "some-new-StringS-var := <-MyKind"
	TryStringS() (dat []string, open bool) // the multi-valued comma-ok receive function - aka "some-new-StringS-var, ok := <-MyKind"
}

type StringSSOnlyChan interface { // send-only channel
	ProvideStringS(dat []string) // the send function - aka "MyKind <- some StringS"
}
