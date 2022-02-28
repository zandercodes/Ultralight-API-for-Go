/*
version.go

Created by JulianZander on 28.02.2022 at 16:14 with GoLand
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

/******************************************************************************
 * Version
 *****************************************************************************/

// UlVersionString Get the version string of the library in MAJOR.MINOR.PATCH format.
func UlVersionString() string {
	return C.GoString(C.ulVersionString())
}

// UlVersionMajor Get the numeric major version of the library.
func UlVersionMajor() uint32 {
	return uint32(C.ulVersionMajor())
}

// UlVersionMinor Get the numeric minor version of the library.
func UlVersionMinor() uint32 {
	return uint32(C.ulVersionMinor())
}

// UlVersionPatch Get the numeric patch version of the library.
func UlVersionPatch() uint32 {
	return uint32(C.ulVersionPatch())
}
