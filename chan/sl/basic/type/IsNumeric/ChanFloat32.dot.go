// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package IsNumeric

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

type Float32Chan interface { // bidirectional channel
	Float32ROnlyChan // aka "<-chan" - receive only
	Float32SOnlyChan // aka "chan<-" - send only
}

type Float32ROnlyChan interface { // receive-only channel
	RequestFloat32() (dat float32)        // the receive function - aka "some-new-Float32-var := <-MyKind"
	TryFloat32() (dat float32, open bool) // the multi-valued comma-ok receive function - aka "some-new-Float32-var, ok := <-MyKind"
}

type Float32SOnlyChan interface { // send-only channel
	ProvideFloat32(dat float32) // the send function - aka "MyKind <- some Float32"
}