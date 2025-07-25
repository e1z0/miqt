package webchannel

/*

#include "gen_qwebchannel.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QWebChannel struct {
	h *C.QWebChannel
	*qt6.QObject
}

func (this *QWebChannel) cPointer() *C.QWebChannel {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebChannel) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebChannel constructs the type using only CGO pointers.
func newQWebChannel(h *C.QWebChannel) *QWebChannel {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QWebChannel_virtbase(h, &outptr_QObject)

	return &QWebChannel{h: h,
		QObject: qt6.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQWebChannel constructs the type using only unsafe pointers.
func UnsafeNewQWebChannel(h unsafe.Pointer) *QWebChannel {
	return newQWebChannel((*C.QWebChannel)(h))
}

// NewQWebChannel constructs a new QWebChannel object.
func NewQWebChannel() *QWebChannel {

	return newQWebChannel(C.QWebChannel_new())
}

// NewQWebChannel2 constructs a new QWebChannel object.
func NewQWebChannel2(parent *qt6.QObject) *QWebChannel {

	return newQWebChannel(C.QWebChannel_new2((*C.QObject)(parent.UnsafePointer())))
}

func (this *QWebChannel) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QWebChannel_metaObject(this.h)))
}

func (this *QWebChannel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QWebChannel_metacast(this.h, param1_Cstring))
}

func QWebChannel_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QWebChannel_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebChannel) RegisterObjects(objects map[string]*qt6.QObject) {
	objects_Keys_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(objects))))
	defer C.free(unsafe.Pointer(objects_Keys_CArray))
	objects_Values_CArray := (*[0xffff]*C.QObject)(C.malloc(C.size_t(8 * len(objects))))
	defer C.free(unsafe.Pointer(objects_Values_CArray))
	objects_ctr := 0
	for objects_k, objects_v := range objects {
		objects_k_ms := C.struct_miqt_string{}
		objects_k_ms.data = C.CString(objects_k)
		objects_k_ms.len = C.size_t(len(objects_k))
		defer C.free(unsafe.Pointer(objects_k_ms.data))
		objects_Keys_CArray[objects_ctr] = objects_k_ms
		objects_Values_CArray[objects_ctr] = (*C.QObject)(objects_v.UnsafePointer())
		objects_ctr++
	}
	objects_mm := C.struct_miqt_map{
		len:    C.size_t(len(objects)),
		keys:   unsafe.Pointer(objects_Keys_CArray),
		values: unsafe.Pointer(objects_Values_CArray),
	}
	C.QWebChannel_registerObjects(this.h, objects_mm)
}

func (this *QWebChannel) RegisteredObjects() map[string]*qt6.QObject {
	var _mm C.struct_miqt_map = C.QWebChannel_registeredObjects(this.h)
	_ret := make(map[string]*qt6.QObject, int(_mm.len))
	_Keys := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*C.QObject)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		var _hashkey_ms C.struct_miqt_string = _Keys[i]
		_hashkey_ret := C.GoStringN(_hashkey_ms.data, C.int(int64(_hashkey_ms.len)))
		C.free(unsafe.Pointer(_hashkey_ms.data))
		_entry_Key := _hashkey_ret
		_entry_Value := qt6.UnsafeNewQObject(unsafe.Pointer(_Values[i]))

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QWebChannel) RegisterObject(id string, object *qt6.QObject) {
	id_ms := C.struct_miqt_string{}
	id_ms.data = C.CString(id)
	id_ms.len = C.size_t(len(id))
	defer C.free(unsafe.Pointer(id_ms.data))
	C.QWebChannel_registerObject(this.h, id_ms, (*C.QObject)(object.UnsafePointer()))
}

func (this *QWebChannel) DeregisterObject(object *qt6.QObject) {
	C.QWebChannel_deregisterObject(this.h, (*C.QObject)(object.UnsafePointer()))
}

func (this *QWebChannel) BlockUpdates() bool {
	return (bool)(C.QWebChannel_blockUpdates(this.h))
}

func (this *QWebChannel) SetBlockUpdates(block bool) {
	C.QWebChannel_setBlockUpdates(this.h, (C.bool)(block))
}

func (this *QWebChannel) PropertyUpdateInterval() int {
	return (int)(C.QWebChannel_propertyUpdateInterval(this.h))
}

func (this *QWebChannel) SetPropertyUpdateInterval(ms int) {
	C.QWebChannel_setPropertyUpdateInterval(this.h, (C.int)(ms))
}

func (this *QWebChannel) BlockUpdatesChanged(block bool) {
	C.QWebChannel_blockUpdatesChanged(this.h, (C.bool)(block))
}
func (this *QWebChannel) OnBlockUpdatesChanged(slot func(block bool)) {
	C.QWebChannel_connect_blockUpdatesChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebChannel_blockUpdatesChanged
func miqt_exec_callback_QWebChannel_blockUpdatesChanged(cb C.intptr_t, block C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(block bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(block)

	gofunc(slotval1)
}

func (this *QWebChannel) ConnectTo(transport *QWebChannelAbstractTransport) {
	C.QWebChannel_connectTo(this.h, transport.cPointer())
}

func (this *QWebChannel) DisconnectFrom(transport *QWebChannelAbstractTransport) {
	C.QWebChannel_disconnectFrom(this.h, transport.cPointer())
}

func QWebChannel_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebChannel_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebChannel_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebChannel_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Sender can only be called from a QWebChannel that was directly constructed.
func (this *QWebChannel) Sender() *qt6.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt6.UnsafeNewQObject(unsafe.Pointer(C.QWebChannel_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QWebChannel that was directly constructed.
func (this *QWebChannel) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QWebChannel_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QWebChannel that was directly constructed.
func (this *QWebChannel) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QWebChannel_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QWebChannel that was directly constructed.
func (this *QWebChannel) IsSignalConnected(signal *qt6.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QWebChannel_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QWebChannel) callVirtualBase_Event(event *qt6.QEvent) bool {

	return (bool)(C.QWebChannel_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QWebChannel) OnEvent(slot func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool) {
	ok := C.QWebChannel_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebChannel_event
func miqt_exec_callback_QWebChannel_event(self *C.QWebChannel, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QWebChannel{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QWebChannel) callVirtualBase_EventFilter(watched *qt6.QObject, event *qt6.QEvent) bool {

	return (bool)(C.QWebChannel_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QWebChannel) OnEventFilter(slot func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool) {
	ok := C.QWebChannel_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebChannel_eventFilter
func miqt_exec_callback_QWebChannel_eventFilter(self *C.QWebChannel, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QWebChannel{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QWebChannel) callVirtualBase_TimerEvent(event *qt6.QTimerEvent) {

	C.QWebChannel_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QWebChannel) OnTimerEvent(slot func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent)) {
	ok := C.QWebChannel_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebChannel_timerEvent
func miqt_exec_callback_QWebChannel_timerEvent(self *C.QWebChannel, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QWebChannel{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QWebChannel) callVirtualBase_ChildEvent(event *qt6.QChildEvent) {

	C.QWebChannel_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QWebChannel) OnChildEvent(slot func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent)) {
	ok := C.QWebChannel_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebChannel_childEvent
func miqt_exec_callback_QWebChannel_childEvent(self *C.QWebChannel, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QWebChannel{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QWebChannel) callVirtualBase_CustomEvent(event *qt6.QEvent) {

	C.QWebChannel_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QWebChannel) OnCustomEvent(slot func(super func(event *qt6.QEvent), event *qt6.QEvent)) {
	ok := C.QWebChannel_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebChannel_customEvent
func miqt_exec_callback_QWebChannel_customEvent(self *C.QWebChannel, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent), event *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QWebChannel{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QWebChannel) callVirtualBase_ConnectNotify(signal *qt6.QMetaMethod) {

	C.QWebChannel_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebChannel) OnConnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QWebChannel_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebChannel_connectNotify
func miqt_exec_callback_QWebChannel_connectNotify(self *C.QWebChannel, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebChannel{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QWebChannel) callVirtualBase_DisconnectNotify(signal *qt6.QMetaMethod) {

	C.QWebChannel_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebChannel) OnDisconnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QWebChannel_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebChannel_disconnectNotify
func miqt_exec_callback_QWebChannel_disconnectNotify(self *C.QWebChannel, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebChannel{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QWebChannel) Delete() {
	C.QWebChannel_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebChannel) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebChannel) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
