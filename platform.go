/*
platform.go

Created by JulianZander on 28.02.2022 at 16:23 with GoLand
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

// UlPlatformSetLogger Set a custom Logger implementation.
//
// This is used to log debug messages to the console or to a log file.
//
// You should call this before ulCreateRenderer() or ulCreateApp().
//
// @note  ulCreateApp() will use the default logger if you never call this.
//
// @note  If you're not using ulCreateApp(), (eg, using ulCreateRenderer())
//        you can still use the default logger by calling
//        ulEnableDefaultLogger() (@see <AppCore/CAPI.h>)
func UlPlatformSetLogger(logger ULLogger) {
	C.ulPlatformSetLogger(logger)
}

// UlPlatformSetFileSystem Set a custom FileSystem implementation.
//
// This is used for loading File URLs (eg, <file:///page.html>). If you don't
// call this, and are not using ulCreateApp() or ulEnablePlatformFileSystem(),
// you will not be able to load any File URLs.
//
// You should call this before ulCreateRenderer() or ulCreateApp().
//
// @note  ulCreateApp() will use the default platform file system if you never
//        call this.
//
// @note  If you're not using ulCreateApp(), (eg, using ulCreateRenderer())
//        you can still use the default platform file system by calling
//        ulEnablePlatformFileSystem() (@see <AppCore/CAPI.h>)
func UlPlatformSetFileSystem(fileSystem ULFileSystem) {
	C.ulPlatformSetFileSystem(fileSystem)
}

// UlPlatformSetSurfaceDefinition Set a custom Surface implementation.
//
// This can be used to wrap a platform-specific GPU texture, Windows DIB,
// macOS CGImage, or any other pixel buffer target for display on screen.
//
// By default, the library uses a bitmap surface for all surfaces but you can
// override this by providing your own surface definition here.
//
// You should call this before ulCreateRenderer() or ulCreateApp().
func UlPlatformSetSurfaceDefinition(surfaceDefinition ULSurfaceDefinition) {
	C.ulPlatformSetSurfaceDefinition(surfaceDefinition)
}

// UlPlatformSetGPUDriver Set a custom GPUDriver implementation.
//
// This should be used if you have enabled the GPU renderer in the Config and
// are using ulCreateRenderer() (which does not provide its own GPUDriver
// implementation).
//
// The GPUDriver interface is used by the library to dispatch GPU calls to
// your native GPU context (eg, D3D11, Metal, OpenGL, Vulkan, etc.) There
// are reference implementations for this interface in the AppCore repo.
//
// You should call this before ulCreateRenderer().
func UlPlatformSetGPUDriver(gpuDriver ULGPUDriver) {
	C.ulPlatformSetGPUDriver(gpuDriver)
}

// UlPlatformSetClipboard Set a custom Clipboard implementation.
//
// This should be used if you are using ulCreateRenderer() (which does not
// provide its own clipboard implementation).
//
// The Clipboard interface is used by the library to make calls to the
// system's native clipboard (eg, cut, copy, paste).
//
// You should call this before ulCreateRenderer().
func UlPlatformSetClipboard(clipboard ULClipboard) {
	C.ulPlatformSetClipboard(clipboard)
}

// UlEnablePlatformFontLoader This is only needed if you are not calling ulCreateApp().
//
// Initializes the platform font loader and sets it as the current FontLoader.
func UlEnablePlatformFontLoader() {
	C.ulEnablePlatformFontLoader()
}

// UlEnablePlatformFileSystem This is only needed if you are not calling ulCreateApp().
//
// Initializes the platform file system (needed for loading file:/// URLs) and
// sets it as the current FileSystem.
//
// You can specify a base directory path to resolve relative paths against.
func UlEnablePlatformFileSystem(baseDir ULString) {
	C.ulEnablePlatformFileSystem(baseDir)
}

// UlEnableDefaultLogger This is only needed if you are not calling ulCreateApp().
//
// Initializes the default logger (writes the log to a file).
//
// You should specify a writable log path to write the log to
// for example "./ultralight.log"
func UlEnableDefaultLogger(logPath ULString) {
	C.ulEnableDefaultLogger(logPath)
}
