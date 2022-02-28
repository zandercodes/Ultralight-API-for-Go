/*
renderer.go

Created by JulianZander on 28.02.2022 at 16:16 with GoLand
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

// UlCreateRenderer Create the Ultralight Renderer directly.
//
// Unlike ulCreateApp(), this does not use any native windows for drawing
// and allows you to manage your own runloop and painting. This method is
// recommended for those wishing to integrate the library into a game.
//
// This singleton manages the lifetime of all Views and coordinates all
// painting, rendering, network requests, and event dispatch.
//
// You should only call this once per process lifetime.
//
// You shoud set up your platform handlers (eg, ulPlatformSetLogger,
// ulPlatformSetFileSystem, etc.) before calling this.
//
// You will also need to define a font loader before calling this--
// as of this writing (v1.2) the only way to do this in C API is by calling
// ulEnablePlatformFontLoader() (available in <AppCore/CAPI.h>).
//
// @NOTE:  You should not call this if you are using ulCreateApp(), it
//         creates its own renderer and provides default implementations for
//         various platform handlers automatically.
func UlCreateRenderer(cfg ULConfig) ULRenderer {
	return (ULRenderer)(C.ulCreateRenderer(cfg))
}

// UlDestroyRenderer Destroy the renderer.
func UlDestroyRenderer(renderer ULRenderer) {
	C.ulDestroyRenderer(renderer)
}

// UlUpdate Update timers and dispatch internal callbacks (JavaScript and network).
func UlUpdate(renderer ULRenderer) {
	C.ulUpdate(renderer)
}

// UlRender Render all active Views.
func UlRender(renderer ULRenderer) {
	C.ulRender(renderer)
}

// UlPurgeMemory Attempt to release as much memory as possible. Don't call this from any
// callbacks or driver code.
func UlPurgeMemory(renderer ULRenderer) {
	C.ulPurgeMemory(renderer)
}

// UlLogMemoryUsage Print detailed memory usage statistics to the log.
// (@see ulPlatformSetLogger)
func UlLogMemoryUsage(renderer ULRenderer) {
	C.ulLogMemoryUsage(renderer)
}
