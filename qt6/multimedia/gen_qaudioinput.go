package multimedia

/*

#include "gen_qaudioinput.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QAudioInput struct {
	h *C.QAudioInput
	*qt6.QObject
}

func (this *QAudioInput) cPointer() *C.QAudioInput {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QAudioInput) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQAudioInput constructs the type using only CGO pointers.
func newQAudioInput(h *C.QAudioInput) *QAudioInput {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QAudioInput_virtbase(h, &outptr_QObject)

	return &QAudioInput{h: h,
		QObject: qt6.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQAudioInput constructs the type using only unsafe pointers.
func UnsafeNewQAudioInput(h unsafe.Pointer) *QAudioInput {
	return newQAudioInput((*C.QAudioInput)(h))
}

// NewQAudioInput constructs a new QAudioInput object.
func NewQAudioInput() *QAudioInput {

	return newQAudioInput(C.QAudioInput_new())
}

// NewQAudioInput2 constructs a new QAudioInput object.
func NewQAudioInput2(deviceInfo *QAudioDevice) *QAudioInput {

	return newQAudioInput(C.QAudioInput_new2(deviceInfo.cPointer()))
}

// NewQAudioInput3 constructs a new QAudioInput object.
func NewQAudioInput3(parent *qt6.QObject) *QAudioInput {

	return newQAudioInput(C.QAudioInput_new3((*C.QObject)(parent.UnsafePointer())))
}

// NewQAudioInput4 constructs a new QAudioInput object.
func NewQAudioInput4(deviceInfo *QAudioDevice, parent *qt6.QObject) *QAudioInput {

	return newQAudioInput(C.QAudioInput_new4(deviceInfo.cPointer(), (*C.QObject)(parent.UnsafePointer())))
}

func (this *QAudioInput) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QAudioInput_metaObject(this.h)))
}

func (this *QAudioInput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QAudioInput_metacast(this.h, param1_Cstring))
}

func QAudioInput_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAudioInput_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAudioInput) Device() *QAudioDevice {
	_goptr := newQAudioDevice(C.QAudioInput_device(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioInput) Volume() float32 {
	return (float32)(C.QAudioInput_volume(this.h))
}

func (this *QAudioInput) IsMuted() bool {
	return (bool)(C.QAudioInput_isMuted(this.h))
}

func (this *QAudioInput) SetDevice(device *QAudioDevice) {
	C.QAudioInput_setDevice(this.h, device.cPointer())
}

func (this *QAudioInput) SetVolume(volume float32) {
	C.QAudioInput_setVolume(this.h, (C.float)(volume))
}

func (this *QAudioInput) SetMuted(muted bool) {
	C.QAudioInput_setMuted(this.h, (C.bool)(muted))
}

func (this *QAudioInput) DeviceChanged() {
	C.QAudioInput_deviceChanged(this.h)
}
func (this *QAudioInput) OnDeviceChanged(slot func()) {
	C.QAudioInput_connect_deviceChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_deviceChanged
func miqt_exec_callback_QAudioInput_deviceChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAudioInput) VolumeChanged(volume float32) {
	C.QAudioInput_volumeChanged(this.h, (C.float)(volume))
}
func (this *QAudioInput) OnVolumeChanged(slot func(volume float32)) {
	C.QAudioInput_connect_volumeChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_volumeChanged
func miqt_exec_callback_QAudioInput_volumeChanged(cb C.intptr_t, volume C.float) {
	gofunc, ok := cgo.Handle(cb).Value().(func(volume float32))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float32)(volume)

	gofunc(slotval1)
}

func (this *QAudioInput) MutedChanged(muted bool) {
	C.QAudioInput_mutedChanged(this.h, (C.bool)(muted))
}
func (this *QAudioInput) OnMutedChanged(slot func(muted bool)) {
	C.QAudioInput_connect_mutedChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioInput_mutedChanged
func miqt_exec_callback_QAudioInput_mutedChanged(cb C.intptr_t, muted C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(muted bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(muted)

	gofunc(slotval1)
}

func QAudioInput_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAudioInput_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAudioInput_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAudioInput_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Sender can only be called from a QAudioInput that was directly constructed.
func (this *QAudioInput) Sender() *qt6.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt6.UnsafeNewQObject(unsafe.Pointer(C.QAudioInput_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QAudioInput that was directly constructed.
func (this *QAudioInput) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QAudioInput_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QAudioInput that was directly constructed.
func (this *QAudioInput) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QAudioInput_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QAudioInput that was directly constructed.
func (this *QAudioInput) IsSignalConnected(signal *qt6.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QAudioInput_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QAudioInput) callVirtualBase_Event(event *qt6.QEvent) bool {

	return (bool)(C.QAudioInput_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QAudioInput) OnEvent(slot func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool) {
	ok := C.QAudioInput_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QAudioInput_event
func miqt_exec_callback_QAudioInput_event(self *C.QAudioInput, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QAudioInput{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QAudioInput) callVirtualBase_EventFilter(watched *qt6.QObject, event *qt6.QEvent) bool {

	return (bool)(C.QAudioInput_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QAudioInput) OnEventFilter(slot func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool) {
	ok := C.QAudioInput_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QAudioInput_eventFilter
func miqt_exec_callback_QAudioInput_eventFilter(self *C.QAudioInput, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QAudioInput{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QAudioInput) callVirtualBase_TimerEvent(event *qt6.QTimerEvent) {

	C.QAudioInput_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QAudioInput) OnTimerEvent(slot func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent)) {
	ok := C.QAudioInput_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QAudioInput_timerEvent
func miqt_exec_callback_QAudioInput_timerEvent(self *C.QAudioInput, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QAudioInput{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAudioInput) callVirtualBase_ChildEvent(event *qt6.QChildEvent) {

	C.QAudioInput_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QAudioInput) OnChildEvent(slot func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent)) {
	ok := C.QAudioInput_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QAudioInput_childEvent
func miqt_exec_callback_QAudioInput_childEvent(self *C.QAudioInput, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QAudioInput{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QAudioInput) callVirtualBase_CustomEvent(event *qt6.QEvent) {

	C.QAudioInput_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QAudioInput) OnCustomEvent(slot func(super func(event *qt6.QEvent), event *qt6.QEvent)) {
	ok := C.QAudioInput_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QAudioInput_customEvent
func miqt_exec_callback_QAudioInput_customEvent(self *C.QAudioInput, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent), event *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QAudioInput{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QAudioInput) callVirtualBase_ConnectNotify(signal *qt6.QMetaMethod) {

	C.QAudioInput_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QAudioInput) OnConnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QAudioInput_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QAudioInput_connectNotify
func miqt_exec_callback_QAudioInput_connectNotify(self *C.QAudioInput, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QAudioInput{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QAudioInput) callVirtualBase_DisconnectNotify(signal *qt6.QMetaMethod) {

	C.QAudioInput_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QAudioInput) OnDisconnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QAudioInput_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QAudioInput_disconnectNotify
func miqt_exec_callback_QAudioInput_disconnectNotify(self *C.QAudioInput, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QAudioInput{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QAudioInput) Delete() {
	C.QAudioInput_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAudioInput) GoGC() {
	runtime.SetFinalizer(this, func(this *QAudioInput) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
