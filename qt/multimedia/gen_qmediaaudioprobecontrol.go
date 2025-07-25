package multimedia

/*

#include "gen_qmediaaudioprobecontrol.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QMediaAudioProbeControl struct {
	h *C.QMediaAudioProbeControl
	*QMediaControl
}

func (this *QMediaAudioProbeControl) cPointer() *C.QMediaAudioProbeControl {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMediaAudioProbeControl) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMediaAudioProbeControl constructs the type using only CGO pointers.
func newQMediaAudioProbeControl(h *C.QMediaAudioProbeControl) *QMediaAudioProbeControl {
	if h == nil {
		return nil
	}
	var outptr_QMediaControl *C.QMediaControl = nil
	C.QMediaAudioProbeControl_virtbase(h, &outptr_QMediaControl)

	return &QMediaAudioProbeControl{h: h,
		QMediaControl: newQMediaControl(outptr_QMediaControl)}
}

// UnsafeNewQMediaAudioProbeControl constructs the type using only unsafe pointers.
func UnsafeNewQMediaAudioProbeControl(h unsafe.Pointer) *QMediaAudioProbeControl {
	return newQMediaAudioProbeControl((*C.QMediaAudioProbeControl)(h))
}

func (this *QMediaAudioProbeControl) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QMediaAudioProbeControl_metaObject(this.h)))
}

func (this *QMediaAudioProbeControl) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QMediaAudioProbeControl_metacast(this.h, param1_Cstring))
}

func QMediaAudioProbeControl_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMediaAudioProbeControl_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaAudioProbeControl_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMediaAudioProbeControl_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaAudioProbeControl) AudioBufferProbed(buffer *QAudioBuffer) {
	C.QMediaAudioProbeControl_audioBufferProbed(this.h, buffer.cPointer())
}
func (this *QMediaAudioProbeControl) OnAudioBufferProbed(slot func(buffer *QAudioBuffer)) {
	C.QMediaAudioProbeControl_connect_audioBufferProbed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaAudioProbeControl_audioBufferProbed
func miqt_exec_callback_QMediaAudioProbeControl_audioBufferProbed(cb C.intptr_t, buffer *C.QAudioBuffer) {
	gofunc, ok := cgo.Handle(cb).Value().(func(buffer *QAudioBuffer))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAudioBuffer(buffer)

	gofunc(slotval1)
}

func (this *QMediaAudioProbeControl) Flush() {
	C.QMediaAudioProbeControl_flush(this.h)
}
func (this *QMediaAudioProbeControl) OnFlush(slot func()) {
	C.QMediaAudioProbeControl_connect_flush(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaAudioProbeControl_flush
func miqt_exec_callback_QMediaAudioProbeControl_flush(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QMediaAudioProbeControl_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaAudioProbeControl_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaAudioProbeControl_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaAudioProbeControl_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaAudioProbeControl_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaAudioProbeControl_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaAudioProbeControl_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaAudioProbeControl_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QMediaAudioProbeControl) Delete() {
	C.QMediaAudioProbeControl_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMediaAudioProbeControl) GoGC() {
	runtime.SetFinalizer(this, func(this *QMediaAudioProbeControl) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
