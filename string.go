/*
string.go

Created by JulianZander on 28.02.2022 at 16:18 with GoLand
Copyright © 2022 ZanderCodes
*/
package main

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/libs/windows/x64
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/libs/darwin/x64
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/libs/linux/x64
#cgo LDFLAGS: -lUltralightCore -lWebCore -lUltralight -lAppCore
#include "Ultralight/CAPI.h"
#include "AppCore/CAPI.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// UlCreateString Create string from null-terminated ASCII C-string.
func UlCreateString(str string) ULString {
	s := C.CString(str)
	defer C.free(unsafe.Pointer(s))
	return (ULString)(C.ulCreateString(s))
}

// UlCreateStringUTF8 Create string from UTF-8 buffer.
func UlCreateStringUTF8(str string, len uint32) ULString {
	s := C.CString(str)
	defer C.free(unsafe.Pointer(s))
	return (ULString)(C.ulCreateStringUTF8(s, (C.size_t)(len)))
}

// UlCreateStringUTF16 Create string from UTF-16 buffer.
func UlCreateStringUTF16(str []ULChar16, len uint32) ULString {
	return (ULString)(C.ulCreateStringUTF16(str, (C.size_t)(len)))
}

// UlCreateStringFromCopy Create string from copy of existing string.
func UlCreateStringFromCopy(str ULString) ULString {
	s := str
	return (ULString)(C.ulCreateStringFromCopy(s))
}

// UlDestroyString Destroy string (you should destroy any strings you explicitly Create).
func UlDestroyString(str ULString) {
	C.ulDestroyString(str)
}

// UlStringGetData Get internal UTF-16 buffer data.
func UlStringGetData(str ULString) *ULChar16 {
	return (*ULChar16)(C.ulStringGetData(str))
}

// UlStringGetLength Get length in UTF-16 characters.
func UlStringGetLength(str ULString) uint32 {
	return uint32(C.ulStringGetLength(str))
}

// UlStringIsEmpty Whether this string is empty or not.
func UlStringIsEmpty(str ULString) bool {
	return bool(C.ulStringIsEmpty(str))
}

// UlStringAssignString Replaces the contents of 'str' with the contents of 'new_str'
func UlStringAssignString(str ULString, newStr ULString) {
	C.ulStringAssignString(str, newStr)
}

// UlStringAssignCString Replaces the contents of 'str' with the contents of a C-string.
func UlStringAssignCString(str ULString, cStr string) {
	s := C.CString(cStr)
	defer C.free(unsafe.Pointer(s))
	C.ulStringAssignCString(str, s)
}
