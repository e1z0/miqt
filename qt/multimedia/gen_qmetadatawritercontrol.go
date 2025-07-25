package multimedia

/*

#include "gen_qmetadatawritercontrol.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QMetaDataWriterControl struct {
	h *C.QMetaDataWriterControl
	*QMediaControl
}

func (this *QMetaDataWriterControl) cPointer() *C.QMetaDataWriterControl {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMetaDataWriterControl) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMetaDataWriterControl constructs the type using only CGO pointers.
func newQMetaDataWriterControl(h *C.QMetaDataWriterControl) *QMetaDataWriterControl {
	if h == nil {
		return nil
	}
	var outptr_QMediaControl *C.QMediaControl = nil
	C.QMetaDataWriterControl_virtbase(h, &outptr_QMediaControl)

	return &QMetaDataWriterControl{h: h,
		QMediaControl: newQMediaControl(outptr_QMediaControl)}
}

// UnsafeNewQMetaDataWriterControl constructs the type using only unsafe pointers.
func UnsafeNewQMetaDataWriterControl(h unsafe.Pointer) *QMetaDataWriterControl {
	return newQMetaDataWriterControl((*C.QMetaDataWriterControl)(h))
}

func (this *QMetaDataWriterControl) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QMetaDataWriterControl_metaObject(this.h)))
}

func (this *QMetaDataWriterControl) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QMetaDataWriterControl_metacast(this.h, param1_Cstring))
}

func QMetaDataWriterControl_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataWriterControl_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataWriterControl_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataWriterControl_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMetaDataWriterControl) IsWritable() bool {
	return (bool)(C.QMetaDataWriterControl_isWritable(this.h))
}

func (this *QMetaDataWriterControl) IsMetaDataAvailable() bool {
	return (bool)(C.QMetaDataWriterControl_isMetaDataAvailable(this.h))
}

func (this *QMetaDataWriterControl) MetaData(key string) *qt.QVariant {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(C.QMetaDataWriterControl_metaData(this.h, key_ms)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaDataWriterControl) SetMetaData(key string, value *qt.QVariant) {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	C.QMetaDataWriterControl_setMetaData(this.h, key_ms, (*C.QVariant)(value.UnsafePointer()))
}

func (this *QMetaDataWriterControl) AvailableMetaData() []string {
	var _ma C.struct_miqt_array = C.QMetaDataWriterControl_availableMetaData(this.h)
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

func (this *QMetaDataWriterControl) MetaDataChanged() {
	C.QMetaDataWriterControl_metaDataChanged(this.h)
}
func (this *QMetaDataWriterControl) OnMetaDataChanged(slot func()) {
	C.QMetaDataWriterControl_connect_metaDataChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMetaDataWriterControl_metaDataChanged
func miqt_exec_callback_QMetaDataWriterControl_metaDataChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMetaDataWriterControl) MetaDataChanged2(key string, value *qt.QVariant) {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	C.QMetaDataWriterControl_metaDataChanged2(this.h, key_ms, (*C.QVariant)(value.UnsafePointer()))
}
func (this *QMetaDataWriterControl) OnMetaDataChanged2(slot func(key string, value *qt.QVariant)) {
	C.QMetaDataWriterControl_connect_metaDataChanged2(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMetaDataWriterControl_metaDataChanged2
func miqt_exec_callback_QMetaDataWriterControl_metaDataChanged2(cb C.intptr_t, key C.struct_miqt_string, value *C.QVariant) {
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

func (this *QMetaDataWriterControl) WritableChanged(writable bool) {
	C.QMetaDataWriterControl_writableChanged(this.h, (C.bool)(writable))
}
func (this *QMetaDataWriterControl) OnWritableChanged(slot func(writable bool)) {
	C.QMetaDataWriterControl_connect_writableChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMetaDataWriterControl_writableChanged
func miqt_exec_callback_QMetaDataWriterControl_writableChanged(cb C.intptr_t, writable C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(writable bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(writable)

	gofunc(slotval1)
}

func (this *QMetaDataWriterControl) MetaDataAvailableChanged(available bool) {
	C.QMetaDataWriterControl_metaDataAvailableChanged(this.h, (C.bool)(available))
}
func (this *QMetaDataWriterControl) OnMetaDataAvailableChanged(slot func(available bool)) {
	C.QMetaDataWriterControl_connect_metaDataAvailableChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMetaDataWriterControl_metaDataAvailableChanged
func miqt_exec_callback_QMetaDataWriterControl_metaDataAvailableChanged(cb C.intptr_t, available C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(available bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(available)

	gofunc(slotval1)
}

func QMetaDataWriterControl_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataWriterControl_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataWriterControl_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataWriterControl_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataWriterControl_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataWriterControl_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaDataWriterControl_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMetaDataWriterControl_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QMetaDataWriterControl) Delete() {
	C.QMetaDataWriterControl_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMetaDataWriterControl) GoGC() {
	runtime.SetFinalizer(this, func(this *QMetaDataWriterControl) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
