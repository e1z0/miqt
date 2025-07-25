package multimedia

/*

#include "gen_qvideoframeformat.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"unsafe"
)

type QVideoFrameFormat__PixelFormat int

const (
	QVideoFrameFormat__Format_Invalid                QVideoFrameFormat__PixelFormat = 0
	QVideoFrameFormat__Format_ARGB8888               QVideoFrameFormat__PixelFormat = 1
	QVideoFrameFormat__Format_ARGB8888_Premultiplied QVideoFrameFormat__PixelFormat = 2
	QVideoFrameFormat__Format_XRGB8888               QVideoFrameFormat__PixelFormat = 3
	QVideoFrameFormat__Format_BGRA8888               QVideoFrameFormat__PixelFormat = 4
	QVideoFrameFormat__Format_BGRA8888_Premultiplied QVideoFrameFormat__PixelFormat = 5
	QVideoFrameFormat__Format_BGRX8888               QVideoFrameFormat__PixelFormat = 6
	QVideoFrameFormat__Format_ABGR8888               QVideoFrameFormat__PixelFormat = 7
	QVideoFrameFormat__Format_XBGR8888               QVideoFrameFormat__PixelFormat = 8
	QVideoFrameFormat__Format_RGBA8888               QVideoFrameFormat__PixelFormat = 9
	QVideoFrameFormat__Format_RGBX8888               QVideoFrameFormat__PixelFormat = 10
	QVideoFrameFormat__Format_AYUV                   QVideoFrameFormat__PixelFormat = 11
	QVideoFrameFormat__Format_AYUV_Premultiplied     QVideoFrameFormat__PixelFormat = 12
	QVideoFrameFormat__Format_YUV420P                QVideoFrameFormat__PixelFormat = 13
	QVideoFrameFormat__Format_YUV422P                QVideoFrameFormat__PixelFormat = 14
	QVideoFrameFormat__Format_YV12                   QVideoFrameFormat__PixelFormat = 15
	QVideoFrameFormat__Format_UYVY                   QVideoFrameFormat__PixelFormat = 16
	QVideoFrameFormat__Format_YUYV                   QVideoFrameFormat__PixelFormat = 17
	QVideoFrameFormat__Format_NV12                   QVideoFrameFormat__PixelFormat = 18
	QVideoFrameFormat__Format_NV21                   QVideoFrameFormat__PixelFormat = 19
	QVideoFrameFormat__Format_IMC1                   QVideoFrameFormat__PixelFormat = 20
	QVideoFrameFormat__Format_IMC2                   QVideoFrameFormat__PixelFormat = 21
	QVideoFrameFormat__Format_IMC3                   QVideoFrameFormat__PixelFormat = 22
	QVideoFrameFormat__Format_IMC4                   QVideoFrameFormat__PixelFormat = 23
	QVideoFrameFormat__Format_Y8                     QVideoFrameFormat__PixelFormat = 24
	QVideoFrameFormat__Format_Y16                    QVideoFrameFormat__PixelFormat = 25
	QVideoFrameFormat__Format_P010                   QVideoFrameFormat__PixelFormat = 26
	QVideoFrameFormat__Format_P016                   QVideoFrameFormat__PixelFormat = 27
	QVideoFrameFormat__Format_SamplerExternalOES     QVideoFrameFormat__PixelFormat = 28
	QVideoFrameFormat__Format_Jpeg                   QVideoFrameFormat__PixelFormat = 29
	QVideoFrameFormat__Format_SamplerRect            QVideoFrameFormat__PixelFormat = 30
	QVideoFrameFormat__Format_YUV420P10              QVideoFrameFormat__PixelFormat = 31
)

type QVideoFrameFormat__Direction int

const (
	QVideoFrameFormat__TopToBottom QVideoFrameFormat__Direction = 0
	QVideoFrameFormat__BottomToTop QVideoFrameFormat__Direction = 1
)

type QVideoFrameFormat__YCbCrColorSpace int

const (
	QVideoFrameFormat__YCbCr_Undefined QVideoFrameFormat__YCbCrColorSpace = 0
	QVideoFrameFormat__YCbCr_BT601     QVideoFrameFormat__YCbCrColorSpace = 1
	QVideoFrameFormat__YCbCr_BT709     QVideoFrameFormat__YCbCrColorSpace = 2
	QVideoFrameFormat__YCbCr_xvYCC601  QVideoFrameFormat__YCbCrColorSpace = 3
	QVideoFrameFormat__YCbCr_xvYCC709  QVideoFrameFormat__YCbCrColorSpace = 4
	QVideoFrameFormat__YCbCr_JPEG      QVideoFrameFormat__YCbCrColorSpace = 5
	QVideoFrameFormat__YCbCr_BT2020    QVideoFrameFormat__YCbCrColorSpace = 6
)

type QVideoFrameFormat__ColorSpace int

const (
	QVideoFrameFormat__ColorSpace_Undefined QVideoFrameFormat__ColorSpace = 0
	QVideoFrameFormat__ColorSpace_BT601     QVideoFrameFormat__ColorSpace = 1
	QVideoFrameFormat__ColorSpace_BT709     QVideoFrameFormat__ColorSpace = 2
	QVideoFrameFormat__ColorSpace_AdobeRgb  QVideoFrameFormat__ColorSpace = 5
	QVideoFrameFormat__ColorSpace_BT2020    QVideoFrameFormat__ColorSpace = 6
)

type QVideoFrameFormat__ColorTransfer int

const (
	QVideoFrameFormat__ColorTransfer_Unknown QVideoFrameFormat__ColorTransfer = 0
	QVideoFrameFormat__ColorTransfer_BT709   QVideoFrameFormat__ColorTransfer = 1
	QVideoFrameFormat__ColorTransfer_BT601   QVideoFrameFormat__ColorTransfer = 2
	QVideoFrameFormat__ColorTransfer_Linear  QVideoFrameFormat__ColorTransfer = 3
	QVideoFrameFormat__ColorTransfer_Gamma22 QVideoFrameFormat__ColorTransfer = 4
	QVideoFrameFormat__ColorTransfer_Gamma28 QVideoFrameFormat__ColorTransfer = 5
	QVideoFrameFormat__ColorTransfer_ST2084  QVideoFrameFormat__ColorTransfer = 6
	QVideoFrameFormat__ColorTransfer_STD_B67 QVideoFrameFormat__ColorTransfer = 7
)

type QVideoFrameFormat__ColorRange int

const (
	QVideoFrameFormat__ColorRange_Unknown QVideoFrameFormat__ColorRange = 0
	QVideoFrameFormat__ColorRange_Video   QVideoFrameFormat__ColorRange = 1
	QVideoFrameFormat__ColorRange_Full    QVideoFrameFormat__ColorRange = 2
)

type QVideoFrameFormat struct {
	h *C.QVideoFrameFormat
}

func (this *QVideoFrameFormat) cPointer() *C.QVideoFrameFormat {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QVideoFrameFormat) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQVideoFrameFormat constructs the type using only CGO pointers.
func newQVideoFrameFormat(h *C.QVideoFrameFormat) *QVideoFrameFormat {
	if h == nil {
		return nil
	}

	return &QVideoFrameFormat{h: h}
}

// UnsafeNewQVideoFrameFormat constructs the type using only unsafe pointers.
func UnsafeNewQVideoFrameFormat(h unsafe.Pointer) *QVideoFrameFormat {
	return newQVideoFrameFormat((*C.QVideoFrameFormat)(h))
}

// NewQVideoFrameFormat constructs a new QVideoFrameFormat object.
func NewQVideoFrameFormat() *QVideoFrameFormat {

	return newQVideoFrameFormat(C.QVideoFrameFormat_new())
}

// NewQVideoFrameFormat2 constructs a new QVideoFrameFormat object.
func NewQVideoFrameFormat2(size *qt6.QSize, pixelFormat QVideoFrameFormat__PixelFormat) *QVideoFrameFormat {

	return newQVideoFrameFormat(C.QVideoFrameFormat_new2((*C.QSize)(size.UnsafePointer()), (C.int)(pixelFormat)))
}

// NewQVideoFrameFormat3 constructs a new QVideoFrameFormat object.
func NewQVideoFrameFormat3(format *QVideoFrameFormat) *QVideoFrameFormat {

	return newQVideoFrameFormat(C.QVideoFrameFormat_new3(format.cPointer()))
}

func (this *QVideoFrameFormat) Swap(other *QVideoFrameFormat) {
	C.QVideoFrameFormat_swap(this.h, other.cPointer())
}

func (this *QVideoFrameFormat) Detach() {
	C.QVideoFrameFormat_detach(this.h)
}

func (this *QVideoFrameFormat) OperatorAssign(format *QVideoFrameFormat) {
	C.QVideoFrameFormat_operatorAssign(this.h, format.cPointer())
}

func (this *QVideoFrameFormat) OperatorEqual(format *QVideoFrameFormat) bool {
	return (bool)(C.QVideoFrameFormat_operatorEqual(this.h, format.cPointer()))
}

func (this *QVideoFrameFormat) OperatorNotEqual(format *QVideoFrameFormat) bool {
	return (bool)(C.QVideoFrameFormat_operatorNotEqual(this.h, format.cPointer()))
}

func (this *QVideoFrameFormat) IsValid() bool {
	return (bool)(C.QVideoFrameFormat_isValid(this.h))
}

func (this *QVideoFrameFormat) PixelFormat() QVideoFrameFormat__PixelFormat {
	return (QVideoFrameFormat__PixelFormat)(C.QVideoFrameFormat_pixelFormat(this.h))
}

func (this *QVideoFrameFormat) FrameSize() *qt6.QSize {
	_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(C.QVideoFrameFormat_frameSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrameFormat) SetFrameSize(size *qt6.QSize) {
	C.QVideoFrameFormat_setFrameSize(this.h, (*C.QSize)(size.UnsafePointer()))
}

func (this *QVideoFrameFormat) SetFrameSize2(width int, height int) {
	C.QVideoFrameFormat_setFrameSize2(this.h, (C.int)(width), (C.int)(height))
}

func (this *QVideoFrameFormat) FrameWidth() int {
	return (int)(C.QVideoFrameFormat_frameWidth(this.h))
}

func (this *QVideoFrameFormat) FrameHeight() int {
	return (int)(C.QVideoFrameFormat_frameHeight(this.h))
}

func (this *QVideoFrameFormat) PlaneCount() int {
	return (int)(C.QVideoFrameFormat_planeCount(this.h))
}

func (this *QVideoFrameFormat) Viewport() *qt6.QRect {
	_goptr := qt6.UnsafeNewQRect(unsafe.Pointer(C.QVideoFrameFormat_viewport(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoFrameFormat) SetViewport(viewport *qt6.QRect) {
	C.QVideoFrameFormat_setViewport(this.h, (*C.QRect)(viewport.UnsafePointer()))
}

func (this *QVideoFrameFormat) ScanLineDirection() QVideoFrameFormat__Direction {
	return (QVideoFrameFormat__Direction)(C.QVideoFrameFormat_scanLineDirection(this.h))
}

func (this *QVideoFrameFormat) SetScanLineDirection(direction QVideoFrameFormat__Direction) {
	C.QVideoFrameFormat_setScanLineDirection(this.h, (C.int)(direction))
}

func (this *QVideoFrameFormat) FrameRate() float64 {
	return (float64)(C.QVideoFrameFormat_frameRate(this.h))
}

func (this *QVideoFrameFormat) SetFrameRate(rate float64) {
	C.QVideoFrameFormat_setFrameRate(this.h, (C.double)(rate))
}

func (this *QVideoFrameFormat) YCbCrColorSpace() QVideoFrameFormat__YCbCrColorSpace {
	return (QVideoFrameFormat__YCbCrColorSpace)(C.QVideoFrameFormat_yCbCrColorSpace(this.h))
}

func (this *QVideoFrameFormat) SetYCbCrColorSpace(colorSpace QVideoFrameFormat__YCbCrColorSpace) {
	C.QVideoFrameFormat_setYCbCrColorSpace(this.h, (C.int)(colorSpace))
}

func (this *QVideoFrameFormat) ColorSpace() QVideoFrameFormat__ColorSpace {
	return (QVideoFrameFormat__ColorSpace)(C.QVideoFrameFormat_colorSpace(this.h))
}

func (this *QVideoFrameFormat) SetColorSpace(colorSpace QVideoFrameFormat__ColorSpace) {
	C.QVideoFrameFormat_setColorSpace(this.h, (C.int)(colorSpace))
}

func (this *QVideoFrameFormat) ColorTransfer() QVideoFrameFormat__ColorTransfer {
	return (QVideoFrameFormat__ColorTransfer)(C.QVideoFrameFormat_colorTransfer(this.h))
}

func (this *QVideoFrameFormat) SetColorTransfer(colorTransfer QVideoFrameFormat__ColorTransfer) {
	C.QVideoFrameFormat_setColorTransfer(this.h, (C.int)(colorTransfer))
}

func (this *QVideoFrameFormat) ColorRange() QVideoFrameFormat__ColorRange {
	return (QVideoFrameFormat__ColorRange)(C.QVideoFrameFormat_colorRange(this.h))
}

func (this *QVideoFrameFormat) SetColorRange(rangeVal QVideoFrameFormat__ColorRange) {
	C.QVideoFrameFormat_setColorRange(this.h, (C.int)(rangeVal))
}

func (this *QVideoFrameFormat) IsMirrored() bool {
	return (bool)(C.QVideoFrameFormat_isMirrored(this.h))
}

func (this *QVideoFrameFormat) SetMirrored(mirrored bool) {
	C.QVideoFrameFormat_setMirrored(this.h, (C.bool)(mirrored))
}

func (this *QVideoFrameFormat) VertexShaderFileName() string {
	var _ms C.struct_miqt_string = C.QVideoFrameFormat_vertexShaderFileName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoFrameFormat) FragmentShaderFileName() string {
	var _ms C.struct_miqt_string = C.QVideoFrameFormat_fragmentShaderFileName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoFrameFormat) MaxLuminance() float32 {
	return (float32)(C.QVideoFrameFormat_maxLuminance(this.h))
}

func (this *QVideoFrameFormat) SetMaxLuminance(lum float32) {
	C.QVideoFrameFormat_setMaxLuminance(this.h, (C.float)(lum))
}

func QVideoFrameFormat_PixelFormatFromImageFormat(format qt6.QImage__Format) QVideoFrameFormat__PixelFormat {
	return (QVideoFrameFormat__PixelFormat)(C.QVideoFrameFormat_pixelFormatFromImageFormat((C.int)(format)))
}

func QVideoFrameFormat_ImageFormatFromPixelFormat(format QVideoFrameFormat__PixelFormat) qt6.QImage__Format {
	return (qt6.QImage__Format)(C.QVideoFrameFormat_imageFormatFromPixelFormat((C.int)(format)))
}

func QVideoFrameFormat_PixelFormatToString(pixelFormat QVideoFrameFormat__PixelFormat) string {
	var _ms C.struct_miqt_string = C.QVideoFrameFormat_pixelFormatToString((C.int)(pixelFormat))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QVideoFrameFormat) Delete() {
	C.QVideoFrameFormat_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QVideoFrameFormat) GoGC() {
	runtime.SetFinalizer(this, func(this *QVideoFrameFormat) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
