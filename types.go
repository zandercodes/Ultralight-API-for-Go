/*
types.go

Created by JulianZander on 28.02.2022 at 16:30 with GoLand
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
import "unsafe"

/******************************************************************************
 * Ultralight/CAPI.h
 *****************************************************************************/

type ULChar16 uint16

type ULConfig C.ULConfig
type ULRenderer C.ULRenderer
type ULSession C.ULSession
type ULView C.ULView
type ULBitmap C.ULBitmap
type ULString C.ULString
type ULBuffer C.ULBuffer
type ULKeyEvent C.ULKeyEvent
type ULMouseEvent C.ULMouseEvent
type ULScrollEvent C.ULScrollEvent
type ULSurface C.ULSurface
type ULBitmapSurface C.ULBitmapSurface

type ULRect struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

type ULIntRect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type ULRenderTarget struct {
	IsEmpty        bool
	Width          uint32
	Height         uint32
	TextureId      uint32
	TextureWidth   uint32
	TextureHeight  uint32
	TextureFormat  ULBitmapFormat
	UvCoords       ULRect
	RenderBufferId uint32
}

type ULChangeTitleCallback func(userData unsafe.Pointer, caller ULView, title ULString)
type ULChangeURLCallback func(userData unsafe.Pointer, caller ULView, url ULString)
type ULChangeTooltipCallback func(userData unsafe.Pointer, caller ULView, tooltip ULString)
type ULChangeCursorCallback func(userData unsafe.Pointer, caller ULView, cursor ULCursor)
type ULAddConsoleMessageCallback func(userData unsafe.Pointer, caller ULView, source ULMessageSource, level ULMessageLevel, message ULString, lineNumber uint32, columnNumber uint32, sourceId ULString)
type ULCreateChildViewCallback func(userData unsafe.Pointer, caller ULView, openerUrl ULString, targetUrl ULString, isPopup bool, popupRect ULIntRect) ULView
type ULBeginLoadingCallback func(userData unsafe.Pointer, caller ULView, frameId uint64, isMainFrame bool, url ULString)
type ULFinishLoadingCallback func(userData unsafe.Pointer, caller ULView, frameId uint64, isMainFrame bool, url ULString)
type ULFailLoadingCallback func(userData unsafe.Pointer, caller ULView, frameId uint64, isMainFrame bool, url ULString, description ULString, errorDomain ULString, errorCode int32)
type ULWindowObjectReadyCallback func(userData unsafe.Pointer, caller ULView, frameId uint64, isMainFrame bool, url ULString)
type ULDOMReadyCallback func(userData unsafe.Pointer, caller ULView, frameId uint64, isMainFrame bool, url ULString)
type ULUpdateHistoryCallback func(userData unsafe.Pointer, caller ULView)
type ULSurfaceDefinitionCreateCallback func(width uint32, height uint32) unsafe.Pointer
type ULSurfaceDefinitionDestroyCallback func(userData unsafe.Pointer)
type ULSurfaceDefinitionGetWidthCallback func(userData unsafe.Pointer) uint32
type ULSurfaceDefinitionGetHeightCallback func(userData unsafe.Pointer) uint32
type ULSurfaceDefinitionGetRowBytesCallback func(userData unsafe.Pointer) uint32
type ULSurfaceDefinitionGetSizeCallback func(userData unsafe.Pointer) uint32
type ULSurfaceDefinitionLockPixelsCallback func(userData unsafe.Pointer) unsafe.Pointer
type ULSurfaceDefinitionUnlockPixelsCallback func(userData unsafe.Pointer)
type ULSurfaceDefinitionResizeCallback func(userData unsafe.Pointer, width uint32, height uint32)

type ULSurfaceDefinition struct {
	Create       ULSurfaceDefinitionCreateCallback
	Destroy      ULSurfaceDefinitionDestroyCallback
	GetWidth     ULSurfaceDefinitionGetWidthCallback
	GetHeight    ULSurfaceDefinitionGetHeightCallback
	GetRowBytes  ULSurfaceDefinitionGetRowBytesCallback
	GetSize      ULSurfaceDefinitionGetSizeCallback
	LockPixels   ULSurfaceDefinitionLockPixelsCallback
	UnlockPixels ULSurfaceDefinitionUnlockPixelsCallback
	Resize       ULSurfaceDefinitionResizeCallback
}

