// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type ByteChan interface { // bidirectional channel
	ByteROnlyChan // aka "<-chan" - receive only
	ByteSOnlyChan // aka "chan<-" - send only
}

type ByteROnlyChan interface { // receive-only channel
	RequestByte() (dat byte)        // the receive function - aka "some-new-Byte-var := <-MyKind"
	TryByte() (dat byte, open bool) // the multi-valued comma-ok receive function - aka "some-new-Byte-var, ok := <-MyKind"
}

type ByteSOnlyChan interface { // send-only channel
	ProvideByte(dat byte) // the send function - aka "MyKind <- some Byte"
}
