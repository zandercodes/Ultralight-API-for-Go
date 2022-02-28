/*
bitmap.go

Created by JulianZander on 28.02.2022 at 16:18 with GoLand
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

// UlCreateEmptyBitmap Create empty bitmap.
func UlCreateEmptyBitmap() ULBitmap {
	return (ULBitmap)(C.ulCreateEmptyBitmap())
}

// UlCreateBitmap Create bitmap with certain dimensions and pixel format.
func UlCreateBitmap(width, height uint32, format ULBitmapFormat) ULBitmap {
	return (ULBitmap)(C.ulCreateBitmap((C.uint)(width), (C.uint)(height), format))
}

// UlCreateBitmapFromPixels Create bitmap from existing pixel buffer. @see Bitmap for help using
// this function.
func UlCreateBitmapFromPixels(width, height uint32, format ULBitmapFormat, rowBytes uint32, pixels unsafe.Pointer, size uint32, shouldCopy bool) ULBitmap {
	return (ULBitmap)(C.ulCreateBitmapFromPixels((C.uint)(width), (C.uint)(height), format, (C.uint)(rowBytes), pixels, (C.size_t)(size), (C._Bool)(shouldCopy)))
}

// UlCreateBitmapFromCopy Create bitmap from copy.
func UlCreateBitmapFromCopy(existingBitmap ULBitmap) ULBitmap {
	return (ULBitmap)(C.ulCreateBitmapFromCopy(existingBitmap))
}

// UlDestroyBitmap Destroy a bitmap (you should only destroy Bitmaps you have explicitly
// created via one of the creation functions above.
func UlDestroyBitmap(bitmap ULBitmap) {
	C.ulDestroyBitmap(bitmap)
}

// UlBitmapGetWidth Get the width in pixels.
func UlBitmapGetWidth(bitmap ULBitmap) uint32 {
	return uint32(C.ulBitmapGetWidth(bitmap))
}

// UlBitmapGetHeight Get the height in pixels.
func UlBitmapGetHeight(bitmap ULBitmap) uint32 {
	return uint32(C.ulBitmapGetHeight(bitmap))
}

// UlBitmapGetFormat Get the pixel format.
func UlBitmapGetFormat(bitmap ULBitmap) ULBitmapFormat {
	return (ULBitmapFormat)(C.ulBitmapGetFormat(bitmap))
}

// UlBitmapGetBpp Get the bytes per pixel.
func UlBitmapGetBpp(bitmap ULBitmap) uint32 {
	return uint32(C.ulBitmapGetBpp(bitmap))
}

// UlBitmapGetRowBytes Get the number of bytes per row.
func UlBitmapGetRowBytes(bitmap ULBitmap) uint32 {
	return uint32(C.ulBitmapGetRowBytes(bitmap))
}

// UlBitmapGetSize Get the size in bytes of the underlying pixel buffer.
func UlBitmapGetSize(bitmap ULBitmap) uint32 {
	return uint32(C.ulBitmapGetSize(bitmap))
}

// UlBitmapOwnsPixels Whether or not this bitmap owns its own pixel buffer.
func UlBitmapOwnsPixels(bitmap ULBitmap) bool {
	return bool(C.ulBitmapOwnsPixels(bitmap))
}

// UlBitmapLockPixels Lock pixels for reading/writing, returns pointer to pixel buffer.
func UlBitmapLockPixels(bitmap ULBitmap) {
	C.ulBitmapLockPixels(bitmap)
}

// UlBitmapUnlockPixels Unlock pixels after locking.
func UlBitmapUnlockPixels(bitmap ULBitmap) {
	C.ulBitmapUnlockPixels(bitmap)
}

// UlBitmapRawPixels Get raw pixel buffer-- you should only call this if Bitmap is already
// locked.
func UlBitmapRawPixels(bitmap ULBitmap) {
	C.ulBitmapRawPixels(bitmap)
}

// UlBitmapIsEmpty Whether or not this bitmap is empty.
func UlBitmapIsEmpty(bitmap ULBitmap) bool {
	return bool(C.ulBitmapIsEmpty(bitmap))
}

// UlBitmapErase Reset bitmap pixels to 0.
func UlBitmapErase(bitmap ULBitmap) {
	C.ulBitmapErase(bitmap)
}

// UlBitmapWritePNG Write bitmap to a PNG on disk.
func UlBitmapWritePNG(bitmap ULBitmap, path string) bool {
	return bool(C.ulBitmapWritePNG(bitmap, C.CString(path)))
}

// UlBitmapSwapRedBlueChannels This converts a BGRA bitmap to RGBA bitmap and vice-versa by swapping
// the red and blue channels.
func UlBitmapSwapRedBlueChannels(bitmap ULBitmap) {
	C.ulBitmapSwapRedBlueChannels(bitmap)
}