type ULFileHandle int32
type ULFileSystemFileExistsCallback func(path ULString) bool
type ULFileSystemGetFileSizeCallback func(handle ULFileHandle, result []int64) bool
type ULFileSystemGetFileMimeTypeCallback func(path ULString, result ULString) bool
type ULFileSystemOpenFileCallback func(path ULString, openForWriting bool) ULFileHandle
type ULFileSystemCloseFileCallback func(handle ULFileHandle)
type ULFileSystemReadFromFileCallback func(handle ULFileHandle, data []byte, length int64) int64

type ULFileSystem struct {
	FileExists      ULFileSystemFileExistsCallback
	GetFileSize     ULFileSystemGetFileSizeCallback
	GetFileMimeType ULFileSystemGetFileMimeTypeCallback
	OpenFile        ULFileSystemOpenFileCallback
	CloseFile       ULFileSystemCloseFileCallback
	ReadFromFile    ULFileSystemReadFromFileCallback
}

type ULLoggerLogMessageCallback func(logLevel ULLogLevel, message ULString)

type ULLogger struct {
	LogMessage ULLoggerLogMessageCallback
}

type ULRenderBuffer struct {
	TextureId        uint32
	Width            uint32
	Height           uint32
	HasStencilBuffer bool
	HasDepthBuffer   bool
}

type ULVertex_2f_4ub_2f struct {
	Pos   [2]float32
	Color [4]byte
	Obj   [2]float32
}

type ULVertex_2f_4ub_2f_2f_28f struct {
	Pos   [2]float32
	Color [4]byte
	Tex   [2]float32
	Obj   [2]float32
	Data0 [4]float32
	Data1 [4]float32
	Data2 [4]float32
	Data3 [4]float32
	Data4 [4]float32
	Data5 [4]float32
	Data6 [4]float32
}

type ULVertexBuffer struct {
	Format ULVertexBufferFormat
	Size   uint32
	Data   []byte
}

type ULIndexType uint32

type ULIndexBuffer struct {
	Size uint32
	Data []byte
}

type ULMatrix4x4 struct {
	Data [16]float32
}

type ULvec4 struct {
	Value [4]float32
}

type ULGPUState struct {
	ViewportWidth   uint32
	ViewportHeight  uint32
	Transform       ULMatrix4x4
	EnableTexturing bool
	EnableBlend     bool
	ShaderType      byte
	RenderBufferId  uint32
	Texture1Id      uint32
	Texture2Id      uint32
	Texture3Id      uint32
	UniformScalar   [8]float32
	UniformVector   [8]ULvec4
	ClipSize        byte
	Clip            [8]ULMatrix4x4
	EnableScissor   bool
	ScissorRect     ULIntRect
}

type ULCommand struct {
	CommandType   byte
	GpuState      ULGPUState
	GeometryId    uint32
	IndicesCount  uint32
	IndicesOffset uint32
}

type ULCommandList struct {
	Size     uint32
	Commands []ULCommand
}

type ULGPUDriverBeginSynchronizeCallback func()
type ULGPUDriverEndSynchronizeCallback func()
type ULGPUDriverNextTextureIdCallback func() uint32
type ULGPUDriverCreateTextureCallback func(textureId uint32, bitmap ULBitmap)
type ULGPUDriverUpdateTextureCallback func(textureId uint32, bitmap ULBitmap)
type ULGPUDriverDestroyTextureCallback func(textureId uint32)
type ULGPUDriverNextRenderBufferIdCallback func() uint32
type ULGPUDriverCreateRenderBufferCallback func(renderBufferId uint32, buffer ULRenderBuffer)
type ULGPUDriverDestroyRenderBufferCallback func(renderBufferId uint32)
type ULGPUDriverNextGeometryIdCallback func() uint32
type ULGPUDriverCreateGeometryCallback func(geometryId uint32, vertices ULVertexBuffer, indices ULIndexBuffer)
type ULGPUDriverUpdateGeometryCallback func(geometryId uint32, vertices ULVertexBuffer, indices ULIndexBuffer)
type ULGPUDriverDestroyGeometryCallback func(geometryId uint32)
type ULGPUDriverUpdateCommandListCallback func(list ULCommandList)

