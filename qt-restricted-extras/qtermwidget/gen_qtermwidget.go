package qtermwidget

/*

#include "gen_qtermwidget.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/e1z0/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QTermWidget__ScrollBarPosition int

const (
	QTermWidget__NoScrollBar    QTermWidget__ScrollBarPosition = 0
	QTermWidget__ScrollBarLeft  QTermWidget__ScrollBarPosition = 1
	QTermWidget__ScrollBarRight QTermWidget__ScrollBarPosition = 2
)

type QTermWidget struct {
	h *C.QTermWidget
	*qt.QWidget
}

func (this *QTermWidget) cPointer() *C.QTermWidget {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QTermWidget) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQTermWidget constructs the type using only CGO pointers.
func newQTermWidget(h *C.QTermWidget) *QTermWidget {
	if h == nil {
		return nil
	}
	var outptr_QWidget *C.QWidget = nil
	C.QTermWidget_virtbase(h, &outptr_QWidget)

	return &QTermWidget{h: h,
		QWidget: qt.UnsafeNewQWidget(unsafe.Pointer(outptr_QWidget))}
}

// UnsafeNewQTermWidget constructs the type using only unsafe pointers.
func UnsafeNewQTermWidget(h unsafe.Pointer) *QTermWidget {
	return newQTermWidget((*C.QTermWidget)(h))
}

// NewQTermWidget constructs a new QTermWidget object.
func NewQTermWidget(parent *qt.QWidget) *QTermWidget {

	return newQTermWidget(C.QTermWidget_new((*C.QWidget)(parent.UnsafePointer())))
}

// NewQTermWidget2 constructs a new QTermWidget object.
func NewQTermWidget2(startnow int) *QTermWidget {

	return newQTermWidget(C.QTermWidget_new2((C.int)(startnow)))
}

// NewQTermWidget3 constructs a new QTermWidget object.
func NewQTermWidget3() *QTermWidget {

	return newQTermWidget(C.QTermWidget_new3())
}

// NewQTermWidget4 constructs a new QTermWidget object.
func NewQTermWidget4(startnow int, parent *qt.QWidget) *QTermWidget {

	return newQTermWidget(C.QTermWidget_new4((C.int)(startnow), (*C.QWidget)(parent.UnsafePointer())))
}

func (this *QTermWidget) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QTermWidget_metaObject(this.h)))
}

func (this *QTermWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QTermWidget_metacast(this.h, param1_Cstring))
}

func QTermWidget_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QTermWidget_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTermWidget_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QTermWidget_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTermWidget) SizeHint() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(C.QTermWidget_sizeHint(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTermWidget) SetTerminalSizeHint(on bool) {
	C.QTermWidget_setTerminalSizeHint(this.h, (C.bool)(on))
}

func (this *QTermWidget) TerminalSizeHint() bool {
	return (bool)(C.QTermWidget_terminalSizeHint(this.h))
}

func (this *QTermWidget) StartShellProgram() {
	C.QTermWidget_startShellProgram(this.h)
}

func (this *QTermWidget) StartTerminalTeletype() {
	C.QTermWidget_startTerminalTeletype(this.h)
}

func (this *QTermWidget) GetShellPID() int {
	return (int)(C.QTermWidget_getShellPID(this.h))
}

func (this *QTermWidget) ChangeDir(dir string) {
	dir_ms := C.struct_miqt_string{}
	dir_ms.data = C.CString(dir)
	dir_ms.len = C.size_t(len(dir))
	defer C.free(unsafe.Pointer(dir_ms.data))
	C.QTermWidget_changeDir(this.h, dir_ms)
}

func (this *QTermWidget) SetTerminalFont(font *qt.QFont) {
	C.QTermWidget_setTerminalFont(this.h, (*C.QFont)(font.UnsafePointer()))
}

func (this *QTermWidget) GetTerminalFont() *qt.QFont {
	_goptr := qt.UnsafeNewQFont(unsafe.Pointer(C.QTermWidget_getTerminalFont(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTermWidget) SetTerminalOpacity(level float64) {
	C.QTermWidget_setTerminalOpacity(this.h, (C.double)(level))
}

func (this *QTermWidget) SetTerminalBackgroundImage(backgroundImage string) {
	backgroundImage_ms := C.struct_miqt_string{}
	backgroundImage_ms.data = C.CString(backgroundImage)
	backgroundImage_ms.len = C.size_t(len(backgroundImage))
	defer C.free(unsafe.Pointer(backgroundImage_ms.data))
	C.QTermWidget_setTerminalBackgroundImage(this.h, backgroundImage_ms)
}

func (this *QTermWidget) SetEnvironment(environment []string) {
	environment_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(environment))))
	defer C.free(unsafe.Pointer(environment_CArray))
	for i := range environment {
		environment_i_ms := C.struct_miqt_string{}
		environment_i_ms.data = C.CString(environment[i])
		environment_i_ms.len = C.size_t(len(environment[i]))
		defer C.free(unsafe.Pointer(environment_i_ms.data))
		environment_CArray[i] = environment_i_ms
	}
	environment_ma := C.struct_miqt_array{len: C.size_t(len(environment)), data: unsafe.Pointer(environment_CArray)}
	C.QTermWidget_setEnvironment(this.h, environment_ma)
}

func (this *QTermWidget) SetShellProgram(progname string) {
	progname_ms := C.struct_miqt_string{}
	progname_ms.data = C.CString(progname)
	progname_ms.len = C.size_t(len(progname))
	defer C.free(unsafe.Pointer(progname_ms.data))
	C.QTermWidget_setShellProgram(this.h, progname_ms)
}

func (this *QTermWidget) SetWorkingDirectory(dir string) {
	dir_ms := C.struct_miqt_string{}
	dir_ms.data = C.CString(dir)
	dir_ms.len = C.size_t(len(dir))
	defer C.free(unsafe.Pointer(dir_ms.data))
	C.QTermWidget_setWorkingDirectory(this.h, dir_ms)
}

func (this *QTermWidget) WorkingDirectory() string {
	var _ms C.struct_miqt_string = C.QTermWidget_workingDirectory(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTermWidget) SetArgs(args []string) {
	args_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(args))))
	defer C.free(unsafe.Pointer(args_CArray))
	for i := range args {
		args_i_ms := C.struct_miqt_string{}
		args_i_ms.data = C.CString(args[i])
		args_i_ms.len = C.size_t(len(args[i]))
		defer C.free(unsafe.Pointer(args_i_ms.data))
		args_CArray[i] = args_i_ms
	}
	args_ma := C.struct_miqt_array{len: C.size_t(len(args)), data: unsafe.Pointer(args_CArray)}
	C.QTermWidget_setArgs(this.h, args_ma)
}

func (this *QTermWidget) SetTextCodec(codec *qt.QTextCodec) {
	C.QTermWidget_setTextCodec(this.h, (*C.QTextCodec)(codec.UnsafePointer()))
}

func (this *QTermWidget) SetColorScheme(name string) {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	C.QTermWidget_setColorScheme(this.h, name_ms)
}

func QTermWidget_AvailableColorSchemes() []string {
	var _ma C.struct_miqt_array = C.QTermWidget_availableColorSchemes()
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

func QTermWidget_AddCustomColorSchemeDir(custom_dir string) {
	custom_dir_ms := C.struct_miqt_string{}
	custom_dir_ms.data = C.CString(custom_dir)
	custom_dir_ms.len = C.size_t(len(custom_dir))
	defer C.free(unsafe.Pointer(custom_dir_ms.data))
	C.QTermWidget_addCustomColorSchemeDir(custom_dir_ms)
}

func (this *QTermWidget) SetHistorySize(lines int) {
	C.QTermWidget_setHistorySize(this.h, (C.int)(lines))
}

func (this *QTermWidget) SetScrollBarPosition(scrollBarPosition QTermWidget__ScrollBarPosition) {
	C.QTermWidget_setScrollBarPosition(this.h, (C.int)(scrollBarPosition))
}

func (this *QTermWidget) ScrollToEnd() {
	C.QTermWidget_scrollToEnd(this.h)
}

func (this *QTermWidget) SendText(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QTermWidget_sendText(this.h, text_ms)
}

func (this *QTermWidget) SetFlowControlEnabled(enabled bool) {
	C.QTermWidget_setFlowControlEnabled(this.h, (C.bool)(enabled))
}

func (this *QTermWidget) FlowControlEnabled() bool {
	return (bool)(C.QTermWidget_flowControlEnabled(this.h))
}

func (this *QTermWidget) SetFlowControlWarningEnabled(enabled bool) {
	C.QTermWidget_setFlowControlWarningEnabled(this.h, (C.bool)(enabled))
}

func QTermWidget_AvailableKeyBindings() []string {
	var _ma C.struct_miqt_array = C.QTermWidget_availableKeyBindings()
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

func (this *QTermWidget) KeyBindings() string {
	var _ms C.struct_miqt_string = C.QTermWidget_keyBindings(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTermWidget) SetMotionAfterPasting(motionAfterPasting int) {
	C.QTermWidget_setMotionAfterPasting(this.h, (C.int)(motionAfterPasting))
}

func (this *QTermWidget) HistoryLinesCount() int {
	return (int)(C.QTermWidget_historyLinesCount(this.h))
}

func (this *QTermWidget) ScreenColumnsCount() int {
	return (int)(C.QTermWidget_screenColumnsCount(this.h))
}

func (this *QTermWidget) ScreenLinesCount() int {
	return (int)(C.QTermWidget_screenLinesCount(this.h))
}

func (this *QTermWidget) SetSelectionStart(row int, column int) {
	C.QTermWidget_setSelectionStart(this.h, (C.int)(row), (C.int)(column))
}

func (this *QTermWidget) SetSelectionEnd(row int, column int) {
	C.QTermWidget_setSelectionEnd(this.h, (C.int)(row), (C.int)(column))
}

func (this *QTermWidget) GetSelectionStart(row *int, column *int) {
	C.QTermWidget_getSelectionStart(this.h, (*C.int)(unsafe.Pointer(row)), (*C.int)(unsafe.Pointer(column)))
}

func (this *QTermWidget) GetSelectionEnd(row *int, column *int) {
	C.QTermWidget_getSelectionEnd(this.h, (*C.int)(unsafe.Pointer(row)), (*C.int)(unsafe.Pointer(column)))
}

func (this *QTermWidget) SelectedText() string {
	var _ms C.struct_miqt_string = C.QTermWidget_selectedText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTermWidget) SetMonitorActivity(monitorActivity bool) {
	C.QTermWidget_setMonitorActivity(this.h, (C.bool)(monitorActivity))
}

func (this *QTermWidget) SetMonitorSilence(monitorSilence bool) {
	C.QTermWidget_setMonitorSilence(this.h, (C.bool)(monitorSilence))
}

func (this *QTermWidget) SetSilenceTimeout(seconds int) {
	C.QTermWidget_setSilenceTimeout(this.h, (C.int)(seconds))
}

func (this *QTermWidget) GetHotSpotAt(pos *qt.QPoint) *Filter__HotSpot {
	int /* TODO  */
}

