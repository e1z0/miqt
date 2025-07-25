package network

/*

#include "gen_qnetworksession.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QNetworkSession__State int

const (
	QNetworkSession__Invalid      QNetworkSession__State = 0
	QNetworkSession__NotAvailable QNetworkSession__State = 1
	QNetworkSession__Connecting   QNetworkSession__State = 2
	QNetworkSession__Connected    QNetworkSession__State = 3
	QNetworkSession__Closing      QNetworkSession__State = 4
	QNetworkSession__Disconnected QNetworkSession__State = 5
	QNetworkSession__Roaming      QNetworkSession__State = 6
)

type QNetworkSession__SessionError int

const (
	QNetworkSession__UnknownSessionError        QNetworkSession__SessionError = 0
	QNetworkSession__SessionAbortedError        QNetworkSession__SessionError = 1
	QNetworkSession__RoamingError               QNetworkSession__SessionError = 2
	QNetworkSession__OperationNotSupportedError QNetworkSession__SessionError = 3
	QNetworkSession__InvalidConfigurationError  QNetworkSession__SessionError = 4
)

type QNetworkSession__UsagePolicy int

const (
	QNetworkSession__NoPolicy                  QNetworkSession__UsagePolicy = 0
	QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = 1
)

type QNetworkSession struct {
	h *C.QNetworkSession
	*qt.QObject
}

func (this *QNetworkSession) cPointer() *C.QNetworkSession {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QNetworkSession) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQNetworkSession constructs the type using only CGO pointers.
func newQNetworkSession(h *C.QNetworkSession) *QNetworkSession {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QNetworkSession_virtbase(h, &outptr_QObject)

	return &QNetworkSession{h: h,
		QObject: qt.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQNetworkSession constructs the type using only unsafe pointers.
func UnsafeNewQNetworkSession(h unsafe.Pointer) *QNetworkSession {
	return newQNetworkSession((*C.QNetworkSession)(h))
}

// NewQNetworkSession constructs a new QNetworkSession object.
func NewQNetworkSession(connConfig *QNetworkConfiguration) *QNetworkSession {

	return newQNetworkSession(C.QNetworkSession_new(connConfig.cPointer()))
}

// NewQNetworkSession2 constructs a new QNetworkSession object.
func NewQNetworkSession2(connConfig *QNetworkConfiguration, parent *qt.QObject) *QNetworkSession {

	return newQNetworkSession(C.QNetworkSession_new2(connConfig.cPointer(), (*C.QObject)(parent.UnsafePointer())))
}

func (this *QNetworkSession) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QNetworkSession_metaObject(this.h)))
}

func (this *QNetworkSession) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QNetworkSession_metacast(this.h, param1_Cstring))
}

func QNetworkSession_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkSession) IsOpen() bool {
	return (bool)(C.QNetworkSession_isOpen(this.h))
}

func (this *QNetworkSession) Configuration() *QNetworkConfiguration {
	_goptr := newQNetworkConfiguration(C.QNetworkSession_configuration(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkSession) Interface() *QNetworkInterface {
	_goptr := newQNetworkInterface(C.QNetworkSession_interface(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkSession) State() QNetworkSession__State {
	return (QNetworkSession__State)(C.QNetworkSession_state(this.h))
}

func (this *QNetworkSession) Error() QNetworkSession__SessionError {
	return (QNetworkSession__SessionError)(C.QNetworkSession_error(this.h))
}

func (this *QNetworkSession) ErrorString() string {
	var _ms C.struct_miqt_string = C.QNetworkSession_errorString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkSession) SessionProperty(key string) *qt.QVariant {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(C.QNetworkSession_sessionProperty(this.h, key_ms)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkSession) SetSessionProperty(key string, value *qt.QVariant) {
	key_ms := C.struct_miqt_string{}
	key_ms.data = C.CString(key)
	key_ms.len = C.size_t(len(key))
	defer C.free(unsafe.Pointer(key_ms.data))
	C.QNetworkSession_setSessionProperty(this.h, key_ms, (*C.QVariant)(value.UnsafePointer()))
}

func (this *QNetworkSession) BytesWritten() uint64 {
	return (uint64)(C.QNetworkSession_bytesWritten(this.h))
}

func (this *QNetworkSession) BytesReceived() uint64 {
	return (uint64)(C.QNetworkSession_bytesReceived(this.h))
}

func (this *QNetworkSession) ActiveTime() uint64 {
	return (uint64)(C.QNetworkSession_activeTime(this.h))
}

func (this *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	return (QNetworkSession__UsagePolicy)(C.QNetworkSession_usagePolicies(this.h))
}

func (this *QNetworkSession) WaitForOpened() bool {
	return (bool)(C.QNetworkSession_waitForOpened(this.h))
}

func (this *QNetworkSession) Open() {
	C.QNetworkSession_open(this.h)
}

func (this *QNetworkSession) Close() {
	C.QNetworkSession_close(this.h)
}

func (this *QNetworkSession) Stop() {
	C.QNetworkSession_stop(this.h)
}

func (this *QNetworkSession) Migrate() {
	C.QNetworkSession_migrate(this.h)
}

func (this *QNetworkSession) Ignore() {
	C.QNetworkSession_ignore(this.h)
}

func (this *QNetworkSession) Accept() {
	C.QNetworkSession_accept(this.h)
}

func (this *QNetworkSession) Reject() {
	C.QNetworkSession_reject(this.h)
}

func (this *QNetworkSession) StateChanged(param1 QNetworkSession__State) {
	C.QNetworkSession_stateChanged(this.h, (C.int)(param1))
}
func (this *QNetworkSession) OnStateChanged(slot func(param1 QNetworkSession__State)) {
	C.QNetworkSession_connect_stateChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_stateChanged
func miqt_exec_callback_QNetworkSession_stateChanged(cb C.intptr_t, param1 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QNetworkSession__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkSession__State)(param1)

	gofunc(slotval1)
}

func (this *QNetworkSession) Opened() {
	C.QNetworkSession_opened(this.h)
}
func (this *QNetworkSession) OnOpened(slot func()) {
	C.QNetworkSession_connect_opened(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_opened
func miqt_exec_callback_QNetworkSession_opened(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkSession) Closed() {
	C.QNetworkSession_closed(this.h)
}
func (this *QNetworkSession) OnClosed(slot func()) {
	C.QNetworkSession_connect_closed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_closed
func miqt_exec_callback_QNetworkSession_closed(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkSession) ErrorWithQNetworkSessionSessionError(param1 QNetworkSession__SessionError) {
	C.QNetworkSession_errorWithQNetworkSessionSessionError(this.h, (C.int)(param1))
}
func (this *QNetworkSession) OnErrorWithQNetworkSessionSessionError(slot func(param1 QNetworkSession__SessionError)) {
	C.QNetworkSession_connect_errorWithQNetworkSessionSessionError(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_errorWithQNetworkSessionSessionError
func miqt_exec_callback_QNetworkSession_errorWithQNetworkSessionSessionError(cb C.intptr_t, param1 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 QNetworkSession__SessionError))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkSession__SessionError)(param1)

	gofunc(slotval1)
}

func (this *QNetworkSession) PreferredConfigurationChanged(config *QNetworkConfiguration, isSeamless bool) {
	C.QNetworkSession_preferredConfigurationChanged(this.h, config.cPointer(), (C.bool)(isSeamless))
}
func (this *QNetworkSession) OnPreferredConfigurationChanged(slot func(config *QNetworkConfiguration, isSeamless bool)) {
	C.QNetworkSession_connect_preferredConfigurationChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_preferredConfigurationChanged
func miqt_exec_callback_QNetworkSession_preferredConfigurationChanged(cb C.intptr_t, config *C.QNetworkConfiguration, isSeamless C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(config *QNetworkConfiguration, isSeamless bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkConfiguration(config)

	slotval2 := (bool)(isSeamless)

	gofunc(slotval1, slotval2)
}

func (this *QNetworkSession) NewConfigurationActivated() {
	C.QNetworkSession_newConfigurationActivated(this.h)
}
func (this *QNetworkSession) OnNewConfigurationActivated(slot func()) {
	C.QNetworkSession_connect_newConfigurationActivated(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_newConfigurationActivated
func miqt_exec_callback_QNetworkSession_newConfigurationActivated(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {
	C.QNetworkSession_usagePoliciesChanged(this.h, (C.int)(usagePolicies))
}
func (this *QNetworkSession) OnUsagePoliciesChanged(slot func(usagePolicies QNetworkSession__UsagePolicy)) {
	C.QNetworkSession_connect_usagePoliciesChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkSession_usagePoliciesChanged
func miqt_exec_callback_QNetworkSession_usagePoliciesChanged(cb C.intptr_t, usagePolicies C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(usagePolicies QNetworkSession__UsagePolicy))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QNetworkSession__UsagePolicy)(usagePolicies)

	gofunc(slotval1)
}

func QNetworkSession_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkSession_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QNetworkSession_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkSession) WaitForOpenedWithMsecs(msecs int) bool {
	return (bool)(C.QNetworkSession_waitForOpenedWithMsecs(this.h, (C.int)(msecs)))
}

// Sender can only be called from a QNetworkSession that was directly constructed.
func (this *QNetworkSession) Sender() *qt.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt.UnsafeNewQObject(unsafe.Pointer(C.QNetworkSession_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QNetworkSession that was directly constructed.
func (this *QNetworkSession) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QNetworkSession_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QNetworkSession that was directly constructed.
func (this *QNetworkSession) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QNetworkSession_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QNetworkSession that was directly constructed.
func (this *QNetworkSession) IsSignalConnected(signal *qt.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QNetworkSession_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QNetworkSession) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	C.QNetworkSession_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QNetworkSession) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QNetworkSession_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QNetworkSession_connectNotify
func miqt_exec_callback_QNetworkSession_connectNotify(self *C.QNetworkSession, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QNetworkSession{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QNetworkSession) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	C.QNetworkSession_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QNetworkSession) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QNetworkSession_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QNetworkSession_disconnectNotify
func miqt_exec_callback_QNetworkSession_disconnectNotify(self *C.QNetworkSession, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QNetworkSession{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

func (this *QNetworkSession) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(C.QNetworkSession_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QNetworkSession) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	ok := C.QNetworkSession_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QNetworkSession_event
func miqt_exec_callback_QNetworkSession_event(self *C.QNetworkSession, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QNetworkSession{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QNetworkSession) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(C.QNetworkSession_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QNetworkSession) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	ok := C.QNetworkSession_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QNetworkSession_eventFilter
func miqt_exec_callback_QNetworkSession_eventFilter(self *C.QNetworkSession, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QNetworkSession{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QNetworkSession) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	C.QNetworkSession_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QNetworkSession) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	ok := C.QNetworkSession_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QNetworkSession_timerEvent
func miqt_exec_callback_QNetworkSession_timerEvent(self *C.QNetworkSession, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QNetworkSession{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QNetworkSession) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	C.QNetworkSession_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QNetworkSession) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	ok := C.QNetworkSession_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QNetworkSession_childEvent
func miqt_exec_callback_QNetworkSession_childEvent(self *C.QNetworkSession, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QNetworkSession{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QNetworkSession) callVirtualBase_CustomEvent(event *qt.QEvent) {

	C.QNetworkSession_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QNetworkSession) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	ok := C.QNetworkSession_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QNetworkSession_customEvent
func miqt_exec_callback_QNetworkSession_customEvent(self *C.QNetworkSession, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QNetworkSession{h: self}).callVirtualBase_CustomEvent, slotval1)

}

// Delete this object from C++ memory.
func (this *QNetworkSession) Delete() {
	C.QNetworkSession_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QNetworkSession) GoGC() {
	runtime.SetFinalizer(this, func(this *QNetworkSession) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
