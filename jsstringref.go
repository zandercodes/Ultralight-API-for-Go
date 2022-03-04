/*
jsstringref.go

Created by JulianZander on 03.03.2022 at 16:10 with GoLand
Copyright © 2022 ZanderCodes
*/
package UltralightAPIforGo

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/libs/windows/x64
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/libs/darwin/x64
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/libs/linux/x64
#cgo LDFLAGS: -lUltralightCore -lWebCore -lUltralight -lAppCore
#include "Ultralight/CAPI.h"
#include "AppCore/CAPI.h"
#include "JavaScriptCore/JavaScript.h"
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// JSStringCreateWithCharacters
// @function
// @abstract         Creates a JavaScript string from a buffer of Unicode characters.
// @param chars      The buffer of Unicode characters to copy into the new JSString.
// @param numChars   The number of characters to copy from the buffer pointed to by chars.
// @result           A JSString containing chars. Ownership follows the Create Rule.
func JSStringCreateWithCharacters(chars []JSChar, numChars uint32) JSStringRef {
	return (JSStringRef)(C.JSStringCreateWithCharacters(chars, (C.size_t)(numChars)))
}

// JSStringCreateWithUTF8CString
// @function
// @abstract         Creates a JavaScript string from a null-terminated UTF8 string.
// @param string     The null-terminated UTF8 string to copy into the new JSString.
// @result           A JSString containing string. Ownership follows the Create Rule.
func JSStringCreateWithUTF8CString(string string) JSStringRef {
	cStr := C.CString(string)
	defer C.free(unsafe.Pointer(cStr))
	runtime.KeepAlive(string) // i don't know for what
	return (JSStringRef)(C.JSStringCreateWithUTF8CString(cStr))
}

// JSStringRetain
// @function
// @abstract         Retains a JavaScript string.
// @param string     The JSString to retain.
// @result           A JSString that is the same as string.
func JSStringRetain(string JSStringRef) JSStringRef {
	return (JSStringRef)(C.JSStringRetain(string))
}

// JSStringRelease
// @function
// @abstract         Releases a JavaScript string.
// @param string     The JSString to release.
func JSStringRelease(string JSStringRef) {
	C.JSStringRelease(string)
}

// JSStringGetLength
// @function
// @abstract         Returns the number of Unicode characters in a JavaScript string.
// @param string     The JSString whose length (in Unicode characters) you want to know.
// @result           The number of Unicode characters stored in string.
func JSStringGetLength(string JSStringRef) uint32 {
	return uint32(C.JSStringGetLength(string))
}

// JSStringGetCharactersPtr
// @function
// @abstract         Returns a pointer to the Unicode character buffer that
//  serves as the backing store for a JavaScript string.
// @param string     The JSString whose backing store you want to access.
// @result           A pointer to the Unicode character buffer that serves as string's
//  backing store, which will be deallocated when string is deallocated.
func JSStringGetCharactersPtr(string JSStringRef) *JSChar {
	return (*JSChar)(C.JSStringGetCharactersPtr(string))
}

// JSStringGetMaximumUTF8CStringSize
// @function
// @abstract Returns the maximum number of bytes a JavaScript string will
//  take up if converted into a null-terminated UTF8 string.
// @param string The JSString whose maximum converted size (in bytes) you
//  want to know.
// @result The maximum number of bytes that could be required to convert string into a
//  null-terminated UTF8 string. The number of bytes that the conversion actually ends
//  up requiring could be less than this, but never more.
func JSStringGetMaximumUTF8CStringSize(string JSStringRef) uint32 {
	return uint32(C.JSStringGetMaximumUTF8CStringSize(string))
}

// JSStringGetUTF8CString
// @function
// @abstract Converts a JavaScript string into a null-terminated UTF8 string,
//  and copies the result into an external byte buffer.
// @param string The source JSString.
// @param buffer The destination byte buffer into which to copy a null-terminated
//  UTF8 representation of string. On return, buffer contains a UTF8 string
//  representation of string. If bufferSize is too small, buffer will contain only
//  partial results. If buffer is not at least bufferSize bytes in size,
//  behavior is undefined.
// @param bufferSize The size of the external buffer in bytes.
// @result The number of bytes written into buffer (including the null-terminator byte).
func JSStringGetUTF8CString(string JSStringRef, buffer []byte, bufferSize uint32) uint32 {
	p := C.malloc(C.size_t(bufferSize))
	defer C.free(p)

	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], buffer)

	return uint32(C.JSStringGetUTF8CString(string, p, (C.size_t)(bufferSize)))
}

// JSStringIsEqual
// @function
// @abstract     Tests whether two JavaScript strings match.
// @param a      The first JSString to test.
// @param b      The second JSString to test.
// @result       true if the two strings match, otherwise false.
func JSStringIsEqual(a JSStringRef, b JSStringRef) bool {
	return bool(C.JSStringIsEqual(a, b))
}

// JSStringIsEqualToUTF8CString
// @function
// @abstract     Tests whether a JavaScript string matches a null-terminated UTF8 string.
// @param a      The JSString to test.
// @param b      The null-terminated UTF8 string to test.
// @result       true if the two strings match, otherwise false.
func JSStringIsEqualToUTF8CString(a JSStringRef, b string) bool {
	cStr := C.CString(b)
	defer C.free(unsafe.Pointer(cStr))
	runtime.KeepAlive(b) // i don't know for what
	return bool(C.JSStringIsEqualToUTF8CString(a, b))
}
