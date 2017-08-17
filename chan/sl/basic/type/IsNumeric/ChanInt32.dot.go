// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Int32Chan interface { // bidirectional channel
	Int32ROnlyChan // aka "<-chan" - receive only
	Int32SOnlyChan // aka "chan<-" - send only
}

type Int32ROnlyChan interface { // receive-only channel
	RequestInt32() (dat int32)        // the receive function - aka "some-new-Int32-var := <-MyKind"
	TryInt32() (dat int32, open bool) // the multi-valued comma-ok receive function - aka "some-new-Int32-var, ok := <-MyKind"
}

type Int32SOnlyChan interface { // send-only channel
	ProvideInt32(dat int32) // the send function - aka "MyKind <- some Int32"
}
