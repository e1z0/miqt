package multimedia

/*

#include "gen_qmetadatareadercontrol.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QMetaDataReaderControl struct {
	h *C.QMetaDataReaderControl
	*QMediaControl
}

func (this *QMetaDataReaderControl) cPointer() *C.QMetaDataReaderControl {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMetaDataReaderControl) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMetaDataReaderControl constructs the type using only CGO pointers.
func newQMetaDataReaderControl(h *C.QMetaDataReaderControl) *QMetaDataReaderControl {
	if h == nil {
		return nil
	}
	var outptr_QMediaControl *C.QMediaControl = nil
	C.QMetaDataReaderControl_virtbase(h, &outptr_QMediaControl)

	return &QMetaDataReaderControl{h: h,
		QMediaControl: newQMediaControl(outptr_QMediaControl)}
}

// UnsafeNewQMetaDataReaderControl constructs the type using only unsafe pointers.
func UnsafeNewQMetaDataReaderControl(h unsafe.Pointer) *QMetaDataReaderControl {
	return newQMetaDataReaderControl((*C.QMetaDataReaderControl)(h))
}

func (this *QMetaDataReaderControl) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QMetaDataReaderControl_metaObject(this.h)))
}

func (this *QMetaDataReaderControl) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QMetaDataReaderControl_metacast(this.h, param1_Cstring))
}

func QMetaDataReaderControl_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataReaderControl_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataReaderControl_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataReaderControl_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMetaDataReaderControl) IsMetaDataAvailable() bool {
	return (bool)(C.QMetaDataReaderControl_isMetaDataAvailable(this.h))
}

func (this *QMetaDataReaderControl) MetaData(key string) *qt.QVariant {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(C.QMetaDataReaderControl_metaData(this.h, key_ms)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaDataReaderControl) AvailableMetaData() []string {
	var _ma C.struct_miqt_array = C.QMetaDataReaderControl_availableMetaData(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QMetaDataReaderControl) MetaDataChanged() {
	C.QMetaDataReaderControl_metaDataChanged(this.h)
}
func (this *QMetaDataReaderControl) OnMetaDataChanged(slot func()) {
	C.QMetaDataReaderControl_connect_metaDataChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMetaDataReaderControl_metaDataChanged
func miqt_exec_callback_QMetaDataReaderControl_metaDataChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMetaDataReaderControl) MetaDataChanged2(key string, value *qt.QVariant) {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	C.QMetaDataReaderControl_metaDataChanged2(this.h, key_ms, (*C.QVariant)(value.UnsafePointer()))
}
func (this *QMetaDataReaderControl) OnMetaDataChanged2(slot func(key string, value *qt.QVariant)) {
	C.QMetaDataReaderControl_connect_metaDataChanged2(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMetaDataReaderControl_metaDataChanged2
func miqt_exec_callback_QMetaDataReaderControl_metaDataChanged2(cb C.intptr_t, key C.struct_miqt_string, value *C.QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(key string, value *qt.QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var key_ms C.struct_miqt_string = key
	key_ret := C.GoStringN(key_ms.data, C.int(int64(key_ms.len)))
	C.free(unsafe.Pointer(key_ms.data))
	slotval1 := key_ret
	slotval2 := qt.UnsafeNewQVariant(unsafe.Pointer(value))

	gofunc(slotval1, slotval2)
}

func (this *QMetaDataReaderControl) MetaDataAvailableChanged(available bool) {
	C.QMetaDataReaderControl_metaDataAvailableChanged(this.h, (C.bool)(available))
}
func (this *QMetaDataReaderControl) OnMetaDataAvailableChanged(slot func(available bool)) {
	C.QMetaDataReaderControl_connect_metaDataAvailableChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMetaDataReaderControl_metaDataAvailableChanged
func miqt_exec_callback_QMetaDataReaderControl_metaDataAvailableChanged(cb C.intptr_t, available C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(available bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(available)

	gofunc(slotval1)
}

func QMetaDataReaderControl_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataReaderControl_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataReaderControl_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataReaderControl_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataReaderControl_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataReaderControl_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataReaderControl_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataReaderControl_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QMetaDataReaderControl) Delete() {
	C.QMetaDataReaderControl_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMetaDataReaderControl) GoGC() {
	runtime.SetFinalizer(this, func(this *QMetaDataReaderControl) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
