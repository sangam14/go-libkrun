//go:build !sev
package krun

/*
#include <libkrun.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// krun_set_root
func SetRoot(ctxId uint32, rootPath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_rootPath := C.CString(rootPath)
	defer C.free(unsafe.Pointer(_rootPath))
	_ret := C.krun_set_root(_ctxId, _rootPath)
	ret = int32(_ret)
	return
}

// krun_set_mapped_volumes
func SetMappedVolumes(ctxId uint32, mappedVolumes []string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_mappedVolumes_ := make([]*C.char, len(mappedVolumes))
	for i := range mappedVolumes {
		_mappedVolumes_[i] = C.CString(mappedVolumes[i])
		defer C.free(unsafe.Pointer(_mappedVolumes_[i]))
	}
	_mappedVolumes := (**C.char)(nil)
	if len(_mappedVolumes_) > 0 {
		_mappedVolumes = (**C.char)(unsafe.Pointer(&_mappedVolumes_[0]))
	}
	_ret := C.krun_set_mapped_volumes(_ctxId, _mappedVolumes)
	ret = int32(_ret)
	return
}