type ULGPUDriver struct {
	BeginSynchronize    ULGPUDriverBeginSynchronizeCallback
	EndSynchronize      ULGPUDriverEndSynchronizeCallback
	NextTextureId       ULGPUDriverNextTextureIdCallback
	CreateTexture       ULGPUDriverCreateTextureCallback
	UpdateTexture       ULGPUDriverUpdateTextureCallback
	DestroyTexture      ULGPUDriverDestroyTextureCallback
	NextRenderBufferId  ULGPUDriverNextRenderBufferIdCallback
	CreateRenderBuffer  ULGPUDriverCreateRenderBufferCallback
	DestroyRenderBuffer ULGPUDriverDestroyRenderBufferCallback
	NextGeometryId      ULGPUDriverNextGeometryIdCallback
	CreateGeometry      ULGPUDriverCreateGeometryCallback
	UpdateGeometry      ULGPUDriverUpdateGeometryCallback
	DestroyGeometry     ULGPUDriverDestroyGeometryCallback
	UpdateCommandList   ULGPUDriverUpdateCommandListCallback
}

type ULClipboardClearCallback func()
type ULClipboardReadPlainTextCallback func(result ULString)
type ULClipboardWritePlainTextCallback func(text ULString)

type ULClipboard struct {
	Clear          ULClipboardClearCallback
	ReadPlainText  ULClipboardReadPlainTextCallback
	WritePlainText ULClipboardWritePlainTextCallback
}

type ULMessageSource int32

const (
	MessageSourceXML            = ULMessageSource(C.kMessageSource_XML)
	MessageSourceJS             = ULMessageSource(C.kMessageSource_JS)
	MessageSourceNetwork        = ULMessageSource(C.kMessageSource_Network)
	MessageSourceConsoleAPI     = ULMessageSource(C.kMessageSource_ConsoleAPI)
	MessageSourceStorage        = ULMessageSource(C.kMessageSource_Storage)
	MessageSourceAppCache       = ULMessageSource(C.kMessageSource_AppCache)
	MessageSourceRendering      = ULMessageSource(C.kMessageSource_Rendering)
	MessageSourceCSS            = ULMessageSource(C.kMessageSource_CSS)
	MessageSourceSecurity       = ULMessageSource(C.kMessageSource_Security)
	MessageSourceContentBlocker = ULMessageSource(C.kMessageSource_ContentBlocker)
	MessageSourceOther          = ULMessageSource(C.kMessageSource_Other)
)

type ULMessageLevel int32

const (
	MessageLevelLog     = ULMessageLevel(C.kMessageLevel_Log)
	MessageLevelWarning = ULMessageLevel(C.kMessageLevel_Warning)
	MessageLevelError   = ULMessageLevel(C.kMessageLevel_Error)
	MessageLevelDebug   = ULMessageLevel(C.kMessageLevel_Debug)
	MessageLevelInfo    = ULMessageLevel(C.kMessageLevel_Info)
)

type ULCursor int32

const (
	CursorPointer                  = ULCursor(C.kCursor_Pointer)
	CursorCross                    = ULCursor(C.kCursor_Cross)
	CursorHand                     = ULCursor(C.kCursor_Hand)
	CursorIBeam                    = ULCursor(C.kCursor_IBeam)
	CursorWait                     = ULCursor(C.kCursor_Wait)
	CursorHelp                     = ULCursor(C.kCursor_Help)
	CursorEastResize               = ULCursor(C.kCursor_EastResize)
	CursorNorthResize              = ULCursor(C.kCursor_NorthResize)
	CursorNorthEastResize          = ULCursor(C.kCursor_NorthEastResize)
	CursorNorthWestResize          = ULCursor(C.kCursor_NorthWestResize)
	CursorSouthResize              = ULCursor(C.kCursor_SouthResize)
	CursorSouthEastResize          = ULCursor(C.kCursor_SouthEastResize)
	CursorSouthWestResize          = ULCursor(C.kCursor_SouthWestResize)
	CursorWestResize               = ULCursor(C.kCursor_WestResize)
	CursorNorthSouthResize         = ULCursor(C.kCursor_NorthSouthResize)
	CursorEastWestResize           = ULCursor(C.kCursor_EastWestResize)
	CursorNorthEastSouthWestResize = ULCursor(C.kCursor_NorthEastSouthWestResize)
	CursorNorthWestSouthEastResize = ULCursor(C.kCursor_NorthWestSouthEastResize)
	CursorColumnResize             = ULCursor(C.kCursor_ColumnResize)
	CursorRowResize                = ULCursor(C.kCursor_RowResize)
	CursorMiddlePanning            = ULCursor(C.kCursor_MiddlePanning)
	CursorEastPanning              = ULCursor(C.kCursor_EastPanning)
	CursorNorthPanning             = ULCursor(C.kCursor_NorthPanning)
	CursorNorthEastPanning         = ULCursor(C.kCursor_NorthEastPanning)
	CursorNorthWestPanning         = ULCursor(C.kCursor_NorthWestPanning)
	CursorSouthPanning             = ULCursor(C.kCursor_SouthPanning)
	CursorSouthEastPanning         = ULCursor(C.kCursor_SouthEastPanning)
	CursorSouthWestPanning         = ULCursor(C.kCursor_SouthWestPanning)
	CursorWestPanning              = ULCursor(C.kCursor_WestPanning)
	CursorMove                     = ULCursor(C.kCursor_Move)
	CursorVerticalText             = ULCursor(C.kCursor_VerticalText)
	CursorCell                     = ULCursor(C.kCursor_Cell)
	CursorContextMenu              = ULCursor(C.kCursor_ContextMenu)
	CursorAlias                    = ULCursor(C.kCursor_Alias)
	CursorProgress                 = ULCursor(C.kCursor_Progress)
	CursorNoDrop                   = ULCursor(C.kCursor_NoDrop)
	CursorCopy                     = ULCursor(C.kCursor_Copy)
	CursorNone                     = ULCursor(C.kCursor_None)
	CursorNotAllowed               = ULCursor(C.kCursor_NotAllowed)
	CursorZoomIn                   = ULCursor(C.kCursor_ZoomIn)
	CursorZoomOut                  = ULCursor(C.kCursor_ZoomOut)
	CursorGrab                     = ULCursor(C.kCursor_Grab)
	CursorGrabbing                 = ULCursor(C.kCursor_Grabbing)
	CursorCustom                   = ULCursor(C.kCursor_Custom)
)

