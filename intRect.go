/*
intRect.go

Created by JulianZander on 28.02.2022 at 16:21 with GoLand
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

// UlIntRectIsEmpty Whether or not a ULIntRect is empty (all members equal to 0)
func UlIntRectIsEmpty(rect ULIntRect) bool {
	return bool(C.ulIntRectIsEmpty())
}

// UlIntRectMakeEmpty Create an empty ULIntRect (all members equal to 0)
func UlIntRectMakeEmpty() ULIntRect {
	return (ULIntRect)(C.ulIntRectMakeEmpty())
}
