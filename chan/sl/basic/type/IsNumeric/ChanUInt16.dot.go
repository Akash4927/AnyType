// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type UInt16Chan interface { // bidirectional channel
	UInt16ROnlyChan // aka "<-chan" - receive only
	UInt16SOnlyChan // aka "chan<-" - send only
}

type UInt16ROnlyChan interface { // receive-only channel
	RequestUInt16() (dat uint16)        // the receive function - aka "some-new-UInt16-var := <-MyKind"
	TryUInt16() (dat uint16, open bool) // the multi-valued comma-ok receive function - aka "some-new-UInt16-var, ok := <-MyKind"
}

type UInt16SOnlyChan interface { // send-only channel
	ProvideUInt16(dat uint16) // the send function - aka "MyKind <- some UInt16"
}
