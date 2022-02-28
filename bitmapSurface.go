/*
bitmapSurface.go

Created by JulianZander on 28.02.2022 at 16:22 with GoLand
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

// UlBitmapSurfaceGetBitmap Get the underlying Bitmap from the default Surface.
//
// @note  Do not call ulDestroyBitmap() on the returned value, it is owned
//        by the surface.
func UlBitmapSurfaceGetBitmap(surface ULBitmapSurface) ULBitmap {
	return (ULBitmap)(C.ulBitmapSurfaceGetBitmap(surface))
}
