package multimedia

/*

#include "gen_qvideodeviceselectorcontrol.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QVideoDeviceSelectorControl struct {
	h *C.QVideoDeviceSelectorControl
	*QMediaControl
}

func (this *QVideoDeviceSelectorControl) cPointer() *C.QVideoDeviceSelectorControl {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QVideoDeviceSelectorControl) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQVideoDeviceSelectorControl constructs the type using only CGO pointers.
func newQVideoDeviceSelectorControl(h *C.QVideoDeviceSelectorControl) *QVideoDeviceSelectorControl {
	if h == nil {
		return nil
	}
	var outptr_QMediaControl *C.QMediaControl = nil
	C.QVideoDeviceSelectorControl_virtbase(h, &outptr_QMediaControl)

	return &QVideoDeviceSelectorControl{h: h,
		QMediaControl: newQMediaControl(outptr_QMediaControl)}
}

// UnsafeNewQVideoDeviceSelectorControl constructs the type using only unsafe pointers.
func UnsafeNewQVideoDeviceSelectorControl(h unsafe.Pointer) *QVideoDeviceSelectorControl {
	return newQVideoDeviceSelectorControl((*C.QVideoDeviceSelectorControl)(h))
}

func (this *QVideoDeviceSelectorControl) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QVideoDeviceSelectorControl_metaObject(this.h)))
}

func (this *QVideoDeviceSelectorControl) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QVideoDeviceSelectorControl_metacast(this.h, param1_Cstring))
}

func QVideoDeviceSelectorControl_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoDeviceSelectorControl_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoDeviceSelectorControl) DeviceCount() int {
	return (int)(C.QVideoDeviceSelectorControl_deviceCount(this.h))
}

func (this *QVideoDeviceSelectorControl) DeviceName(index int) string {
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_deviceName(this.h, (C.int)(index))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoDeviceSelectorControl) DeviceDescription(index int) string {
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_deviceDescription(this.h, (C.int)(index))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoDeviceSelectorControl) DefaultDevice() int {
	return (int)(C.QVideoDeviceSelectorControl_defaultDevice(this.h))
}

func (this *QVideoDeviceSelectorControl) SelectedDevice() int {
	return (int)(C.QVideoDeviceSelectorControl_selectedDevice(this.h))
}

func (this *QVideoDeviceSelectorControl) SetSelectedDevice(index int) {
	C.QVideoDeviceSelectorControl_setSelectedDevice(this.h, (C.int)(index))
}

func (this *QVideoDeviceSelectorControl) SelectedDeviceChanged(index int) {
	C.QVideoDeviceSelectorControl_selectedDeviceChanged(this.h, (C.int)(index))
}
func (this *QVideoDeviceSelectorControl) OnSelectedDeviceChanged(slot func(index int)) {
	C.QVideoDeviceSelectorControl_connect_selectedDeviceChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoDeviceSelectorControl_selectedDeviceChanged
func miqt_exec_callback_QVideoDeviceSelectorControl_selectedDeviceChanged(cb C.intptr_t, index C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QVideoDeviceSelectorControl) SelectedDeviceChangedWithName(name string) {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	C.QVideoDeviceSelectorControl_selectedDeviceChangedWithName(this.h, name_ms)
}
func (this *QVideoDeviceSelectorControl) OnSelectedDeviceChangedWithName(slot func(name string)) {
	C.QVideoDeviceSelectorControl_connect_selectedDeviceChangedWithName(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoDeviceSelectorControl_selectedDeviceChangedWithName
func miqt_exec_callback_QVideoDeviceSelectorControl_selectedDeviceChangedWithName(cb C.intptr_t, name C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(name string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var name_ms C.struct_miqt_string = name
	name_ret := C.GoStringN(name_ms.data, C.int(int64(name_ms.len)))
	C.free(unsafe.Pointer(name_ms.data))
	slotval1 := name_ret

	gofunc(slotval1)
}

func (this *QVideoDeviceSelectorControl) DevicesChanged() {
	C.QVideoDeviceSelectorControl_devicesChanged(this.h)
}
func (this *QVideoDeviceSelectorControl) OnDevicesChanged(slot func()) {
	C.QVideoDeviceSelectorControl_connect_devicesChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoDeviceSelectorControl_devicesChanged
func miqt_exec_callback_QVideoDeviceSelectorControl_devicesChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QVideoDeviceSelectorControl_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoDeviceSelectorControl_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoDeviceSelectorControl_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoDeviceSelectorControl_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QVideoDeviceSelectorControl_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QVideoDeviceSelectorControl) Delete() {
	C.QVideoDeviceSelectorControl_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QVideoDeviceSelectorControl) GoGC() {
	runtime.SetFinalizer(this, func(this *QVideoDeviceSelectorControl) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
