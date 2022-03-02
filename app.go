/*
app.go

Created by JulianZander on 28.02.2022 at 16:24 with GoLand
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
import "unsafe"

// UlCreateApp Create the app singleton.
//
// @param  settings  ULSettings to customize app runtime behavior. You can pass
//                   NULL for this parameter to use default settings.
//
// @param  cfg  ULConfig options for the Ultralight renderer. You can pass
//              NULL for this parameter to use default config.
//
// @note  You should only create one of these per application lifetime.
//
// @note  Certain Config options may be overridden during app creation,
//        most commonly Config::face_winding and Config::device_scale_hint.
func UlCreateApp(stgs ULSettings, cfg ULConfig) ULApp {
	return (ULApp)(C.ulCreateApp(stgs, cfg))
}

// UlDestroyApp Destroy the App instance.
func UlDestroyApp(app ULApp) {
	C.ulDestroyApp(app)
}

// UlAppSetWindow Set the main window, you must set this before calling ulAppRun.
//
// @param  window  The window to use for all rendering.
//
// @note  We currently only support one window per app, this will change
//         later once we add support for multiple driver instances.
func UlAppSetWindow(app ULApp, window ULWindow) {
	C.ulAppSetWindow(app, window)
}

// UlAppGetWindow Get the main window.
func UlAppGetWindow(app ULApp) ULWindow {
	return (ULWindow)(C.ulAppGetWindow(app))
}

// UlAppSetUpdateCallback Set a callback for whenever the App updates. You should update all app
// logic here.
//
// @note  This event is fired right before the run loop calls
//        Renderer::Update and Renderer::Render.
func UlAppSetUpdateCallback(app ULApp, callback ULUpdateCallback, userData unsafe.Pointer) {
	cCallback := *(*C.ULUpdateCallback)(unsafe.Pointer(&callback))
	C.ulAppSetUpdateCallback(app, cCallback, userData)
}

// UlAppIsRunning Whether or not the app is running.
func UlAppIsRunning(app ULApp) bool {
	return bool(C.ulAppIsRunning(app))
}

// UlAppGetMainMonitor Get the main monitor (this is never NULL).
//
// @note  We'll add monitor enumeration later.
func UlAppGetMainMonitor(app ULApp) ULMonitor {
	return (ULMonitor)(C.ulAppGetMainMonitor(app))
}

// UlAppGetRenderer Get the underlying Renderer instance.
func UlAppGetRenderer(app ULApp) ULRenderer {
	return (ULRenderer)(C.ulAppGetRenderer(app))
}

// UlAppRun Run the main loop, make sure to call ulAppSetWindow before calling this.
func UlAppRun(app ULApp) {
	C.ulAppRun(app)
}

// UlAppQuit Quit the application.
func UlAppQuit(app ULApp) {
	C.ulAppQuit(app)
}
