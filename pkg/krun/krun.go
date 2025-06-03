package krun

/*
#include <libkrun.h>
#include <stdlib.h>
#cgo pkg-config: libkrun
*/
import "C"

import (
	"unsafe"
)

// krun_add_disk
func AddDisk(ctxId uint32, blockId string, diskPath string, readOnly bool) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_blockId := C.CString(blockId)
	defer C.free(unsafe.Pointer(_blockId))
	_diskPath := C.CString(diskPath)
	defer C.free(unsafe.Pointer(_diskPath))
	_readOnly := C.bool(readOnly)
	_ret := C.krun_add_disk(_ctxId, _blockId, _diskPath, _readOnly)
	ret = int32(_ret)
	return
}

// krun_add_disk2
func AddDisk2(ctxId uint32, blockId string, diskPath string, diskFormat uint32, readOnly bool) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_blockId := C.CString(blockId)
	defer C.free(unsafe.Pointer(_blockId))
	_diskPath := C.CString(diskPath)
	defer C.free(unsafe.Pointer(_diskPath))
	_diskFormat := C.uint32_t(diskFormat)
	_readOnly := C.bool(readOnly)
	_ret := C.krun_add_disk2(_ctxId, _blockId, _diskPath, _diskFormat, _readOnly)
	ret = int32(_ret)
	return
}

// krun_add_virtiofs
func AddVirtiofs(ctxId uint32, cTag string, cPath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_cTag := C.CString(cTag)
	defer C.free(unsafe.Pointer(_cTag))
	_cPath := C.CString(cPath)
	defer C.free(unsafe.Pointer(_cPath))
	_ret := C.krun_add_virtiofs(_ctxId, _cTag, _cPath)
	ret = int32(_ret)
	return
}

// krun_add_virtiofs2
func AddVirtiofs2(ctxId uint32, cTag string, cPath string, shmSize uint64) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_cTag := C.CString(cTag)
	defer C.free(unsafe.Pointer(_cTag))
	_cPath := C.CString(cPath)
	defer C.free(unsafe.Pointer(_cPath))
	_shmSize := C.uint64_t(shmSize)
	_ret := C.krun_add_virtiofs2(_ctxId, _cTag, _cPath, _shmSize)
	ret = int32(_ret)
	return
}

// krun_add_vsock_port
func AddVsockPort(ctxId uint32, port uint32, cFilepath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_port := C.uint32_t(port)
	_cFilepath := C.CString(cFilepath)
	defer C.free(unsafe.Pointer(_cFilepath))
	_ret := C.krun_add_vsock_port(_ctxId, _port, _cFilepath)
	ret = int32(_ret)
	return
}

// krun_add_vsock_port2
func AddVsockPort2(ctxId uint32, port uint32, cFilepath string, listen bool) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_port := C.uint32_t(port)
	_cFilepath := C.CString(cFilepath)
	defer C.free(unsafe.Pointer(_cFilepath))
	_listen := C.bool(listen)
	_ret := C.krun_add_vsock_port2(_ctxId, _port, _cFilepath, _listen)
	ret = int32(_ret)
	return
}

// krun_check_nested_virt
func CheckNestedVirt() (ret int32) {
	_ret := C.krun_check_nested_virt()
	ret = int32(_ret)
	return
}

// krun_create_ctx
func CreateCtx() (ret int32) {
	_ret := C.krun_create_ctx()
	ret = int32(_ret)
	return
}

// krun_free_ctx
func FreeCtx(ctxId uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_ret := C.krun_free_ctx(_ctxId)
	ret = int32(_ret)
	return
}

// krun_get_shutdown_eventfd
func GetShutdownEventfd(ctxId uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_ret := C.krun_get_shutdown_eventfd(_ctxId)
	ret = int32(_ret)
	return
}

// krun_set_console_output
func SetConsoleOutput(ctxId uint32, cFilepath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_cFilepath := C.CString(cFilepath)
	defer C.free(unsafe.Pointer(_cFilepath))
	_ret := C.krun_set_console_output(_ctxId, _cFilepath)
	ret = int32(_ret)
	return
}

// krun_set_data_disk
func SetDataDisk(ctxId uint32, diskPath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_diskPath := C.CString(diskPath)
	defer C.free(unsafe.Pointer(_diskPath))
	_ret := C.krun_set_data_disk(_ctxId, _diskPath)
	ret = int32(_ret)
	return
}

