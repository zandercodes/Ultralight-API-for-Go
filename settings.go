/*
settings.go

Created by JulianZander on 28.02.2022 at 16:23 with GoLand
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

// UlCreateSettings Create settings with default values (see <AppCore/App.h>).
func UlCreateSettings() ULSettings {
	return (ULSettings)(C.ulCreateSettings())
}

// UlDestroySettings Destroy settings.
func UlDestroySettings(stgs ULSettings) {
	C.ulDestroySettings(stgs)
}

// UlSettingsSetDeveloperName Set the name of the developer of this app.
//
// This is used to generate a unique path to store local application data
// on the user's machine.
//
//Default is "MyCompany"
func UlSettingsSetDeveloperName(stgs ULSettings, str ULString) {
	C.ulSettingsSetDeveloperName(stgs, str)
}

// UlSettingsSetAppName Set the name of this app.
//
// This is used to generate a unique path to store local application data
// on the user's machine.
//
// Default is "MyApp"
func UlSettingsSetAppName(stgs ULSettings, name ULString) {
	C.ulSettingsSetAppName(stgs, name)
}

// UlSettingsSetFileSystemPath Set the root file path for our file system, you should set this to the
// relative path where all of your app data is.
//
// This will be used to resolve all file URLs, eg file:///page.html
//
// @note  The default path is "./assets/"
//
//        This relative path is resolved using the following logic:
//         - Windows: relative to the executable path
//         - Linux:   relative to the executable path
//         - macOS:   relative to YourApp.app/Contents/Resources/
func UlSettingsSetFileSystemPath(stgs ULSettings, path ULString) {
	C.ulSettingsSetFileSystemPath(stgs, path)
}

// UlSettingsSetLoadShadersFromFileSystem Set whether or not we should load and compile shaders from the file system
// (eg, from the /shaders/ path, relative to file_system_path).
//
// If this is false (the default), we will instead load pre-compiled shaders
// from memory which speeds up application startup time.
func UlSettingsSetLoadShadersFromFileSystem(stgs ULSettings, enabled bool) {
	cEnabled := (C._Bool)(enabled)
	C.ulSettingsSetLoadShadersFromFileSystem(stgs, cEnabled)
}

// UlSettingsSetForceCPURenderer We try to use the GPU renderer when a compatible GPU is detected.
//
// Set this to true to force the engine to always use the CPU renderer.
func UlSettingsSetForceCPURenderer(stgs ULSettings, forceCPU bool) {
	cForceCPU := (C._Bool)(forceCPU)
	C.ulSettingsSetForceCPURenderer(stgs, cForceCPU)
}
