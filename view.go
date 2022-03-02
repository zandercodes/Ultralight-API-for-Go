/*
view.go

Created by JulianZander on 28.02.2022 at 16:17 with GoLand
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
#include <stdlib.h>
*/
import "C"
import "unsafe"

// UlCreateView Create a View with certain size (in pixels).
//
// @note  You can pass null to 'session' to use the default session.
func UlCreateView(renderer ULRenderer, width, height uint32, transparent bool, session ULSession, forceCpuRenderer bool) ULView {
	return (ULView)(C.ulCreateView(renderer, (C.uint)(width), (C.uint)(height), (C._Bool)(transparent), session, (C._Bool)(forceCpuRenderer)))
}

// UlDestroyView Destroy a View.
func UlDestroyView(view ULView) {
	C.ulDestroyView(view)
}

// UlViewGetURL Get current URL.
//
// @note Don't destroy the returned string, it is owned by the View.
func UlViewGetURL(view ULView) ULString {
	return (ULString)(C.ulViewGetURL(view))
}

// UlViewGetTitle Get current title.
//
// @note Don't destroy the returned string, it is owned by the View.
func UlViewGetTitle(view ULView) ULString {
	return (ULString)(C.ulViewGetTitle(view))
}

// UlViewGetWidth Get the width, in pixels.
func UlViewGetWidth(view ULView) uint32 {
	return uint32(C.ulViewGetWidth(view))
}

// UlViewGetHeight Get the height, in pixels.
func UlViewGetHeight(view ULView) uint32 {
	return uint32(C.ulViewGetHeight(view))
}

// UlViewIsLoading Check if main frame is loading.
func UlViewIsLoading(view ULView) bool {
	return bool(C.ulViewIsLoading(view))
}

// UlViewGetRenderTarget Get the RenderTarget for the View.
//
// @note  Only valid when the GPU renderer is enabled in Config.
func UlViewGetRenderTarget(view ULView) ULRenderTarget {
	return (ULRenderTarget)(C.ulViewGetRenderTarget(view))
}

// UlViewGetSurface Get the Surface for the View (native pixel buffer container).
//
// @note  Only valid when the GPU renderer is disabled in Config.
//
//        (Will return a nullptr when the GPU renderer is enabled.)
//
//        The default Surface is BitmapSurface but you can provide your
//        own Surface implementation via ulPlatformSetSurfaceDefinition.
//
//        When using the default Surface, you can retrieve the underlying
//        bitmap by casting ULSurface to ULBitmapSurface and calling
//        ulBitmapSurfaceGetBitmap().
func UlViewGetSurface(view ULView) ULSurface {
	return (ULSurface)(C.ulViewGetSurface(view))
}

// UlViewLoadHTML Load a raw string of HTML.
func UlViewLoadHTML(view ULView, html ULString) {
	C.ulViewLoadHTML(view, html)
}

// UlViewLoadURL Load a URL into main frame.
func UlViewLoadURL(view ULView, url ULString) {
	C.ulViewLoadURL(view, url)
}

// UlViewResize Resize view to a certain width and height (in pixels).
func UlViewResize(view ULView, width, height uint32) {
	C.ulViewResize(view, (C.uint)(width), (C.uint)(height))
}

// UlViewLockJSContext Acquire the page's JSContext for use with JavaScriptCore API.
//
// @note  This call locks the context for the current thread. You should
//        call ulViewUnlockJSContext() after using the context so other
//        worker threads can modify JavaScript state.
//
// @note  The lock is recusive, it's okay to call this multiple times as long
//        as you call ulViewUnlockJSContext() the same number of times.
func UlViewLockJSContext(view ULView) JSContextRef {
	return (JSContextRef)(C.ulViewLockJSContext(view))
}

// UlViewUnlockJSContext Unlock the page's JSContext after a previous call to ulViewLockJSContext().
func UlViewUnlockJSContext(view ULView) {
	C.ulViewUnlockJSContext(view)
}

