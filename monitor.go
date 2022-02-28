/*
monitor.go

Created by JulianZander on 28.02.2022 at 16:24 with GoLand
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

// UlMonitorGetScale Get the monitor's DPI scale (1.0 = 100%).
func UlMonitorGetScale(mon ULMonitor) float64 {
	return float64(C.ulMonitorGetScale(mon))
}

// UlMonitorGetWidth Get the width of the monitor (in pixels).
func UlMonitorGetWidth(mon ULMonitor) uint32 {
	return uint32(C.ulMonitorGetWidth(mon))
}

// UlMonitorGetHeight Get the height of the monitor (in pixels).
func UlMonitorGetHeight(mon ULMonitor) uint32 {
	return uint32(C.ulMonitorGetHeight(mon))
}
