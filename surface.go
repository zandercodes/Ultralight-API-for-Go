/*
surface.go

Created by JulianZander on 28.02.2022 at 16:21 with GoLand
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

// UlSurfaceGetWidth Width (in pixels).
func UlSurfaceGetWidth(surface ULSurface) uint32 {
	return uint32(C.ulSurfaceGetWidth(surface))
}

// UlSurfaceGetHeight Height (in pixels).
func UlSurfaceGetHeight(surface ULSurface) uint32 {
	return uint32(C.ulSurfaceGetHeight(surface))
}

// UlSurfaceGetRowBytes Number of bytes between rows (usually width * 4)
func UlSurfaceGetRowBytes(surface ULSurface) uint32 {
	return uint32(C.ulSurfaceGetRowBytes(surface))
}

// UlSurfaceGetSize Size in bytes.
func UlSurfaceGetSize(surface ULSurface) uint32 {
	return uint32(C.ulSurfaceGetSize(surface))
}

// UlSurfaceLockPixels Lock the pixel buffer and get a pointer to the beginning of the data
// for reading/writing.
//
// Native pixel format is premultiplied BGRA 32-bit (8 bits per channel).
func UlSurfaceLockPixels(surface ULSurface) {
	C.ulSurfaceLockPixels(surface)
}

// UlSurfaceUnlockPixels Unlock the pixel buffer.
func UlSurfaceUnlockPixels(surface ULSurface) {
	C.ulSurfaceUnlockPixels(surface)
}

// UlSurfaceResize Resize the pixel buffer to a certain width and height (both in pixels).
//
// This should never be called while pixels are locked.
func ulSurfaceResize(surface ULSurface, width, height uint32) {
	C.ulSurfaceResize(surface, (C.int)(width), (C.int)(height))
}

// UlSurfaceSetDirtyBounds Set the dirty bounds to a certain value.
//
// This is called after the Renderer paints to an area of the pixel buffer.
// (The new value will be joined with the existing dirty_bounds())
func UlSurfaceSetDirtyBounds(surface ULSurface, bounds ULIntRect) {
	C.ulSurfaceSetDirtyBounds(surface, bounds)
}

// UlSurfaceGetDirtyBounds Get the dirty bounds.
//
// This value can be used to determine which portion of the pixel buffer has
// been updated since the last call to ulSurfaceClearDirtyBounds().
//
// The general algorithm to determine if a Surface needs display is:
// <pre>
//   if (!ulIntRectIsEmpty(ulSurfaceGetDirtyBounds(surface))) {
//       // Surface pixels are dirty and needs display.
//       // Cast Surface to native Surface and use it here (pseudo code)
//       DisplaySurface(surface);
//
//       // Once you're done, clear the dirty bounds:
//       ulSurfaceClearDirtyBounds(surface);
//  }
//  </pre>
func UlSurfaceGetDirtyBounds(surface ULSurface) ULIntRect {
	return (ULIntRect)(C.ulSurfaceGetDirtyBounds(surface))
}

// UlSurfaceClearDirtyBounds Clear the dirty bounds.
//
// You should call this after you're done displaying the Surface.
func UlSurfaceClearDirtyBounds(surface ULSurface) {
	C.ulSurfaceClearDirtyBounds(surface)
}

// UlSurfaceGetUserData Get the underlying user data pointer (this is only valid if you have
// set a custom surface implementation via ulPlatformSetSurfaceDefinition).
//
// This will return nullptr if this surface is the default ULBitmapSurface.
func UlSurfaceGetUserData(surface ULSurface) {
	C.ulSurfaceGetUserData(surface)
}
