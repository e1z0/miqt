package multimedia

/*

#include "gen_qcameradevice.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"unsafe"
)

type QCameraDevice__Position int

const (
	QCameraDevice__UnspecifiedPosition QCameraDevice__Position = 0
	QCameraDevice__BackFace            QCameraDevice__Position = 1
	QCameraDevice__FrontFace           QCameraDevice__Position = 2
)

type QCameraFormat struct {
	h *C.QCameraFormat
}

func (this *QCameraFormat) cPointer() *C.QCameraFormat {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QCameraFormat) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQCameraFormat constructs the type using only CGO pointers.
func newQCameraFormat(h *C.QCameraFormat) *QCameraFormat {
	if h == nil {
		return nil
	}

	return &QCameraFormat{h: h}
}

// UnsafeNewQCameraFormat constructs the type using only unsafe pointers.
func UnsafeNewQCameraFormat(h unsafe.Pointer) *QCameraFormat {
	return newQCameraFormat((*C.QCameraFormat)(h))
}

// NewQCameraFormat constructs a new QCameraFormat object.
func NewQCameraFormat() *QCameraFormat {

	return newQCameraFormat(C.QCameraFormat_new())
}

// NewQCameraFormat2 constructs a new QCameraFormat object.
func NewQCameraFormat2(other *QCameraFormat) *QCameraFormat {

	return newQCameraFormat(C.QCameraFormat_new2(other.cPointer()))
}

func (this *QCameraFormat) OperatorAssign(other *QCameraFormat) {
	C.QCameraFormat_operatorAssign(this.h, other.cPointer())
}

func (this *QCameraFormat) PixelFormat() QVideoFrameFormat__PixelFormat {
	return (QVideoFrameFormat__PixelFormat)(C.QCameraFormat_pixelFormat(this.h))
}

func (this *QCameraFormat) Resolution() *qt6.QSize {
	_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(C.QCameraFormat_resolution(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCameraFormat) MinFrameRate() float32 {
	return (float32)(C.QCameraFormat_minFrameRate(this.h))
}

func (this *QCameraFormat) MaxFrameRate() float32 {
	return (float32)(C.QCameraFormat_maxFrameRate(this.h))
}

func (this *QCameraFormat) IsNull() bool {
	return (bool)(C.QCameraFormat_isNull(this.h))
}

func (this *QCameraFormat) OperatorEqual(other *QCameraFormat) bool {
	return (bool)(C.QCameraFormat_operatorEqual(this.h, other.cPointer()))
}

func (this *QCameraFormat) OperatorNotEqual(other *QCameraFormat) bool {
	return (bool)(C.QCameraFormat_operatorNotEqual(this.h, other.cPointer()))
}

// Delete this object from C++ memory.
func (this *QCameraFormat) Delete() {
	C.QCameraFormat_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCameraFormat) GoGC() {
	runtime.SetFinalizer(this, func(this *QCameraFormat) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QCameraDevice struct {
	h *C.QCameraDevice
}

func (this *QCameraDevice) cPointer() *C.QCameraDevice {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QCameraDevice) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQCameraDevice constructs the type using only CGO pointers.
func newQCameraDevice(h *C.QCameraDevice) *QCameraDevice {
	if h == nil {
		return nil
	}

	return &QCameraDevice{h: h}
}

// UnsafeNewQCameraDevice constructs the type using only unsafe pointers.
func UnsafeNewQCameraDevice(h unsafe.Pointer) *QCameraDevice {
	return newQCameraDevice((*C.QCameraDevice)(h))
}

// NewQCameraDevice constructs a new QCameraDevice object.
func NewQCameraDevice() *QCameraDevice {

	return newQCameraDevice(C.QCameraDevice_new())
}

// NewQCameraDevice2 constructs a new QCameraDevice object.
func NewQCameraDevice2(other *QCameraDevice) *QCameraDevice {

	return newQCameraDevice(C.QCameraDevice_new2(other.cPointer()))
}

func (this *QCameraDevice) OperatorAssign(other *QCameraDevice) {
	C.QCameraDevice_operatorAssign(this.h, other.cPointer())
}

func (this *QCameraDevice) OperatorEqual(other *QCameraDevice) bool {
	return (bool)(C.QCameraDevice_operatorEqual(this.h, other.cPointer()))
}

func (this *QCameraDevice) OperatorNotEqual(other *QCameraDevice) bool {
	return (bool)(C.QCameraDevice_operatorNotEqual(this.h, other.cPointer()))
}

func (this *QCameraDevice) IsNull() bool {
	return (bool)(C.QCameraDevice_isNull(this.h))
}

func (this *QCameraDevice) Id() []byte {
	var _bytearray C.struct_miqt_string = C.QCameraDevice_id(this.h)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QCameraDevice) Description() string {
	var _ms C.struct_miqt_string = C.QCameraDevice_description(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCameraDevice) IsDefault() bool {
	return (bool)(C.QCameraDevice_isDefault(this.h))
}

func (this *QCameraDevice) Position() QCameraDevice__Position {
	return (QCameraDevice__Position)(C.QCameraDevice_position(this.h))
}

func (this *QCameraDevice) PhotoResolutions() []qt6.QSize {
	var _ma C.struct_miqt_array = C.QCameraDevice_photoResolutions(this.h)
	_ret := make([]qt6.QSize, int(_ma.len))
	_outCast := (*[0xffff]*C.QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(_outCast[i]))
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QCameraDevice) VideoFormats() []QCameraFormat {
	var _ma C.struct_miqt_array = C.QCameraDevice_videoFormats(this.h)
	_ret := make([]QCameraFormat, int(_ma.len))
	_outCast := (*[0xffff]*C.QCameraFormat)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQCameraFormat(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

// Delete this object from C++ memory.
func (this *QCameraDevice) Delete() {
	C.QCameraDevice_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCameraDevice) GoGC() {
	runtime.SetFinalizer(this, func(this *QCameraDevice) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
