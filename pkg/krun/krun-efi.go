//go:build efi
package krun

/*
#include <libkrun.h>
#include <stdlib.h>
#cgo LDFLAGS: -lkrun-efi
*/
import "C"

// krun_get_shutdown_eventfd
func GetShutdownEventfd(ctxId uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_ret := C.krun_get_shutdown_eventfd(_ctxId)
	ret = int32(_ret)
	return
}


