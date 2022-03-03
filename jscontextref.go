/*
jscontextref.go

Created by JulianZander on 02.03.2022 at 13:17 with GoLand
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

// JSContextGroupCreate
// @function
// @abstract Creates a JavaScript context group.
// @discussion A JSContextGroup associates JavaScript contexts with one another.
//  Contexts in the same group may share and exchange JavaScript objects. Sharing and/or exchanging
//  JavaScript objects between contexts in different groups will produce undefined behavior.
//  When objects from the same context group are used in multiple threads, explicit
//  synchronization is required.
//
//  A JSContextGroup may need to run deferred tasks on a run loop, such as garbage collection
//  or resolving WebAssembly compilations. By default, calling JSContextGroupCreate will use
//  the run loop of the thread it was called on. Currently, there is no API to change a
//  JSContextGroup's run loop once it has been created.
// @result The created JSContextGroup.
func JSContextGroupCreate() JSContextGroupRef {
	return (JSContextGroupRef)(C.JSContextGroupCreate())
}

// JSContextGroupRetain
// @function
// @abstract Retains a JavaScript context group.
// @param group The JSContextGroup to retain.
// @result A JSContextGroup that is the same as group.
func JSContextGroupRetain(group JSContextGroupRef) JSContextGroupRef {
	return (JSContextGroupRef)(C.JSContextGroupRetain(group))
}

// JSContextGroupRelease
// @function
// @abstract Releases a JavaScript context group.
// @param group The JSContextGroup to release.
func JSContextGroupRelease(group JSContextGroupRef) {
	C.JSContextGroupRelease(group)
}

// JSGlobalContextCreate
// @function
// @abstract Creates a global JavaScript execution context.
// @discussion JSGlobalContextCreate allocates a global object and populates it with all the
//  built-in JavaScript objects, such as Object, Function, String, and Array.
//
//  In WebKit version 4.0 and later, the context is created in a unique context group.
//  Therefore, scripts may execute in it concurrently with scripts executing in other contexts.
//  However, you may not use values created in the context in other contexts.
// @param globalObjectClass The class to use when creating the global object. Pass
//  NULL to use the default object class.
// @result A JSGlobalContext with a global object of class globalObjectClass.
func JSGlobalContextCreate(globalObjectClass JSClassRef) JSGlobalContextRef {
	return (JSGlobalContextRef)(C.JSGlobalContextCreate(globalObjectClass))
}

// JSGlobalContextCreateInGroup
// @function
// @abstract Creates a global JavaScript execution context in the context group provided.
// @discussion JSGlobalContextCreateInGroup allocates a global object and populates it with
//  all the built-in JavaScript objects, such as Object, Function, String, and Array.
// @param globalObjectClass The class to use when creating the global object. Pass
//  NULL to use the default object class.
// @param group The context group to use. The created global context retains the group.
//  Pass NULL to create a unique group for the context.
// @result A JSGlobalContext with a global object of class globalObjectClass and a context
//  group equal to group.
func JSGlobalContextCreateInGroup(group JSContextGroupRef, globalObjectClass JSClassRef) JSGlobalContextRef {
	return (JSGlobalContextRef)(C.JSGlobalContextCreateInGroup(group, globalObjectClass))
}

// JSGlobalContextRetain
// @function
// @abstract Retains a global JavaScript execution context.
// @param ctx The JSGlobalContext to retain.
// @result A JSGlobalContext that is the same as ctx.
func JSGlobalContextRetain(ctx JSGlobalContextRef) JSGlobalContextRef {
	return (JSGlobalContextRef)(C.JSGlobalContextRetain(ctx))
}

// JSGlobalContextRelease
// @function
// @abstract Releases a global JavaScript execution context.
// @param ctx The JSGlobalContext to release.
func JSGlobalContextRelease(ctx JSGlobalContextRef) {
	C.JSGlobalContextRelease(ctx)
}

// JSContextGetGlobalObject
// @function
// @abstract Gets the global object of a JavaScript execution context.
// @param ctx The JSContext whose global object you want to get.
// @result ctx's global object.
func JSContextGetGlobalObject(ctx JSContextRef) JSObjectRef {
	return (JSObjectRef)(C.JSContextGetGlobalObject(ctx))
}

// JSContextGetGroup
// @function
// @abstract Gets the context group to which a JavaScript execution context belongs.
// @param ctx The JSContext whose group you want to get.
// @result ctx's group.
func JSContextGetGroup(ctx JSContextRef) JSContextGroupRef {
	return (JSContextGroupRef)(C.JSContextGetGroup(ctx))
}

// JSContextGetGlobalContext
// @function
// @abstract Gets the global context of a JavaScript execution context.
// @param ctx The JSContext whose global context you want to get.
// @result ctx's global context.
func JSContextGetGlobalContext(ctx JSContextRef) JSGlobalContextRef {
	return (JSGlobalContextRef)(C.JSContextGetGlobalContext(ctx))
}

// JSGlobalContextCopyName
// @function
// @abstract Gets a copy of the name of a context.
// @param ctx The JSGlobalContext whose name you want to get.
// @result The name for ctx.
// @discussion A JSGlobalContext's name is exposed for remote debugging to make it
// easier to identify the context you would like to attach to.
func JSGlobalContextCopyName(ctx JSGlobalContextRef) JSStringRef {
	return (JSStringRef)(C.JSGlobalContextCopyName(ctx))
}

// JSGlobalContextSetName
// @function
// @abstract Sets the remote debugging name for a context.
// @param ctx The JSGlobalContext that you want to name.
// @param name The remote debugging name to set on ctx.
func JSGlobalContextSetName(ctx JSGlobalContextRef, name JSStringRef) {
	C.JSGlobalContextSetName(ctx, name)
}
