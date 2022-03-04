/*
jstypedarray.go

Created by JulianZander on 04.03.2022 at 14:03 with GoLand
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
import "unsafe"

// JSObjectMakeTypedArray
// @function
// @abstract           Creates a JavaScript Typed Array object with the given number of elements.
// @param ctx          The execution context to use.
// @param arrayType    A value identifying the type of array to create. If arrayType is kJSTypedArrayTypeNone or kJSTypedArrayTypeArrayBuffer then NULL will be returned.
// @param length       The number of elements to be in the new Typed Array.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             A JSObjectRef that is a Typed Array with all elements set to zero or NULL if there was an error.
func JSObjectMakeTypedArray(ctx JSContextRef, arrayType JSTypedArrayType, length uint32, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeTypedArray(ctx, arrayType, (C.size_t)(length), exception))
}

// JSObjectMakeTypedArrayWithBytesNoCopy
// @function
// @abstract                 Creates a JavaScript Typed Array object from an existing pointer.
// @param ctx                The execution context to use.
// @param arrayType          A value identifying the type of array to create. If arrayType is kJSTypedArrayTypeNone or kJSTypedArrayTypeArrayBuffer then NULL will be returned.
// @param bytes              A pointer to the byte buffer to be used as the backing store of the Typed Array object.
// @param byteLength         The number of bytes pointed to by the parameter bytes.
// @param bytesDeallocator   The allocator to use to deallocate the external buffer when the JSTypedArrayData object is deallocated.
// @param deallocatorContext A pointer to pass back to the deallocator.
// @param exception          A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result                   A JSObjectRef Typed Array whose backing store is the same as the one pointed to by bytes or NULL if there was an error.
// @discussion               If an exception is thrown during this function the bytesDeallocator will always be called.
func JSObjectMakeTypedArrayWithBytesNoCopy(ctx JSContextRef, arrayType JSTypedArrayType, bytes unsafe.Pointer, byteLength uint32, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeTypedArrayWithBytesNoCopy(ctx, arrayType, bytes, (C.size_t)(byteLength), bytesDeallocator, deallocatorContext, exception))
}

// JSObjectMakeTypedArrayWithArrayBuffer
// @function
// @abstract           Creates a JavaScript Typed Array object from an existing JavaScript Array Buffer object.
// @param ctx          The execution context to use.
// @param arrayType    A value identifying the type of array to create. If arrayType is kJSTypedArrayTypeNone or kJSTypedArrayTypeArrayBuffer then NULL will be returned.
// @param buffer       An Array Buffer object that should be used as the backing store for the created JavaScript Typed Array object.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             A JSObjectRef that is a Typed Array or NULL if there was an error. The backing store of the Typed Array will be buffer.
func JSObjectMakeTypedArrayWithArrayBuffer(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeTypedArrayWithArrayBuffer(ctx, arrayType, buffer, exception))
}

// JSObjectMakeTypedArrayWithArrayBufferAndOffset
// @function
// @abstract           Creates a JavaScript Typed Array object from an existing JavaScript Array Buffer object with the given offset and length.
// @param ctx          The execution context to use.
// @param arrayType    A value identifying the type of array to create. If arrayType is kJSTypedArrayTypeNone or kJSTypedArrayTypeArrayBuffer then NULL will be returned.
// @param buffer       An Array Buffer object that should be used as the backing store for the created JavaScript Typed Array object.
// @param byteOffset   The byte offset for the created Typed Array. byteOffset should aligned with the element size of arrayType.
// @param length       The number of elements to include in the Typed Array.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             A JSObjectRef that is a Typed Array or NULL if there was an error. The backing store of the Typed Array will be buffer.
func JSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, byteOffset uint32, length uint32, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx, arrayType, buffer, (C.size_t)(byteOffset), (C.size_t)(length), exception))
}

// JSObjectGetTypedArrayBytesPtr
// @function
// @abstract           Returns a temporary pointer to the backing store of a JavaScript Typed Array object.
// @param ctx          The execution context to use.
// @param object       The Typed Array object whose backing store pointer to return.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             A pointer to the raw data buffer that serves as object's backing store or NULL if object is not a Typed Array object.
// @discussion         The pointer returned by this function is temporary and is not guaranteed to remain valid across JavaScriptCore API calls.
func JSObjectGetTypedArrayBytesPtr(ctx JSContextRef, object JSObjectRef, exception []JSValueRef) unsafe.Pointer {
	return unsafe.Pointer(C.JSObjectGetTypedArrayBytesPtr(ctx, object, exception))
}

// JSObjectGetTypedArrayLength
// @function
// @abstract           Returns the length of a JavaScript Typed Array object.
// @param ctx          The execution context to use.
// @param object       The Typed Array object whose length to return.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             The length of the Typed Array object or 0 if the object is not a Typed Array object.
func JSObjectGetTypedArrayLength(ctx JSContextRef, object JSObjectRef, exception []JSValueRef) uint32 {
	return uint32(C.JSObjectGetTypedArrayLength(ctx, object, exception))
}

// JSObjectGetTypedArrayByteLength
// @function
// @abstract           Returns the byte length of a JavaScript Typed Array object.
// @param ctx          The execution context to use.
// @param object       The Typed Array object whose byte length to return.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             The byte length of the Typed Array object or 0 if the object is not a Typed Array object.
func JSObjectGetTypedArrayByteLength(ctx JSContextRef, object JSObjectRef, exception []JSValueRef) uint32 {
	return uint32(C.JSObjectGetTypedArrayByteLength(ctx, object, exception))
}

// JSObjectGetTypedArrayByteOffset
// @function
// @abstract           Returns the byte offset of a JavaScript Typed Array object.
// @param ctx          The execution context to use.
// @param object       The Typed Array object whose byte offset to return.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             The byte offset of the Typed Array object or 0 if the object is not a Typed Array object.
func JSObjectGetTypedArrayByteOffset(ctx JSContextRef, object JSObjectRef, exception []JSValueRef) uint32 {
	return uint32(C.JSObjectGetTypedArrayByteOffset(ctx, object, exception))
}

// JSObjectGetTypedArrayBuffer
// @function
// @abstract           Returns the JavaScript Array Buffer object that is used as the backing of a JavaScript Typed Array object.
// @param ctx          The execution context to use.
// @param object       The JSObjectRef whose Typed Array type data pointer to obtain.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             A JSObjectRef with a JSTypedArrayType of kJSTypedArrayTypeArrayBuffer or NULL if object is not a Typed Array.
func JSObjectGetTypedArrayBuffer(ctx JSContextRef, object JSObjectRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectGetTypedArrayBuffer(ctx, object, exception))
}

// JSObjectMakeArrayBufferWithBytesNoCopy
// @function
// @abstract                 Creates a JavaScript Array Buffer object from an existing pointer.
// @param ctx                The execution context to use.
// @param bytes              A pointer to the byte buffer to be used as the backing store of the Typed Array object.
// @param byteLength         The number of bytes pointed to by the parameter bytes.
// @param bytesDeallocator   The allocator to use to deallocate the external buffer when the Typed Array data object is deallocated.
// @param deallocatorContext A pointer to pass back to the deallocator.
// @param exception          A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result                   A JSObjectRef Array Buffer whose backing store is the same as the one pointed to by bytes or NULL if there was an error.
// @discussion               If an exception is thrown during this function the bytesDeallocator will always be called.
func JSObjectMakeArrayBufferWithBytesNoCopy(ctx JSContextRef, bytes unsafe.Pointer, byteLength uint32, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeArrayBufferWithBytesNoCopy(ctx, bytes, (C.size_t)(byteLength), bytesDeallocator, deallocatorContext, exception))
}

// JSObjectGetArrayBufferBytesPtr
// @function
// @abstract         Returns a pointer to the data buffer that serves as the backing store for a JavaScript Typed Array object.
// @param object     The Array Buffer object whose internal backing store pointer to return.
// @param exception  A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result           A pointer to the raw data buffer that serves as object's backing store or NULL if object is not an Array Buffer object.
// @discussion       The pointer returned by this function is temporary and is not guaranteed to remain valid across JavaScriptCore API calls.
func JSObjectGetArrayBufferBytesPtr(ctx JSContextRef, object JSObjectRef, exception []JSValueRef) unsafe.Pointer {
	return unsafe.Pointer(C.JSObjectGetArrayBufferBytesPtr(ctx, object, exception))
}

// JSObjectGetArrayBufferByteLength
// @function
// @abstract         Returns the number of bytes in a JavaScript data object.
// @param ctx        The execution context to use.
// @param object     The JS Arary Buffer object whose length in bytes to return.
// @param exception  A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result           The number of bytes stored in the data object.
func JSObjectGetArrayBufferByteLength(ctx JSContextRef, object JSObjectRef, exception []JSValueRef) uint32 {
	return uint32(C.JSObjectGetArrayBufferByteLength(ctx, object, exception))
}