// krun_set_env
func SetEnv(ctxId uint32, envp []string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_envp_ := make([]*C.char, len(envp))
	for i := range envp {
		_envp_[i] = C.CString(envp[i])
		defer C.free(unsafe.Pointer(_envp_[i]))
	}
	_envp := (**C.char)(nil)
	if len(_envp_) > 0 {
		_envp = (**C.char)(unsafe.Pointer(&_envp_[0]))
	}
	_ret := C.krun_set_env(_ctxId, _envp)
	ret = int32(_ret)
	return
}

// krun_set_exec
func SetExec(ctxId uint32, execPath string, argv []string, envp []string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_execPath := C.CString(execPath)
	defer C.free(unsafe.Pointer(_execPath))
	_argv_ := make([]*C.char, len(argv))
	for i := range argv {
		_argv_[i] = C.CString(argv[i])
		defer C.free(unsafe.Pointer(_argv_[i]))
	}
	_argv := (**C.char)(nil)
	if len(_argv_) > 0 {
		_argv = (**C.char)(unsafe.Pointer(&_argv_[0]))
	}
	_envp_ := make([]*C.char, len(envp))
	for i := range envp {
		_envp_[i] = C.CString(envp[i])
		defer C.free(unsafe.Pointer(_envp_[i]))
	}
	_envp := (**C.char)(nil)
	if len(_envp_) > 0 {
		_envp = (**C.char)(unsafe.Pointer(&_envp_[0]))
	}
	_ret := C.krun_set_exec(_ctxId, _execPath, _argv, _envp)
	ret = int32(_ret)
	return
}

// krun_set_gpu_options
func SetGpuOptions(ctxId uint32, virglFlags uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_virglFlags := C.uint32_t(virglFlags)
	_ret := C.krun_set_gpu_options(_ctxId, _virglFlags)
	ret = int32(_ret)
	return
}

// krun_set_gpu_options2
func SetGpuOptions2(ctxId uint32, virglFlags uint32, shmSize uint64) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_virglFlags := C.uint32_t(virglFlags)
	_shmSize := C.uint64_t(shmSize)
	_ret := C.krun_set_gpu_options2(_ctxId, _virglFlags, _shmSize)
	ret = int32(_ret)
	return
}

// krun_set_gvproxy_path
func SetGvproxyPath(ctxId uint32, cPath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_cPath := C.CString(cPath)
	defer C.free(unsafe.Pointer(_cPath))
	_ret := C.krun_set_gvproxy_path(_ctxId, _cPath)
	ret = int32(_ret)
	return
}

// krun_set_kernel
func SetKernel(ctxId uint32, kernelPath string, kernelFormat uint32, initramfs string, cmdline string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_kernelPath := C.CString(kernelPath)
	defer C.free(unsafe.Pointer(_kernelPath))
	_kernelFormat := C.uint32_t(kernelFormat)
	_initramfs := C.CString(initramfs)
	defer C.free(unsafe.Pointer(_initramfs))
	_cmdline := C.CString(cmdline)
	defer C.free(unsafe.Pointer(_cmdline))
	_ret := C.krun_set_kernel(_ctxId, _kernelPath, _kernelFormat, _initramfs, _cmdline)
	ret = int32(_ret)
	return
}

// krun_set_log_level
func SetLogLevel(level uint32) (ret int32) {
	_level := C.uint32_t(level)
	_ret := C.krun_set_log_level(_level)
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

// krun_set_nested_virt
func SetNestedVirt(ctxId uint32, enabled bool) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_enabled := C.bool(enabled)
	_ret := C.krun_set_nested_virt(_ctxId, _enabled)
	ret = int32(_ret)
	return
}

// krun_set_net_mac
func SetNetMac(ctxId uint32, cMac *byte) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_cMac := (*C.uint8_t)(unsafe.Pointer(cMac))
	_ret := C.krun_set_net_mac(_ctxId, _cMac)
	ret = int32(_ret)
	return
}

// krun_set_passt_fd
func SetPasstFd(ctxId uint32, fd int32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_fd := C.int(fd)
	_ret := C.krun_set_passt_fd(_ctxId, _fd)
	ret = int32(_ret)
	return
}

