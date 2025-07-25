package multimedia

/*

#include "gen_qvideosink.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QVideoSink struct {
	h *C.QVideoSink
	*qt6.QObject
}

func (this *QVideoSink) cPointer() *C.QVideoSink {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QVideoSink) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQVideoSink constructs the type using only CGO pointers.
func newQVideoSink(h *C.QVideoSink) *QVideoSink {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QVideoSink_virtbase(h, &outptr_QObject)

	return &QVideoSink{h: h,
		QObject: qt6.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQVideoSink constructs the type using only unsafe pointers.
func UnsafeNewQVideoSink(h unsafe.Pointer) *QVideoSink {
	return newQVideoSink((*C.QVideoSink)(h))
}

// NewQVideoSink constructs a new QVideoSink object.
func NewQVideoSink() *QVideoSink {

	return newQVideoSink(C.QVideoSink_new())
}

// NewQVideoSink2 constructs a new QVideoSink object.
func NewQVideoSink2(parent *qt6.QObject) *QVideoSink {

	return newQVideoSink(C.QVideoSink_new2((*C.QObject)(parent.UnsafePointer())))
}

func (this *QVideoSink) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QVideoSink_metaObject(this.h)))
}

func (this *QVideoSink) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QVideoSink_metacast(this.h, param1_Cstring))
}

func QVideoSink_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QVideoSink_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoSink) VideoSize() *qt6.QSize {
	_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(C.QVideoSink_videoSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoSink) SubtitleText() string {
	var _ms C.struct_miqt_string = C.QVideoSink_subtitleText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoSink) SetSubtitleText(subtitle string) {
	subtitle_ms := C.struct_miqt_string{}
	subtitle_ms.data = C.CString(subtitle)
	subtitle_ms.len = C.size_t(len(subtitle))
	defer C.free(unsafe.Pointer(subtitle_ms.data))
	C.QVideoSink_setSubtitleText(this.h, subtitle_ms)
}

func (this *QVideoSink) SetVideoFrame(frame *QVideoFrame) {
	C.QVideoSink_setVideoFrame(this.h, frame.cPointer())
}

func (this *QVideoSink) VideoFrame() *QVideoFrame {
	_goptr := newQVideoFrame(C.QVideoSink_videoFrame(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoSink) VideoFrameChanged(frame *QVideoFrame) {
	C.QVideoSink_videoFrameChanged(this.h, frame.cPointer())
}
func (this *QVideoSink) OnVideoFrameChanged(slot func(frame *QVideoFrame)) {
	C.QVideoSink_connect_videoFrameChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_videoFrameChanged
func miqt_exec_callback_QVideoSink_videoFrameChanged(cb C.intptr_t, frame *C.QVideoFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(frame *QVideoFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVideoFrame(frame)

	gofunc(slotval1)
}

func (this *QVideoSink) SubtitleTextChanged(subtitleText string) {
	subtitleText_ms := C.struct_miqt_string{}
	subtitleText_ms.data = C.CString(subtitleText)
	subtitleText_ms.len = C.size_t(len(subtitleText))
	defer C.free(unsafe.Pointer(subtitleText_ms.data))
	C.QVideoSink_subtitleTextChanged(this.h, subtitleText_ms)
}
func (this *QVideoSink) OnSubtitleTextChanged(slot func(subtitleText string)) {
	C.QVideoSink_connect_subtitleTextChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_subtitleTextChanged
func miqt_exec_callback_QVideoSink_subtitleTextChanged(cb C.intptr_t, subtitleText C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(subtitleText string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var subtitleText_ms C.struct_miqt_string = subtitleText
	subtitleText_ret := C.GoStringN(subtitleText_ms.data, C.int(int64(subtitleText_ms.len)))
	C.free(unsafe.Pointer(subtitleText_ms.data))
	slotval1 := subtitleText_ret

	gofunc(slotval1)
}

func (this *QVideoSink) VideoSizeChanged() {
	C.QVideoSink_videoSizeChanged(this.h)
}
func (this *QVideoSink) OnVideoSizeChanged(slot func()) {
	C.QVideoSink_connect_videoSizeChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_videoSizeChanged
func miqt_exec_callback_QVideoSink_videoSizeChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QVideoSink_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QVideoSink_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoSink_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QVideoSink_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Sender can only be called from a QVideoSink that was directly constructed.
func (this *QVideoSink) Sender() *qt6.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt6.UnsafeNewQObject(unsafe.Pointer(C.QVideoSink_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QVideoSink that was directly constructed.
func (this *QVideoSink) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QVideoSink_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QVideoSink that was directly constructed.
func (this *QVideoSink) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QVideoSink_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QVideoSink that was directly constructed.
func (this *QVideoSink) IsSignalConnected(signal *qt6.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QVideoSink_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QVideoSink) callVirtualBase_Event(event *qt6.QEvent) bool {

	return (bool)(C.QVideoSink_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QVideoSink) OnEvent(slot func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool) {
	ok := C.QVideoSink_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QVideoSink_event
func miqt_exec_callback_QVideoSink_event(self *C.QVideoSink, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QVideoSink{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QVideoSink) callVirtualBase_EventFilter(watched *qt6.QObject, event *qt6.QEvent) bool {

	return (bool)(C.QVideoSink_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QVideoSink) OnEventFilter(slot func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool) {
	ok := C.QVideoSink_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QVideoSink_eventFilter
func miqt_exec_callback_QVideoSink_eventFilter(self *C.QVideoSink, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QVideoSink{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QVideoSink) callVirtualBase_TimerEvent(event *qt6.QTimerEvent) {

	C.QVideoSink_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QVideoSink) OnTimerEvent(slot func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent)) {
	ok := C.QVideoSink_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QVideoSink_timerEvent
func miqt_exec_callback_QVideoSink_timerEvent(self *C.QVideoSink, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QVideoSink{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QVideoSink) callVirtualBase_ChildEvent(event *qt6.QChildEvent) {

	C.QVideoSink_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QVideoSink) OnChildEvent(slot func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent)) {
	ok := C.QVideoSink_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QVideoSink_childEvent
func miqt_exec_callback_QVideoSink_childEvent(self *C.QVideoSink, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QVideoSink{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QVideoSink) callVirtualBase_CustomEvent(event *qt6.QEvent) {

	C.QVideoSink_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QVideoSink) OnCustomEvent(slot func(super func(event *qt6.QEvent), event *qt6.QEvent)) {
	ok := C.QVideoSink_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QVideoSink_customEvent
func miqt_exec_callback_QVideoSink_customEvent(self *C.QVideoSink, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent), event *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QVideoSink{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QVideoSink) callVirtualBase_ConnectNotify(signal *qt6.QMetaMethod) {

	C.QVideoSink_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QVideoSink) OnConnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QVideoSink_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QVideoSink_connectNotify
func miqt_exec_callback_QVideoSink_connectNotify(self *C.QVideoSink, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QVideoSink{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QVideoSink) callVirtualBase_DisconnectNotify(signal *qt6.QMetaMethod) {

	C.QVideoSink_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QVideoSink) OnDisconnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QVideoSink_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QVideoSink_disconnectNotify
func miqt_exec_callback_QVideoSink_disconnectNotify(self *C.QVideoSink, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QVideoSink{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QVideoSink) Delete() {
	C.QVideoSink_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QVideoSink) GoGC() {
	runtime.SetFinalizer(this, func(this *QVideoSink) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