func (this *QTermWidget) GetHotSpotAt2(row int, column int) *Filter__HotSpot {
	int /* TODO  */
}

func (this *QTermWidget) FilterActions(position *qt.QPoint) []*qt.QAction {
	var _ma C.struct_miqt_array = C.QTermWidget_filterActions(this.h, (*C.QPoint)(position.UnsafePointer()))
	_ret := make([]*qt.QAction, int(_ma.len))
	_outCast := (*[0xffff]*C.QAction)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = qt.UnsafeNewQAction(unsafe.Pointer(_outCast[i]))
	}
	return _ret
}

func (this *QTermWidget) GetPtySlaveFd() int {
	return (int)(C.QTermWidget_getPtySlaveFd(this.h))
}

func (this *QTermWidget) SetKeyboardCursorShape(shape Konsole__Emulation__KeyboardCursorShape) {
	C.QTermWidget_setKeyboardCursorShape(this.h, shape)
}

func (this *QTermWidget) SetBlinkingCursor(blink bool) {
	C.QTermWidget_setBlinkingCursor(this.h, (C.bool)(blink))
}

func (this *QTermWidget) SetBidiEnabled(enabled bool) {
	C.QTermWidget_setBidiEnabled(this.h, (C.bool)(enabled))
}

func (this *QTermWidget) IsBidiEnabled() bool {
	return (bool)(C.QTermWidget_isBidiEnabled(this.h))
}