// krun_set_port_map
func SetPortMap(ctxId uint32, portMap []string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_portMap_ := make([]*C.char, len(portMap))
	for i := range portMap {
		_portMap_[i] = C.CString(portMap[i])
		defer C.free(unsafe.Pointer(_portMap_[i]))
	}
	_portMap := (**C.char)(nil)
	if len(_portMap_) > 0 {
		_portMap = (**C.char)(unsafe.Pointer(&_portMap_[0]))
	}
	_ret := C.krun_set_port_map(_ctxId, _portMap)
	ret = int32(_ret)
	return
}

// krun_set_rlimits
func SetRlimits(ctxId uint32, rlimits []string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_rlimits_ := make([]*C.char, len(rlimits))
	for i := range rlimits {
		_rlimits_[i] = C.CString(rlimits[i])
		defer C.free(unsafe.Pointer(_rlimits_[i]))
	}
	_rlimits := (**C.char)(nil)
	if len(_rlimits_) > 0 {
		_rlimits = (**C.char)(unsafe.Pointer(&_rlimits_[0]))
	}
	_ret := C.krun_set_rlimits(_ctxId, _rlimits)
	ret = int32(_ret)
	return
}

// krun_set_root
func SetRoot(ctxId uint32, rootPath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_rootPath := C.CString(rootPath)
	defer C.free(unsafe.Pointer(_rootPath))
	_ret := C.krun_set_root(_ctxId, _rootPath)
	ret = int32(_ret)
	return
}

// krun_set_root_disk
func SetRootDisk(ctxId uint32, diskPath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_diskPath := C.CString(diskPath)
	defer C.free(unsafe.Pointer(_diskPath))
	_ret := C.krun_set_root_disk(_ctxId, _diskPath)
	ret = int32(_ret)
	return
}

// krun_set_smbios_oem_strings
func SetSmbiosOemStrings(ctxId uint32, oemStrings []string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_oemStrings_ := make([]*C.char, len(oemStrings))
	for i := range oemStrings {
		_oemStrings_[i] = C.CString(oemStrings[i])
		defer C.free(unsafe.Pointer(_oemStrings_[i]))
	}
	_oemStrings := (**C.char)(nil)
	if len(_oemStrings_) > 0 {
		_oemStrings = (**C.char)(unsafe.Pointer(&_oemStrings_[0]))
	}
	_ret := C.krun_set_smbios_oem_strings(_ctxId, _oemStrings)
	ret = int32(_ret)
	return
}

// krun_set_snd_device
func SetSndDevice(ctxId uint32, enable bool) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_enable := C.bool(enable)
	_ret := C.krun_set_snd_device(_ctxId, _enable)
	ret = int32(_ret)
	return
}

// krun_set_tee_config_file
func SetTeeConfigFile(ctxId uint32, filepath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_filepath := C.CString(filepath)
	defer C.free(unsafe.Pointer(_filepath))
	_ret := C.krun_set_tee_config_file(_ctxId, _filepath)
	ret = int32(_ret)
	return
}

// krun_set_vm_config
func SetVmConfig(ctxId uint32, numVcpus byte, ramMib uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_numVcpus := C.uint8_t(numVcpus)
	_ramMib := C.uint32_t(ramMib)
	_ret := C.krun_set_vm_config(_ctxId, _numVcpus, _ramMib)
	ret = int32(_ret)
	return
}

// krun_set_workdir
func SetWorkdir(ctxId uint32, workdirPath string) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_workdirPath := C.CString(workdirPath)
	defer C.free(unsafe.Pointer(_workdirPath))
	_ret := C.krun_set_workdir(_ctxId, _workdirPath)
	ret = int32(_ret)
	return
}

// krun_setgid
func Setgid(ctxId uint32, gid uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_gid := C.gid_t(gid)
	_ret := C.krun_setgid(_ctxId, _gid)
	ret = int32(_ret)
	return
}

// krun_setuid
func Setuid(ctxId uint32, uid uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_uid := C.uid_t(uid)
	_ret := C.krun_setuid(_ctxId, _uid)
	ret = int32(_ret)
	return
}

// krun_split_irqchip
func SplitIrqchip(ctxId uint32, enable bool) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_enable := C.bool(enable)
	_ret := C.krun_split_irqchip(_ctxId, _enable)
	ret = int32(_ret)
	return
}

// krun_start_enter
func StartEnter(ctxId uint32) (ret int32) {
	_ctxId := C.uint32_t(ctxId)
	_ret := C.krun_start_enter(_ctxId)
	ret = int32(_ret)
	return
}
