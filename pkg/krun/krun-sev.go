//go:build sev
package krun

/*
#include <libkrun.h>
#include <stdlib.h>
#cgo pkg-config: libkrun-sev
*/
import "C"

import (
	"unsafe"
)

// krun_set_tee_config_file
func SetTeeConfigFile(ctxId uint32, filepath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_filepath := C.CString(filepath)
	defer C.free(unsafe.Pointer(_filepath))
	_ret := C.krun_set_tee_config_file(_ctxId, _filepath)
	ret = int32(_ret)
	return
}