func (this *QTermWidget) SetAutoClose(autoClose bool) {
	C.QTermWidget_setAutoClose(this.h, (C.bool)(autoClose))
}

func (this *QTermWidget) Title() string {
	var _ms C.struct_miqt_string = C.QTermWidget_title(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTermWidget) Icon() string {
	var _ms C.struct_miqt_string = C.QTermWidget_icon(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTermWidget) IsTitleChanged() bool {
	return (bool)(C.QTermWidget_isTitleChanged(this.h))
}

func (this *QTermWidget) BracketText(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QTermWidget_bracketText(this.h, text_ms)
}

func (this *QTermWidget) Finished() {
	C.QTermWidget_finished(this.h)
}
func (this *QTermWidget) OnFinished(slot func()) {
	C.QTermWidget_connect_finished(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_finished
func miqt_exec_callback_QTermWidget_finished(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTermWidget) CopyAvailable(param1 bool) {
	C.QTermWidget_copyAvailable(this.h, (C.bool)(param1))
}
func (this *QTermWidget) OnCopyAvailable(slot func(param1 bool)) {
	C.QTermWidget_connect_copyAvailable(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_copyAvailable
func miqt_exec_callback_QTermWidget_copyAvailable(cb C.intptr_t, param1 C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QTermWidget) TermGetFocus() {
	C.QTermWidget_termGetFocus(this.h)
}
func (this *QTermWidget) OnTermGetFocus(slot func()) {
	C.QTermWidget_connect_termGetFocus(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_termGetFocus
func miqt_exec_callback_QTermWidget_termGetFocus(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTermWidget) TermLostFocus() {
	C.QTermWidget_termLostFocus(this.h)
}
func (this *QTermWidget) OnTermLostFocus(slot func()) {
	C.QTermWidget_connect_termLostFocus(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_termLostFocus
func miqt_exec_callback_QTermWidget_termLostFocus(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTermWidget) TermKeyPressed(param1 *qt.QKeyEvent) {
	C.QTermWidget_termKeyPressed(this.h, (*C.QKeyEvent)(param1.UnsafePointer()))
}
func (this *QTermWidget) OnTermKeyPressed(slot func(param1 *qt.QKeyEvent)) {
	C.QTermWidget_connect_termKeyPressed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_termKeyPressed
func miqt_exec_callback_QTermWidget_termKeyPressed(cb C.intptr_t, param1 *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QTermWidget) UrlActivated(param1 *qt.QUrl, fromContextMenu bool) {
	C.QTermWidget_urlActivated(this.h, (*C.QUrl)(param1.UnsafePointer()), (C.bool)(fromContextMenu))
}
func (this *QTermWidget) OnUrlActivated(slot func(param1 *qt.QUrl, fromContextMenu bool)) {
	C.QTermWidget_connect_urlActivated(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_urlActivated
func miqt_exec_callback_QTermWidget_urlActivated(cb C.intptr_t, param1 *C.QUrl, fromContextMenu C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt.QUrl, fromContextMenu bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(param1))

	slotval2 := (bool)(fromContextMenu)

	gofunc(slotval1, slotval2)
}

func (this *QTermWidget) Bell(message string) {
	message_ms := C.struct_miqt_string{}
	message_ms.data = C.CString(message)
	message_ms.len = C.size_t(len(message))
	defer C.free(unsafe.Pointer(message_ms.data))
	C.QTermWidget_bell(this.h, message_ms)
}
func (this *QTermWidget) OnBell(slot func(message string)) {
	C.QTermWidget_connect_bell(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_bell
func miqt_exec_callback_QTermWidget_bell(cb C.intptr_t, message C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(message string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var message_ms C.struct_miqt_string = message
	message_ret := C.GoStringN(message_ms.data, C.int(int64(message_ms.len)))
	C.free(unsafe.Pointer(message_ms.data))
	slotval1 := message_ret

	gofunc(slotval1)
}

func (this *QTermWidget) Activity() {
	C.QTermWidget_activity(this.h)
}
func (this *QTermWidget) OnActivity(slot func()) {
	C.QTermWidget_connect_activity(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_activity
func miqt_exec_callback_QTermWidget_activity(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTermWidget) Silence() {
	C.QTermWidget_silence(this.h)
}
func (this *QTermWidget) OnSilence(slot func()) {
	C.QTermWidget_connect_silence(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_silence
func miqt_exec_callback_QTermWidget_silence(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTermWidget) SendData(param1 string, param2 int) {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	C.QTermWidget_sendData(this.h, param1_Cstring, (C.int)(param2))
}
func (this *QTermWidget) OnSendData(slot func(param1 string, param2 int)) {
	C.QTermWidget_connect_sendData(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_sendData
func miqt_exec_callback_QTermWidget_sendData(cb C.intptr_t, param1 *C.const_char, param2 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string, param2 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := C.GoString(param1_ret)

	slotval2 := (int)(param2)

	gofunc(slotval1, slotval2)
}

func (this *QTermWidget) ProfileChanged(profile string) {
	profile_ms := C.struct_miqt_string{}
	profile_ms.data = C.CString(profile)
	profile_ms.len = C.size_t(len(profile))
	defer C.free(unsafe.Pointer(profile_ms.data))
	C.QTermWidget_profileChanged(this.h, profile_ms)
}
func (this *QTermWidget) OnProfileChanged(slot func(profile string)) {
	C.QTermWidget_connect_profileChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_profileChanged
func miqt_exec_callback_QTermWidget_profileChanged(cb C.intptr_t, profile C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(profile string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var profile_ms C.struct_miqt_string = profile
	profile_ret := C.GoStringN(profile_ms.data, C.int(int64(profile_ms.len)))
	C.free(unsafe.Pointer(profile_ms.data))
	slotval1 := profile_ret

	gofunc(slotval1)
}

func (this *QTermWidget) TitleChanged() {
	C.QTermWidget_titleChanged(this.h)
}
func (this *QTermWidget) OnTitleChanged(slot func()) {
	C.QTermWidget_connect_titleChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_titleChanged
func miqt_exec_callback_QTermWidget_titleChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTermWidget) RightMouseButtonPressed() {
	C.QTermWidget_rightMouseButtonPressed(this.h)
}
func (this *QTermWidget) OnRightMouseButtonPressed(slot func()) {
	C.QTermWidget_connect_rightMouseButtonPressed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_rightMouseButtonPressed
func miqt_exec_callback_QTermWidget_rightMouseButtonPressed(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTermWidget) ReceivedData(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QTermWidget_receivedData(this.h, text_ms)
}
func (this *QTermWidget) OnReceivedData(slot func(text string)) {
	C.QTermWidget_connect_receivedData(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTermWidget_receivedData
func miqt_exec_callback_QTermWidget_receivedData(cb C.intptr_t, text C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms C.struct_miqt_string = text
	text_ret := C.GoStringN(text_ms.data, C.int(int64(text_ms.len)))
	C.free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)
}

func (this *QTermWidget) ReceivedRemoteData(data []byte) {
	data_alias := C.struct_miqt_string{}
	if len(data) > 0 {
		data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		data_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	data_alias.len = C.size_t(len(data))
	C.QTermWidget_receivedRemoteData(this.h, data_alias)
}

func (this *QTermWidget) CopyClipboard() {
	C.QTermWidget_copyClipboard(this.h)
}

func (this *QTermWidget) PasteClipboard() {
	C.QTermWidget_pasteClipboard(this.h)
}

func (this *QTermWidget) PasteSelection() {
	C.QTermWidget_pasteSelection(this.h)
}

func (this *QTermWidget) ZoomIn() {
	C.QTermWidget_zoomIn(this.h)
}

func (this *QTermWidget) ZoomOut() {
	C.QTermWidget_zoomOut(this.h)
}

func (this *QTermWidget) SetSize(size *qt.QSize) {
	C.QTermWidget_setSize(this.h, (*C.QSize)(size.UnsafePointer()))
}

func (this *QTermWidget) SetKeyBindings(kb string) {
	kb_ms := C.struct_miqt_string{}
	kb_ms.data = C.CString(kb)
	kb_ms.len = C.size_t(len(kb))
	defer C.free(unsafe.Pointer(kb_ms.data))
	C.QTermWidget_setKeyBindings(this.h, kb_ms)
}

func (this *QTermWidget) Clear() {
	C.QTermWidget_clear(this.h)
}

func (this *QTermWidget) ToggleShowSearchBar() {
	C.QTermWidget_toggleShowSearchBar(this.h)
}

func (this *QTermWidget) MouseEventEmiter() {
	C.QTermWidget_mouseEventEmiter(this.h)
}

func QTermWidget_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTermWidget_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTermWidget_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTermWidget_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTermWidget_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTermWidget_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTermWidget_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QTermWidget_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTermWidget) SelectedTextWithPreserveLineBreaks(preserveLineBreaks bool) string {
	var _ms C.struct_miqt_string = C.QTermWidget_selectedTextWithPreserveLineBreaks(this.h, (C.bool)(preserveLineBreaks))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// SessionFinished can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) SessionFinished() {

	var _dynamic_cast_ok C.bool = false
	C.QTermWidget_protectedbase_sessionFinished(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// SelectionChanged can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) SelectionChanged(textSelected bool) {

	var _dynamic_cast_ok C.bool = false
	C.QTermWidget_protectedbase_selectionChanged(&_dynamic_cast_ok, unsafe.Pointer(this.h), (C.bool)(textSelected))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// UpdateMicroFocus can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) UpdateMicroFocus() {

	var _dynamic_cast_ok C.bool = false
	C.QTermWidget_protectedbase_updateMicroFocus(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// Create can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) Create() {

	var _dynamic_cast_ok C.bool = false
	C.QTermWidget_protectedbase_create(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// Destroy can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) Destroy() {

	var _dynamic_cast_ok C.bool = false
	C.QTermWidget_protectedbase_destroy(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// FocusNextChild can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) FocusNextChild() bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QTermWidget_protectedbase_focusNextChild(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// FocusPreviousChild can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) FocusPreviousChild() bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QTermWidget_protectedbase_focusPreviousChild(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Sender can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) Sender() *qt.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt.UnsafeNewQObject(unsafe.Pointer(C.QTermWidget_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QTermWidget_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QTermWidget_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QTermWidget that was directly constructed.
func (this *QTermWidget) IsSignalConnected(signal *qt.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QTermWidget_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QTermWidget) callVirtualBase_ResizeEvent(param1 *qt.QResizeEvent) {

	C.QTermWidget_virtualbase_resizeEvent(unsafe.Pointer(this.h), (*C.QResizeEvent)(param1.UnsafePointer()))

}
func (this *QTermWidget) OnResizeEvent(slot func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent)) {
	ok := C.QTermWidget_override_virtual_resizeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_resizeEvent
func miqt_exec_callback_QTermWidget_resizeEvent(self *C.QTermWidget, cb C.intptr_t, param1 *C.QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQResizeEvent(unsafe.Pointer(param1))

	gofunc((&QTermWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_KeyPressEvent(event *qt.QKeyEvent) {

	C.QTermWidget_virtualbase_keyPressEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnKeyPressEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	ok := C.QTermWidget_override_virtual_keyPressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_keyPressEvent
func miqt_exec_callback_QTermWidget_keyPressEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_DevType() int {

	return (int)(C.QTermWidget_virtualbase_devType(unsafe.Pointer(this.h)))

}
func (this *QTermWidget) OnDevType(slot func(super func() int) int) {
	ok := C.QTermWidget_override_virtual_devType(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_devType
func miqt_exec_callback_QTermWidget_devType(self *C.QTermWidget, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_DevType)

	return (C.int)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_SetVisible(visible bool) {

	C.QTermWidget_virtualbase_setVisible(unsafe.Pointer(this.h), (C.bool)(visible))

}
func (this *QTermWidget) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	ok := C.QTermWidget_override_virtual_setVisible(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_setVisible
func miqt_exec_callback_QTermWidget_setVisible(self *C.QTermWidget, cb C.intptr_t, visible C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QTermWidget{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QTermWidget) callVirtualBase_SizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(C.QTermWidget_virtualbase_sizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTermWidget) OnSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	ok := C.QTermWidget_override_virtual_sizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_sizeHint
func miqt_exec_callback_QTermWidget_sizeHint(self *C.QTermWidget, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_SizeHint)

	return (*C.QSize)(virtualReturn.UnsafePointer())

}

func (this *QTermWidget) callVirtualBase_MinimumSizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(C.QTermWidget_virtualbase_minimumSizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTermWidget) OnMinimumSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	ok := C.QTermWidget_override_virtual_minimumSizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_minimumSizeHint
func miqt_exec_callback_QTermWidget_minimumSizeHint(self *C.QTermWidget, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_MinimumSizeHint)

	return (*C.QSize)(virtualReturn.UnsafePointer())

}

func (this *QTermWidget) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(C.QTermWidget_virtualbase_heightForWidth(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QTermWidget) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	ok := C.QTermWidget_override_virtual_heightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_heightForWidth
func miqt_exec_callback_QTermWidget_heightForWidth(self *C.QTermWidget, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(C.QTermWidget_virtualbase_hasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QTermWidget) OnHasHeightForWidth(slot func(super func() bool) bool) {
	ok := C.QTermWidget_override_virtual_hasHeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_hasHeightForWidth
func miqt_exec_callback_QTermWidget_hasHeightForWidth(self *C.QTermWidget, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_HasHeightForWidth)

	return (C.bool)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_PaintEngine() *qt.QPaintEngine {

	return qt.UnsafeNewQPaintEngine(unsafe.Pointer(C.QTermWidget_virtualbase_paintEngine(unsafe.Pointer(this.h))))

}
func (this *QTermWidget) OnPaintEngine(slot func(super func() *qt.QPaintEngine) *qt.QPaintEngine) {
	ok := C.QTermWidget_override_virtual_paintEngine(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_paintEngine
func miqt_exec_callback_QTermWidget_paintEngine(self *C.QTermWidget, cb C.intptr_t) *C.QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPaintEngine) *qt.QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_PaintEngine)

	return (*C.QPaintEngine)(virtualReturn.UnsafePointer())

}

func (this *QTermWidget) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(C.QTermWidget_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QTermWidget) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	ok := C.QTermWidget_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_event
func miqt_exec_callback_QTermWidget_event(self *C.QTermWidget, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_MousePressEvent(event *qt.QMouseEvent) {

	C.QTermWidget_virtualbase_mousePressEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnMousePressEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	ok := C.QTermWidget_override_virtual_mousePressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_mousePressEvent
func miqt_exec_callback_QTermWidget_mousePressEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_MouseReleaseEvent(event *qt.QMouseEvent) {

	C.QTermWidget_virtualbase_mouseReleaseEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnMouseReleaseEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	ok := C.QTermWidget_override_virtual_mouseReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_mouseReleaseEvent
func miqt_exec_callback_QTermWidget_mouseReleaseEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_MouseDoubleClickEvent(event *qt.QMouseEvent) {

	C.QTermWidget_virtualbase_mouseDoubleClickEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnMouseDoubleClickEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	ok := C.QTermWidget_override_virtual_mouseDoubleClickEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_mouseDoubleClickEvent
func miqt_exec_callback_QTermWidget_mouseDoubleClickEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_MouseMoveEvent(event *qt.QMouseEvent) {

	C.QTermWidget_virtualbase_mouseMoveEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnMouseMoveEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	ok := C.QTermWidget_override_virtual_mouseMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_mouseMoveEvent
func miqt_exec_callback_QTermWidget_mouseMoveEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_WheelEvent(event *qt.QWheelEvent) {

	C.QTermWidget_virtualbase_wheelEvent(unsafe.Pointer(this.h), (*C.QWheelEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnWheelEvent(slot func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent)) {
	ok := C.QTermWidget_override_virtual_wheelEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_wheelEvent
func miqt_exec_callback_QTermWidget_wheelEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQWheelEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_KeyReleaseEvent(event *qt.QKeyEvent) {

	C.QTermWidget_virtualbase_keyReleaseEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnKeyReleaseEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	ok := C.QTermWidget_override_virtual_keyReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_keyReleaseEvent
func miqt_exec_callback_QTermWidget_keyReleaseEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_FocusInEvent(event *qt.QFocusEvent) {

	C.QTermWidget_virtualbase_focusInEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnFocusInEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	ok := C.QTermWidget_override_virtual_focusInEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_focusInEvent
func miqt_exec_callback_QTermWidget_focusInEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_FocusOutEvent(event *qt.QFocusEvent) {

	C.QTermWidget_virtualbase_focusOutEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnFocusOutEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	ok := C.QTermWidget_override_virtual_focusOutEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_focusOutEvent
func miqt_exec_callback_QTermWidget_focusOutEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_EnterEvent(event *qt.QEvent) {

	C.QTermWidget_virtualbase_enterEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnEnterEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	ok := C.QTermWidget_override_virtual_enterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_enterEvent
func miqt_exec_callback_QTermWidget_enterEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_LeaveEvent(event *qt.QEvent) {

	C.QTermWidget_virtualbase_leaveEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnLeaveEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	ok := C.QTermWidget_override_virtual_leaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_leaveEvent
func miqt_exec_callback_QTermWidget_leaveEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_PaintEvent(event *qt.QPaintEvent) {

	C.QTermWidget_virtualbase_paintEvent(unsafe.Pointer(this.h), (*C.QPaintEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnPaintEvent(slot func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent)) {
	ok := C.QTermWidget_override_virtual_paintEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_paintEvent
func miqt_exec_callback_QTermWidget_paintEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPaintEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_MoveEvent(event *qt.QMoveEvent) {

	C.QTermWidget_virtualbase_moveEvent(unsafe.Pointer(this.h), (*C.QMoveEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnMoveEvent(slot func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent)) {
	ok := C.QTermWidget_override_virtual_moveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_moveEvent
func miqt_exec_callback_QTermWidget_moveEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMoveEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_CloseEvent(event *qt.QCloseEvent) {

	C.QTermWidget_virtualbase_closeEvent(unsafe.Pointer(this.h), (*C.QCloseEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnCloseEvent(slot func(super func(event *qt.QCloseEvent), event *qt.QCloseEvent)) {
	ok := C.QTermWidget_override_virtual_closeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_closeEvent
func miqt_exec_callback_QTermWidget_closeEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QCloseEvent), event *qt.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQCloseEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_ContextMenuEvent(event *qt.QContextMenuEvent) {

	C.QTermWidget_virtualbase_contextMenuEvent(unsafe.Pointer(this.h), (*C.QContextMenuEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnContextMenuEvent(slot func(super func(event *qt.QContextMenuEvent), event *qt.QContextMenuEvent)) {
	ok := C.QTermWidget_override_virtual_contextMenuEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_contextMenuEvent
func miqt_exec_callback_QTermWidget_contextMenuEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QContextMenuEvent), event *qt.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQContextMenuEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_TabletEvent(event *qt.QTabletEvent) {

	C.QTermWidget_virtualbase_tabletEvent(unsafe.Pointer(this.h), (*C.QTabletEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnTabletEvent(slot func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent)) {
	ok := C.QTermWidget_override_virtual_tabletEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_tabletEvent
func miqt_exec_callback_QTermWidget_tabletEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTabletEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_ActionEvent(event *qt.QActionEvent) {

	C.QTermWidget_virtualbase_actionEvent(unsafe.Pointer(this.h), (*C.QActionEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnActionEvent(slot func(super func(event *qt.QActionEvent), event *qt.QActionEvent)) {
	ok := C.QTermWidget_override_virtual_actionEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_actionEvent
func miqt_exec_callback_QTermWidget_actionEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QActionEvent), event *qt.QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQActionEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_DragEnterEvent(event *qt.QDragEnterEvent) {

	C.QTermWidget_virtualbase_dragEnterEvent(unsafe.Pointer(this.h), (*C.QDragEnterEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnDragEnterEvent(slot func(super func(event *qt.QDragEnterEvent), event *qt.QDragEnterEvent)) {
	ok := C.QTermWidget_override_virtual_dragEnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_dragEnterEvent
func miqt_exec_callback_QTermWidget_dragEnterEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragEnterEvent), event *qt.QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragEnterEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_DragMoveEvent(event *qt.QDragMoveEvent) {

	C.QTermWidget_virtualbase_dragMoveEvent(unsafe.Pointer(this.h), (*C.QDragMoveEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnDragMoveEvent(slot func(super func(event *qt.QDragMoveEvent), event *qt.QDragMoveEvent)) {
	ok := C.QTermWidget_override_virtual_dragMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_dragMoveEvent
func miqt_exec_callback_QTermWidget_dragMoveEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragMoveEvent), event *qt.QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragMoveEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_DragLeaveEvent(event *qt.QDragLeaveEvent) {

	C.QTermWidget_virtualbase_dragLeaveEvent(unsafe.Pointer(this.h), (*C.QDragLeaveEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnDragLeaveEvent(slot func(super func(event *qt.QDragLeaveEvent), event *qt.QDragLeaveEvent)) {
	ok := C.QTermWidget_override_virtual_dragLeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_dragLeaveEvent
func miqt_exec_callback_QTermWidget_dragLeaveEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragLeaveEvent), event *qt.QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_DropEvent(event *qt.QDropEvent) {

	C.QTermWidget_virtualbase_dropEvent(unsafe.Pointer(this.h), (*C.QDropEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnDropEvent(slot func(super func(event *qt.QDropEvent), event *qt.QDropEvent)) {
	ok := C.QTermWidget_override_virtual_dropEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_dropEvent
func miqt_exec_callback_QTermWidget_dropEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDropEvent), event *qt.QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDropEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_ShowEvent(event *qt.QShowEvent) {

	C.QTermWidget_virtualbase_showEvent(unsafe.Pointer(this.h), (*C.QShowEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnShowEvent(slot func(super func(event *qt.QShowEvent), event *qt.QShowEvent)) {
	ok := C.QTermWidget_override_virtual_showEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_showEvent
func miqt_exec_callback_QTermWidget_showEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QShowEvent), event *qt.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQShowEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_HideEvent(event *qt.QHideEvent) {

	C.QTermWidget_virtualbase_hideEvent(unsafe.Pointer(this.h), (*C.QHideEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnHideEvent(slot func(super func(event *qt.QHideEvent), event *qt.QHideEvent)) {
	ok := C.QTermWidget_override_virtual_hideEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_hideEvent
func miqt_exec_callback_QTermWidget_hideEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QHideEvent), event *qt.QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQHideEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *int64) bool {
	eventType_alias := C.struct_miqt_string{}
	if len(eventType) > 0 {
		eventType_alias.data = (*C.char)(unsafe.Pointer(&eventType[0]))
	} else {
		eventType_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	eventType_alias.len = C.size_t(len(eventType))

	return (bool)(C.QTermWidget_virtualbase_nativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*C.long)(unsafe.Pointer(result))))

}
func (this *QTermWidget) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *int64) bool, eventType []byte, message unsafe.Pointer, result *int64) bool) {
	ok := C.QTermWidget_override_virtual_nativeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_nativeEvent
func miqt_exec_callback_QTermWidget_nativeEvent(self *C.QTermWidget, cb C.intptr_t, eventType C.struct_miqt_string, message unsafe.Pointer, result *C.long) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *int64) bool, eventType []byte, message unsafe.Pointer, result *int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray C.struct_miqt_string = eventType
	eventType_ret := C.GoBytes(unsafe.Pointer(eventType_bytearray.data), C.int(int64(eventType_bytearray.len)))
	C.free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*int64)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (C.bool)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_ChangeEvent(param1 *qt.QEvent) {

	C.QTermWidget_virtualbase_changeEvent(unsafe.Pointer(this.h), (*C.QEvent)(param1.UnsafePointer()))

}
func (this *QTermWidget) OnChangeEvent(slot func(super func(param1 *qt.QEvent), param1 *qt.QEvent)) {
	ok := C.QTermWidget_override_virtual_changeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_changeEvent
func miqt_exec_callback_QTermWidget_changeEvent(self *C.QTermWidget, cb C.intptr_t, param1 *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QEvent), param1 *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(param1))

	gofunc((&QTermWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_Metric(param1 qt.QPaintDevice__PaintDeviceMetric) int {

	return (int)(C.QTermWidget_virtualbase_metric(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QTermWidget) OnMetric(slot func(super func(param1 qt.QPaintDevice__PaintDeviceMetric) int, param1 qt.QPaintDevice__PaintDeviceMetric) int) {
	ok := C.QTermWidget_override_virtual_metric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_metric
func miqt_exec_callback_QTermWidget_metric(self *C.QTermWidget, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt.QPaintDevice__PaintDeviceMetric) int, param1 qt.QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QPaintDevice__PaintDeviceMetric)(param1)

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_Metric, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_InitPainter(painter *qt.QPainter) {

	C.QTermWidget_virtualbase_initPainter(unsafe.Pointer(this.h), (*C.QPainter)(painter.UnsafePointer()))

}
func (this *QTermWidget) OnInitPainter(slot func(super func(painter *qt.QPainter), painter *qt.QPainter)) {
	ok := C.QTermWidget_override_virtual_initPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_initPainter
func miqt_exec_callback_QTermWidget_initPainter(self *C.QTermWidget, cb C.intptr_t, painter *C.QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter), painter *qt.QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))

	gofunc((&QTermWidget{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QTermWidget) callVirtualBase_Redirected(offset *qt.QPoint) *qt.QPaintDevice {

	return qt.UnsafeNewQPaintDevice(unsafe.Pointer(C.QTermWidget_virtualbase_redirected(unsafe.Pointer(this.h), (*C.QPoint)(offset.UnsafePointer()))))

}
func (this *QTermWidget) OnRedirected(slot func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice) {
	ok := C.QTermWidget_override_virtual_redirected(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_redirected
func miqt_exec_callback_QTermWidget_redirected(self *C.QTermWidget, cb C.intptr_t, offset *C.QPoint) *C.QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPoint(unsafe.Pointer(offset))

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_Redirected, slotval1)

	return (*C.QPaintDevice)(virtualReturn.UnsafePointer())

}

func (this *QTermWidget) callVirtualBase_SharedPainter() *qt.QPainter {

	return qt.UnsafeNewQPainter(unsafe.Pointer(C.QTermWidget_virtualbase_sharedPainter(unsafe.Pointer(this.h))))

}
func (this *QTermWidget) OnSharedPainter(slot func(super func() *qt.QPainter) *qt.QPainter) {
	ok := C.QTermWidget_override_virtual_sharedPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_sharedPainter
func miqt_exec_callback_QTermWidget_sharedPainter(self *C.QTermWidget, cb C.intptr_t) *C.QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPainter) *qt.QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_SharedPainter)

	return (*C.QPainter)(virtualReturn.UnsafePointer())

}

func (this *QTermWidget) callVirtualBase_InputMethodEvent(param1 *qt.QInputMethodEvent) {

	C.QTermWidget_virtualbase_inputMethodEvent(unsafe.Pointer(this.h), (*C.QInputMethodEvent)(param1.UnsafePointer()))

}
func (this *QTermWidget) OnInputMethodEvent(slot func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent)) {
	ok := C.QTermWidget_override_virtual_inputMethodEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_inputMethodEvent
func miqt_exec_callback_QTermWidget_inputMethodEvent(self *C.QTermWidget, cb C.intptr_t, param1 *C.QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQInputMethodEvent(unsafe.Pointer(param1))

	gofunc((&QTermWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_InputMethodQuery(param1 qt.InputMethodQuery) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(C.QTermWidget_virtualbase_inputMethodQuery(unsafe.Pointer(this.h), (C.int)(param1))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTermWidget) OnInputMethodQuery(slot func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant) {
	ok := C.QTermWidget_override_virtual_inputMethodQuery(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_inputMethodQuery
func miqt_exec_callback_QTermWidget_inputMethodQuery(self *C.QTermWidget, cb C.intptr_t, param1 C.int) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.InputMethodQuery)(param1)

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return (*C.QVariant)(virtualReturn.UnsafePointer())

}

func (this *QTermWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(C.QTermWidget_virtualbase_focusNextPrevChild(unsafe.Pointer(this.h), (C.bool)(next)))

}
func (this *QTermWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	ok := C.QTermWidget_override_virtual_focusNextPrevChild(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_focusNextPrevChild
func miqt_exec_callback_QTermWidget_focusNextPrevChild(self *C.QTermWidget, cb C.intptr_t, next C.bool) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(C.QTermWidget_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QTermWidget) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	ok := C.QTermWidget_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_eventFilter
func miqt_exec_callback_QTermWidget_eventFilter(self *C.QTermWidget, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QTermWidget{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QTermWidget) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	C.QTermWidget_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	ok := C.QTermWidget_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_timerEvent
func miqt_exec_callback_QTermWidget_timerEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	C.QTermWidget_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	ok := C.QTermWidget_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_childEvent
func miqt_exec_callback_QTermWidget_childEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_CustomEvent(event *qt.QEvent) {

	C.QTermWidget_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QTermWidget) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	ok := C.QTermWidget_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_customEvent
func miqt_exec_callback_QTermWidget_customEvent(self *C.QTermWidget, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QTermWidget{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QTermWidget) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	C.QTermWidget_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QTermWidget) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QTermWidget_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_connectNotify
func miqt_exec_callback_QTermWidget_connectNotify(self *C.QTermWidget, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QTermWidget{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QTermWidget) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	C.QTermWidget_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QTermWidget) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	ok := C.QTermWidget_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QTermWidget_disconnectNotify
func miqt_exec_callback_QTermWidget_disconnectNotify(self *C.QTermWidget, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QTermWidget{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QTermWidget) Delete() {
	C.QTermWidget_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QTermWidget) GoGC() {
	runtime.SetFinalizer(this, func(this *QTermWidget) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
