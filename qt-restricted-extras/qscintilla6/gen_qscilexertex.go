package qscintilla6

/*

#include "gen_qscilexertex.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QsciLexerTeX__ int

const (
	QsciLexerTeX__Default QsciLexerTeX__ = 0
	QsciLexerTeX__Special QsciLexerTeX__ = 1
	QsciLexerTeX__Group   QsciLexerTeX__ = 2
	QsciLexerTeX__Symbol  QsciLexerTeX__ = 3
	QsciLexerTeX__Command QsciLexerTeX__ = 4
	QsciLexerTeX__Text    QsciLexerTeX__ = 5
)

type QsciLexerTeX struct {
	h *C.QsciLexerTeX
	*QsciLexer
}

func (this *QsciLexerTeX) cPointer() *C.QsciLexerTeX {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QsciLexerTeX) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQsciLexerTeX constructs the type using only CGO pointers.
func newQsciLexerTeX(h *C.QsciLexerTeX) *QsciLexerTeX {
	if h == nil {
		return nil
	}
	var outptr_QsciLexer *C.QsciLexer = nil
	C.QsciLexerTeX_virtbase(h, &outptr_QsciLexer)

	return &QsciLexerTeX{h: h,
		QsciLexer: newQsciLexer(outptr_QsciLexer)}
}

// UnsafeNewQsciLexerTeX constructs the type using only unsafe pointers.
func UnsafeNewQsciLexerTeX(h unsafe.Pointer) *QsciLexerTeX {
	return newQsciLexerTeX((*C.QsciLexerTeX)(h))
}

// NewQsciLexerTeX constructs a new QsciLexerTeX object.
func NewQsciLexerTeX() *QsciLexerTeX {

	return newQsciLexerTeX(C.QsciLexerTeX_new())
}

// NewQsciLexerTeX2 constructs a new QsciLexerTeX object.
func NewQsciLexerTeX2(parent *qt6.QObject) *QsciLexerTeX {

	return newQsciLexerTeX(C.QsciLexerTeX_new2((*C.QObject)(parent.UnsafePointer())))
}

func (this *QsciLexerTeX) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QsciLexerTeX_metaObject(this.h)))
}

func (this *QsciLexerTeX) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QsciLexerTeX_metacast(this.h, param1_Cstring))
}

func QsciLexerTeX_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerTeX_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciLexerTeX) Language() string {
	_ret := C.QsciLexerTeX_language(this.h)
	return C.GoString(_ret)
}

func (this *QsciLexerTeX) Lexer() string {
	_ret := C.QsciLexerTeX_lexer(this.h)
	return C.GoString(_ret)
}

func (this *QsciLexerTeX) WordCharacters() string {
	_ret := C.QsciLexerTeX_wordCharacters(this.h)
	return C.GoString(_ret)
}

func (this *QsciLexerTeX) DefaultColor(style int) *qt6.QColor {
	_goptr := qt6.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerTeX_defaultColor(this.h, (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QsciLexerTeX) Keywords(set int) string {
	_ret := C.QsciLexerTeX_keywords(this.h, (C.int)(set))
	return C.GoString(_ret)
}

func (this *QsciLexerTeX) Description(style int) string {
	var _ms C.struct_miqt_string = C.QsciLexerTeX_description(this.h, (C.int)(style))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QsciLexerTeX) RefreshProperties() {
	C.QsciLexerTeX_refreshProperties(this.h)
}

func (this *QsciLexerTeX) SetFoldComments(fold bool) {
	C.QsciLexerTeX_setFoldComments(this.h, (C.bool)(fold))
}

func (this *QsciLexerTeX) FoldComments() bool {
	return (bool)(C.QsciLexerTeX_foldComments(this.h))
}

func (this *QsciLexerTeX) SetFoldCompact(fold bool) {
	C.QsciLexerTeX_setFoldCompact(this.h, (C.bool)(fold))
}

func (this *QsciLexerTeX) FoldCompact() bool {
	return (bool)(C.QsciLexerTeX_foldCompact(this.h))
}

func (this *QsciLexerTeX) SetProcessComments(enable bool) {
	C.QsciLexerTeX_setProcessComments(this.h, (C.bool)(enable))
}

func (this *QsciLexerTeX) ProcessComments() bool {
	return (bool)(C.QsciLexerTeX_processComments(this.h))
}

func (this *QsciLexerTeX) SetProcessIf(enable bool) {
	C.QsciLexerTeX_setProcessIf(this.h, (C.bool)(enable))
}

func (this *QsciLexerTeX) ProcessIf() bool {
	return (bool)(C.QsciLexerTeX_processIf(this.h))
}

func QsciLexerTeX_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerTeX_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QsciLexerTeX_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QsciLexerTeX_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// ReadProperties can only be called from a QsciLexerTeX that was directly constructed.
func (this *QsciLexerTeX) ReadProperties(qs *qt6.QSettings, prefix string) bool {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QsciLexerTeX_protectedbase_readProperties(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QSettings)(qs.UnsafePointer()), prefix_ms))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// WriteProperties can only be called from a QsciLexerTeX that was directly constructed.
func (this *QsciLexerTeX) WriteProperties(qs *qt6.QSettings, prefix string) bool {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QsciLexerTeX_protectedbase_writeProperties(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QSettings)(qs.UnsafePointer()), prefix_ms))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Sender can only be called from a QsciLexerTeX that was directly constructed.
func (this *QsciLexerTeX) Sender() *qt6.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt6.UnsafeNewQObject(unsafe.Pointer(C.QsciLexerTeX_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QsciLexerTeX that was directly constructed.
func (this *QsciLexerTeX) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QsciLexerTeX_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QsciLexerTeX that was directly constructed.
func (this *QsciLexerTeX) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QsciLexerTeX_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QsciLexerTeX that was directly constructed.
func (this *QsciLexerTeX) IsSignalConnected(signal *qt6.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QsciLexerTeX_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}
func (this *QsciLexerTeX) OnLanguage(slot func() string) {
	ok := C.QsciLexerTeX_override_virtual_language(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_language
func miqt_exec_callback_QsciLexerTeX_language(self *C.QsciLexerTeX, cb C.intptr_t) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func() string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_Lexer() string {

	_ret := C.QsciLexerTeX_virtualbase_lexer(unsafe.Pointer(this.h))
	return C.GoString(_ret)

}
func (this *QsciLexerTeX) OnLexer(slot func(super func() string) string) {
	ok := C.QsciLexerTeX_override_virtual_lexer(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_lexer
func miqt_exec_callback_QsciLexerTeX_lexer(self *C.QsciLexerTeX, cb C.intptr_t) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_Lexer)
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_LexerId() int {

	return (int)(C.QsciLexerTeX_virtualbase_lexerId(unsafe.Pointer(this.h)))

}
func (this *QsciLexerTeX) OnLexerId(slot func(super func() int) int) {
	ok := C.QsciLexerTeX_override_virtual_lexerId(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_lexerId
func miqt_exec_callback_QsciLexerTeX_lexerId(self *C.QsciLexerTeX, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_LexerId)

	return (C.int)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_AutoCompletionFillups() string {

	_ret := C.QsciLexerTeX_virtualbase_autoCompletionFillups(unsafe.Pointer(this.h))
	return C.GoString(_ret)

}
func (this *QsciLexerTeX) OnAutoCompletionFillups(slot func(super func() string) string) {
	ok := C.QsciLexerTeX_override_virtual_autoCompletionFillups(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_autoCompletionFillups
func miqt_exec_callback_QsciLexerTeX_autoCompletionFillups(self *C.QsciLexerTeX, cb C.intptr_t) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_AutoCompletionFillups)
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_AutoCompletionWordSeparators() []string {

	var _ma C.struct_miqt_array = C.QsciLexerTeX_virtualbase_autoCompletionWordSeparators(unsafe.Pointer(this.h))
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
func (this *QsciLexerTeX) OnAutoCompletionWordSeparators(slot func(super func() []string) []string) {
	ok := C.QsciLexerTeX_override_virtual_autoCompletionWordSeparators(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_autoCompletionWordSeparators
func miqt_exec_callback_QsciLexerTeX_autoCompletionWordSeparators(self *C.QsciLexerTeX, cb C.intptr_t) C.struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_AutoCompletionWordSeparators)
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

func (this *QsciLexerTeX) callVirtualBase_BlockEnd(style *int) string {

	_ret := C.QsciLexerTeX_virtualbase_blockEnd(unsafe.Pointer(this.h), (*C.int)(unsafe.Pointer(style)))
	return C.GoString(_ret)

}
func (this *QsciLexerTeX) OnBlockEnd(slot func(super func(style *int) string, style *int) string) {
	ok := C.QsciLexerTeX_override_virtual_blockEnd(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_blockEnd
func miqt_exec_callback_QsciLexerTeX_blockEnd(self *C.QsciLexerTeX, cb C.intptr_t, style *C.int) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style *int) string, style *int) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*int)(unsafe.Pointer(style))

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_BlockEnd, slotval1)
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_BlockLookback() int {

	return (int)(C.QsciLexerTeX_virtualbase_blockLookback(unsafe.Pointer(this.h)))

}
func (this *QsciLexerTeX) OnBlockLookback(slot func(super func() int) int) {
	ok := C.QsciLexerTeX_override_virtual_blockLookback(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_blockLookback
func miqt_exec_callback_QsciLexerTeX_blockLookback(self *C.QsciLexerTeX, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_BlockLookback)

	return (C.int)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_BlockStart(style *int) string {

	_ret := C.QsciLexerTeX_virtualbase_blockStart(unsafe.Pointer(this.h), (*C.int)(unsafe.Pointer(style)))
	return C.GoString(_ret)

}
func (this *QsciLexerTeX) OnBlockStart(slot func(super func(style *int) string, style *int) string) {
	ok := C.QsciLexerTeX_override_virtual_blockStart(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_blockStart
func miqt_exec_callback_QsciLexerTeX_blockStart(self *C.QsciLexerTeX, cb C.intptr_t, style *C.int) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style *int) string, style *int) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*int)(unsafe.Pointer(style))

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_BlockStart, slotval1)
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_BlockStartKeyword(style *int) string {

	_ret := C.QsciLexerTeX_virtualbase_blockStartKeyword(unsafe.Pointer(this.h), (*C.int)(unsafe.Pointer(style)))
	return C.GoString(_ret)

}
func (this *QsciLexerTeX) OnBlockStartKeyword(slot func(super func(style *int) string, style *int) string) {
	ok := C.QsciLexerTeX_override_virtual_blockStartKeyword(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_blockStartKeyword
func miqt_exec_callback_QsciLexerTeX_blockStartKeyword(self *C.QsciLexerTeX, cb C.intptr_t, style *C.int) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style *int) string, style *int) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*int)(unsafe.Pointer(style))

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_BlockStartKeyword, slotval1)
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_BraceStyle() int {

	return (int)(C.QsciLexerTeX_virtualbase_braceStyle(unsafe.Pointer(this.h)))

}
func (this *QsciLexerTeX) OnBraceStyle(slot func(super func() int) int) {
	ok := C.QsciLexerTeX_override_virtual_braceStyle(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_braceStyle
func miqt_exec_callback_QsciLexerTeX_braceStyle(self *C.QsciLexerTeX, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_BraceStyle)

	return (C.int)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_CaseSensitive() bool {

	return (bool)(C.QsciLexerTeX_virtualbase_caseSensitive(unsafe.Pointer(this.h)))

}
func (this *QsciLexerTeX) OnCaseSensitive(slot func(super func() bool) bool) {
	ok := C.QsciLexerTeX_override_virtual_caseSensitive(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_caseSensitive
func miqt_exec_callback_QsciLexerTeX_caseSensitive(self *C.QsciLexerTeX, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_CaseSensitive)

	return (C.bool)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_Color(style int) *qt6.QColor {

	_goptr := qt6.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerTeX_virtualbase_color(unsafe.Pointer(this.h), (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QsciLexerTeX) OnColor(slot func(super func(style int) *qt6.QColor, style int) *qt6.QColor) {
	ok := C.QsciLexerTeX_override_virtual_color(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_color
func miqt_exec_callback_QsciLexerTeX_color(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) *C.QColor {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) *qt6.QColor, style int) *qt6.QColor)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_Color, slotval1)

	return (*C.QColor)(virtualReturn.UnsafePointer())

}

func (this *QsciLexerTeX) callVirtualBase_EolFill(style int) bool {

	return (bool)(C.QsciLexerTeX_virtualbase_eolFill(unsafe.Pointer(this.h), (C.int)(style)))

}
func (this *QsciLexerTeX) OnEolFill(slot func(super func(style int) bool, style int) bool) {
	ok := C.QsciLexerTeX_override_virtual_eolFill(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_eolFill
func miqt_exec_callback_QsciLexerTeX_eolFill(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) bool, style int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_EolFill, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_Font(style int) *qt6.QFont {

	_goptr := qt6.UnsafeNewQFont(unsafe.Pointer(C.QsciLexerTeX_virtualbase_font(unsafe.Pointer(this.h), (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QsciLexerTeX) OnFont(slot func(super func(style int) *qt6.QFont, style int) *qt6.QFont) {
	ok := C.QsciLexerTeX_override_virtual_font(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_font
func miqt_exec_callback_QsciLexerTeX_font(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) *C.QFont {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) *qt6.QFont, style int) *qt6.QFont)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_Font, slotval1)

	return (*C.QFont)(virtualReturn.UnsafePointer())

}

func (this *QsciLexerTeX) callVirtualBase_IndentationGuideView() int {

	return (int)(C.QsciLexerTeX_virtualbase_indentationGuideView(unsafe.Pointer(this.h)))

}
func (this *QsciLexerTeX) OnIndentationGuideView(slot func(super func() int) int) {
	ok := C.QsciLexerTeX_override_virtual_indentationGuideView(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_indentationGuideView
func miqt_exec_callback_QsciLexerTeX_indentationGuideView(self *C.QsciLexerTeX, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_IndentationGuideView)

	return (C.int)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_Keywords(set int) string {

	_ret := C.QsciLexerTeX_virtualbase_keywords(unsafe.Pointer(this.h), (C.int)(set))
	return C.GoString(_ret)

}
func (this *QsciLexerTeX) OnKeywords(slot func(super func(set int) string, set int) string) {
	ok := C.QsciLexerTeX_override_virtual_keywords(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_keywords
func miqt_exec_callback_QsciLexerTeX_keywords(self *C.QsciLexerTeX, cb C.intptr_t, set C.int) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(set int) string, set int) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(set)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_Keywords, slotval1)
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_DefaultStyle() int {

	return (int)(C.QsciLexerTeX_virtualbase_defaultStyle(unsafe.Pointer(this.h)))

}
func (this *QsciLexerTeX) OnDefaultStyle(slot func(super func() int) int) {
	ok := C.QsciLexerTeX_override_virtual_defaultStyle(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_defaultStyle
func miqt_exec_callback_QsciLexerTeX_defaultStyle(self *C.QsciLexerTeX, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_DefaultStyle)

	return (C.int)(virtualReturn)

}
func (this *QsciLexerTeX) OnDescription(slot func(style int) string) {
	ok := C.QsciLexerTeX_override_virtual_description(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_description
func miqt_exec_callback_QsciLexerTeX_description(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) C.struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(style int) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc(slotval1)
	virtualReturn_ms := C.struct_miqt_string{}
	virtualReturn_ms.data = C.CString(virtualReturn)
	virtualReturn_ms.len = C.size_t(len(virtualReturn))
	defer C.free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QsciLexerTeX) callVirtualBase_Paper(style int) *qt6.QColor {

	_goptr := qt6.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerTeX_virtualbase_paper(unsafe.Pointer(this.h), (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QsciLexerTeX) OnPaper(slot func(super func(style int) *qt6.QColor, style int) *qt6.QColor) {
	ok := C.QsciLexerTeX_override_virtual_paper(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_paper
func miqt_exec_callback_QsciLexerTeX_paper(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) *C.QColor {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) *qt6.QColor, style int) *qt6.QColor)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_Paper, slotval1)

	return (*C.QColor)(virtualReturn.UnsafePointer())

}

func (this *QsciLexerTeX) callVirtualBase_DefaultColorWithStyle(style int) *qt6.QColor {

	_goptr := qt6.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerTeX_virtualbase_defaultColorWithStyle(unsafe.Pointer(this.h), (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QsciLexerTeX) OnDefaultColorWithStyle(slot func(super func(style int) *qt6.QColor, style int) *qt6.QColor) {
	ok := C.QsciLexerTeX_override_virtual_defaultColorWithStyle(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_defaultColorWithStyle
func miqt_exec_callback_QsciLexerTeX_defaultColorWithStyle(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) *C.QColor {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) *qt6.QColor, style int) *qt6.QColor)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_DefaultColorWithStyle, slotval1)

	return (*C.QColor)(virtualReturn.UnsafePointer())

}

func (this *QsciLexerTeX) callVirtualBase_DefaultEolFill(style int) bool {

	return (bool)(C.QsciLexerTeX_virtualbase_defaultEolFill(unsafe.Pointer(this.h), (C.int)(style)))

}
func (this *QsciLexerTeX) OnDefaultEolFill(slot func(super func(style int) bool, style int) bool) {
	ok := C.QsciLexerTeX_override_virtual_defaultEolFill(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_defaultEolFill
func miqt_exec_callback_QsciLexerTeX_defaultEolFill(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) bool, style int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_DefaultEolFill, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_DefaultFontWithStyle(style int) *qt6.QFont {

	_goptr := qt6.UnsafeNewQFont(unsafe.Pointer(C.QsciLexerTeX_virtualbase_defaultFontWithStyle(unsafe.Pointer(this.h), (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QsciLexerTeX) OnDefaultFontWithStyle(slot func(super func(style int) *qt6.QFont, style int) *qt6.QFont) {
	ok := C.QsciLexerTeX_override_virtual_defaultFontWithStyle(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_defaultFontWithStyle
func miqt_exec_callback_QsciLexerTeX_defaultFontWithStyle(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) *C.QFont {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) *qt6.QFont, style int) *qt6.QFont)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_DefaultFontWithStyle, slotval1)

	return (*C.QFont)(virtualReturn.UnsafePointer())

}

func (this *QsciLexerTeX) callVirtualBase_DefaultPaperWithStyle(style int) *qt6.QColor {

	_goptr := qt6.UnsafeNewQColor(unsafe.Pointer(C.QsciLexerTeX_virtualbase_defaultPaperWithStyle(unsafe.Pointer(this.h), (C.int)(style))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QsciLexerTeX) OnDefaultPaperWithStyle(slot func(super func(style int) *qt6.QColor, style int) *qt6.QColor) {
	ok := C.QsciLexerTeX_override_virtual_defaultPaperWithStyle(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_defaultPaperWithStyle
func miqt_exec_callback_QsciLexerTeX_defaultPaperWithStyle(self *C.QsciLexerTeX, cb C.intptr_t, style C.int) *C.QColor {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(style int) *qt6.QColor, style int) *qt6.QColor)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(style)

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_DefaultPaperWithStyle, slotval1)

	return (*C.QColor)(virtualReturn.UnsafePointer())

}

func (this *QsciLexerTeX) callVirtualBase_SetEditor(editor *QsciScintilla) {

	C.QsciLexerTeX_virtualbase_setEditor(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QsciLexerTeX) OnSetEditor(slot func(super func(editor *QsciScintilla), editor *QsciScintilla)) {
	ok := C.QsciLexerTeX_override_virtual_setEditor(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_setEditor
func miqt_exec_callback_QsciLexerTeX_setEditor(self *C.QsciLexerTeX, cb C.intptr_t, editor *C.QsciScintilla) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QsciScintilla), editor *QsciScintilla))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQsciScintilla(editor)

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_SetEditor, slotval1)

}

func (this *QsciLexerTeX) callVirtualBase_RefreshProperties() {

	C.QsciLexerTeX_virtualbase_refreshProperties(unsafe.Pointer(this.h))

}
func (this *QsciLexerTeX) OnRefreshProperties(slot func(super func())) {
	ok := C.QsciLexerTeX_override_virtual_refreshProperties(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_refreshProperties
func miqt_exec_callback_QsciLexerTeX_refreshProperties(self *C.QsciLexerTeX, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_RefreshProperties)

}

func (this *QsciLexerTeX) callVirtualBase_StyleBitsNeeded() int {

	return (int)(C.QsciLexerTeX_virtualbase_styleBitsNeeded(unsafe.Pointer(this.h)))

}
func (this *QsciLexerTeX) OnStyleBitsNeeded(slot func(super func() int) int) {
	ok := C.QsciLexerTeX_override_virtual_styleBitsNeeded(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_styleBitsNeeded
func miqt_exec_callback_QsciLexerTeX_styleBitsNeeded(self *C.QsciLexerTeX, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_StyleBitsNeeded)

	return (C.int)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_WordCharacters() string {

	_ret := C.QsciLexerTeX_virtualbase_wordCharacters(unsafe.Pointer(this.h))
	return C.GoString(_ret)

}
func (this *QsciLexerTeX) OnWordCharacters(slot func(super func() string) string) {
	ok := C.QsciLexerTeX_override_virtual_wordCharacters(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_wordCharacters
func miqt_exec_callback_QsciLexerTeX_wordCharacters(self *C.QsciLexerTeX, cb C.intptr_t) *C.const_char {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_WordCharacters)
	virtualReturn_Cstring := C.CString(virtualReturn)
	defer C.free(unsafe.Pointer(virtualReturn_Cstring))

	return virtualReturn_Cstring

}

func (this *QsciLexerTeX) callVirtualBase_SetAutoIndentStyle(autoindentstyle int) {

	C.QsciLexerTeX_virtualbase_setAutoIndentStyle(unsafe.Pointer(this.h), (C.int)(autoindentstyle))

}
func (this *QsciLexerTeX) OnSetAutoIndentStyle(slot func(super func(autoindentstyle int), autoindentstyle int)) {
	ok := C.QsciLexerTeX_override_virtual_setAutoIndentStyle(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_setAutoIndentStyle
func miqt_exec_callback_QsciLexerTeX_setAutoIndentStyle(self *C.QsciLexerTeX, cb C.intptr_t, autoindentstyle C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(autoindentstyle int), autoindentstyle int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(autoindentstyle)

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_SetAutoIndentStyle, slotval1)

}

func (this *QsciLexerTeX) callVirtualBase_SetColor(c *qt6.QColor, style int) {

	C.QsciLexerTeX_virtualbase_setColor(unsafe.Pointer(this.h), (*C.QColor)(c.UnsafePointer()), (C.int)(style))

}
func (this *QsciLexerTeX) OnSetColor(slot func(super func(c *qt6.QColor, style int), c *qt6.QColor, style int)) {
	ok := C.QsciLexerTeX_override_virtual_setColor(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_setColor
func miqt_exec_callback_QsciLexerTeX_setColor(self *C.QsciLexerTeX, cb C.intptr_t, c *C.QColor, style C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(c *qt6.QColor, style int), c *qt6.QColor, style int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQColor(unsafe.Pointer(c))

	slotval2 := (int)(style)

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_SetColor, slotval1, slotval2)

}

func (this *QsciLexerTeX) callVirtualBase_SetEolFill(eoffill bool, style int) {

	C.QsciLexerTeX_virtualbase_setEolFill(unsafe.Pointer(this.h), (C.bool)(eoffill), (C.int)(style))

}
func (this *QsciLexerTeX) OnSetEolFill(slot func(super func(eoffill bool, style int), eoffill bool, style int)) {
	ok := C.QsciLexerTeX_override_virtual_setEolFill(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_setEolFill
func miqt_exec_callback_QsciLexerTeX_setEolFill(self *C.QsciLexerTeX, cb C.intptr_t, eoffill C.bool, style C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eoffill bool, style int), eoffill bool, style int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(eoffill)

	slotval2 := (int)(style)

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_SetEolFill, slotval1, slotval2)

}

func (this *QsciLexerTeX) callVirtualBase_SetFont(f *qt6.QFont, style int) {

	C.QsciLexerTeX_virtualbase_setFont(unsafe.Pointer(this.h), (*C.QFont)(f.UnsafePointer()), (C.int)(style))

}
func (this *QsciLexerTeX) OnSetFont(slot func(super func(f *qt6.QFont, style int), f *qt6.QFont, style int)) {
	ok := C.QsciLexerTeX_override_virtual_setFont(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_setFont
func miqt_exec_callback_QsciLexerTeX_setFont(self *C.QsciLexerTeX, cb C.intptr_t, f *C.QFont, style C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(f *qt6.QFont, style int), f *qt6.QFont, style int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQFont(unsafe.Pointer(f))

	slotval2 := (int)(style)

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_SetFont, slotval1, slotval2)

}

func (this *QsciLexerTeX) callVirtualBase_SetPaper(c *qt6.QColor, style int) {

	C.QsciLexerTeX_virtualbase_setPaper(unsafe.Pointer(this.h), (*C.QColor)(c.UnsafePointer()), (C.int)(style))

}
func (this *QsciLexerTeX) OnSetPaper(slot func(super func(c *qt6.QColor, style int), c *qt6.QColor, style int)) {
	ok := C.QsciLexerTeX_override_virtual_setPaper(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_setPaper
func miqt_exec_callback_QsciLexerTeX_setPaper(self *C.QsciLexerTeX, cb C.intptr_t, c *C.QColor, style C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(c *qt6.QColor, style int), c *qt6.QColor, style int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQColor(unsafe.Pointer(c))

	slotval2 := (int)(style)

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_SetPaper, slotval1, slotval2)

}

func (this *QsciLexerTeX) callVirtualBase_ReadProperties(qs *qt6.QSettings, prefix string) bool {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))

	return (bool)(C.QsciLexerTeX_virtualbase_readProperties(unsafe.Pointer(this.h), (*C.QSettings)(qs.UnsafePointer()), prefix_ms))

}
func (this *QsciLexerTeX) OnReadProperties(slot func(super func(qs *qt6.QSettings, prefix string) bool, qs *qt6.QSettings, prefix string) bool) {
	ok := C.QsciLexerTeX_override_virtual_readProperties(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_readProperties
func miqt_exec_callback_QsciLexerTeX_readProperties(self *C.QsciLexerTeX, cb C.intptr_t, qs *C.QSettings, prefix C.struct_miqt_string) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(qs *qt6.QSettings, prefix string) bool, qs *qt6.QSettings, prefix string) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQSettings(unsafe.Pointer(qs))

	var prefix_ms C.struct_miqt_string = prefix
	prefix_ret := C.GoStringN(prefix_ms.data, C.int(int64(prefix_ms.len)))
	C.free(unsafe.Pointer(prefix_ms.data))
	slotval2 := prefix_ret

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_ReadProperties, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_WriteProperties(qs *qt6.QSettings, prefix string) bool {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))

	return (bool)(C.QsciLexerTeX_virtualbase_writeProperties(unsafe.Pointer(this.h), (*C.QSettings)(qs.UnsafePointer()), prefix_ms))

}
func (this *QsciLexerTeX) OnWriteProperties(slot func(super func(qs *qt6.QSettings, prefix string) bool, qs *qt6.QSettings, prefix string) bool) {
	ok := C.QsciLexerTeX_override_virtual_writeProperties(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_writeProperties
func miqt_exec_callback_QsciLexerTeX_writeProperties(self *C.QsciLexerTeX, cb C.intptr_t, qs *C.QSettings, prefix C.struct_miqt_string) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(qs *qt6.QSettings, prefix string) bool, qs *qt6.QSettings, prefix string) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQSettings(unsafe.Pointer(qs))

	var prefix_ms C.struct_miqt_string = prefix
	prefix_ret := C.GoStringN(prefix_ms.data, C.int(int64(prefix_ms.len)))
	C.free(unsafe.Pointer(prefix_ms.data))
	slotval2 := prefix_ret

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_WriteProperties, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_Event(event *qt6.QEvent) bool {

	return (bool)(C.QsciLexerTeX_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QsciLexerTeX) OnEvent(slot func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool) {
	ok := C.QsciLexerTeX_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_event
func miqt_exec_callback_QsciLexerTeX_event(self *C.QsciLexerTeX, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_EventFilter(watched *qt6.QObject, event *qt6.QEvent) bool {

	return (bool)(C.QsciLexerTeX_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QsciLexerTeX) OnEventFilter(slot func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool) {
	ok := C.QsciLexerTeX_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_eventFilter
func miqt_exec_callback_QsciLexerTeX_eventFilter(self *C.QsciLexerTeX, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QsciLexerTeX{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QsciLexerTeX) callVirtualBase_TimerEvent(event *qt6.QTimerEvent) {

	C.QsciLexerTeX_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QsciLexerTeX) OnTimerEvent(slot func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent)) {
	ok := C.QsciLexerTeX_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_timerEvent
func miqt_exec_callback_QsciLexerTeX_timerEvent(self *C.QsciLexerTeX, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QsciLexerTeX) callVirtualBase_ChildEvent(event *qt6.QChildEvent) {

	C.QsciLexerTeX_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QsciLexerTeX) OnChildEvent(slot func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent)) {
	ok := C.QsciLexerTeX_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_childEvent
func miqt_exec_callback_QsciLexerTeX_childEvent(self *C.QsciLexerTeX, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QsciLexerTeX) callVirtualBase_CustomEvent(event *qt6.QEvent) {

	C.QsciLexerTeX_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QsciLexerTeX) OnCustomEvent(slot func(super func(event *qt6.QEvent), event *qt6.QEvent)) {
	ok := C.QsciLexerTeX_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_customEvent
func miqt_exec_callback_QsciLexerTeX_customEvent(self *C.QsciLexerTeX, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent), event *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QsciLexerTeX) callVirtualBase_ConnectNotify(signal *qt6.QMetaMethod) {

	C.QsciLexerTeX_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QsciLexerTeX) OnConnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QsciLexerTeX_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_connectNotify
func miqt_exec_callback_QsciLexerTeX_connectNotify(self *C.QsciLexerTeX, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QsciLexerTeX) callVirtualBase_DisconnectNotify(signal *qt6.QMetaMethod) {

	C.QsciLexerTeX_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QsciLexerTeX) OnDisconnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QsciLexerTeX_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QsciLexerTeX_disconnectNotify
func miqt_exec_callback_QsciLexerTeX_disconnectNotify(self *C.QsciLexerTeX, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QsciLexerTeX{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QsciLexerTeX) Delete() {
	C.QsciLexerTeX_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QsciLexerTeX) GoGC() {
	runtime.SetFinalizer(this, func(this *QsciLexerTeX) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
