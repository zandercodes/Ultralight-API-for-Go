/*
jsobjectref.go

Created by JulianZander on 02.03.2022 at 13:50 with GoLand
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

// JSClassCreate
// @function
// @abstract Creates a JavaScript class suitable for use with JSObjectMake.
// @param definition A JSClassDefinition that defines the class.
// @result A JSClass with the given definition. Ownership follows the Create Rule.
func JSClassCreate(definition []JSClassDefinition) JSClassRef {
	return (JSClassRef)(C.JSClassCreate(definition))
}

// JSClassRetain
// @function
// @abstract Retains a JavaScript class.
// @param jsClass The JSClass to retain.
// @result A JSClass that is the same as jsClass.
func JSClassRetain(jsClass JSClassRef) JSClassRef {
	return (JSClassRef)(C.JSClassRetain(jsClass))
}

// JSClassRelease
// @function
// @abstract Releases a JavaScript class.
// @param jsClass The JSClass to release.
func JSClassRelease(jsClass JSClassRef) {
	C.JSClassRelease(jsClass)
}

// JSClassGetPrivate
// @function
// @abstract Retrieves the private data from a class reference, only possible with classes created with version 1000 (extended callbacks).
// @param jsClass The class to get the data from
// @result The private data on the class, or NULL, if not set
// @discussion Only classes with version 1000 (extended callbacks) can store private data, for other classes always NULL will always be returned.
func JSClassGetPrivate(jsClass JSClassRef) {
	C.JSClassGetPrivate(jsClass)
}

// JSClassSetPrivate
// @function
// @abstract Sets the private data on a class, only possible with classes created with version 1000 (extended callbacks).
// @param jsClass The class to set the data on
// @param data A void* to set as the private data for the class
// @result true if the data has been set on the class, false if the class has not been created with version 1000 (extended callbacks)
// @discussion Only classes with version 1000 (extended callbacks) can store private data, for other classes the function always fails. The set pointer is not touched by the engine.
func JSClassSetPrivate(jsClass JSClassRef, data unsafe.Pointer) bool {
	return bool(C.JSClassSetPrivate(jsClass, data))
}

// JSObjectMake
// @function
// @abstract Creates a JavaScript object.
// @param ctx The execution context to use.
// @param jsClass The JSClass to assign to the object. Pass NULL to use the default object class.
// @param data A void* to set as the object's private data. Pass NULL to specify no private data.
// @result A JSObject with the given class and private data.
// @discussion The default object class does not allocate storage for private data, so you must provide a non-NULL jsClass to JSObjectMake if you want your object to be able to store private data.
//
// data is set on the created object before the intialize methods in its class chain are called. This enables the initialize methods to retrieve and manipulate data through JSObjectGetPrivate.
func JSObjectMake(ctx JSContextRef, jsClass JSClassRef, data unsafe.Pointer) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMake(ctx, jsClass, data))
}

// JSObjectMakeFunctionWithCallback
// @function
// @abstract Convenience method for creating a JavaScript function with a given callback as its implementation.
// @param ctx The execution context to use.
// @param name A JSString containing the function's name. This will be used when converting the function to string. Pass NULL to create an anonymous function.
// @param callAsFunction The JSObjectCallAsFunctionCallback to invoke when the function is called.
// @result A JSObject that is a function. The object's prototype will be the default function prototype.
func JSObjectMakeFunctionWithCallback(ctx JSContextRef, name JSStringRef, callAsFunction JSObjectCallAsFunctionCallback) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeFunctionWithCallback(ctx, name, callAsFunction))
}

// JSObjectMakeConstructor
// @function
// @abstract Convenience method for creating a JavaScript constructor.
// @param ctx The execution context to use.
// @param jsClass A JSClass that is the class your constructor will assign to the objects its constructs. jsClass will be used to set the constructor's .prototype property, and to evaluate 'instanceof' expressions. Pass NULL to use the default object class.
// @param callAsConstructor A JSObjectCallAsConstructorCallback to invoke when your constructor is used in a 'new' expression. Pass NULL to use the default object constructor.
// @result A JSObject that is a constructor. The object's prototype will be the default object prototype.
// @discussion The default object constructor takes no arguments and constructs an object of class jsClass with no private data.
func JSObjectMakeConstructor(ctx JSContextRef, jsClass JSClassRef, callAsConstructor JSObjectCallAsConstructorCallback) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeConstructor(ctx, jsClass, callAsConstructor))
}

// JSObjectMakeArray
// @function
// @abstract Creates a JavaScript Array object.
// @param ctx The execution context to use.
// @param argumentCount An integer count of the number of arguments in arguments.
// @param arguments A JSValue array of data to populate the Array with. Pass NULL if argumentCount is 0.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result A JSObject that is an Array.
// @discussion The behavior of this function does not exactly match the behavior of the built-in Array constructor. Specifically, if one argument
// is supplied, this function returns an array with one element.
func JSObjectMakeArray(ctx JSContextRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeArray(ctx, (C.size_t)(argumentCount), arguments, exception))
}

// JSObjectMakeDate
// @function
// @abstract Creates a JavaScript Date object, as if by invoking the built-in Date constructor.
// @param ctx The execution context to use.
// @param argumentCount An integer count of the number of arguments in arguments.
// @param arguments A JSValue array of arguments to pass to the Date Constructor. Pass NULL if argumentCount is 0.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result A JSObject that is a Date.
func JSObjectMakeDate(ctx JSContextRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeDate(ctx, (C.size_t)(argumentCount), arguments, exception))
}

// JSObjectMakeError
// @function
// @abstract Creates a JavaScript Error object, as if by invoking the built-in Error constructor.
// @param ctx The execution context to use.
// @param argumentCount An integer count of the number of arguments in arguments.
// @param arguments A JSValue array of arguments to pass to the Error Constructor. Pass NULL if argumentCount is 0.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result A JSObject that is a Error.
func JSObjectMakeError(ctx JSContextRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeError(ctx, (C.size_t)(argumentCount), arguments, exception))
}

// JSObjectMakeRegExp
// @function
// @abstract Creates a JavaScript RegExp object, as if by invoking the built-in RegExp constructor.
// @param ctx The execution context to use.
// @param argumentCount An integer count of the number of arguments in arguments.
// @param arguments A JSValue array of arguments to pass to the RegExp Constructor. Pass NULL if argumentCount is 0.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result A JSObject that is a RegExp.
func JSObjectMakeRegExp(ctx JSContextRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeRegExp(ctx, (C.size_t)(argumentCount), arguments, exception))
}

// JSObjectMakeDeferredPromise
// @function
// @abstract Creates a JavaScript promise object by invoking the provided executor.
// @param ctx The execution context to use.
// @param resolve A pointer to a JSObjectRef in which to store the resolve function for the new promise. Pass NULL if you do not care to store the resolve callback.
// @param reject A pointer to a JSObjectRef in which to store the reject function for the new promise. Pass NULL if you do not care to store the reject callback.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result A JSObject that is a promise or NULL if an exception occurred.
func JSObjectMakeDeferredPromise(ctx JSContextRef, resolve []JSObjectRef, reject []JSObjectRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeDeferredPromise(ctx, resolve, reject, exception))
}

// JSObjectMakeFunction
// @function
// @abstract Creates a function with a given script as its body.
// @param ctx The execution context to use.
// @param name A JSString containing the function's name. This will be used when converting the function to string. Pass NULL to create an anonymous function.
// @param parameterCount An integer count of the number of parameter names in parameterNames.
// @param parameterNames A JSString array containing the names of the function's parameters. Pass NULL if parameterCount is 0.
// @param body A JSString containing the script to use as the function's body.
// @param sourceURL A JSString containing a URL for the script's source file. This is only used when reporting exceptions. Pass NULL if you do not care to include source file information in exceptions.
// @param startingLineNumber An integer value specifying the script's starting line number in the file located at sourceURL. This is only used when reporting exceptions. The value is one-based, so the first line is line 1 and invalid values are clamped to 1.
// @param exception A pointer to a JSValueRef in which to store a syntax error exception, if any. Pass NULL if you do not care to store a syntax error exception.
// @result A JSObject that is a function, or NULL if either body or parameterNames contains a syntax error. The object's prototype will be the default function prototype.
// @discussion Use this method when you want to execute a script repeatedly, to avoid the cost of re-parsing the script before each execution.
func JSObjectMakeFunction(ctx JSContextRef, name JSStringRef, parameterCount uint32, parameterNames []JSStringRef, body JSStringRef, sourceURL JSStringRef, startingLineNumber int32, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectMakeFunction(ctx, name, (C.uint)(parameterCount), parameterNames, body, sourceURL, (C.int)(startingLineNumber), exception))
}

// JSObjectGetPrototype
// @function
// @abstract Gets an object's prototype.
// @param ctx  The execution context to use.
// @param object A JSObject whose prototype you want to get.
// @result A JSValue that is the object's prototype.
func JSObjectGetPrototype(ctx JSContextRef, object JSObjectRef) JSValueRef {
	return (JSValueRef)(C.JSObjectGetPrototype(ctx, object))
}

// JSObjectSetPrototype
// @function
// @abstract Sets an object's prototype.
// @param ctx  The execution context to use.
// @param object The JSObject whose prototype you want to set.
// @param value A JSValue to set as the object's prototype.
func JSObjectSetPrototype(ctx JSContextRef, object JSObjectRef, value JSValueRef) {
	C.JSObjectSetPrototype(ctx, object, value)
}

// JSObjectHasProperty
// @function
// @abstract Tests whether an object has a given property.
// @param object The JSObject to test.
// @param propertyName A JSString containing the property's name.
// @result true if the object has a property whose name matches propertyName, otherwise false.
func JSObjectHasProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef) bool {
	return bool(C.JSObjectHasProperty(ctx, object, propertyName))
}

// JSObjectGetProperty
// @function
// @abstract Gets a property from an object.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to get.
// @param propertyName A JSString containing the property's name.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result The property's value if object has the property, otherwise the undefined value.
func JSObjectGetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception []JSValueRef) JSValueRef {
	return (JSValueRef)(C.JSObjectGetProperty(ctx, object, propertyName, exception))
}

// JSObjectSetProperty
// @function
// @abstract Sets a property on an object.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to set.
// @param propertyName A JSString containing the property's name.
// @param value A JSValueRef to use as the property's value.
// @param attributes A logically ORed set of JSPropertyAttributes to give to the property.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
func JSObjectSetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, value JSValueRef, attributes JSPropertyAttributes, exception []JSValueRef) {
	C.JSObjectSetProperty(ctx, object, propertyName, value, attributes, exception)
}

// JSObjectDeleteProperty
// @function
// @abstract Deletes a property from an object.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to delete.
// @param propertyName A JSString containing the property's name.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result true if the delete operation succeeds, otherwise false (for example, if the property has the kJSPropertyAttributeDontDelete attribute set).
func JSObjectDeleteProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception []JSValueRef) bool {
	return bool(C.JSObjectDeleteProperty(ctx, object, propertyName, exception))
}

// JSObjectHasPropertyForKey
// @function
// @abstract Tests whether an object has a given property using a JSValueRef as the property key.
// @param object The JSObject to test.
// @param propertyKey A JSValueRef containing the property key to use when looking up the property.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result true if the object has a property whose name matches propertyKey, otherwise false.
// @discussion This function is the same as performing "propertyKey in object" from JavaScript.
func JSObjectHasPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception []JSValueRef) bool {
	return bool(C.JSObjectHasPropertyForKey(ctx, object, propertyKey, exception))
}

// JSObjectGetPropertyForKey
// @function
// @abstract Gets a property from an object using a JSValueRef as the property key.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to get.
// @param propertyKey A JSValueRef containing the property key to use when looking up the property.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result The property's value if object has the property key, otherwise the undefined value.
// @discussion This function is the same as performing "object[propertyKey]" from JavaScript.
func JSObjectGetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception []JSValueRef) JSValueRef {
	return (JSValueRef)(C.JSObjectGetPropertyForKey(ctx, object, propertyKey, exception))
}

// JSObjectSetPropertyForKey
// @function
// @abstract Sets a property on an object using a JSValueRef as the property key.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to set.
// @param propertyKey A JSValueRef containing the property key to use when looking up the property.
// @param value A JSValueRef to use as the property's value.
// @param attributes A logically ORed set of JSPropertyAttributes to give to the property.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @discussion This function is the same as performing "object[propertyKey] = value" from JavaScript.
func JSObjectSetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, value JSValueRef, attributes JSPropertyAttributes, exception []JSValueRef) {
	C.JSObjectSetPropertyForKey(ctx, object, propertyKey, value, attributes, exception)
}

// JSObjectDeletePropertyForKey
// @function
// @abstract Deletes a property from an object using a JSValueRef as the property key.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to delete.
// @param propertyKey A JSValueRef containing the property key to use when looking up the property.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result true if the delete operation succeeds, otherwise false (for example, if the property has the kJSPropertyAttributeDontDelete attribute set).
// @discussion This function is the same as performing "delete object[propertyKey]" from JavaScript.
func JSObjectDeletePropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception []JSValueRef) bool {
	return bool(C.JSObjectDeletePropertyForKey(ctx, object, propertyKey, exception))
}

// JSObjectGetPropertyAtIndex
// @function
// @abstract Gets a property from an object by numeric index.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to get.
// @param propertyIndex An integer value that is the property's name.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result The property's value if object has the property, otherwise the undefined value.
// @discussion Calling JSObjectGetPropertyAtIndex is equivalent to calling JSObjectGetProperty with a string containing propertyIndex, but JSObjectGetPropertyAtIndex provides optimized access to numeric properties.
func JSObjectGetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex uint32, exception []JSValueRef) JSValueRef {
	return (JSValueRef)(C.JSObjectGetPropertyAtIndex(ctx, object, propertyIndex, exception))
}

// JSObjectSetPropertyAtIndex
// @function
// @abstract Sets a property on an object by numeric index.
// @param ctx The execution context to use.
// @param object The JSObject whose property you want to set.
// @param propertyIndex The property's name as a number.
// @param value A JSValue to use as the property's value.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @discussion Calling JSObjectSetPropertyAtIndex is equivalent to calling JSObjectSetProperty with a string containing propertyIndex, but JSObjectSetPropertyAtIndex provides optimized access to numeric properties.
func JSObjectSetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex uint32, value JSValueRef, exception []JSValueRef) {
	C.JSObjectSetPropertyAtIndex(ctx, object, (C.uint)(propertyIndex), value, exception)
}

// JSObjectGetPrivate
// @function
// @abstract Gets an object's private data.
// @param object A JSObject whose private data you want to get.
// @result A void* that is the object's private data, if the object has private data, otherwise NULL.
func JSObjectGetPrivate(object JSObjectRef) unsafe.Pointer {
	return C.JSObjectGetPrivate(object)
}

// JSObjectSetPrivate
// @function
// @abstract Sets a pointer to private data on an object.
// @param object The JSObject whose private data you want to set.
// @param data A void* to set as the object's private data.
// @result true if object can store private data, otherwise false.
// @discussion The default object class does not allocate storage for private data. Only objects created with a non-NULL JSClass can store private data.
func JSObjectSetPrivate(object JSObjectRef, data unsafe.Pointer) bool {
	return bool(C.JSObjectSetPrivate(object, data))
}

// JSObjectIsFunction
// @function
// @abstract Tests whether an object can be called as a function.
// @param ctx  The execution context to use.
// @param object The JSObject to test.
// @result true if the object can be called as a function, otherwise false.
func JSObjectIsFunction(ctx JSContextRef, object JSObjectRef) bool {
	return bool(C.JSObjectIsFunction(ctx, object))
}

// JSObjectCallAsFunction
// @function
// @abstract Calls an object as a function.
// @param ctx The execution context to use.
// @param object The JSObject to call as a function.
// @param thisObject The object to use as "this," or NULL to use the global object as "this."
// @param argumentCount An integer count of the number of arguments in arguments.
// @param arguments A JSValue array of arguments to pass to the function. Pass NULL if argumentCount is 0.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result The JSValue that results from calling object as a function, or NULL if an exception is thrown or object is not a function.
func JSObjectCallAsFunction(ctx JSContextRef, object JSObjectRef, thisObject JSObjectRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSValueRef {
	return (JSValueRef)(C.JSObjectCallAsFunction(ctx, object, thisObject, (C.size_t)(argumentCount), arguments, exception))
}

// JSObjectIsConstructor
// @function
// @abstract Tests whether an object can be called as a constructor.
// @param ctx  The execution context to use.
// @param object The JSObject to test.
// @result true if the object can be called as a constructor, otherwise false.
func JSObjectIsConstructor(ctx JSContextRef, object JSObjectRef) bool {
	return bool(C.JSObjectIsConstructor(ctx, object))
}

// JSObjectCallAsConstructor
// @function
// @abstract Calls an object as a constructor.
// @param ctx The execution context to use.
// @param object The JSObject to call as a constructor.
// @param argumentCount An integer count of the number of arguments in arguments.
// @param arguments A JSValue array of arguments to pass to the constructor. Pass NULL if argumentCount is 0.
// @param exception A pointer to a JSValueRef in which to store an exception, if any. Pass NULL if you do not care to store an exception.
// @result The JSObject that results from calling object as a constructor, or NULL if an exception is thrown or object is not a constructor.
func JSObjectCallAsConstructor(ctx JSContextRef, object JSObjectRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSObjectRef {
	return (JSObjectRef)(C.JSObjectCallAsConstructor(ctx, object, (C.size_t)(argumentCount), arguments, exception))
}

// JSObjectCopyPropertyNames
// @function
// @abstract Gets the names of an object's enumerable properties.
// @param ctx The execution context to use.
// @param object The object whose property names you want to get.
// @result A JSPropertyNameArray containing the names object's enumerable properties. Ownership follows the Create Rule.
func JSObjectCopyPropertyNames(ctx JSContextRef, object JSObjectRef) JSPropertyNameArrayRef {
	return (JSPropertyNameArrayRef)(C.JSObjectCopyPropertyNames(ctx, object))
}

// JSPropertyNameArrayRetain
// @function
// @abstract Retains a JavaScript property name array.
// @param array The JSPropertyNameArray to retain.
// @result A JSPropertyNameArray that is the same as array.
func JSPropertyNameArrayRetain(array JSPropertyNameArrayRef) JSPropertyNameArrayRef {
	return (JSPropertyNameArrayRef)(C.JSPropertyNameArrayRetain(array))
}

// JSPropertyNameArrayRelease
// @function
// @abstract Releases a JavaScript property name array.
// @param array The JSPropetyNameArray to release.
func JSPropertyNameArrayRelease(array JSPropertyNameArrayRef) {
	C.JSPropertyNameArrayRelease(array)
}

// JSPropertyNameArrayGetCount
// @function
// @abstract Gets a count of the number of items in a JavaScript property name array.
// @param array The array from which to retrieve the count.
// @result An integer count of the number of names in array.
func JSPropertyNameArrayGetCount(array JSPropertyNameArrayRef) uint32 {
	return uint32(C.JSPropertyNameArrayGetCount(array))
}

// JSPropertyNameArrayGetNameAtIndex
// @function
// @abstract Gets a property name at a given index in a JavaScript property name array.
// @param array The array from which to retrieve the property name.
// @param index The index of the property name to retrieve.
// @result A JSStringRef containing the property name.
func JSPropertyNameArrayGetNameAtIndex(array JSPropertyNameArrayRef, index uint32) JSStringRef {
	return (JSStringRef)(C.JSPropertyNameArrayGetNameAtIndex(array, (C.size_t)(index)))
}

// JSPropertyNameAccumulatorAddName
// @function
// @abstract Adds a property name to a JavaScript property name accumulator.
// @param accumulator The accumulator object to which to add the property name.
// @param propertyName The property name to add.
func JSPropertyNameAccumulatorAddName(accumulator JSPropertyNameAccumulatorRef, propertyName JSStringRef) {
	C.JSPropertyNameAccumulatorAddName(accumulator, propertyName)
}
