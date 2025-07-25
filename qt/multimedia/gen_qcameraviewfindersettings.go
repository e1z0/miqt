package multimedia

/*

#include "gen_qcameraviewfindersettings.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"unsafe"
)

type QCameraViewfinderSettings struct {
	h *C.QCameraViewfinderSettings
}

func (this *QCameraViewfinderSettings) cPointer() *C.QCameraViewfinderSettings {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QCameraViewfinderSettings) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQCameraViewfinderSettings constructs the type using only CGO pointers.
func newQCameraViewfinderSettings(h *C.QCameraViewfinderSettings) *QCameraViewfinderSettings {
	if h == nil {
		return nil
	}

	return &QCameraViewfinderSettings{h: h}
}

// UnsafeNewQCameraViewfinderSettings constructs the type using only unsafe pointers.
func UnsafeNewQCameraViewfinderSettings(h unsafe.Pointer) *QCameraViewfinderSettings {
	return newQCameraViewfinderSettings((*C.QCameraViewfinderSettings)(h))
}

// NewQCameraViewfinderSettings constructs a new QCameraViewfinderSettings object.
func NewQCameraViewfinderSettings() *QCameraViewfinderSettings {

	return newQCameraViewfinderSettings(C.QCameraViewfinderSettings_new())
}

// NewQCameraViewfinderSettings2 constructs a new QCameraViewfinderSettings object.
func NewQCameraViewfinderSettings2(other *QCameraViewfinderSettings) *QCameraViewfinderSettings {

	return newQCameraViewfinderSettings(C.QCameraViewfinderSettings_new2(other.cPointer()))
}

func (this *QCameraViewfinderSettings) OperatorAssign(other *QCameraViewfinderSettings) {
	C.QCameraViewfinderSettings_operatorAssign(this.h, other.cPointer())
}

func (this *QCameraViewfinderSettings) Swap(other *QCameraViewfinderSettings) {
	C.QCameraViewfinderSettings_swap(this.h, other.cPointer())
}

func (this *QCameraViewfinderSettings) IsNull() bool {
	return (bool)(C.QCameraViewfinderSettings_isNull(this.h))
}

func (this *QCameraViewfinderSettings) Resolution() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(C.QCameraViewfinderSettings_resolution(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCameraViewfinderSettings) SetResolution(resolution *qt.QSize) {
	C.QCameraViewfinderSettings_setResolution(this.h, (*C.QSize)(resolution.UnsafePointer()))
}

func (this *QCameraViewfinderSettings) SetResolution2(width int, height int) {
	C.QCameraViewfinderSettings_setResolution2(this.h, (C.int)(width), (C.int)(height))
}

func (this *QCameraViewfinderSettings) MinimumFrameRate() float64 {
	return (float64)(C.QCameraViewfinderSettings_minimumFrameRate(this.h))
}

func (this *QCameraViewfinderSettings) SetMinimumFrameRate(rate float64) {
	C.QCameraViewfinderSettings_setMinimumFrameRate(this.h, (C.double)(rate))
}

func (this *QCameraViewfinderSettings) MaximumFrameRate() float64 {
	return (float64)(C.QCameraViewfinderSettings_maximumFrameRate(this.h))
}

func (this *QCameraViewfinderSettings) SetMaximumFrameRate(rate float64) {
	C.QCameraViewfinderSettings_setMaximumFrameRate(this.h, (C.double)(rate))
}

func (this *QCameraViewfinderSettings) PixelFormat() QVideoFrame__PixelFormat {
	return (QVideoFrame__PixelFormat)(C.QCameraViewfinderSettings_pixelFormat(this.h))
}

func (this *QCameraViewfinderSettings) SetPixelFormat(format QVideoFrame__PixelFormat) {
	C.QCameraViewfinderSettings_setPixelFormat(this.h, (C.int)(format))
}

func (this *QCameraViewfinderSettings) PixelAspectRatio() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(C.QCameraViewfinderSettings_pixelAspectRatio(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCameraViewfinderSettings) SetPixelAspectRatio(ratio *qt.QSize) {
	C.QCameraViewfinderSettings_setPixelAspectRatio(this.h, (*C.QSize)(ratio.UnsafePointer()))
}

func (this *QCameraViewfinderSettings) SetPixelAspectRatio2(horizontal int, vertical int) {
	C.QCameraViewfinderSettings_setPixelAspectRatio2(this.h, (C.int)(horizontal), (C.int)(vertical))
}

// Delete this object from C++ memory.
func (this *QCameraViewfinderSettings) Delete() {
	C.QCameraViewfinderSettings_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCameraViewfinderSettings) GoGC() {
	runtime.SetFinalizer(this, func(this *QCameraViewfinderSettings) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