type ULBitmapFormat int32

const (
	BitmapFormat_A8_UNORM         = ULBitmapFormat(C.kBitmapFormat_A8_UNORM)
	BitmapFormat_BGRA8_UNORM_SRGB = ULBitmapFormat(C.kBitmapFormat_BGRA8_UNORM_SRGB)
)

type ULKeyEventType int32

const (
	KeyEventTypeKeyDown    = ULKeyEventType(C.kKeyEventType_KeyDown)
	KeyEventTypeKeyUp      = ULKeyEventType(C.kKeyEventType_KeyUp)
	KeyEventTypeRawKeyDown = ULKeyEventType(C.kKeyEventType_RawKeyDown)
	KeyEventTypeChar       = ULKeyEventType(C.kKeyEventType_Char)
)

type ULMouseEventType int32

const (
	MouseEventTypeMouseMoved = ULMouseEventType(C.kMouseEventType_MouseMoved)
	MouseEventTypeMouseDown  = ULMouseEventType(C.kMouseEventType_MouseDown)
	MouseEventTypeMouseUp    = ULMouseEventType(C.kMouseEventType_MouseUp)
)

type ULMouseButton int32

const (
	MouseButtonNone   = ULMouseButton(C.kMouseButton_None)
	MouseButtonLeft   = ULMouseButton(C.kMouseButton_Left)
	MouseButtonMiddle = ULMouseButton(C.kMouseButton_Middle)
	MouseButtonRight  = ULMouseButton(C.kMouseButton_Right)
)

type ULScrollEventType int32

const (
	ScrollEventTypeScrollByPixel = ULScrollEventType(C.kScrollEventType_ScrollByPixel)
	ScrollEventTypeScrollByPage  = ULScrollEventType(C.kScrollEventType_ScrollByPage)
)

type ULFaceWinding int32

const (
	FaceWindingClockwise       = ULFaceWinding(C.kFaceWinding_Clockwise)
	FaceWindowCounterClockwise = ULFaceWinding(C.kFaceWindow_CounterClockwise)
)

type ULFontHinting int32

const (
	FontHintingSmooth     = ULFontHinting(C.kFontHinting_Smooth)
	FontHintingNormal     = ULFontHinting(C.kFontHinting_Normal)
	FontHintingMonochrome = ULFontHinting(C.kFontHinting_Monochrome)
)

type ULLogLevel int32

const (
	LogLevelError   = ULLogLevel(C.kLogLevel_Error)
	LogLevelWarning = ULLogLevel(C.kLogLevel_Warning)
	LogLevelInfo    = ULLogLevel(C.kLogLevel_Info)
)

type ULVertexBufferFormat int32

const (
	VertexBufferFormat_2f_4ub_2f        = ULVertexBufferFormat(C.kVertexBufferFormat_2f_4ub_2f)
	VertexBufferFormat_2f_4ub_2f_2f_28f = ULVertexBufferFormat(C.kVertexBufferFormat_2f_4ub_2f_2f_28f)
)

