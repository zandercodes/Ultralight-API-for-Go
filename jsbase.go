/*
jsbase.go

Created by JulianZander on 02.03.2022 at 13:04 with GoLand
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

// JSEvaluateScript
// @function JSEvaluateScript
// @abstract Evaluates a string of JavaScript.
// @param ctx The execution context to use.
// @param script A JSString containing the script to evaluate.
// @param thisObject The object to use as "this," or NULL to use the global object as "this."
// @param sourceURL A JSString containing a URL for the script's source file. This is used by debuggers and when reporting exceptions. Pass NULL if you do not care to include source file information.
// @param startingLineNumber An integer value specifying the script's starting line number in the file located at sourceURL. This is only used when reporting exceptions. The value is one-based, so the first line is line 1 and invalid values are clamped to 1.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result The JSValue that results from evaluating script, or NULL if an exception is thrown.
func JSEvaluateScript(ctx JSContextRef, script JSStringRef, thisObject JSObjectRef, sourceURL JSStringRef, startingLineNumber int32, exception []JSValueRef) JSValueRef {
	return (JSValueRef)(C.JSEvaluateScript(ctx, script, thisObject, sourceURL, (C.int)(startingLineNumber), exception))
}

// JSCheckScriptSyntax
// @function JSCheckScriptSyntax
// @abstract Checks for syntax errors in a string of JavaScript.
// @param ctx The execution context to use.
// @param script A JSString containing the script to check for syntax errors.
// @param sourceURL A JSString containing a URL for the script's source file. This is only used when reporting exceptions. Pass NULL if you do not care to include source file information in exceptions.
// @param startingLineNumber An integer value specifying the script's starting line number in the file located at sourceURL. This is only used when reporting exceptions. The value is one-based, so the first line is line 1 and invalid values are clamped to 1.
// @param exception A pointer to a JSValueRef in which to store a syntax error exception, if any. Pass NULL if you do not care to store a syntax error exception.
// @result true if the script is syntactically correct, otherwise false.
func JSCheckScriptSyntax(ctx JSContextRef, script JSStringRef, sourceURL JSStringRef, startingLineNumber int32, exception []JSValueRef) bool {
	return bool(C.JSCheckScriptSyntax(ctx, script, sourceURL, (C.int)(startingLineNumber), exception))
}

// JSGarbageCollect
// @function JSGarbageCollect
// @abstract Performs a JavaScript garbage collection.
// @param ctx The execution context to use.
// @discussion JavaScript values that are on the machine stack, in a register,
//  protected by JSValueProtect, set as the global object of an execution context,
//  or reachable from any such value will not be collected.
//
//  During JavaScript execution, you are not required to call this function; the
//  JavaScript engine will garbage collect as needed. JavaScript values created
//  within a context group are automatically destroyed when the last reference
//  to the context group is released.
func JSGarbageCollect(ctx JSContextRef) {
	C.JSGarbageCollect(ctx)
}
