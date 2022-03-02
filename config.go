/*
config.go

Created by JulianZander on 28.02.2022 at 16:15 with GoLand
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

// UlCreateConfig Create config with default values (see <Ultralight/platform/Config.h>).
func UlCreateConfig() ULConfig {
	// return *(*ULConfig)(unsafe.Pointer(C.ulCreateConfig()))
	return (ULConfig)(C.ulCreateConfig())
}

// UlDestroyConfig Destroy config.
func UlDestroyConfig(config ULConfig) {
	C.ulDestroyConfig(config)
}

// UlConfigSetResourcePath Set the file path to the directory that contains Ultralight's bundled
// resources (eg, cacert.pem and other localized resources).
func UlConfigSetResourcePath(config ULConfig, resourcePath ULString) {
	C.ulConfigSetResourcePath(config, resourcePath)
}

// UlConfigSetCachePath Set the file path to a writable directory that will be used to store
// cookies, cached resources, and other persistent data.
func UlConfigSetCachePath(config ULConfig, cachePath ULString) {
	C.ulConfigSetCachePath(config, cachePath)
}

// UlConfigSetUseGPURenderer When enabled, each view will be rendered to an offscreen GPU texture
// using the GPU driver set in ulPlatformSetGPUDriver. You can fetch
// details for the texture via ulViewGetRenderTarget.
//
// When disabled (the default), each view will be rendered to an offscreen
// pixel buffer. This pixel buffer can optionally be provided by the user--
// for more info see ulViewGetSurface.
func UlConfigSetUseGPURenderer(cfg ULConfig, useGPU bool) {
	C.ulConfigSetUseGPURenderer(cfg, (C._Bool)(useGPU))
}

// UlConfigSetDeviceScale Set the amount that the application DPI has been scaled, used for
// scaling device coordinates to pixels and oversampling raster shapes
// (Default = 1.0).
func UlConfigSetDeviceScale(cfg ULConfig, value float64) {
	C.ulConfigSetDeviceScale(cfg, (C.double)(value))
}

// UlConfigSetFaceWinding The winding order for front-facing triangles. @see FaceWinding
//
// Note: This is only used with custom GPUDrivers
func UlConfigSetFaceWinding(cfg ULConfig, fontHinting ULFontHinting) {
	C.ulConfigSetFaceWinding(cfg, (C.int)(fontHinting))
}

// UlConfigSetEnableImages Set whether images should be enabled (Default = True).
func UlConfigSetEnableImages(cfg ULConfig, enable bool) {
	C.ulConfigSetEnableImages(cfg, (C._Bool)(enable))
}

// UlConfigSetEnableJavaScript Set whether JavaScript should be enabled (Default = True).
func UlConfigSetEnableJavaScript(cfg ULConfig, enable bool) {
	C.ulConfigSetEnableJavaScript(cfg, (C._Bool)(enable))
}

// UlConfigSetFontHinting The hinting algorithm to use when rendering fonts. (Default = kFontHinting_Normal)
// @see ULFontHinting
func UlConfigSetFontHinting(cfg ULConfig, fontHinting ULFontHinting) {
	C.ulConfigSetFontHinting(cfg, fontHinting)
}

// UlConfigSetFontGamma The gamma to use when compositing font glyphs, change this value to
// adjust contrast (Adobe and Apple prefer 1.8, others may prefer 2.2).
// (Default = 1.8)
func UlConfigSetFontGamma(cfg ULConfig, fontGamma float64) {
	C.ulConfigSetFontGamma(cfg, (C.double)(fontGamma))
}

// UlConfigSetFontFamilyStandard Set default font-family to use (Default = Times New Roman).
func UlConfigSetFontFamilyStandard(cfg ULConfig, fontName ULString) {
	C.ulConfigSetFontFamilyStandard(cfg, fontName)
}

// UlConfigSetFontFamilyFixed Set default font-family to use for fixed fonts, eg <pre> and <code>
// (Default = Courier New).
func UlConfigSetFontFamilyFixed(cfg ULConfig, fontName ULString) {
	C.ulConfigSetFontFamilyFixed(cfg, fontName)
}

// UlConfigSetFontFamilySerif Set default font-family to use for serif fonts (Default = Times New Roman).
func UlConfigSetFontFamilySerif(cfg ULConfig, fontName ULString) {
	C.ulConfigSetFontFamilySerif(cfg, fontName)
}

// UlConfigSetFontFamilySansSerif Set default font-family to use for sans-serif fonts (Default = Arial).
func UlConfigSetFontFamilySansSerif(cfg ULConfig, fontName ULString) {
	C.ulConfigSetFontFamilySansSerif(cfg, fontName)
}

// UlConfigSetUserAgent Set user agent string (See <Ultralight/platform/Config.h> for the default).
func UlConfigSetUserAgent(cfg ULConfig, agent ULString) {
	C.ulConfigSetUserAgent(cfg, agent)
}

// UlConfigSetUserStylesheet Set user stylesheet (CSS) (Default = Empty).
func UlConfigSetUserStylesheet(cfg ULConfig, css ULString) {
	C.ulConfigSetUserStylesheet(cfg, css)
}

// UlConfigSetForceRepaint Set whether or not we should continuously repaint any Views or compositor
// layers, regardless if they are dirty or not. This is mainly used to
// diagnose painting/shader issues. (Default = False)
func UlConfigSetForceRepaint(cfg ULConfig, enable bool) {
	C.ulConfigSetForceRepaint(cfg, (C._Bool)(enable))
}

// UlConfigSetAnimationTimerDelay Set the amount of time to wait before triggering another repaint when a
// CSS animation is active. (Default = 1.0 / 60.0)
func UlConfigSetAnimationTimerDelay(cfg ULConfig, delay float64) {
	C.ulConfigSetAnimationTimerDelay(cfg, (C.double)(delay))
}

// UlConfigSetScrollTimerDelay When a smooth scroll animation is active, the amount of time (in seconds)
// to wait before triggering another repaint. Default is 60 Hz.
func UlConfigSetScrollTimerDelay(cfg ULConfig, delay float64) {
	C.ulConfigSetScrollTimerDelay(cfg, (C.double)(delay))
}

// UlConfigSetRecycleDelay The amount of time (in seconds) to wait before running the recycler (will
// attempt to return excess memory back to the system). (Default = 4.0)
func UlConfigSetRecycleDelay(cfg ULConfig, delay float64) {
	C.ulConfigSetRecycleDelay(cfg, (C.double)(delay))
}

// UlConfigSetMemoryCacheSize Set the size of WebCore's memory cache for decoded images, scripts, and
// other assets in bytes. (Default = 64 * 1024 * 1024)
func UlConfigSetMemoryCacheSize(cfg ULConfig, size uint32) {
	C.ulConfigSetMemoryCacheSize(cfg, (C.uint)(size))
}

// UlConfigSetPageCacheSize Set the number of pages to keep in the cache. (Default = 0)
func UlConfigSetPageCacheSize(cfg ULConfig, size uint32) {
	C.ulConfigSetPageCacheSize(cfg, (C.uint)(size))
}

// UlConfigSetOverrideRAMSize JavaScriptCore tries to detect the system's physical RAM size to set
// reasonable allocation limits. Set this to anything other than 0 to
// override the detected value. Size is in bytes.
//
// This can be used to force JavaScriptCore to be more conservative with
// its allocation strategy (at the cost of some performance).
func UlConfigSetOverrideRAMSize(cfg ULConfig, size uint32) {
	C.ulConfigSetOverrideRAMSize(cfg, (C.uint)(size))
}

// UlConfigSetMinLargeHeapSize The minimum size of large VM heaps in JavaScriptCore. Set this to a
// lower value to make these heaps start with a smaller initial value.
func UlConfigSetMinLargeHeapSize(cfg ULConfig, size uint32) {
	C.ulConfigSetMinLargeHeapSize(cfg, (C.uint)(size))
}

// UlConfigSetMinSmallHeapSize The minimum size of small VM heaps in JavaScriptCore. Set this to a
// lower value to make these heaps start with a smaller initial value.
func UlConfigSetMinSmallHeapSize(cfg ULConfig, size uint32) {
	C.ulConfigSetMinSmallHeapSize(cfg, (C.uint)(size))
}
