/*
rect.go

Created by JulianZander on 28.02.2022 at 16:20 with GoLand
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

// UlRectIsEmpty Whether or not a ULRect is empty (all members equal to 0)
func UlRectIsEmpty(rect ULRect) bool {
	return bool(C.ulRectIsEmpty(rect))
}

// UlRectMakeEmpty Create an empty ULRect (all members equal to 0)
func UlRectMakeEmpty() ULRect {
	return (ULRect)(C.ulRectMakeEmpty())
}
