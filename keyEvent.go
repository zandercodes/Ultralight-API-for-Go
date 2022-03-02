/*
keyEvent.go

Created by JulianZander on 28.02.2022 at 16:19 with GoLand
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

// UlCreateKeyEvent Create a key event, @see KeyEvent for help with the following parameters.
func UlCreateKeyEvent(typ ULKeyEventType, modifiers uint32, virtualKeyCode, nativeKeyCode int32, text, unmodifiedText ULString, isKeypad, isAutoRepeat, isSystemKey bool) ULKeyEvent {
	return (ULKeyEvent)(C.ulCreateKeyEvent(typ, (C.uint)(modifiers), (C.int)(virtualKeyCode), (C.int)(nativeKeyCode), text, unmodifiedText, (C._Bool)(isKeypad), (C._Bool)(isAutoRepeat), (C._Bool)(isSystemKey)))
}

// TODO: ulCreateKeyEventWindows

// TODO: ulCreateKeyEventMacOS

// UlDestroyKeyEvent Destroy a key event.
func UlDestroyKeyEvent(event ULKeyEvent) {
	C.ulDestroyKeyEvent(event)
}
