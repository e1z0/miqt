package multimedia

/*

#include "gen_qmediagaplessplaybackcontrol.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QMediaGaplessPlaybackControl struct {
	h *C.QMediaGaplessPlaybackControl
	*QMediaControl
}

func (this *QMediaGaplessPlaybackControl) cPointer() *C.QMediaGaplessPlaybackControl {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMediaGaplessPlaybackControl) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMediaGaplessPlaybackControl constructs the type using only CGO pointers.
func newQMediaGaplessPlaybackControl(h *C.QMediaGaplessPlaybackControl) *QMediaGaplessPlaybackControl {
	if h == nil {
		return nil
	}
	var outptr_QMediaControl *C.QMediaControl = nil
	C.QMediaGaplessPlaybackControl_virtbase(h, &outptr_QMediaControl)

	return &QMediaGaplessPlaybackControl{h: h,
		QMediaControl: newQMediaControl(outptr_QMediaControl)}
}

// UnsafeNewQMediaGaplessPlaybackControl constructs the type using only unsafe pointers.
func UnsafeNewQMediaGaplessPlaybackControl(h unsafe.Pointer) *QMediaGaplessPlaybackControl {
	return newQMediaGaplessPlaybackControl((*C.QMediaGaplessPlaybackControl)(h))
}

func (this *QMediaGaplessPlaybackControl) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QMediaGaplessPlaybackControl_metaObject(this.h)))
}

func (this *QMediaGaplessPlaybackControl) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QMediaGaplessPlaybackControl_metacast(this.h, param1_Cstring))
}

func QMediaGaplessPlaybackControl_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMediaGaplessPlaybackControl_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaGaplessPlaybackControl_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMediaGaplessPlaybackControl_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMediaGaplessPlaybackControl) NextMedia() *QMediaContent {
	_goptr := newQMediaContent(C.QMediaGaplessPlaybackControl_nextMedia(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMediaGaplessPlaybackControl) SetNextMedia(media *QMediaContent) {
	C.QMediaGaplessPlaybackControl_setNextMedia(this.h, media.cPointer())
}

func (this *QMediaGaplessPlaybackControl) IsCrossfadeSupported() bool {
	return (bool)(C.QMediaGaplessPlaybackControl_isCrossfadeSupported(this.h))
}

func (this *QMediaGaplessPlaybackControl) CrossfadeTime() float64 {
	return (float64)(C.QMediaGaplessPlaybackControl_crossfadeTime(this.h))
}

func (this *QMediaGaplessPlaybackControl) SetCrossfadeTime(crossfadeTime float64) {
	C.QMediaGaplessPlaybackControl_setCrossfadeTime(this.h, (C.double)(crossfadeTime))
}

func (this *QMediaGaplessPlaybackControl) CrossfadeTimeChanged(crossfadeTime float64) {
	C.QMediaGaplessPlaybackControl_crossfadeTimeChanged(this.h, (C.double)(crossfadeTime))
}
func (this *QMediaGaplessPlaybackControl) OnCrossfadeTimeChanged(slot func(crossfadeTime float64)) {
	C.QMediaGaplessPlaybackControl_connect_crossfadeTimeChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaGaplessPlaybackControl_crossfadeTimeChanged
func miqt_exec_callback_QMediaGaplessPlaybackControl_crossfadeTimeChanged(cb C.intptr_t, crossfadeTime C.double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(crossfadeTime float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(crossfadeTime)

	gofunc(slotval1)
}

func (this *QMediaGaplessPlaybackControl) NextMediaChanged(media *QMediaContent) {
	C.QMediaGaplessPlaybackControl_nextMediaChanged(this.h, media.cPointer())
}
func (this *QMediaGaplessPlaybackControl) OnNextMediaChanged(slot func(media *QMediaContent)) {
	C.QMediaGaplessPlaybackControl_connect_nextMediaChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaGaplessPlaybackControl_nextMediaChanged
func miqt_exec_callback_QMediaGaplessPlaybackControl_nextMediaChanged(cb C.intptr_t, media *C.QMediaContent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(media *QMediaContent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMediaContent(media)

	gofunc(slotval1)
}

func (this *QMediaGaplessPlaybackControl) AdvancedToNextMedia() {
	C.QMediaGaplessPlaybackControl_advancedToNextMedia(this.h)
}
func (this *QMediaGaplessPlaybackControl) OnAdvancedToNextMedia(slot func()) {
	C.QMediaGaplessPlaybackControl_connect_advancedToNextMedia(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMediaGaplessPlaybackControl_advancedToNextMedia
func miqt_exec_callback_QMediaGaplessPlaybackControl_advancedToNextMedia(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QMediaGaplessPlaybackControl_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaGaplessPlaybackControl_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaGaplessPlaybackControl_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaGaplessPlaybackControl_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaGaplessPlaybackControl_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaGaplessPlaybackControl_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMediaGaplessPlaybackControl_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMediaGaplessPlaybackControl_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QMediaGaplessPlaybackControl) Delete() {
	C.QMediaGaplessPlaybackControl_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMediaGaplessPlaybackControl) GoGC() {
	runtime.SetFinalizer(this, func(this *QMediaGaplessPlaybackControl) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
