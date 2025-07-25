package svg

/*

#include "gen_qgraphicssvgitem.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QGraphicsSvgItem__ int

const (
	QGraphicsSvgItem__Type QGraphicsSvgItem__ = 13
)

type QGraphicsSvgItem struct {
	h *C.QGraphicsSvgItem
	*qt.QGraphicsObject
}

func (this *QGraphicsSvgItem) cPointer() *C.QGraphicsSvgItem {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QGraphicsSvgItem) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQGraphicsSvgItem constructs the type using only CGO pointers.
func newQGraphicsSvgItem(h *C.QGraphicsSvgItem) *QGraphicsSvgItem {
	if h == nil {
		return nil
	}
	var outptr_QGraphicsObject *C.QGraphicsObject = nil
	C.QGraphicsSvgItem_virtbase(h, &outptr_QGraphicsObject)

	return &QGraphicsSvgItem{h: h,
		QGraphicsObject: qt.UnsafeNewQGraphicsObject(unsafe.Pointer(outptr_QGraphicsObject))}
}

// UnsafeNewQGraphicsSvgItem constructs the type using only unsafe pointers.
func UnsafeNewQGraphicsSvgItem(h unsafe.Pointer) *QGraphicsSvgItem {
	return newQGraphicsSvgItem((*C.QGraphicsSvgItem)(h))
}

// NewQGraphicsSvgItem constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem() *QGraphicsSvgItem {

	return newQGraphicsSvgItem(C.QGraphicsSvgItem_new())
}

// NewQGraphicsSvgItem2 constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem2(fileName string) *QGraphicsSvgItem {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))

	return newQGraphicsSvgItem(C.QGraphicsSvgItem_new2(fileName_ms))
}

// NewQGraphicsSvgItem3 constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem3(parentItem *qt.QGraphicsItem) *QGraphicsSvgItem {

	return newQGraphicsSvgItem(C.QGraphicsSvgItem_new3((*C.QGraphicsItem)(parentItem.UnsafePointer())))
}

// NewQGraphicsSvgItem4 constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem4(fileName string, parentItem *qt.QGraphicsItem) *QGraphicsSvgItem {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))

	return newQGraphicsSvgItem(C.QGraphicsSvgItem_new4(fileName_ms, (*C.QGraphicsItem)(parentItem.UnsafePointer())))
}

func (this *QGraphicsSvgItem) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QGraphicsSvgItem_metaObject(this.h)))
}

func (this *QGraphicsSvgItem) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QGraphicsSvgItem_metacast(this.h, param1_Cstring))
}

func QGraphicsSvgItem_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QGraphicsSvgItem_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsSvgItem_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QGraphicsSvgItem_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsSvgItem) SetSharedRenderer(renderer *QSvgRenderer) {
	C.QGraphicsSvgItem_setSharedRenderer(this.h, renderer.cPointer())
}

func (this *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	return newQSvgRenderer(C.QGraphicsSvgItem_renderer(this.h))
}

func (this *QGraphicsSvgItem) SetElementId(id string) {
	id_ms := C.struct_miqt_string{}
	id_ms.data = C.CString(id)
	id_ms.len = C.size_t(len(id))
	defer C.free(unsafe.Pointer(id_ms.data))
	C.QGraphicsSvgItem_setElementId(this.h, id_ms)
}

func (this *QGraphicsSvgItem) ElementId() string {
	var _ms C.struct_miqt_string = C.QGraphicsSvgItem_elementId(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsSvgItem) SetCachingEnabled(cachingEnabled bool) {
	C.QGraphicsSvgItem_setCachingEnabled(this.h, (C.bool)(cachingEnabled))
}

func (this *QGraphicsSvgItem) IsCachingEnabled() bool {
	return (bool)(C.QGraphicsSvgItem_isCachingEnabled(this.h))
}

func (this *QGraphicsSvgItem) SetMaximumCacheSize(size *qt.QSize) {
	C.QGraphicsSvgItem_setMaximumCacheSize(this.h, (*C.QSize)(size.UnsafePointer()))
}

func (this *QGraphicsSvgItem) MaximumCacheSize() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(C.QGraphicsSvgItem_maximumCacheSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSvgItem) BoundingRect() *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(C.QGraphicsSvgItem_boundingRect(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSvgItem) Paint(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget) {
	C.QGraphicsSvgItem_paint(this.h, (*C.QPainter)(painter.UnsafePointer()), (*C.QStyleOptionGraphicsItem)(option.UnsafePointer()), (*C.QWidget)(widget.UnsafePointer()))
}

func (this *QGraphicsSvgItem) Type() int {
	return (int)(C.QGraphicsSvgItem_type(this.h))
}

func QGraphicsSvgItem_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QGraphicsSvgItem_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsSvgItem_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QGraphicsSvgItem_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsSvgItem_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QGraphicsSvgItem_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsSvgItem_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QGraphicsSvgItem_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// UpdateMicroFocus can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) UpdateMicroFocus() {

	var _dynamic_cast_ok C.bool = false
	C.QGraphicsSvgItem_protectedbase_updateMicroFocus(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// Sender can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) Sender() *qt.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt.UnsafeNewQObject(unsafe.Pointer(C.QGraphicsSvgItem_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QGraphicsSvgItem_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QGraphicsSvgItem_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) IsSignalConnected(signal *qt.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QGraphicsSvgItem_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// AddToIndex can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) AddToIndex() {

	var _dynamic_cast_ok C.bool = false
	C.QGraphicsSvgItem_protectedbase_addToIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// RemoveFromIndex can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) RemoveFromIndex() {

	var _dynamic_cast_ok C.bool = false
	C.QGraphicsSvgItem_protectedbase_removeFromIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// PrepareGeometryChange can only be called from a QGraphicsSvgItem that was directly constructed.
func (this *QGraphicsSvgItem) PrepareGeometryChange() {

	var _dynamic_cast_ok C.bool = false
	C.QGraphicsSvgItem_protectedbase_prepareGeometryChange(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

func (this *QGraphicsSvgItem) callVirtualBase_BoundingRect() *qt.QRectF {

	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(C.QGraphicsSvgItem_virtualbase_boundingRect(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsSvgItem) OnBoundingRect(slot func(super func() *qt.QRectF) *qt.QRectF) {
	ok := C.QGraphicsSvgItem_override_virtual_boundingRect(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_boundingRect
func miqt_exec_callback_QGraphicsSvgItem_boundingRect(self *C.QGraphicsSvgItem, cb C.intptr_t) *C.QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QRectF) *qt.QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_BoundingRect)

	return (*C.QRectF)(virtualReturn.UnsafePointer())

}

func (this *QGraphicsSvgItem) callVirtualBase_Paint(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget) {

	C.QGraphicsSvgItem_virtualbase_paint(unsafe.Pointer(this.h), (*C.QPainter)(painter.UnsafePointer()), (*C.QStyleOptionGraphicsItem)(option.UnsafePointer()), (*C.QWidget)(widget.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnPaint(slot func(super func(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget), painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget)) {
	ok := C.QGraphicsSvgItem_override_virtual_paint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_paint
func miqt_exec_callback_QGraphicsSvgItem_paint(self *C.QGraphicsSvgItem, cb C.intptr_t, painter *C.QPainter, option *C.QStyleOptionGraphicsItem, widget *C.QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget), painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))

	slotval2 := qt.UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option))

	slotval3 := qt.UnsafeNewQWidget(unsafe.Pointer(widget))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Paint, slotval1, slotval2, slotval3)

}

func (this *QGraphicsSvgItem) callVirtualBase_Type() int {

	return (int)(C.QGraphicsSvgItem_virtualbase_type(unsafe.Pointer(this.h)))

}
func (this *QGraphicsSvgItem) OnType(slot func(super func() int) int) {
	ok := C.QGraphicsSvgItem_override_virtual_type(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_type
func miqt_exec_callback_QGraphicsSvgItem_type(self *C.QGraphicsSvgItem, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Type)

	return (C.int)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_Event(ev *qt.QEvent) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(ev.UnsafePointer())))

}
func (this *QGraphicsSvgItem) OnEvent(slot func(super func(ev *qt.QEvent) bool, ev *qt.QEvent) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_event
func miqt_exec_callback_QGraphicsSvgItem_event(self *C.QGraphicsSvgItem, cb C.intptr_t, ev *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *qt.QEvent) bool, ev *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(ev))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QGraphicsSvgItem) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_eventFilter
func miqt_exec_callback_QGraphicsSvgItem_eventFilter(self *C.QGraphicsSvgItem, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	C.QGraphicsSvgItem_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_timerEvent
func miqt_exec_callback_QGraphicsSvgItem_timerEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	C.QGraphicsSvgItem_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_childEvent
func miqt_exec_callback_QGraphicsSvgItem_childEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_CustomEvent(event *qt.QEvent) {

	C.QGraphicsSvgItem_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_customEvent
func miqt_exec_callback_QGraphicsSvgItem_customEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	C.QGraphicsSvgItem_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QGraphicsSvgItem_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_connectNotify
func miqt_exec_callback_QGraphicsSvgItem_connectNotify(self *C.QGraphicsSvgItem, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	C.QGraphicsSvgItem_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QGraphicsSvgItem_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_disconnectNotify
func miqt_exec_callback_QGraphicsSvgItem_disconnectNotify(self *C.QGraphicsSvgItem, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_Advance(phase int) {

	C.QGraphicsSvgItem_virtualbase_advance(unsafe.Pointer(this.h), (C.int)(phase))

}
func (this *QGraphicsSvgItem) OnAdvance(slot func(super func(phase int), phase int)) {
	ok := C.QGraphicsSvgItem_override_virtual_advance(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_advance
func miqt_exec_callback_QGraphicsSvgItem_advance(self *C.QGraphicsSvgItem, cb C.intptr_t, phase C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(phase int), phase int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(phase)

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Advance, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_Shape() *qt.QPainterPath {

	_goptr := qt.UnsafeNewQPainterPath(unsafe.Pointer(C.QGraphicsSvgItem_virtualbase_shape(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsSvgItem) OnShape(slot func(super func() *qt.QPainterPath) *qt.QPainterPath) {
	ok := C.QGraphicsSvgItem_override_virtual_shape(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_shape
func miqt_exec_callback_QGraphicsSvgItem_shape(self *C.QGraphicsSvgItem, cb C.intptr_t) *C.QPainterPath {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPainterPath) *qt.QPainterPath)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Shape)

	return (*C.QPainterPath)(virtualReturn.UnsafePointer())

}

func (this *QGraphicsSvgItem) callVirtualBase_Contains(point *qt.QPointF) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_contains(unsafe.Pointer(this.h), (*C.QPointF)(point.UnsafePointer())))

}
func (this *QGraphicsSvgItem) OnContains(slot func(super func(point *qt.QPointF) bool, point *qt.QPointF) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_contains(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_contains
func miqt_exec_callback_QGraphicsSvgItem_contains(self *C.QGraphicsSvgItem, cb C.intptr_t, point *C.QPointF) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(point *qt.QPointF) bool, point *qt.QPointF) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPointF(unsafe.Pointer(point))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Contains, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_CollidesWithItem(other *qt.QGraphicsItem, mode qt.ItemSelectionMode) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_collidesWithItem(unsafe.Pointer(this.h), (*C.QGraphicsItem)(other.UnsafePointer()), (C.int)(mode)))

}
func (this *QGraphicsSvgItem) OnCollidesWithItem(slot func(super func(other *qt.QGraphicsItem, mode qt.ItemSelectionMode) bool, other *qt.QGraphicsItem, mode qt.ItemSelectionMode) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_collidesWithItem(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_collidesWithItem
func miqt_exec_callback_QGraphicsSvgItem_collidesWithItem(self *C.QGraphicsSvgItem, cb C.intptr_t, other *C.QGraphicsItem, mode C.int) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(other *qt.QGraphicsItem, mode qt.ItemSelectionMode) bool, other *qt.QGraphicsItem, mode qt.ItemSelectionMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsItem(unsafe.Pointer(other))

	slotval2 := (qt.ItemSelectionMode)(mode)

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_CollidesWithItem, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_CollidesWithPath(path *qt.QPainterPath, mode qt.ItemSelectionMode) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_collidesWithPath(unsafe.Pointer(this.h), (*C.QPainterPath)(path.UnsafePointer()), (C.int)(mode)))

}
func (this *QGraphicsSvgItem) OnCollidesWithPath(slot func(super func(path *qt.QPainterPath, mode qt.ItemSelectionMode) bool, path *qt.QPainterPath, mode qt.ItemSelectionMode) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_collidesWithPath(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_collidesWithPath
func miqt_exec_callback_QGraphicsSvgItem_collidesWithPath(self *C.QGraphicsSvgItem, cb C.intptr_t, path *C.QPainterPath, mode C.int) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(path *qt.QPainterPath, mode qt.ItemSelectionMode) bool, path *qt.QPainterPath, mode qt.ItemSelectionMode) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainterPath(unsafe.Pointer(path))

	slotval2 := (qt.ItemSelectionMode)(mode)

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_CollidesWithPath, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_IsObscuredBy(item *qt.QGraphicsItem) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_isObscuredBy(unsafe.Pointer(this.h), (*C.QGraphicsItem)(item.UnsafePointer())))

}
func (this *QGraphicsSvgItem) OnIsObscuredBy(slot func(super func(item *qt.QGraphicsItem) bool, item *qt.QGraphicsItem) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_isObscuredBy(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_isObscuredBy
func miqt_exec_callback_QGraphicsSvgItem_isObscuredBy(self *C.QGraphicsSvgItem, cb C.intptr_t, item *C.QGraphicsItem) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(item *qt.QGraphicsItem) bool, item *qt.QGraphicsItem) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsItem(unsafe.Pointer(item))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_IsObscuredBy, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_OpaqueArea() *qt.QPainterPath {

	_goptr := qt.UnsafeNewQPainterPath(unsafe.Pointer(C.QGraphicsSvgItem_virtualbase_opaqueArea(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsSvgItem) OnOpaqueArea(slot func(super func() *qt.QPainterPath) *qt.QPainterPath) {
	ok := C.QGraphicsSvgItem_override_virtual_opaqueArea(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_opaqueArea
func miqt_exec_callback_QGraphicsSvgItem_opaqueArea(self *C.QGraphicsSvgItem, cb C.intptr_t) *C.QPainterPath {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPainterPath) *qt.QPainterPath)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_OpaqueArea)

	return (*C.QPainterPath)(virtualReturn.UnsafePointer())

}

func (this *QGraphicsSvgItem) callVirtualBase_SceneEventFilter(watched *qt.QGraphicsItem, event *qt.QEvent) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_sceneEventFilter(unsafe.Pointer(this.h), (*C.QGraphicsItem)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QGraphicsSvgItem) OnSceneEventFilter(slot func(super func(watched *qt.QGraphicsItem, event *qt.QEvent) bool, watched *qt.QGraphicsItem, event *qt.QEvent) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_sceneEventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_sceneEventFilter
func miqt_exec_callback_QGraphicsSvgItem_sceneEventFilter(self *C.QGraphicsSvgItem, cb C.intptr_t, watched *C.QGraphicsItem, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QGraphicsItem, event *qt.QEvent) bool, watched *qt.QGraphicsItem, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsItem(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_SceneEventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_SceneEvent(event *qt.QEvent) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_sceneEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QGraphicsSvgItem) OnSceneEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_sceneEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_sceneEvent
func miqt_exec_callback_QGraphicsSvgItem_sceneEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_SceneEvent, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_ContextMenuEvent(event *qt.QGraphicsSceneContextMenuEvent) {

	C.QGraphicsSvgItem_virtualbase_contextMenuEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneContextMenuEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnContextMenuEvent(slot func(super func(event *qt.QGraphicsSceneContextMenuEvent), event *qt.QGraphicsSceneContextMenuEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_contextMenuEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_contextMenuEvent
func miqt_exec_callback_QGraphicsSvgItem_contextMenuEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneContextMenuEvent), event *qt.QGraphicsSceneContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_DragEnterEvent(event *qt.QGraphicsSceneDragDropEvent) {

	C.QGraphicsSvgItem_virtualbase_dragEnterEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneDragDropEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnDragEnterEvent(slot func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_dragEnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_dragEnterEvent
func miqt_exec_callback_QGraphicsSvgItem_dragEnterEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_DragLeaveEvent(event *qt.QGraphicsSceneDragDropEvent) {

	C.QGraphicsSvgItem_virtualbase_dragLeaveEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneDragDropEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnDragLeaveEvent(slot func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_dragLeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_dragLeaveEvent
func miqt_exec_callback_QGraphicsSvgItem_dragLeaveEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_DragMoveEvent(event *qt.QGraphicsSceneDragDropEvent) {

	C.QGraphicsSvgItem_virtualbase_dragMoveEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneDragDropEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnDragMoveEvent(slot func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_dragMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_dragMoveEvent
func miqt_exec_callback_QGraphicsSvgItem_dragMoveEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_DropEvent(event *qt.QGraphicsSceneDragDropEvent) {

	C.QGraphicsSvgItem_virtualbase_dropEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneDragDropEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnDropEvent(slot func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_dropEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_dropEvent
func miqt_exec_callback_QGraphicsSvgItem_dropEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneDragDropEvent), event *qt.QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_FocusInEvent(event *qt.QFocusEvent) {

	C.QGraphicsSvgItem_virtualbase_focusInEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnFocusInEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_focusInEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_focusInEvent
func miqt_exec_callback_QGraphicsSvgItem_focusInEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_FocusOutEvent(event *qt.QFocusEvent) {

	C.QGraphicsSvgItem_virtualbase_focusOutEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnFocusOutEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_focusOutEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_focusOutEvent
func miqt_exec_callback_QGraphicsSvgItem_focusOutEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_HoverEnterEvent(event *qt.QGraphicsSceneHoverEvent) {

	C.QGraphicsSvgItem_virtualbase_hoverEnterEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneHoverEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnHoverEnterEvent(slot func(super func(event *qt.QGraphicsSceneHoverEvent), event *qt.QGraphicsSceneHoverEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_hoverEnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_hoverEnterEvent
func miqt_exec_callback_QGraphicsSvgItem_hoverEnterEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneHoverEvent), event *qt.QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_HoverEnterEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_HoverMoveEvent(event *qt.QGraphicsSceneHoverEvent) {

	C.QGraphicsSvgItem_virtualbase_hoverMoveEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneHoverEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnHoverMoveEvent(slot func(super func(event *qt.QGraphicsSceneHoverEvent), event *qt.QGraphicsSceneHoverEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_hoverMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_hoverMoveEvent
func miqt_exec_callback_QGraphicsSvgItem_hoverMoveEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneHoverEvent), event *qt.QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_HoverMoveEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_HoverLeaveEvent(event *qt.QGraphicsSceneHoverEvent) {

	C.QGraphicsSvgItem_virtualbase_hoverLeaveEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneHoverEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnHoverLeaveEvent(slot func(super func(event *qt.QGraphicsSceneHoverEvent), event *qt.QGraphicsSceneHoverEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_hoverLeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_hoverLeaveEvent
func miqt_exec_callback_QGraphicsSvgItem_hoverLeaveEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneHoverEvent), event *qt.QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_HoverLeaveEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_KeyPressEvent(event *qt.QKeyEvent) {

	C.QGraphicsSvgItem_virtualbase_keyPressEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnKeyPressEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_keyPressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_keyPressEvent
func miqt_exec_callback_QGraphicsSvgItem_keyPressEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_KeyReleaseEvent(event *qt.QKeyEvent) {

	C.QGraphicsSvgItem_virtualbase_keyReleaseEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnKeyReleaseEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_keyReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_keyReleaseEvent
func miqt_exec_callback_QGraphicsSvgItem_keyReleaseEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_MousePressEvent(event *qt.QGraphicsSceneMouseEvent) {

	C.QGraphicsSvgItem_virtualbase_mousePressEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneMouseEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnMousePressEvent(slot func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_mousePressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_mousePressEvent
func miqt_exec_callback_QGraphicsSvgItem_mousePressEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_MouseMoveEvent(event *qt.QGraphicsSceneMouseEvent) {

	C.QGraphicsSvgItem_virtualbase_mouseMoveEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneMouseEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnMouseMoveEvent(slot func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_mouseMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_mouseMoveEvent
func miqt_exec_callback_QGraphicsSvgItem_mouseMoveEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_MouseReleaseEvent(event *qt.QGraphicsSceneMouseEvent) {

	C.QGraphicsSvgItem_virtualbase_mouseReleaseEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneMouseEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnMouseReleaseEvent(slot func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_mouseReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_mouseReleaseEvent
func miqt_exec_callback_QGraphicsSvgItem_mouseReleaseEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_MouseDoubleClickEvent(event *qt.QGraphicsSceneMouseEvent) {

	C.QGraphicsSvgItem_virtualbase_mouseDoubleClickEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneMouseEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnMouseDoubleClickEvent(slot func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_mouseDoubleClickEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_mouseDoubleClickEvent
func miqt_exec_callback_QGraphicsSvgItem_mouseDoubleClickEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_WheelEvent(event *qt.QGraphicsSceneWheelEvent) {

	C.QGraphicsSvgItem_virtualbase_wheelEvent(unsafe.Pointer(this.h), (*C.QGraphicsSceneWheelEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnWheelEvent(slot func(super func(event *qt.QGraphicsSceneWheelEvent), event *qt.QGraphicsSceneWheelEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_wheelEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_wheelEvent
func miqt_exec_callback_QGraphicsSvgItem_wheelEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QGraphicsSceneWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QGraphicsSceneWheelEvent), event *qt.QGraphicsSceneWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_InputMethodEvent(event *qt.QInputMethodEvent) {

	C.QGraphicsSvgItem_virtualbase_inputMethodEvent(unsafe.Pointer(this.h), (*C.QInputMethodEvent)(event.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnInputMethodEvent(slot func(super func(event *qt.QInputMethodEvent), event *qt.QInputMethodEvent)) {
	ok := C.QGraphicsSvgItem_override_virtual_inputMethodEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_inputMethodEvent
func miqt_exec_callback_QGraphicsSvgItem_inputMethodEvent(self *C.QGraphicsSvgItem, cb C.intptr_t, event *C.QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QInputMethodEvent), event *qt.QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQInputMethodEvent(unsafe.Pointer(event))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QGraphicsSvgItem) callVirtualBase_InputMethodQuery(query qt.InputMethodQuery) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(C.QGraphicsSvgItem_virtualbase_inputMethodQuery(unsafe.Pointer(this.h), (C.int)(query))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsSvgItem) OnInputMethodQuery(slot func(super func(query qt.InputMethodQuery) *qt.QVariant, query qt.InputMethodQuery) *qt.QVariant) {
	ok := C.QGraphicsSvgItem_override_virtual_inputMethodQuery(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_inputMethodQuery
func miqt_exec_callback_QGraphicsSvgItem_inputMethodQuery(self *C.QGraphicsSvgItem, cb C.intptr_t, query C.int) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query qt.InputMethodQuery) *qt.QVariant, query qt.InputMethodQuery) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.InputMethodQuery)(query)

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return (*C.QVariant)(virtualReturn.UnsafePointer())

}

func (this *QGraphicsSvgItem) callVirtualBase_ItemChange(change qt.QGraphicsItem__GraphicsItemChange, value *qt.QVariant) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(C.QGraphicsSvgItem_virtualbase_itemChange(unsafe.Pointer(this.h), (C.int)(change), (*C.QVariant)(value.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsSvgItem) OnItemChange(slot func(super func(change qt.QGraphicsItem__GraphicsItemChange, value *qt.QVariant) *qt.QVariant, change qt.QGraphicsItem__GraphicsItemChange, value *qt.QVariant) *qt.QVariant) {
	ok := C.QGraphicsSvgItem_override_virtual_itemChange(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_itemChange
func miqt_exec_callback_QGraphicsSvgItem_itemChange(self *C.QGraphicsSvgItem, cb C.intptr_t, change C.int, value *C.QVariant) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(change qt.QGraphicsItem__GraphicsItemChange, value *qt.QVariant) *qt.QVariant, change qt.QGraphicsItem__GraphicsItemChange, value *qt.QVariant) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QGraphicsItem__GraphicsItemChange)(change)

	slotval2 := qt.UnsafeNewQVariant(unsafe.Pointer(value))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_ItemChange, slotval1, slotval2)

	return (*C.QVariant)(virtualReturn.UnsafePointer())

}

func (this *QGraphicsSvgItem) callVirtualBase_SupportsExtension(extension qt.QGraphicsItem__Extension) bool {

	return (bool)(C.QGraphicsSvgItem_virtualbase_supportsExtension(unsafe.Pointer(this.h), (C.int)(extension)))

}
func (this *QGraphicsSvgItem) OnSupportsExtension(slot func(super func(extension qt.QGraphicsItem__Extension) bool, extension qt.QGraphicsItem__Extension) bool) {
	ok := C.QGraphicsSvgItem_override_virtual_supportsExtension(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_supportsExtension
func miqt_exec_callback_QGraphicsSvgItem_supportsExtension(self *C.QGraphicsSvgItem, cb C.intptr_t, extension C.int) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(extension qt.QGraphicsItem__Extension) bool, extension qt.QGraphicsItem__Extension) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QGraphicsItem__Extension)(extension)

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_SupportsExtension, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_SetExtension(extension qt.QGraphicsItem__Extension, variant *qt.QVariant) {

	C.QGraphicsSvgItem_virtualbase_setExtension(unsafe.Pointer(this.h), (C.int)(extension), (*C.QVariant)(variant.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnSetExtension(slot func(super func(extension qt.QGraphicsItem__Extension, variant *qt.QVariant), extension qt.QGraphicsItem__Extension, variant *qt.QVariant)) {
	ok := C.QGraphicsSvgItem_override_virtual_setExtension(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_setExtension
func miqt_exec_callback_QGraphicsSvgItem_setExtension(self *C.QGraphicsSvgItem, cb C.intptr_t, extension C.int, variant *C.QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(extension qt.QGraphicsItem__Extension, variant *qt.QVariant), extension qt.QGraphicsItem__Extension, variant *qt.QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QGraphicsItem__Extension)(extension)

	slotval2 := qt.UnsafeNewQVariant(unsafe.Pointer(variant))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_SetExtension, slotval1, slotval2)

}

func (this *QGraphicsSvgItem) callVirtualBase_Extension(variant *qt.QVariant) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(C.QGraphicsSvgItem_virtualbase_extension(unsafe.Pointer(this.h), (*C.QVariant)(variant.UnsafePointer()))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsSvgItem) OnExtension(slot func(super func(variant *qt.QVariant) *qt.QVariant, variant *qt.QVariant) *qt.QVariant) {
	ok := C.QGraphicsSvgItem_override_virtual_extension(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QGraphicsSvgItem_extension
func miqt_exec_callback_QGraphicsSvgItem_extension(self *C.QGraphicsSvgItem, cb C.intptr_t, variant *C.QVariant) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(variant *qt.QVariant) *qt.QVariant, variant *qt.QVariant) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQVariant(unsafe.Pointer(variant))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Extension, slotval1)

	return (*C.QVariant)(virtualReturn.UnsafePointer())

}

// Delete this object from C++ memory.
func (this *QGraphicsSvgItem) Delete() {
	C.QGraphicsSvgItem_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QGraphicsSvgItem) GoGC() {
	runtime.SetFinalizer(this, func(this *QGraphicsSvgItem) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
