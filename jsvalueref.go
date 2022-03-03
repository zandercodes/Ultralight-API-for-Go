/*
jsvalueref.go

Created by JulianZander on 03.03.2022 at 14:39 with GoLand
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

// JSValueGetType
// @function
// @abstract       Returns a JavaScript value's type.
// @param ctx  The execution context to use.
// @param value    The JSValue whose type you want to obtain.
// @result         A value of type JSType that identifies value's type.
func JSValueGetType(ctx JSContextRef, value JSValueRef) JSType {
	return (JSType)(C.JSValueGetType(ctx, value))
}

// JSValueIsUndefined
// @function
// @abstract       Tests whether a JavaScript value's type is the undefined type.
// @param ctx  The execution context to use.
// @param value    The JSValue to test.
// @result         true if value's type is the undefined type, otherwise false.
func JSValueIsUndefined(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsUndefined(ctx, value))
}

// JSValueIsNull
// @function
// @abstract       Tests whether a JavaScript value's type is the null type.
// @param ctx  The execution context to use.
// @param value    The JSValue to test.
// @result         true if value's type is the null type, otherwise false.
func JSValueIsNull(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsNull(ctx, value))
}

// JSValueIsBoolean
// @function
// @abstract       Tests whether a JavaScript value's type is the boolean type.
// @param ctx  The execution context to use.
// @param value    The JSValue to test.
// @result         true if value's type is the boolean type, otherwise false.
func JSValueIsBoolean(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsBoolean(ctx, value))
}

// JSValueIsNumber
// @function
// @abstract       Tests whether a JavaScript value's type is the number type.
// @param ctx  The execution context to use.
// @param value    The JSValue to test.
// @result         true if value's type is the number type, otherwise false.
func JSValueIsNumber(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsNumber(ctx, value))
}

// JSValueIsString
// @function
// @abstract       Tests whether a JavaScript value's type is the string type.
// @param ctx  The execution context to use.
// @param value    The JSValue to test.
// @result         true if value's type is the string type, otherwise false.
func JSValueIsString(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsString(ctx, value))
}

// JSValueIsSymbol
// @function
// @abstract       Tests whether a JavaScript value's type is the symbol type.
// @param ctx      The execution context to use.
// @param value    The JSValue to test.
// @result         true if value's type is the symbol type, otherwise false.
func JSValueIsSymbol(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsSymbol(ctx, value))
}

// JSValueIsObject
// @function
// @abstract       Tests whether a JavaScript value's type is the object type.
// @param ctx  The execution context to use.
// @param value    The JSValue to test.
// @result         true if value's type is the object type, otherwise false.
func JSValueIsObject(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsObject(ctx, value))
}

// JSValueIsObjectOfClass
// @function
// @abstract Tests whether a JavaScript value is an object with a given class in its class chain.
// @param ctx The execution context to use.
// @param value The JSValue to test.
// @param jsClass The JSClass to test against.
// @result true if value is an object and has jsClass in its class chain, otherwise false.
func JSValueIsObjectOfClass(ctx JSContextRef, value JSValueRef, jsClass JSClassRef) bool {
	return bool(C.JSValueIsObjectOfClass(ctx, value, jsClass))
}

// JSValueIsArray
// @function
// @abstract       Tests whether a JavaScript value is an array.
// @param ctx      The execution context to use.
// @param value    The JSValue to test.
// @result         true if value is an array, otherwise false.
func JSValueIsArray(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsArray(ctx, value))
}

// JSValueIsDate
// @function
// @abstract       Tests whether a JavaScript value is a date.
// @param ctx      The execution context to use.
// @param value    The JSValue to test.
// @result         true if value is a date, otherwise false.
func JSValueIsDate(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueIsDate(ctx, value))
}

// JSValueGetTypedArrayType
// @function
// @abstract           Returns a JavaScript value's Typed Array type.
// @param ctx          The execution context to use.
// @param value        The JSValue whose Typed Array type to return.
// @param exception    A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result             A value of type JSTypedArrayType that identifies value's Typed Array type, or kJSTypedArrayTypeNone if the value is not a Typed Array object.
func JSValueGetTypedArrayType(ctx JSContextRef, value JSValueRef, exception []JSValueRef) JSTypedArrayType {
	return (JSTypedArrayType)(C.JSValueGetTypedArrayType(ctx, value, exception))
}

// JSValueIsEqual
// @function
// @abstract Tests whether two JavaScript values are equal, as compared by the JS == operator.
// @param ctx The execution context to use.
// @param a The first value to test.
// @param b The second value to test.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result true if the two values are equal, false if they are not equal or an exception is thrown.
func JSValueIsEqual(ctx JSContextRef, a JSValueRef, b JSValueRef, exception []JSValueRef) bool {
	return bool(C.JSValueIsEqual(ctx, a, b, exception))
}

// JSValueIsStrictEqual
// @function
// @abstract       Tests whether two JavaScript values are strict equal, as compared by the JS === operator.
// @param ctx  The execution context to use.
// @param a        The first value to test.
// @param b        The second value to test.
// @result         true if the two values are strict equal, otherwise false.
func JSValueIsStrictEqual(ctx JSContextRef, a JSValueRef, b JSValueRef) bool {
	return bool(C.JSValueIsStrictEqual(ctx, a, b))
}

// JSValueIsInstanceOfConstructor
// @function
// @abstract Tests whether a JavaScript value is an object constructed by a given constructor, as compared by the JS instanceof operator.
// @param ctx The execution context to use.
// @param value The JSValue to test.
// @param constructor The constructor to test against.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result true if value is an object constructed by constructor, as compared by the JS instanceof operator, otherwise false.
func JSValueIsInstanceOfConstructor(ctx JSContextRef, value JSValueRef, constructor JSObjectRef, exception []JSValueRef) bool {
	return bool(C.JSValueIsInstanceOfConstructor(ctx, value, constructor, exception))
}

// JSValueMakeUndefined
// @function
// @abstract       Creates a JavaScript value of the undefined type.
// @param ctx  The execution context to use.
// @result         The unique undefined value.
func JSValueMakeUndefined(ctx JSContextRef) JSValueRef {
	return (JSValueRef)(C.JSValueMakeUndefined(ctx))
}

// JSValueMakeNull
// @function
// @abstract       Creates a JavaScript value of the null type.
// @param ctx  The execution context to use.
// @result         The unique null value.
func JSValueMakeNull(ctx JSContextRef) JSValueRef {
	return (JSValueRef)(C.JSValueMakeNull(ctx))
}

// JSValueMakeBoolean
// @function
// @abstract       Creates a JavaScript value of the boolean type.
// @param ctx  The execution context to use.
// @param boolean  The bool to assign to the newly created JSValue.
// @result         A JSValue of the boolean type, representing the value of boolean.
func JSValueMakeBoolean(ctx JSContextRef, boolean bool) JSValueRef {
	return (JSValueRef)(C.JSValueMakeBoolean(ctx, (C._Bool)(boolean)))
}

// JSValueMakeNumber
// @function
// @abstract       Creates a JavaScript value of the number type.
// @param ctx  The execution context to use.
// @param number   The double to assign to the newly created JSValue.
// @result         A JSValue of the number type, representing the value of number.
func JSValueMakeNumber(ctx JSContextRef, number float64) JSValueRef {
	return (JSValueRef)(C.JSValueMakeNumber(ctx, (C.double)(number)))
}

// JSValueMakeString
// @function
// @abstract       Creates a JavaScript value of the string type.
// @param ctx  The execution context to use.
// @param string   The JSString to assign to the newly created JSValue. The
//  newly created JSValue retains string, and releases it upon garbage collection.
// @result         A JSValue of the string type, representing the value of string.
func JSValueMakeString(ctx JSContextRef, string JSStringRef) JSValueRef {
	return (JSValueRef)(C.JSValueMakeString(ctx, string))
}

// JSValueMakeSymbol
// @function
// @abstract            Creates a JavaScript value of the symbol type.
// @param ctx           The execution context to use.
// @param description   A description of the newly created symbol value.
// @result              A unique JSValue of the symbol type, whose description matches the one provided.
func JSValueMakeSymbol(ctx JSContextRef, description JSStringRef) JSValueRef {
	return (JSValueRef)(C.JSValueMakeSymbol(ctx, description))
}

// JSValueMakeFromJSONString
// @function
// @abstract       Creates a JavaScript value from a JSON formatted string.
// @param ctx      The execution context to use.
// @param string   The JSString containing the JSON string to be parsed.
// @result         A JSValue containing the parsed value, or NULL if the input is invalid.
func JSValueMakeFromJSONString(ctx JSContextRef, string JSStringRef) JSValueRef {
	return (JSValueRef)(C.JSValueMakeFromJSONString(ctx, string))
}

// JSValueCreateJSONString
// @function
// @abstract       Creates a JavaScript string containing the JSON serialized representation of a JS value.
// @param ctx      The execution context to use.
// @param value    The value to serialize.
// @param indent   The number of spaces to indent when nesting.  If 0, the resulting JSON will not contains newlines.  The size of the indent is clamped to 10 spaces.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result         A JSString with the result of serialization, or NULL if an exception is thrown.
func JSValueCreateJSONString(ctx JSContextRef, value JSValueRef, indent uint32, exception []JSValueRef) JSStringRef {
	return (JSStringRef)(C.JSValueCreateJSONString(ctx, value, (C.uint)(indent), exception))
}

// JSValueToBoolean
// @function
// @abstract       Converts a JavaScript value to boolean and returns the resulting boolean.
// @param ctx  The execution context to use.
// @param value    The JSValue to convert.
// @result         The boolean result of conversion.
func JSValueToBoolean(ctx JSContextRef, value JSValueRef) bool {
	return bool(C.JSValueToBoolean(ctx, value))
}

// JSValueToNumber
// @function
// @abstract       Converts a JavaScript value to number and returns the resulting number.
// @param ctx  The execution context to use.
// @param value    The JSValue to convert.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result         The numeric result of conversion, or NaN if an exception is thrown.
func JSValueToNumber(ctx JSContextRef, value JSValueRef, exception []JSValueRef) float64 {
	return float64(C.JSValueToNumber(ctx, value, exception))
}

// JSValueToStringCopy
// @function
// @abstract       Converts a JavaScript value to string and copies the result into a JavaScript string.
// @param ctx  The execution context to use.
// @param value    The JSValue to convert.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result         A JSString with the result of conversion, or NULL if an exception is thrown. Ownership follows the Create Rule.
func JSValueToStringCopy(ctx JSContextRef, value JSValueRef, exception []JSValueRef) JSStringRef {
	return (JSStringRef)(C.JSValueToStringCopy(ctx, value, exception))
}

// JSValueToObject
// @function
// @abstract Converts a JavaScript value to object and returns the resulting object.
// @param ctx  The execution context to use.
// @param value    The JSValue to convert.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result         The JSObject result of conversion, or NULL if an exception is thrown.
func JSValueToObject(ctx JSContextRef, value JSValueRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSValueToObject(ctx, value, exception))
}

// JSValueProtect
// @function
// @abstract Protects a JavaScript value from garbage collection.
// @param ctx The execution context to use.
// @param value The JSValue to protect.
// @discussion Use this method when you want to store a JSValue in a global or on the heap, where the garbage collector will not be able to discover your reference to it.
//
// A value may be protected multiple times and must be unprotected an equal number of times before becoming eligible for garbage collection.
func JSValueProtect(ctx JSContextRef, value JSValueRef) {
	C.JSValueProtect(ctx, value)
}

// JSValueUnprotect
// @function
// @abstract       Unprotects a JavaScript value from garbage collection.
// @param ctx      The execution context to use.
// @param value    The JSValue to unprotect.
// @discussion     A value may be protected multiple times and must be unprotected an
//  equal number of times before becoming eligible for garbage collection.
func JSValueUnprotect(ctx JSContextRef, value JSValueRef) {
	C.JSValueUnprotect(ctx, value)
}
