// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

var hardDiv bool // TODO: set if a hardware divider is available

func checkgoarm() {
	// TODO(minux): FP checks like in os_linux_arm.go.

	// NaCl/ARM only supports ARMv7
	if goarm != 7 {
		print("runtime: NaCl requires ARMv7. Recompile using GOARM=7.\n")
		exit(1)
	}
}

//go:nosplit
func cputicks() int64 {
	// Currently cputicks() is used in blocking profiler and to seed runtime·fastrand().
	// runtime·nanotime() is a poor approximation of CPU ticks that is enough for the profiler.
	// TODO: need more entropy to better seed fastrand.
	return nanotime()
}
