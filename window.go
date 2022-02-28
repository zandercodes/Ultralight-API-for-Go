/*
window.go

Created by JulianZander on 28.02.2022 at 16:25 with GoLand
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

// UlCreateWindow Create a new Window.
//
// @param  monitor       The monitor to create the Window on.
//
// @param  width         The width (in device coordinates).
//
// @param  height        The height (in device coordinates).
//
// @param  fullscreen    Whether or not the window is fullscreen.
//
// @param  window_flags  Various window flags.
func UlCreateWindow(mon ULMonitor, width uint32, height uint32, fullscreen bool, windowFlags ULWindowFlags) ULWindow {
	cWidth := (C.uint)(width)
	cHeight := (C.uint)(height)
	cFullscreen := (C._Bool)(fullscreen)
	cWindowFlags := (C.uint)(windowFlags)
	return (ULWindow)(C.ulCreateWindow(mon, cWidth, cHeight, cFullscreen, cWindowFlags))
}

// UlDestroyWindow Destroy a Window.
func UlDestroyWindow(win ULWindow) {
	C.ulDestroyWindow(win)
}

// UlWindowSetCloseCallback Set a callback to be notified when a window closes.
func UlWindowSetCloseCallback(win ULWindow, callback ULCloseCallback, userData unsafe.Pointer) {
	cCallback := *(*C.ULCloseCallback)(unsafe.Pointer(&callback))
	C.ulWindowSetCloseCallback(win, cCallback, userData)
}

// UlWindowSetResizeCallback Set a callback to be notified when a window resizes
// (parameters are passed back in pixels).
func UlWindowSetResizeCallback(win ULWindow, callback ULResizeCallback, userData unsafe.Pointer) {
	cCallback := *(*C.ULResizeCallback)(unsafe.Pointer(&callback))
	C.ulWindowSetResizeCallback(win, cCallback, userData)
}

// UlWindowGetWidth Get window width (in pixels).
func UlWindowGetWidth(win ULWindow) uint32 {
	return uint32(C.ulWindowGetWidth(win))
}

// UlWindowGetHeight Get window height (in pixels).
func UlWindowGetHeight(win ULWindow) uint32 {
	return uint32(C.ulWindowGetHeight(win))
}

// UlWindowIsFullscreen Get whether or not a window is fullscreen.
func UlWindowIsFullscreen(win ULWindow) bool {
	return bool(C.ulWindowIsFullscreen(win))
}

// UlWindowGetScale Get the DPI scale of a window.
func UlWindowGetScale(win ULWindow) float64 {
	return float64(C.ulWindowGetScale(win))
}

// UlWindowSetTitle Set the window title.
func UlWindowSetTitle(win ULWindow, title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.ulWindowSetTitle(win, t)
}

// UlWindowSetCursor Set the cursor for a window.
func UlWindowSetCursor(win ULWindow, cursor ULCursor) {
	cULCursor := (C.ULCursor)(cursor)
	C.ulWindowSetCursor(win, cULCursor)
}

// UlWindowClose Close a window.
func UlWindowClose(win ULWindow) {
	C.ulWindowClose(win)
}

// UlWindowDeviceToPixel Convert device coordinates to pixels using the current DPI scale.
func UlWindowDeviceToPixel(win ULWindow, val int32) int32 {
	cVal := (C.int)(val)
	return int32(C.ulWindowDeviceToPixel(win, cVal))
}

// UlWindowPixelsToDevice Convert pixels to device coordinates using the current DPI scale.
func UlWindowPixelsToDevice(win ULWindow, val int32) int32 {
	cVal := (C.int)(val)
	return int32(C.ulWindowPixelsToDevice(win, cVal))
}

// UlWindowGetNativeHandle Get the underlying native window handle.
//
// @note This is:  - HWND on Windows
//                 - NSWindow* on macOS
//                 - GLFWwindow* on Linux
func UlWindowGetNativeHandle(win ULWindow) {
	C.ulWindowGetNativeHandle(win)
}
