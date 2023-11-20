package sdk

// #include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type EventFunc func(string) string

var Event EventFunc

func Log(message string) {
	ptr, size := StringToPtr(message)
	_log(ptr, size)
	runtime.KeepAlive(message) // keep message alive until ptr is no longer needed.
}

//go:wasmimport env log
func _log(ptr, size uint32)

func PtrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func StringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}

func StringToLeakedPtr(s string) (uint32, uint32) {
	size := C.ulong(len(s))
	ptr := unsafe.Pointer(C.malloc(size))
	copy(unsafe.Slice((*byte)(ptr), size), s)
	return uint32(uintptr(ptr)), uint32(size)
}

//export event
func _event(ptr, size uint32) (ptrSize uint64) {
	data := PtrToString(ptr, size)
	result := Event(data)
	ptr, size = StringToLeakedPtr(result)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