// UlViewEvaluateScript Evaluate a string of JavaScript and return result.
//
// @param  js_string  The string of JavaScript to evaluate.
//
// @param  exception  The address of a ULString to store a description of the
//                    last exception. Pass NULL to ignore this. Don't destroy
//                    the exception string returned, it's owned by the View.
//
// @note
//     Don't destroy the returned string, it's owned by the View. This value
//     is reset with every call-- if you want to retain it you should copy
//     the result to a new string via ulCreateStringFromCopy().
//
// @note An example of using this API:
//       <pre>
//         ULString script = ulCreateString("1 + 1");
//         ULString exception;
//         ULString result = ulViewEvaluateScript(view, script, &exception);
//         /* Use the result ("2") and exception description (if any) here.*\/
//         ulDestroyString(script);
//       </pre>
func UlViewEvaluateScript(view ULView, js ULString, exception []ULString) ULString {
	return (ULString)(C.ulViewEvaluateScript(view, js, exception))
}

// UlViewCanGoBack Check if can navigate backwards in history.
func UlViewCanGoBack(view ULView) bool {
	return bool(C.ulViewCanGoBack(view))
}

// UlViewCanGoForward Check if can navigate forwards in history.
func UlViewCanGoForward(view ULView) bool {
	return bool(C.ulViewCanGoForward(view))
}

// UlViewGoBack Navigate backwards in history.
func UlViewGoBack(view ULView) {
	C.ulViewGoBack(view)
}

// UlViewGoForward Navigate forwards in history.
func UlViewGoForward(view ULView) {
	C.ulViewGoForward(view)
}

// UlViewGoToHistoryOffset Navigate to arbitrary offset in history.
func UlViewGoToHistoryOffset(view ULView, offset int32) {
	C.ulViewGoToHistoryOffset(view, (C.int)(offset))
}

// UlViewReload Reload current page.
func UlViewReload(view ULView) {
	C.ulViewReload(view)
}

// UlViewStop Stop all page loads.
func UlViewStop(view ULView) {
	C.ulViewStop(view)
}

// UlViewFocus Give focus to the View.
//
// You should call this to give visual indication that the View has input
// focus (changes active text selection colors, for example).
func UlViewFocus(view ULView) {
	C.ulViewFocus(view)
}

// UlViewUnfocus Remove focus from the View and unfocus any focused input elements.
//
// You should call this to give visual indication that the View has lost
// input focus.
func UlViewUnfocus(view ULView) {
	C.ulViewUnfocus(view)
}

// UlViewHasFocus Whether or not the View has focus.
func UlViewHasFocus(view ULView) bool {
	return bool(C.ulViewHasFocus(view))
}

// UlViewHasInputFocus Whether or not the View has an input element with visible keyboard focus
// (indicated by a blinking caret).
//
// You can use this to decide whether or not the View should consume
// keyboard input events (useful in games with mixed UI and key handling).
func UlViewHasInputFocus(view ULView) bool {
	return bool(C.ulViewHasInputFocus(view))
}

// UlViewFireKeyEvent Fire a keyboard event.
func UlViewFireKeyEvent(view ULView, keyEvent ULKeyEvent) {
	C.ulViewFireKeyEvent(view, keyEvent)
}

// UlViewFireMouseEvent Fire a mouse event.
func UlViewFireMouseEvent(view ULView, mouseEvent ULMouseEvent) {
	C.ulViewFireMouseEvent(view, mouseEvent)
}

// UlViewFireScrollEvent Fire a scroll event.
func UlViewFireScrollEvent(view ULView, scrollEvent ULScrollEvent) {
	C.ulViewFireScrollEvent(view, scrollEvent)
}

// UlViewSetChangeTitleCallback Set callback for when the page title changes.
func UlViewSetChangeTitleCallback(view ULView, callback ULChangeTitleCallback, userData unsafe.Pointer) {
	C.ulViewSetChangeTitleCallback(view, callback, userData)
}

// UlViewSetChangeURLCallback Set callback for when the page URL changes.
func UlViewSetChangeURLCallback(view ULView, callback ULChangeURLCallback, userData unsafe.Pointer) {
	C.ulViewSetChangeURLCallback(view, callback, userData)
}

// UlViewSetChangeTooltipCallback Set callback for when the tooltip changes (usually result of a mouse hover).
func UlViewSetChangeTooltipCallback(view ULView, callback ULChangeTooltipCallback, userData unsafe.Pointer) {
	C.ulViewSetChangeTooltipCallback(view, callback, userData)
}

// UlViewSetChangeCursorCallback Set callback for when the mouse cursor changes.
func UlViewSetChangeCursorCallback(view ULView, callback ULChangeCursorCallback, userData unsafe.Pointer) {
	C.ulViewSetChangeCursorCallback(view, callback, userData)
}

