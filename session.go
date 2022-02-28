/*
session.go

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

// UlCreateSession Create a Session to store local data in (such as cookies, local storage,
// application cache, indexed db, etc).
func UlCreateSession(renderer ULRenderer, isPersistent bool, name ULString) {
	C.ulCreateSession(renderer, (C._Bool)(isPersistent), name)
}

// UlDestroySession Destroy a Session.
func UlDestroySession(session ULSession) {
	C.ulDestroySession(session)
}

// UlDefaultSession Get the default session (persistent session named "default").
//
// @note  This session is owned by the Renderer, you shouldn't destroy it.
func UlDefaultSession(renderer ULRenderer) ULSession {
	return (ULSession)(C.ulDefaultSession(renderer))
}

// UlSessionIsPersistent Whether or not is persistent (backed to disk).
func UlSessionIsPersistent(session ULSession) bool {
	return bool(C.ulSessionIsPersistent(session))
}

// UlSessionGetName Unique name identifying the session (used for unique disk path).
func UlSessionGetName(session ULSession) ULString {
	return (ULString)(C.ulSessionGetName(session))
}

// UlSessionGetId Unique numeric Id for the session.
func UlSessionGetId(session ULSession) uint64 {
	return uint64(C.ulSessionGetId(session))
}

// UlSessionGetDiskPath The disk path to write to (used by persistent sessions only).
func UlSessionGetDiskPath(session ULSession) ULString {
	return (ULString)(C.ulSessionGetDiskPath(session))
}
