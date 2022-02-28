/*
mouseEvent.go

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

// UlCreateMouseEvent Create a mouse event, @see MouseEvent for help using this function.
func UlCreateMouseEvent(typ ULKeyEventType, x, y int32, button ULMouseButton) ULMouseEvent {
	return (ULMouseEvent)(C.ulCreateMouseEvent(typ, (C.int)(x), (C.int)(y), button))
}

// UlDestroyMouseEvent Destroy a mouse event.
func UlDestroyMouseEvent(event ULMouseEvent) {
	C.ulDestroyMouseEvent(event)
}