type ULShaderType int32

const (
	ShaderTypeFill     = ULShaderType(C.kShaderType_Fill)
	ShaderTypeFillPath = ULShaderType(C.kShaderType_FillPath)
)

type ULCommandType int32

const (
	CommandTypeClearRenderBuffer = ULCommandType(C.kCommandType_ClearRenderBuffer)
	CommandTypeDrawGeometry      = ULCommandType(C.kCommandType_DrawGeometry)
)

/******************************************************************************
 * AppCore/CAPI.h
 *****************************************************************************/

type ULSettings C.ULSettings
type ULApp C.ULApp
type ULWindow C.ULWindow
type ULMonitor C.ULMonitor
type ULOverlay C.ULOverlay

type ULUpdateCallback func(userData unsafe.Pointer)
type ULCloseCallback func(userData unsafe.Pointer)
type ULResizeCallback func(userData unsafe.Pointer, width uint32, height uint32)

type ULWindowFlags uint32

const (
	WindowFlagsBorderless  = ULWindowFlags(C.kWindowFlags_Borderless)
	WindowFlagsTitled      = ULWindowFlags(C.kWindowFlags_Titled)
	WindowFlagsResizable   = ULWindowFlags(C.kWindowFlags_Resizable)
	WindowFlagsMaximizable = ULWindowFlags(C.kWindowFlags_Maximizable)
)

/******************************************************************************
 * JavaScriptCore
 *****************************************************************************/

/* JavaScript engine interface */
type JSContextGroupRef C.JSContextGroupRef
type JSContextRef C.JSContextRef
type JSGlobalContextRef C.JSGlobalContextRef
type JSStringRef C.JSStringRef
type JSClassRef C.JSClassRef
type JSPropertyNameArrayRef C.JSPropertyNameArrayRef
type JSPropertyNameAccumulatorRef C.JSPropertyNameAccumulatorRef

type JSTypedArrayBytesDeallocator func(bytes unsafe.Pointer, deallocatorContext unsafe.Pointer)

/* JavaScript data types */
type JSValueRef C.JSValueRef
type JSObjectRef C.JSObjectRef

type JSPropertyAttributes uint32

type JSType int32

const (
	JSTypeUndefined = JSType(C.kJSTypeUndefined)
	JSTypeNull      = JSType(C.kJSTypeNull)
	JSTypeBoolean   = JSType(C.kJSTypeBoolean)
	JSTypeNumber    = JSType(C.kJSTypeNumber)
	JSTypeString    = JSType(C.kJSTypeString)
	JSTypeObject    = JSType(C.kJSTypeObject)
	JSTypeSymbol    = JSType(C.kJSTypeSymbol)
)

type JSTypedArrayType int32
type JSClassAttributes uint32