// UlViewSetAddConsoleMessageCallback Set callback for when a message is added to the console (useful for
// JavaScript / network errors and debugging).
func UlViewSetAddConsoleMessageCallback(view ULView, callback ULAddConsoleMessageCallback, userData unsafe.Pointer) {
	C.ulViewSetAddConsoleMessageCallback(view, callback, userData)
}

// UlViewSetCreateChildViewCallback Set callback for when the page wants to create a new View.
//
// This is usually the result of a user clicking a link with target="_blank"
// or by JavaScript calling window.open(url).
//
// To allow creation of these new Views, you should create a new View in
// this callback, resize it to your container,
// and return it. You are responsible for displaying the returned View.
//
// You should return NULL if you want to block the action.
func UlViewSetCreateChildViewCallback(view ULView, callback ULCreateChildViewCallback, userData unsafe.Pointer) {
	C.ulViewSetCreateChildViewCallback(view, callback, userData)
}

// UlViewSetBeginLoadingCallback Set callback for when the page begins loading a new URL into a frame.
func UlViewSetBeginLoadingCallback(view ULView, callback ULBeginLoadingCallback, userData unsafe.Pointer) {
	C.ulViewSetBeginLoadingCallback(view, callback, userData)
}

// UlViewSetFinishLoadingCallback Set callback for when the page finishes loading a URL into a frame.
func UlViewSetFinishLoadingCallback(view ULView, callback ULFinishLoadingCallback, userData unsafe.Pointer) {
	C.ulViewSetFinishLoadingCallback(view, callback, userData)
}

// UlViewSetFailLoadingCallback Set callback for when an error occurs while loading a URL into a frame.
func UlViewSetFailLoadingCallback(view ULView, callback ULFailLoadingCallback, userData unsafe.Pointer) {
	C.ulViewSetFailLoadingCallback(view, callback, userData)
}

// UlViewSetWindowObjectReadyCallback Set callback for when the JavaScript window object is reset for a new
// page load.
//
// This is called before any scripts are executed on the page and is the
// earliest time to setup any initial JavaScript state or bindings.
//
// The document is not guaranteed to be loaded/parsed at this point. If
// you need to make any JavaScript calls that are dependent on DOM elements
// or scripts on the page, use DOMReady instead.
//
// The window object is lazily initialized (this will not be called on pages
// with no scripts).
func UlViewSetWindowObjectReadyCallback(view ULView, callback ULWindowObjectReadyCallback, userData unsafe.Pointer) {
	C.ulViewSetWindowObjectReadyCallback(view, callback, userData)
}

// UlViewSetDOMReadyCallback Set callback for when all JavaScript has been parsed and the document is
// ready.
//
// This is the best time to make any JavaScript calls that are dependent on
// DOM elements or scripts on the page.
func UlViewSetDOMReadyCallback(view ULView, callback ULDOMReadyCallback, userData unsafe.Pointer) {
	C.ulViewSetDOMReadyCallback(view, callback, userData)
}

// UlViewSetUpdateHistoryCallback Set callback for when the history (back/forward state) is modified.
func UlViewSetUpdateHistoryCallback(view ULView, callback ULUpdateHistoryCallback, userData unsafe.Pointer) {
	C.ulViewSetUpdateHistoryCallback(view, callback, userData)
}

// UlViewSetNeedsPaint Set whether or not a view should be repainted during the next call to
// ulRender.
//
// @note  This flag is automatically set whenever the page content changes
//        but you can set it directly in case you need to force a repaint.
func UlViewSetNeedsPaint(view ULView, needsPaint bool) {
	C.ulViewSetNeedsPaint(view, needsPaint)
}

// UlViewGetNeedsPaint Whether or not a view should be painted during the next call to ulRender.
func UlViewGetNeedsPaint(view ULView) bool {
	return bool(C.ulViewGetNeedsPaint(view))
}

// UlViewCreateInspectorView Create an inspector for this View, this is useful for debugging and
// inspecting pages locally. This will only succeed if you have the
// inspector assets in your filesystem-- the inspector will look for
// file:///inspector/Main.html when it loads.
//
// @note  The initial dimensions of the returned View are 10x10, you should
//        call ulViewResize on the returned View to resize it to your desired
//        dimensions.
//
// @note  You will need to call ulDestroyView on the returned instance
//        when you're done using it.
func UlViewCreateInspectorView(view ULView) ULView {
	return (ULView)(C.ulViewCreateInspectorView(view))
}
