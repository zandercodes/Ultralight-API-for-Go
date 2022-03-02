/*
overlay.go

Created by JulianZander on 28.02.2022 at 16:25 with GoLand
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

// UlCreateOverlay Create a new Overlay.
//
// @param  window  The window to create the Overlay in. (we currently only
//                 support one window per application)
//
// @param  width   The width in device coordinates.
//
// @param  height  The height in device coordinates.
//
// @param  x       The x-position (offset from the left of the Window), in
//                 pixels.
//
// @param  y       The y-position (offset from the top of the Window), in
//                 pixels.
//
// @note  Each Overlay is essentially a View and an on-screen quad. You should
//        create the Overlay then load content into the underlying View.
func UlCreateOverlay(win ULWindow, width, height uint32, x, y int32) ULOverlay {
	return (ULOverlay)(C.ulCreateOverlay(win, (C.uint)(width), (C.uint)(height), (C.int)(x), (C.int)(y)))
}

// UlCreateOverlayWithView Create a new Overlay, wrapping an existing View.
//
// @param  window  The window to create the Overlay in. (we currently only
//                 support one window per application)
//
// @param  view    The View to wrap (will use its width and height).
//
// @param  x       The x-position (offset from the left of the Window), in
//                 pixels.
//
// @param  y       The y-position (offset from the top of the Window), in
//                 pixels.
//
// @note  Each Overlay is essentially a View and an on-screen quad. You should
//        create the Overlay then load content into the underlying View.
func UlCreateOverlayWithView(win ULWindow, view ULView, x, y int32) ULOverlay {
	return (ULOverlay)(C.ulCreateOverlayWithView(win, view, (C.int)(x), (C.int)(y)))
}

// UlDestroyOverlay Destroy an overlay.
func UlDestroyOverlay(overlay ULOverlay) {
	C.ulDestroyOverlay(overlay)
}

// UlOverlayGetView Get the underlying View.
func UlOverlayGetView(overlay ULOverlay) ULView {
	return (ULView)(C.ulOverlayGetView(overlay))
}

// UlOverlayGetWidth Get the width (in pixels).
func UlOverlayGetWidth(overlay ULOverlay) uint32 {
	return uint32(C.ulOverlayGetWidth(overlay))
}

// UlOverlayGetHeight Get the height (in pixels).
func UlOverlayGetHeight(overlay ULOverlay) uint32 {
	return uint32(C.ulOverlayGetHeight(overlay))
}

// UlOverlayGetX Get the x-position (offset from the left of the Window), in pixels.
func UlOverlayGetX(overlay ULOverlay) int32 {
	return int32(C.ulOverlayGetX(overlay))
}

// UlOverlayGetY Get the y-position (offset from the top of the Window), in pixels.
func UlOverlayGetY(overlay ULOverlay) int32 {
	return int32(C.ulOverlayGetY(overlay))
}

// UlOverlayMoveTo Move the overlay to a new position (in pixels).
func UlOverlayMoveTo(overlay ULOverlay, x, y int32) {
	C.ulOverlayMoveTo(overlay, (C.int)(x), (C.int)(y))
}

// UlOverlayResize Resize the overlay (and underlying View), dimensions should be
// specified in pixels.
func UlOverlayResize(overlay ULOverlay, width, height uint32) {
	C.ulOverlayResize(overlay, (C.uint)(width), (C.uint)(height))
}

// UlOverlayIsHidden Whether or not the overlay is hidden (not drawn).
func UlOverlayIsHidden(overlay ULOverlay) bool {
	return bool(C.ulOverlayIsHidden(overlay))
}

// UlOverlayHide Hide the overlay (will no longer be drawn).
func UlOverlayHide(overlay ULOverlay) {
	C.ulOverlayHide(overlay)
}

// UlOverlayShow Show the overlay.
func UlOverlayShow(overlay ULOverlay) {
	C.ulOverlayShow(overlay)
}

// UlOverlayHasFocus Whether or not an overlay has keyboard focus.
func UlOverlayHasFocus(overlay ULOverlay) bool {
	return bool(C.ulOverlayHasFocus(overlay))
}

// UlOverlayFocus Grant this overlay exclusive keyboard focus.
func UlOverlayFocus(overlay ULOverlay) {
	C.ulOverlayFocus(overlay)
}

// UlOverlayUnfocus Remove keyboard focus.
func UlOverlayUnfocus(overlay ULOverlay) {
	C.ulOverlayUnfocus(overlay)
}