type JSObjectInitializeCallback func(ctx JSContextRef, object JSObjectRef)
type JSObjectInitializeCallbackEx func(ctx JSContextRef, jsClass JSClassRef, object JSObjectRef)
type JSObjectFinalizeCallback func(object JSObjectRef)
type JSObjectFinalizeCallbackEx func(jsClass JSClassRef, object JSObjectRef)
type JSObjectHasPropertyCallback func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef) bool
type JSObjectHasPropertyCallbackEx func(ctx JSContextRef, jsClass JSClassRef, object JSObjectRef, propertyName JSStringRef) bool
type JSObjectGetPropertyCallback func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception []JSValueRef) JSValueRef
type JSObjectGetPropertyCallbackEx func(ctx JSContextRef, jsClass JSClassRef, object JSObjectRef, propertyName JSStringRef, exception []JSValueRef) JSValueRef
type JSObjectSetPropertyCallback func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, value JSValueRef, exception []JSValueRef) bool
type JSObjectSetPropertyCallbackEx func(ctx JSContextRef, jsClass JSClassRef, object JSObjectRef, propertyName JSStringRef, value JSValueRef, exception []JSValueRef) bool
type JSObjectDeletePropertyCallback func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception []JSValueRef) bool
type JSObjectDeletePropertyCallbackEx func(ctx JSContextRef, jsClass JSClassRef, object JSObjectRef, propertyName JSStringRef, exception []JSValueRef) bool
type JSObjectGetPropertyNamesCallback func(ctx JSContextRef, object JSObjectRef, propertyNames JSPropertyNameAccumulatorRef)
type JSObjectGetPropertyNamesCallbackEx func(ctx JSContextRef, jsClass JSClassRef, object JSObjectRef, propertyNames JSPropertyNameAccumulatorRef)
type JSObjectCallAsFunctionCallback func(ctx JSContextRef, function JSObjectRef, thisObject JSObjectRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSValueRef
type JSObjectCallAsFunctionCallbackEx func(ctx JSContextRef, jsClass JSClassRef, className JSStringRef, function JSObjectRef, thisObject JSObjectRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSValueRef
type JSObjectCallAsConstructorCallback func(ctx JSContextRef, constructor JSObjectRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSObjectRef
type JSObjectCallAsConstructorCallbackEx func(ctx JSContextRef, jsClass JSClassRef, constructor JSObjectRef, argumentCount uint32, arguments []JSValueRef, exception []JSValueRef) JSObjectRef
type JSObjectHasInstanceCallback func(ctx JSContextRef, constructor JSObjectRef, possibleInstance JSValueRef, exception []JSValueRef) bool
type JSObjectHasInstanceCallbackEx func(ctx JSContextRef, jsClass JSClassRef, constructor JSObjectRef, possibleInstance JSValueRef, exception []JSValueRef) bool
type JSObjectConvertToTypeCallback func(ctx JSContextRef, object JSObjectRef, _type JSType, exception []JSValueRef) JSValueRef
type JSObjectConvertToTypeCallbackEx func(ctx JSContextRef, jsClass JSClassRef, object JSObjectRef, _type JSType, exception []JSValueRef) JSValueRef

type JSStaticValue struct {
	Name        string
	GetProperty JSObjectGetPropertyCallback
	SetProperty JSObjectSetPropertyCallback
	Attributes  JSPropertyAttributes
}

type JSStaticValueEx struct {
	Name          string
	GetPropertyEx JSObjectGetPropertyCallbackEx
	SetPropertyEx JSObjectSetPropertyCallbackEx
	Attributes    JSPropertyAttributes
}

type JSStaticFunction struct {
	Name           string
	CallAsFunction JSObjectCallAsFunctionCallback
	Attributes     JSPropertyAttributes
}

type JSStaticFunctionEx struct {
	Name             string
	CallAsFunctionEx JSObjectCallAsFunctionCallbackEx
	Attributes       JSPropertyAttributes
}

type JSClassDefinition struct {
	Version     int32
	Attributes  JSClassAttributes
	ClassName   string
	ParentClass JSClassRef
	PrivateData unsafe.Pointer
}

type JSChar uint16

const (
	JSTypedArrayTypeInt8Array         = JSTypedArrayType(C.kJSTypedArrayTypeInt8Array)
	JSTypedArrayTypeInt16Array        = JSTypedArrayType(C.kJSTypedArrayTypeInt16Array)
	JSTypedArrayTypeInt32Array        = JSTypedArrayType(C.kJSTypedArrayTypeInt32Array)
	JSTypedArrayTypeUint8Array        = JSTypedArrayType(C.kJSTypedArrayTypeUint8Array)
	JSTypedArrayTypeUint8ClampedArray = JSTypedArrayType(C.kJSTypedArrayTypeUint8ClampedArray)
	JSTypedArrayTypeUint16Array       = JSTypedArrayType(C.kJSTypedArrayTypeUint16Array)
	JSTypedArrayTypeUint32Array       = JSTypedArrayType(C.kJSTypedArrayTypeUint32Array)
	JSTypedArrayTypeFloat32Array      = JSTypedArrayType(C.kJSTypedArrayTypeFloat32Array)
	JSTypedArrayTypeFloat64Array      = JSTypedArrayType(C.kJSTypedArrayTypeFloat64Array)
	JSTypedArrayTypeArrayBuffer       = JSTypedArrayType(C.kJSTypedArrayTypeArrayBuffer)
	JSTypedArrayTypeNone              = JSTypedArrayType(C.kJSTypedArrayTypeNone)
)

const (
	JSPropertyAttributeNone       = int32(C.kJSPropertyAttributeNone)
	JSPropertyAttributeReadOnly   = int32(C.kJSPropertyAttributeReadOnly)
	JSPropertyAttributeDontEnum   = int32(C.kJSPropertyAttributeDontEnum)
	JSPropertyAttributeDontDelete = int32(C.kJSPropertyAttributeDontDelete)
)

const (
	JSClassAttributeNone                 = int32(C.kJSClassAttributeNone)
	JSClassAttributeNoAutomaticPrototype = int32(C.kJSClassAttributeNoAutomaticPrototype)
)
