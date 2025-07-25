package qscintilla

/*

#include "gen_qsciabstractapis.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QsciAbstractAPIs struct {
	h *C.QsciAbstractAPIs
	*qt.QObject
}

func (this *QsciAbstractAPIs) cPointer() *C.QsciAbstractAPIs {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QsciAbstractAPIs) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQsciAbstractAPIs constructs the type using only CGO pointers.
func newQsciAbstractAPIs(h *C.QsciAbstractAPIs) *QsciAbstractAPIs {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QsciAbstractAPIs_virtbase(h, &outptr_QObject)

	return &QsciAbstractAPIs{h: h,
		QObject: qt.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQsciAbstractAPIs constructs the type using only unsafe pointers.
func UnsafeNewQsciAbstractAPIs(h unsafe.Pointer) *QsciAbstractAPIs {
	return newQsciAbstractAPIs((*C.QsciAbstractAPIs)(h))
}

// NewQsciAbstractAPIs constructs a new QsciAbstractAPIs object.
func NewQsciAbstractAPIs(lexer *QsciLexer) *QsciAbstractAPIs {

	return newQsciAbstractAPIs(C.QsciAbstractAPIs_new(lexer.cPointer()))
}

func (this *QsciAbstractAPIs) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QsciAbstractAPIs_metaObject(this.h)))
}

func (this *QsciAbstractAPIs) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QsciAbstractAPIs_metacast(this.h, param1_Cstring))
}

func QsciAbstractAPIs_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciAbstractAPIs_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciAbstractAPIs_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciAbstractAPIs_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciAbstractAPIs) Lexer() *QsciLexer {
	return newQsciLexer(C.QsciAbstractAPIs_lexer(this.h))
}

func (this *QsciAbstractAPIs) UpdateAutoCompletionList(context []string, list []string) {
	context_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(context))))
	defer C.free(unsafe.Pointer(context_CArray))
	for i := range context {
		context_i_ms := C.struct_miqt_string{}
		context_i_ms.data = C.CString(context[i])
		context_i_ms.len = C.size_t(len(context[i]))
		defer C.free(unsafe.Pointer(context_i_ms.data))
		context_CArray[i] = context_i_ms
	}
	context_ma := C.struct_miqt_array{len: C.size_t(len(context)), data: unsafe.Pointer(context_CArray)}
	list_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(list))))
	defer C.free(unsafe.Pointer(list_CArray))
	for i := range list {
		list_i_ms := C.struct_miqt_string{}
		list_i_ms.data = C.CString(list[i])
		list_i_ms.len = C.size_t(len(list[i]))
		defer C.free(unsafe.Pointer(list_i_ms.data))
		list_CArray[i] = list_i_ms
	}
	list_ma := C.struct_miqt_array{len: C.size_t(len(list)), data: unsafe.Pointer(list_CArray)}
	C.QsciAbstractAPIs_updateAutoCompletionList(this.h, context_ma, list_ma)
}

func (this *QsciAbstractAPIs) AutoCompletionSelected(selection string) {
	selection_ms := C.struct_miqt_string{}
	selection_ms.data = C.CString(selection)
	selection_ms.len = C.size_t(len(selection))
	defer C.free(unsafe.Pointer(selection_ms.data))
	C.QsciAbstractAPIs_autoCompletionSelected(this.h, selection_ms)
}

func (this *QsciAbstractAPIs) CallTips(context []string, commas int, style QsciScintilla__CallTipsStyle, shifts []int) []string {
	context_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(context))))
	defer C.free(unsafe.Pointer(context_CArray))
	for i := range context {
		context_i_ms := C.struct_miqt_string{}
		context_i_ms.data = C.CString(context[i])
		context_i_ms.len = C.size_t(len(context[i]))
		defer C.free(unsafe.Pointer(context_i_ms.data))
		context_CArray[i] = context_i_ms
	}
	context_ma := C.struct_miqt_array{len: C.size_t(len(context)), data: unsafe.Pointer(context_CArray)}
	shifts_CArray := (*[0xffff]C.int)(C.malloc(C.size_t(8 * len(shifts))))
	defer C.free(unsafe.Pointer(shifts_CArray))
	for i := range shifts {
		shifts_CArray[i] = (C.int)(shifts[i])
	}
	shifts_ma := C.struct_miqt_array{len: C.size_t(len(shifts)), data: unsafe.Pointer(shifts_CArray)}
	var _ma C.struct_miqt_array = C.QsciAbstractAPIs_callTips(this.h, context_ma, (C.int)(commas), (C.int)(style), shifts_ma)
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

func QsciAbstractAPIs_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciAbstractAPIs_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciAbstractAPIs_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciAbstractAPIs_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciAbstractAPIs_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciAbstractAPIs_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciAbstractAPIs_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciAbstractAPIs_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Sender can only be called from a QsciAbstractAPIs that was directly constructed.
func (this *QsciAbstractAPIs) Sender() *qt.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt.UnsafeNewQObject(unsafe.Pointer(C.QsciAbstractAPIs_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QsciAbstractAPIs that was directly constructed.
func (this *QsciAbstractAPIs) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QsciAbstractAPIs_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QsciAbstractAPIs that was directly constructed.
func (this *QsciAbstractAPIs) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QsciAbstractAPIs_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QsciAbstractAPIs that was directly constructed.
func (this *QsciAbstractAPIs) IsSignalConnected(signal *qt.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QsciAbstractAPIs_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}
func (this *QsciAbstractAPIs) OnUpdateAutoCompletionList(slot func(context []string, list []string)) {
	ok := C.QsciAbstractAPIs_override_virtual_updateAutoCompletionList(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_updateAutoCompletionList
func miqt_exec_callback_QsciAbstractAPIs_updateAutoCompletionList(self *C.QsciAbstractAPIs, cb C.intptr_t, context C.struct_miqt_array, list C.struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(context []string, list []string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var context_ma C.struct_miqt_array = context
	context_ret := make([]string, int(context_ma.len))
	context_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(context_ma.data)) // hey ya
	for i := 0; i < int(context_ma.len); i++ {
		var context_lv_ms C.struct_miqt_string = context_outCast[i]
		context_lv_ret := C.GoStringN(context_lv_ms.data, C.int(int64(context_lv_ms.len)))
		C.free(unsafe.Pointer(context_lv_ms.data))
		context_ret[i] = context_lv_ret
	}
	slotval1 := context_ret

	var list_ma C.struct_miqt_array = list
	list_ret := make([]string, int(list_ma.len))
	list_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(list_ma.data)) // hey ya
	for i := 0; i < int(list_ma.len); i++ {
		var list_lv_ms C.struct_miqt_string = list_outCast[i]
		list_lv_ret := C.GoStringN(list_lv_ms.data, C.int(int64(list_lv_ms.len)))
		C.free(unsafe.Pointer(list_lv_ms.data))
		list_ret[i] = list_lv_ret
	}
	slotval2 := list_ret

	gofunc(slotval1, slotval2)

}

func (this *QsciAbstractAPIs) callVirtualBase_AutoCompletionSelected(selection string) {
	selection_ms := C.struct_miqt_string{}
	selection_ms.data = C.CString(selection)
	selection_ms.len = C.size_t(len(selection))
	defer C.free(unsafe.Pointer(selection_ms.data))

	C.QsciAbstractAPIs_virtualbase_autoCompletionSelected(unsafe.Pointer(this.h), selection_ms)

}
func (this *QsciAbstractAPIs) OnAutoCompletionSelected(slot func(super func(selection string), selection string)) {
	ok := C.QsciAbstractAPIs_override_virtual_autoCompletionSelected(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_autoCompletionSelected
func miqt_exec_callback_QsciAbstractAPIs_autoCompletionSelected(self *C.QsciAbstractAPIs, cb C.intptr_t, selection C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection string), selection string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var selection_ms C.struct_miqt_string = selection
	selection_ret := C.GoStringN(selection_ms.data, C.int(int64(selection_ms.len)))
	C.free(unsafe.Pointer(selection_ms.data))
	slotval1 := selection_ret

	gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_AutoCompletionSelected, slotval1)

}
func (this *QsciAbstractAPIs) OnCallTips(slot func(context []string, commas int, style QsciScintilla__CallTipsStyle, shifts []int) []string) {
	ok := C.QsciAbstractAPIs_override_virtual_callTips(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_callTips
func miqt_exec_callback_QsciAbstractAPIs_callTips(self *C.QsciAbstractAPIs, cb C.intptr_t, context C.struct_miqt_array, commas C.int, style C.int, shifts C.struct_miqt_array) C.struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(context []string, commas int, style QsciScintilla__CallTipsStyle, shifts []int) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var context_ma C.struct_miqt_array = context
	context_ret := make([]string, int(context_ma.len))
	context_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(context_ma.data)) // hey ya
	for i := 0; i < int(context_ma.len); i++ {
		var context_lv_ms C.struct_miqt_string = context_outCast[i]
		context_lv_ret := C.GoStringN(context_lv_ms.data, C.int(int64(context_lv_ms.len)))
		C.free(unsafe.Pointer(context_lv_ms.data))
		context_ret[i] = context_lv_ret
	}
	slotval1 := context_ret

	slotval2 := (int)(commas)

	slotval3 := (QsciScintilla__CallTipsStyle)(style)

	var shifts_ma C.struct_miqt_array = shifts
	shifts_ret := make([]int, int(shifts_ma.len))
	shifts_outCast := (*[0xffff]C.int)(unsafe.Pointer(shifts_ma.data)) // hey ya
	for i := 0; i < int(shifts_ma.len); i++ {
		shifts_ret[i] = (int)(shifts_outCast[i])
	}
	slotval4 := shifts_ret

	virtualReturn := gofunc(slotval1, slotval2, slotval3, slotval4)
	virtualReturn_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(virtualReturn))))
	defer C.free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_i_ms := C.struct_miqt_string{}
		virtualReturn_i_ms.data = C.CString(virtualReturn[i])
		virtualReturn_i_ms.len = C.size_t(len(virtualReturn[i]))
		defer C.free(unsafe.Pointer(virtualReturn_i_ms.data))
		virtualReturn_CArray[i] = virtualReturn_i_ms
	}
	virtualReturn_ma := C.struct_miqt_array{len: C.size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QsciAbstractAPIs) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(C.QsciAbstractAPIs_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QsciAbstractAPIs) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	ok := C.QsciAbstractAPIs_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_event
func miqt_exec_callback_QsciAbstractAPIs_event(self *C.QsciAbstractAPIs, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QsciAbstractAPIs) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(C.QsciAbstractAPIs_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QsciAbstractAPIs) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	ok := C.QsciAbstractAPIs_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_eventFilter
func miqt_exec_callback_QsciAbstractAPIs_eventFilter(self *C.QsciAbstractAPIs, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QsciAbstractAPIs) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	C.QsciAbstractAPIs_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QsciAbstractAPIs) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	ok := C.QsciAbstractAPIs_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_timerEvent
func miqt_exec_callback_QsciAbstractAPIs_timerEvent(self *C.QsciAbstractAPIs, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QsciAbstractAPIs) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	C.QsciAbstractAPIs_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QsciAbstractAPIs) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	ok := C.QsciAbstractAPIs_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_childEvent
func miqt_exec_callback_QsciAbstractAPIs_childEvent(self *C.QsciAbstractAPIs, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QsciAbstractAPIs) callVirtualBase_CustomEvent(event *qt.QEvent) {

	C.QsciAbstractAPIs_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QsciAbstractAPIs) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	ok := C.QsciAbstractAPIs_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_customEvent
func miqt_exec_callback_QsciAbstractAPIs_customEvent(self *C.QsciAbstractAPIs, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QsciAbstractAPIs) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	C.QsciAbstractAPIs_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QsciAbstractAPIs) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QsciAbstractAPIs_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_connectNotify
func miqt_exec_callback_QsciAbstractAPIs_connectNotify(self *C.QsciAbstractAPIs, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QsciAbstractAPIs) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	C.QsciAbstractAPIs_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QsciAbstractAPIs) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QsciAbstractAPIs_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciAbstractAPIs_disconnectNotify
func miqt_exec_callback_QsciAbstractAPIs_disconnectNotify(self *C.QsciAbstractAPIs, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QsciAbstractAPIs{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QsciAbstractAPIs) Delete() {
	C.QsciAbstractAPIs_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QsciAbstractAPIs) GoGC() {
	runtime.SetFinalizer(this, func(this *QsciAbstractAPIs) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
