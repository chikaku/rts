package rts

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

// error: go.info.runtime.allp: relocation target go.info.[]*pkg.p not defined
// var allp []*P

//go:linkname allpptr runtime.allp
var allpptr **P

//go:linkname allpLock runtime.allpLock
var allpLock mutex

//go:linkname lock2 runtime.lock2
func lock2(*mutex)

//go:linkname unlock2 runtime.unlock2
func unlock2(*mutex)

// NumTimersP returns the number of timers in each P that currently exist
func NumTimersP() []uint32 {
	n := runtime.GOMAXPROCS(0)
	count := make([]uint32, 0, n)

	lock2(&allpLock)
	allp := unsafe.Slice(allpptr, n)
	for _, p := range allp {
		count = append(count, atomic.LoadUint32(&p.numTimers))
	}
	unlock2(&allpLock)

	return count
}

// NumTimers returns the number of timers currently exist in the process
func NumTimers() uint32 {
	var total uint32
	for _, v := range NumTimersP() {
		total += v
	}
	return total
}
