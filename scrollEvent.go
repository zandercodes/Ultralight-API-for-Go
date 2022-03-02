/*
scrollEvent.go

Created by JulianZander on 28.02.2022 at 16:20 with GoLand
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

// UlCreateScrollEvent Create a scroll event, @see ScrollEvent for help using this function.
func UlCreateScrollEvent(typ ULScrollEvent, deltaX, deltaY int32) ULScrollEvent {
	return (ULScrollEvent)(C.ulCreateScrollEvent(typ, (C.int)(deltaX), (C.int)(deltaY)))
}

// UlDestroyScrollEvent Destroy a scroll event.
func UlDestroyScrollEvent(event ULScrollEvent) {
	C.ulDestroyScrollEvent(event)
}
