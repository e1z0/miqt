package webengine

/*

#include "gen_qwebengineview.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/e1z0/miqt/qt6"
	"github.com/e1z0/miqt/qt6/printsupport"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QWebEngineView struct {
	h *C.QWebEngineView
	*qt6.QWidget
}

func (this *QWebEngineView) cPointer() *C.QWebEngineView {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineView) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineView constructs the type using only CGO pointers.
func newQWebEngineView(h *C.QWebEngineView) *QWebEngineView {
	if h == nil {
		return nil
	}
	var outptr_QWidget *C.QWidget = nil
	C.QWebEngineView_virtbase(h, &outptr_QWidget)

	return &QWebEngineView{h: h,
		QWidget: qt6.UnsafeNewQWidget(unsafe.Pointer(outptr_QWidget))}
}

// UnsafeNewQWebEngineView constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineView(h unsafe.Pointer) *QWebEngineView {
	return newQWebEngineView((*C.QWebEngineView)(h))
}

// NewQWebEngineView constructs a new QWebEngineView object.
func NewQWebEngineView(parent *qt6.QWidget) *QWebEngineView {

	return newQWebEngineView(C.QWebEngineView_new((*C.QWidget)(parent.UnsafePointer())))
}

// NewQWebEngineView2 constructs a new QWebEngineView object.
func NewQWebEngineView2() *QWebEngineView {

	return newQWebEngineView(C.QWebEngineView_new2())
}

// NewQWebEngineView3 constructs a new QWebEngineView object.
func NewQWebEngineView3(profile *QWebEngineProfile) *QWebEngineView {

	return newQWebEngineView(C.QWebEngineView_new3(profile.cPointer()))
}

// NewQWebEngineView4 constructs a new QWebEngineView object.
func NewQWebEngineView4(page *QWebEnginePage) *QWebEngineView {

	return newQWebEngineView(C.QWebEngineView_new4(page.cPointer()))
}

// NewQWebEngineView5 constructs a new QWebEngineView object.
func NewQWebEngineView5(profile *QWebEngineProfile, parent *qt6.QWidget) *QWebEngineView {

	return newQWebEngineView(C.QWebEngineView_new5(profile.cPointer(), (*C.QWidget)(parent.UnsafePointer())))
}

// NewQWebEngineView6 constructs a new QWebEngineView object.
func NewQWebEngineView6(page *QWebEnginePage, parent *qt6.QWidget) *QWebEngineView {

	return newQWebEngineView(C.QWebEngineView_new6(page.cPointer(), (*C.QWidget)(parent.UnsafePointer())))
}

func (this *QWebEngineView) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QWebEngineView_metaObject(this.h)))
}

func (this *QWebEngineView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QWebEngineView_metacast(this.h, param1_Cstring))
}

func QWebEngineView_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineView_ForPage(page *QWebEnginePage) *QWebEngineView {
	return newQWebEngineView(C.QWebEngineView_forPage(page.cPointer()))
}

func (this *QWebEngineView) Page() *QWebEnginePage {
	return newQWebEnginePage(C.QWebEngineView_page(this.h))
}

func (this *QWebEngineView) SetPage(page *QWebEnginePage) {
	C.QWebEngineView_setPage(this.h, page.cPointer())
}

func (this *QWebEngineView) Load(url *qt6.QUrl) {
	C.QWebEngineView_load(this.h, (*C.QUrl)(url.UnsafePointer()))
}

func (this *QWebEngineView) LoadWithRequest(request *QWebEngineHttpRequest) {
	C.QWebEngineView_loadWithRequest(this.h, request.cPointer())
}

func (this *QWebEngineView) SetHtml(html string) {
	html_ms := C.struct_miqt_string{}
	html_ms.data = C.CString(html)
	html_ms.len = C.size_t(len(html))
	defer C.free(unsafe.Pointer(html_ms.data))
	C.QWebEngineView_setHtml(this.h, html_ms)
}

func (this *QWebEngineView) SetContent(data []byte) {
	data_alias := C.struct_miqt_string{}
	if len(data) > 0 {
		data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		data_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	data_alias.len = C.size_t(len(data))
	C.QWebEngineView_setContent(this.h, data_alias)
}

func (this *QWebEngineView) History() *QWebEngineHistory {
	return newQWebEngineHistory(C.QWebEngineView_history(this.h))
}

func (this *QWebEngineView) Title() string {
	var _ms C.struct_miqt_string = C.QWebEngineView_title(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineView) SetUrl(url *qt6.QUrl) {
	C.QWebEngineView_setUrl(this.h, (*C.QUrl)(url.UnsafePointer()))
}

func (this *QWebEngineView) Url() *qt6.QUrl {
	_goptr := qt6.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineView_url(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) IconUrl() *qt6.QUrl {
	_goptr := qt6.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineView_iconUrl(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) Icon() *qt6.QIcon {
	_goptr := qt6.UnsafeNewQIcon(unsafe.Pointer(C.QWebEngineView_icon(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) HasSelection() bool {
	return (bool)(C.QWebEngineView_hasSelection(this.h))
}

func (this *QWebEngineView) SelectedText() string {
	var _ms C.struct_miqt_string = C.QWebEngineView_selectedText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineView) PageAction(action QWebEnginePage__WebAction) *qt6.QAction {
	return qt6.UnsafeNewQAction(unsafe.Pointer(C.QWebEngineView_pageAction(this.h, (C.int)(action))))
}

func (this *QWebEngineView) TriggerPageAction(action QWebEnginePage__WebAction) {
	C.QWebEngineView_triggerPageAction(this.h, (C.int)(action))
}

func (this *QWebEngineView) ZoomFactor() float64 {
	return (float64)(C.QWebEngineView_zoomFactor(this.h))
}

func (this *QWebEngineView) SetZoomFactor(factor float64) {
	C.QWebEngineView_setZoomFactor(this.h, (C.double)(factor))
}

func (this *QWebEngineView) SizeHint() *qt6.QSize {
	_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(C.QWebEngineView_sizeHint(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineView) Settings() *QWebEngineSettings {
	return newQWebEngineSettings(C.QWebEngineView_settings(this.h))
}

func (this *QWebEngineView) CreateStandardContextMenu() *qt6.QMenu {
	return qt6.UnsafeNewQMenu(unsafe.Pointer(C.QWebEngineView_createStandardContextMenu(this.h)))
}

func (this *QWebEngineView) LastContextMenuRequest() *QWebEngineContextMenuRequest {
	return newQWebEngineContextMenuRequest(C.QWebEngineView_lastContextMenuRequest(this.h))
}

func (this *QWebEngineView) PrintToPdf(filePath string) {
	filePath_ms := C.struct_miqt_string{}
	filePath_ms.data = C.CString(filePath)
	filePath_ms.len = C.size_t(len(filePath))
	defer C.free(unsafe.Pointer(filePath_ms.data))
	C.QWebEngineView_printToPdf(this.h, filePath_ms)
}

func (this *QWebEngineView) Print(printer *printsupport.QPrinter) {
	C.QWebEngineView_print(this.h, (*C.QPrinter)(printer.UnsafePointer()))
}

func (this *QWebEngineView) Stop() {
	C.QWebEngineView_stop(this.h)
}

func (this *QWebEngineView) Back() {
	C.QWebEngineView_back(this.h)
}

func (this *QWebEngineView) Forward() {
	C.QWebEngineView_forward(this.h)
}

func (this *QWebEngineView) Reload() {
	C.QWebEngineView_reload(this.h)
}

func (this *QWebEngineView) LoadStarted() {
	C.QWebEngineView_loadStarted(this.h)
}
func (this *QWebEngineView) OnLoadStarted(slot func()) {
	C.QWebEngineView_connect_loadStarted(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_loadStarted
func miqt_exec_callback_QWebEngineView_loadStarted(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWebEngineView) LoadProgress(progress int) {
	C.QWebEngineView_loadProgress(this.h, (C.int)(progress))
}
func (this *QWebEngineView) OnLoadProgress(slot func(progress int)) {
	C.QWebEngineView_connect_loadProgress(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_loadProgress
func miqt_exec_callback_QWebEngineView_loadProgress(cb C.intptr_t, progress C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(progress int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(progress)

	gofunc(slotval1)
}

func (this *QWebEngineView) LoadFinished(param1 bool) {
	C.QWebEngineView_loadFinished(this.h, (C.bool)(param1))
}
func (this *QWebEngineView) OnLoadFinished(slot func(param1 bool)) {
	C.QWebEngineView_connect_loadFinished(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_loadFinished
func miqt_exec_callback_QWebEngineView_loadFinished(cb C.intptr_t, param1 C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QWebEngineView) TitleChanged(title string) {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	C.QWebEngineView_titleChanged(this.h, title_ms)
}
func (this *QWebEngineView) OnTitleChanged(slot func(title string)) {
	C.QWebEngineView_connect_titleChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_titleChanged
func miqt_exec_callback_QWebEngineView_titleChanged(cb C.intptr_t, title C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(title string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var title_ms C.struct_miqt_string = title
	title_ret := C.GoStringN(title_ms.data, C.int(int64(title_ms.len)))
	C.free(unsafe.Pointer(title_ms.data))
	slotval1 := title_ret

	gofunc(slotval1)
}

func (this *QWebEngineView) SelectionChanged() {
	C.QWebEngineView_selectionChanged(this.h)
}
func (this *QWebEngineView) OnSelectionChanged(slot func()) {
	C.QWebEngineView_connect_selectionChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_selectionChanged
func miqt_exec_callback_QWebEngineView_selectionChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWebEngineView) UrlChanged(param1 *qt6.QUrl) {
	C.QWebEngineView_urlChanged(this.h, (*C.QUrl)(param1.UnsafePointer()))
}
func (this *QWebEngineView) OnUrlChanged(slot func(param1 *qt6.QUrl)) {
	C.QWebEngineView_connect_urlChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_urlChanged
func miqt_exec_callback_QWebEngineView_urlChanged(cb C.intptr_t, param1 *C.QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt6.QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQUrl(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QWebEngineView) IconUrlChanged(param1 *qt6.QUrl) {
	C.QWebEngineView_iconUrlChanged(this.h, (*C.QUrl)(param1.UnsafePointer()))
}
func (this *QWebEngineView) OnIconUrlChanged(slot func(param1 *qt6.QUrl)) {
	C.QWebEngineView_connect_iconUrlChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_iconUrlChanged
func miqt_exec_callback_QWebEngineView_iconUrlChanged(cb C.intptr_t, param1 *C.QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt6.QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQUrl(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QWebEngineView) IconChanged(param1 *qt6.QIcon) {
	C.QWebEngineView_iconChanged(this.h, (*C.QIcon)(param1.UnsafePointer()))
}
func (this *QWebEngineView) OnIconChanged(slot func(param1 *qt6.QIcon)) {
	C.QWebEngineView_connect_iconChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_iconChanged
func miqt_exec_callback_QWebEngineView_iconChanged(cb C.intptr_t, param1 *C.QIcon) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt6.QIcon))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQIcon(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QWebEngineView) RenderProcessTerminated(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int) {
	C.QWebEngineView_renderProcessTerminated(this.h, (C.int)(terminationStatus), (C.int)(exitCode))
}
func (this *QWebEngineView) OnRenderProcessTerminated(slot func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int)) {
	C.QWebEngineView_connect_renderProcessTerminated(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_renderProcessTerminated
func miqt_exec_callback_QWebEngineView_renderProcessTerminated(cb C.intptr_t, terminationStatus C.int, exitCode C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(terminationStatus QWebEnginePage__RenderProcessTerminationStatus, exitCode int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QWebEnginePage__RenderProcessTerminationStatus)(terminationStatus)

	slotval2 := (int)(exitCode)

	gofunc(slotval1, slotval2)
}

func (this *QWebEngineView) PdfPrintingFinished(filePath string, success bool) {
	filePath_ms := C.struct_miqt_string{}
	filePath_ms.data = C.CString(filePath)
	filePath_ms.len = C.size_t(len(filePath))
	defer C.free(unsafe.Pointer(filePath_ms.data))
	C.QWebEngineView_pdfPrintingFinished(this.h, filePath_ms, (C.bool)(success))
}
func (this *QWebEngineView) OnPdfPrintingFinished(slot func(filePath string, success bool)) {
	C.QWebEngineView_connect_pdfPrintingFinished(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_pdfPrintingFinished
func miqt_exec_callback_QWebEngineView_pdfPrintingFinished(cb C.intptr_t, filePath C.struct_miqt_string, success C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(filePath string, success bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var filePath_ms C.struct_miqt_string = filePath
	filePath_ret := C.GoStringN(filePath_ms.data, C.int(int64(filePath_ms.len)))
	C.free(unsafe.Pointer(filePath_ms.data))
	slotval1 := filePath_ret
	slotval2 := (bool)(success)

	gofunc(slotval1, slotval2)
}

func (this *QWebEngineView) PrintRequested() {
	C.QWebEngineView_printRequested(this.h)
}
func (this *QWebEngineView) OnPrintRequested(slot func()) {
	C.QWebEngineView_connect_printRequested(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_printRequested
func miqt_exec_callback_QWebEngineView_printRequested(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWebEngineView) PrintFinished(success bool) {
	C.QWebEngineView_printFinished(this.h, (C.bool)(success))
}
func (this *QWebEngineView) OnPrintFinished(slot func(success bool)) {
	C.QWebEngineView_connect_printFinished(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineView_printFinished
func miqt_exec_callback_QWebEngineView_printFinished(cb C.intptr_t, success C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(success bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(success)

	gofunc(slotval1)
}

func QWebEngineView_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineView_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineView_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineView) SetHtml2(html string, baseUrl *qt6.QUrl) {
	html_ms := C.struct_miqt_string{}
	html_ms.data = C.CString(html)
	html_ms.len = C.size_t(len(html))
	defer C.free(unsafe.Pointer(html_ms.data))
	C.QWebEngineView_setHtml2(this.h, html_ms, (*C.QUrl)(baseUrl.UnsafePointer()))
}

func (this *QWebEngineView) SetContent2(data []byte, mimeType string) {
	data_alias := C.struct_miqt_string{}
	if len(data) > 0 {
		data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		data_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	data_alias.len = C.size_t(len(data))
	mimeType_ms := C.struct_miqt_string{}
	mimeType_ms.data = C.CString(mimeType)
	mimeType_ms.len = C.size_t(len(mimeType))
	defer C.free(unsafe.Pointer(mimeType_ms.data))
	C.QWebEngineView_setContent2(this.h, data_alias, mimeType_ms)
}

func (this *QWebEngineView) SetContent3(data []byte, mimeType string, baseUrl *qt6.QUrl) {
	data_alias := C.struct_miqt_string{}
	if len(data) > 0 {
		data_alias.data = (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		data_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	data_alias.len = C.size_t(len(data))
	mimeType_ms := C.struct_miqt_string{}
	mimeType_ms.data = C.CString(mimeType)
	mimeType_ms.len = C.size_t(len(mimeType))
	defer C.free(unsafe.Pointer(mimeType_ms.data))
	C.QWebEngineView_setContent3(this.h, data_alias, mimeType_ms, (*C.QUrl)(baseUrl.UnsafePointer()))
}

func (this *QWebEngineView) TriggerPageAction2(action QWebEnginePage__WebAction, checked bool) {
	C.QWebEngineView_triggerPageAction2(this.h, (C.int)(action), (C.bool)(checked))
}

func (this *QWebEngineView) PrintToPdf2(filePath string, layout *qt6.QPageLayout) {
	filePath_ms := C.struct_miqt_string{}
	filePath_ms.data = C.CString(filePath)
	filePath_ms.len = C.size_t(len(filePath))
	defer C.free(unsafe.Pointer(filePath_ms.data))
	C.QWebEngineView_printToPdf2(this.h, filePath_ms, (*C.QPageLayout)(layout.UnsafePointer()))
}

func (this *QWebEngineView) PrintToPdf3(filePath string, layout *qt6.QPageLayout, ranges *qt6.QPageRanges) {
	filePath_ms := C.struct_miqt_string{}
	filePath_ms.data = C.CString(filePath)
	filePath_ms.len = C.size_t(len(filePath))
	defer C.free(unsafe.Pointer(filePath_ms.data))
	C.QWebEngineView_printToPdf3(this.h, filePath_ms, (*C.QPageLayout)(layout.UnsafePointer()), (*C.QPageRanges)(ranges.UnsafePointer()))
}

// UpdateMicroFocus can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) UpdateMicroFocus() {

	var _dynamic_cast_ok C.bool = false
	C.QWebEngineView_protectedbase_updateMicroFocus(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// Create can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) Create() {

	var _dynamic_cast_ok C.bool = false
	C.QWebEngineView_protectedbase_create(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// Destroy can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) Destroy() {

	var _dynamic_cast_ok C.bool = false
	C.QWebEngineView_protectedbase_destroy(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// FocusNextChild can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) FocusNextChild() bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QWebEngineView_protectedbase_focusNextChild(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// FocusPreviousChild can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) FocusPreviousChild() bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QWebEngineView_protectedbase_focusPreviousChild(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Sender can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) Sender() *qt6.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt6.UnsafeNewQObject(unsafe.Pointer(C.QWebEngineView_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QWebEngineView_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QWebEngineView_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QWebEngineView that was directly constructed.
func (this *QWebEngineView) IsSignalConnected(signal *qt6.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QWebEngineView_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QWebEngineView) callVirtualBase_SizeHint() *qt6.QSize {

	_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(C.QWebEngineView_virtualbase_sizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWebEngineView) OnSizeHint(slot func(super func() *qt6.QSize) *qt6.QSize) {
	ok := C.QWebEngineView_override_virtual_sizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_sizeHint
func miqt_exec_callback_QWebEngineView_sizeHint(self *C.QWebEngineView, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt6.QSize) *qt6.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_SizeHint)

	return (*C.QSize)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_CreateWindow(typeVal QWebEnginePage__WebWindowType) *QWebEngineView {

	return newQWebEngineView(C.QWebEngineView_virtualbase_createWindow(unsafe.Pointer(this.h), (C.int)(typeVal)))

}
func (this *QWebEngineView) OnCreateWindow(slot func(super func(typeVal QWebEnginePage__WebWindowType) *QWebEngineView, typeVal QWebEnginePage__WebWindowType) *QWebEngineView) {
	ok := C.QWebEngineView_override_virtual_createWindow(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_createWindow
func miqt_exec_callback_QWebEngineView_createWindow(self *C.QWebEngineView, cb C.intptr_t, typeVal C.int) *C.QWebEngineView {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal QWebEnginePage__WebWindowType) *QWebEngineView, typeVal QWebEnginePage__WebWindowType) *QWebEngineView)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QWebEnginePage__WebWindowType)(typeVal)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_CreateWindow, slotval1)

	return virtualReturn.cPointer()

}

func (this *QWebEngineView) callVirtualBase_ContextMenuEvent(param1 *qt6.QContextMenuEvent) {

	C.QWebEngineView_virtualbase_contextMenuEvent(unsafe.Pointer(this.h), (*C.QContextMenuEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnContextMenuEvent(slot func(super func(param1 *qt6.QContextMenuEvent), param1 *qt6.QContextMenuEvent)) {
	ok := C.QWebEngineView_override_virtual_contextMenuEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_contextMenuEvent
func miqt_exec_callback_QWebEngineView_contextMenuEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt6.QContextMenuEvent), param1 *qt6.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQContextMenuEvent(unsafe.Pointer(param1))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_Event(param1 *qt6.QEvent) bool {

	return (bool)(C.QWebEngineView_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(param1.UnsafePointer())))

}
func (this *QWebEngineView) OnEvent(slot func(super func(param1 *qt6.QEvent) bool, param1 *qt6.QEvent) bool) {
	ok := C.QWebEngineView_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_event
func miqt_exec_callback_QWebEngineView_event(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt6.QEvent) bool, param1 *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(param1))

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_ShowEvent(param1 *qt6.QShowEvent) {

	C.QWebEngineView_virtualbase_showEvent(unsafe.Pointer(this.h), (*C.QShowEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnShowEvent(slot func(super func(param1 *qt6.QShowEvent), param1 *qt6.QShowEvent)) {
	ok := C.QWebEngineView_override_virtual_showEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_showEvent
func miqt_exec_callback_QWebEngineView_showEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt6.QShowEvent), param1 *qt6.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQShowEvent(unsafe.Pointer(param1))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_HideEvent(param1 *qt6.QHideEvent) {

	C.QWebEngineView_virtualbase_hideEvent(unsafe.Pointer(this.h), (*C.QHideEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnHideEvent(slot func(super func(param1 *qt6.QHideEvent), param1 *qt6.QHideEvent)) {
	ok := C.QWebEngineView_override_virtual_hideEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_hideEvent
func miqt_exec_callback_QWebEngineView_hideEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt6.QHideEvent), param1 *qt6.QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQHideEvent(unsafe.Pointer(param1))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_CloseEvent(param1 *qt6.QCloseEvent) {

	C.QWebEngineView_virtualbase_closeEvent(unsafe.Pointer(this.h), (*C.QCloseEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnCloseEvent(slot func(super func(param1 *qt6.QCloseEvent), param1 *qt6.QCloseEvent)) {
	ok := C.QWebEngineView_override_virtual_closeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_closeEvent
func miqt_exec_callback_QWebEngineView_closeEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt6.QCloseEvent), param1 *qt6.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQCloseEvent(unsafe.Pointer(param1))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DragEnterEvent(e *qt6.QDragEnterEvent) {

	C.QWebEngineView_virtualbase_dragEnterEvent(unsafe.Pointer(this.h), (*C.QDragEnterEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDragEnterEvent(slot func(super func(e *qt6.QDragEnterEvent), e *qt6.QDragEnterEvent)) {
	ok := C.QWebEngineView_override_virtual_dragEnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_dragEnterEvent
func miqt_exec_callback_QWebEngineView_dragEnterEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt6.QDragEnterEvent), e *qt6.QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQDragEnterEvent(unsafe.Pointer(e))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DragLeaveEvent(e *qt6.QDragLeaveEvent) {

	C.QWebEngineView_virtualbase_dragLeaveEvent(unsafe.Pointer(this.h), (*C.QDragLeaveEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDragLeaveEvent(slot func(super func(e *qt6.QDragLeaveEvent), e *qt6.QDragLeaveEvent)) {
	ok := C.QWebEngineView_override_virtual_dragLeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_dragLeaveEvent
func miqt_exec_callback_QWebEngineView_dragLeaveEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt6.QDragLeaveEvent), e *qt6.QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DragMoveEvent(e *qt6.QDragMoveEvent) {

	C.QWebEngineView_virtualbase_dragMoveEvent(unsafe.Pointer(this.h), (*C.QDragMoveEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDragMoveEvent(slot func(super func(e *qt6.QDragMoveEvent), e *qt6.QDragMoveEvent)) {
	ok := C.QWebEngineView_override_virtual_dragMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_dragMoveEvent
func miqt_exec_callback_QWebEngineView_dragMoveEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt6.QDragMoveEvent), e *qt6.QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQDragMoveEvent(unsafe.Pointer(e))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DropEvent(e *qt6.QDropEvent) {

	C.QWebEngineView_virtualbase_dropEvent(unsafe.Pointer(this.h), (*C.QDropEvent)(e.UnsafePointer()))

}
func (this *QWebEngineView) OnDropEvent(slot func(super func(e *qt6.QDropEvent), e *qt6.QDropEvent)) {
	ok := C.QWebEngineView_override_virtual_dropEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_dropEvent
func miqt_exec_callback_QWebEngineView_dropEvent(self *C.QWebEngineView, cb C.intptr_t, e *C.QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *qt6.QDropEvent), e *qt6.QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQDropEvent(unsafe.Pointer(e))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DevType() int {

	return (int)(C.QWebEngineView_virtualbase_devType(unsafe.Pointer(this.h)))

}
func (this *QWebEngineView) OnDevType(slot func(super func() int) int) {
	ok := C.QWebEngineView_override_virtual_devType(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_devType
func miqt_exec_callback_QWebEngineView_devType(self *C.QWebEngineView, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_DevType)

	return (C.int)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_SetVisible(visible bool) {

	C.QWebEngineView_virtualbase_setVisible(unsafe.Pointer(this.h), (C.bool)(visible))

}
func (this *QWebEngineView) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	ok := C.QWebEngineView_override_virtual_setVisible(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_setVisible
func miqt_exec_callback_QWebEngineView_setVisible(self *C.QWebEngineView, cb C.intptr_t, visible C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QWebEngineView{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MinimumSizeHint() *qt6.QSize {

	_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(C.QWebEngineView_virtualbase_minimumSizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWebEngineView) OnMinimumSizeHint(slot func(super func() *qt6.QSize) *qt6.QSize) {
	ok := C.QWebEngineView_override_virtual_minimumSizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_minimumSizeHint
func miqt_exec_callback_QWebEngineView_minimumSizeHint(self *C.QWebEngineView, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt6.QSize) *qt6.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_MinimumSizeHint)

	return (*C.QSize)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(C.QWebEngineView_virtualbase_heightForWidth(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QWebEngineView) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	ok := C.QWebEngineView_override_virtual_heightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_heightForWidth
func miqt_exec_callback_QWebEngineView_heightForWidth(self *C.QWebEngineView, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(C.QWebEngineView_virtualbase_hasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QWebEngineView) OnHasHeightForWidth(slot func(super func() bool) bool) {
	ok := C.QWebEngineView_override_virtual_hasHeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_hasHeightForWidth
func miqt_exec_callback_QWebEngineView_hasHeightForWidth(self *C.QWebEngineView, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_HasHeightForWidth)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_PaintEngine() *qt6.QPaintEngine {

	return qt6.UnsafeNewQPaintEngine(unsafe.Pointer(C.QWebEngineView_virtualbase_paintEngine(unsafe.Pointer(this.h))))

}
func (this *QWebEngineView) OnPaintEngine(slot func(super func() *qt6.QPaintEngine) *qt6.QPaintEngine) {
	ok := C.QWebEngineView_override_virtual_paintEngine(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_paintEngine
func miqt_exec_callback_QWebEngineView_paintEngine(self *C.QWebEngineView, cb C.intptr_t) *C.QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt6.QPaintEngine) *qt6.QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_PaintEngine)

	return (*C.QPaintEngine)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_MousePressEvent(event *qt6.QMouseEvent) {

	C.QWebEngineView_virtualbase_mousePressEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMousePressEvent(slot func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent)) {
	ok := C.QWebEngineView_override_virtual_mousePressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_mousePressEvent
func miqt_exec_callback_QWebEngineView_mousePressEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MouseReleaseEvent(event *qt6.QMouseEvent) {

	C.QWebEngineView_virtualbase_mouseReleaseEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMouseReleaseEvent(slot func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent)) {
	ok := C.QWebEngineView_override_virtual_mouseReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_mouseReleaseEvent
func miqt_exec_callback_QWebEngineView_mouseReleaseEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MouseDoubleClickEvent(event *qt6.QMouseEvent) {

	C.QWebEngineView_virtualbase_mouseDoubleClickEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMouseDoubleClickEvent(slot func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent)) {
	ok := C.QWebEngineView_override_virtual_mouseDoubleClickEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_mouseDoubleClickEvent
func miqt_exec_callback_QWebEngineView_mouseDoubleClickEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MouseMoveEvent(event *qt6.QMouseEvent) {

	C.QWebEngineView_virtualbase_mouseMoveEvent(unsafe.Pointer(this.h), (*C.QMouseEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMouseMoveEvent(slot func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent)) {
	ok := C.QWebEngineView_override_virtual_mouseMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_mouseMoveEvent
func miqt_exec_callback_QWebEngineView_mouseMoveEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QMouseEvent), event *qt6.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_WheelEvent(event *qt6.QWheelEvent) {

	C.QWebEngineView_virtualbase_wheelEvent(unsafe.Pointer(this.h), (*C.QWheelEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnWheelEvent(slot func(super func(event *qt6.QWheelEvent), event *qt6.QWheelEvent)) {
	ok := C.QWebEngineView_override_virtual_wheelEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_wheelEvent
func miqt_exec_callback_QWebEngineView_wheelEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QWheelEvent), event *qt6.QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQWheelEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_KeyPressEvent(event *qt6.QKeyEvent) {

	C.QWebEngineView_virtualbase_keyPressEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnKeyPressEvent(slot func(super func(event *qt6.QKeyEvent), event *qt6.QKeyEvent)) {
	ok := C.QWebEngineView_override_virtual_keyPressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_keyPressEvent
func miqt_exec_callback_QWebEngineView_keyPressEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QKeyEvent), event *qt6.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_KeyReleaseEvent(event *qt6.QKeyEvent) {

	C.QWebEngineView_virtualbase_keyReleaseEvent(unsafe.Pointer(this.h), (*C.QKeyEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnKeyReleaseEvent(slot func(super func(event *qt6.QKeyEvent), event *qt6.QKeyEvent)) {
	ok := C.QWebEngineView_override_virtual_keyReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_keyReleaseEvent
func miqt_exec_callback_QWebEngineView_keyReleaseEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QKeyEvent), event *qt6.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_FocusInEvent(event *qt6.QFocusEvent) {

	C.QWebEngineView_virtualbase_focusInEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnFocusInEvent(slot func(super func(event *qt6.QFocusEvent), event *qt6.QFocusEvent)) {
	ok := C.QWebEngineView_override_virtual_focusInEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_focusInEvent
func miqt_exec_callback_QWebEngineView_focusInEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QFocusEvent), event *qt6.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_FocusOutEvent(event *qt6.QFocusEvent) {

	C.QWebEngineView_virtualbase_focusOutEvent(unsafe.Pointer(this.h), (*C.QFocusEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnFocusOutEvent(slot func(super func(event *qt6.QFocusEvent), event *qt6.QFocusEvent)) {
	ok := C.QWebEngineView_override_virtual_focusOutEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_focusOutEvent
func miqt_exec_callback_QWebEngineView_focusOutEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QFocusEvent), event *qt6.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_EnterEvent(event *qt6.QEnterEvent) {

	C.QWebEngineView_virtualbase_enterEvent(unsafe.Pointer(this.h), (*C.QEnterEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnEnterEvent(slot func(super func(event *qt6.QEnterEvent), event *qt6.QEnterEvent)) {
	ok := C.QWebEngineView_override_virtual_enterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_enterEvent
func miqt_exec_callback_QWebEngineView_enterEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEnterEvent), event *qt6.QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEnterEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_LeaveEvent(event *qt6.QEvent) {

	C.QWebEngineView_virtualbase_leaveEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnLeaveEvent(slot func(super func(event *qt6.QEvent), event *qt6.QEvent)) {
	ok := C.QWebEngineView_override_virtual_leaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_leaveEvent
func miqt_exec_callback_QWebEngineView_leaveEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent), event *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_PaintEvent(event *qt6.QPaintEvent) {

	C.QWebEngineView_virtualbase_paintEvent(unsafe.Pointer(this.h), (*C.QPaintEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnPaintEvent(slot func(super func(event *qt6.QPaintEvent), event *qt6.QPaintEvent)) {
	ok := C.QWebEngineView_override_virtual_paintEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_paintEvent
func miqt_exec_callback_QWebEngineView_paintEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QPaintEvent), event *qt6.QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQPaintEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_MoveEvent(event *qt6.QMoveEvent) {

	C.QWebEngineView_virtualbase_moveEvent(unsafe.Pointer(this.h), (*C.QMoveEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnMoveEvent(slot func(super func(event *qt6.QMoveEvent), event *qt6.QMoveEvent)) {
	ok := C.QWebEngineView_override_virtual_moveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_moveEvent
func miqt_exec_callback_QWebEngineView_moveEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QMoveEvent), event *qt6.QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMoveEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_ResizeEvent(event *qt6.QResizeEvent) {

	C.QWebEngineView_virtualbase_resizeEvent(unsafe.Pointer(this.h), (*C.QResizeEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnResizeEvent(slot func(super func(event *qt6.QResizeEvent), event *qt6.QResizeEvent)) {
	ok := C.QWebEngineView_override_virtual_resizeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_resizeEvent
func miqt_exec_callback_QWebEngineView_resizeEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QResizeEvent), event *qt6.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQResizeEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_TabletEvent(event *qt6.QTabletEvent) {

	C.QWebEngineView_virtualbase_tabletEvent(unsafe.Pointer(this.h), (*C.QTabletEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnTabletEvent(slot func(super func(event *qt6.QTabletEvent), event *qt6.QTabletEvent)) {
	ok := C.QWebEngineView_override_virtual_tabletEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_tabletEvent
func miqt_exec_callback_QWebEngineView_tabletEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QTabletEvent), event *qt6.QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQTabletEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_ActionEvent(event *qt6.QActionEvent) {

	C.QWebEngineView_virtualbase_actionEvent(unsafe.Pointer(this.h), (*C.QActionEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnActionEvent(slot func(super func(event *qt6.QActionEvent), event *qt6.QActionEvent)) {
	ok := C.QWebEngineView_override_virtual_actionEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_actionEvent
func miqt_exec_callback_QWebEngineView_actionEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QActionEvent), event *qt6.QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQActionEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := C.struct_miqt_string{}
	if len(eventType) > 0 {
		eventType_alias.data = (*C.char)(unsafe.Pointer(&eventType[0]))
	} else {
		eventType_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	eventType_alias.len = C.size_t(len(eventType))

	return (bool)(C.QWebEngineView_virtualbase_nativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*C.intptr_t)(unsafe.Pointer(result))))

}
func (this *QWebEngineView) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	ok := C.QWebEngineView_override_virtual_nativeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_nativeEvent
func miqt_exec_callback_QWebEngineView_nativeEvent(self *C.QWebEngineView, cb C.intptr_t, eventType C.struct_miqt_string, message unsafe.Pointer, result *C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray C.struct_miqt_string = eventType
	eventType_ret := C.GoBytes(unsafe.Pointer(eventType_bytearray.data), C.int(int64(eventType_bytearray.len)))
	C.free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_ChangeEvent(param1 *qt6.QEvent) {

	C.QWebEngineView_virtualbase_changeEvent(unsafe.Pointer(this.h), (*C.QEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnChangeEvent(slot func(super func(param1 *qt6.QEvent), param1 *qt6.QEvent)) {
	ok := C.QWebEngineView_override_virtual_changeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_changeEvent
func miqt_exec_callback_QWebEngineView_changeEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt6.QEvent), param1 *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(param1))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_Metric(param1 qt6.QPaintDevice__PaintDeviceMetric) int {

	return (int)(C.QWebEngineView_virtualbase_metric(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QWebEngineView) OnMetric(slot func(super func(param1 qt6.QPaintDevice__PaintDeviceMetric) int, param1 qt6.QPaintDevice__PaintDeviceMetric) int) {
	ok := C.QWebEngineView_override_virtual_metric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_metric
func miqt_exec_callback_QWebEngineView_metric(self *C.QWebEngineView, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt6.QPaintDevice__PaintDeviceMetric) int, param1 qt6.QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt6.QPaintDevice__PaintDeviceMetric)(param1)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_Metric, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_InitPainter(painter *qt6.QPainter) {

	C.QWebEngineView_virtualbase_initPainter(unsafe.Pointer(this.h), (*C.QPainter)(painter.UnsafePointer()))

}
func (this *QWebEngineView) OnInitPainter(slot func(super func(painter *qt6.QPainter), painter *qt6.QPainter)) {
	ok := C.QWebEngineView_override_virtual_initPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_initPainter
func miqt_exec_callback_QWebEngineView_initPainter(self *C.QWebEngineView, cb C.intptr_t, painter *C.QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt6.QPainter), painter *qt6.QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQPainter(unsafe.Pointer(painter))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QWebEngineView) callVirtualBase_Redirected(offset *qt6.QPoint) *qt6.QPaintDevice {

	return qt6.UnsafeNewQPaintDevice(unsafe.Pointer(C.QWebEngineView_virtualbase_redirected(unsafe.Pointer(this.h), (*C.QPoint)(offset.UnsafePointer()))))

}
func (this *QWebEngineView) OnRedirected(slot func(super func(offset *qt6.QPoint) *qt6.QPaintDevice, offset *qt6.QPoint) *qt6.QPaintDevice) {
	ok := C.QWebEngineView_override_virtual_redirected(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_redirected
func miqt_exec_callback_QWebEngineView_redirected(self *C.QWebEngineView, cb C.intptr_t, offset *C.QPoint) *C.QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *qt6.QPoint) *qt6.QPaintDevice, offset *qt6.QPoint) *qt6.QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQPoint(unsafe.Pointer(offset))

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_Redirected, slotval1)

	return (*C.QPaintDevice)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_SharedPainter() *qt6.QPainter {

	return qt6.UnsafeNewQPainter(unsafe.Pointer(C.QWebEngineView_virtualbase_sharedPainter(unsafe.Pointer(this.h))))

}
func (this *QWebEngineView) OnSharedPainter(slot func(super func() *qt6.QPainter) *qt6.QPainter) {
	ok := C.QWebEngineView_override_virtual_sharedPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_sharedPainter
func miqt_exec_callback_QWebEngineView_sharedPainter(self *C.QWebEngineView, cb C.intptr_t) *C.QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt6.QPainter) *qt6.QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_SharedPainter)

	return (*C.QPainter)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_InputMethodEvent(param1 *qt6.QInputMethodEvent) {

	C.QWebEngineView_virtualbase_inputMethodEvent(unsafe.Pointer(this.h), (*C.QInputMethodEvent)(param1.UnsafePointer()))

}
func (this *QWebEngineView) OnInputMethodEvent(slot func(super func(param1 *qt6.QInputMethodEvent), param1 *qt6.QInputMethodEvent)) {
	ok := C.QWebEngineView_override_virtual_inputMethodEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_inputMethodEvent
func miqt_exec_callback_QWebEngineView_inputMethodEvent(self *C.QWebEngineView, cb C.intptr_t, param1 *C.QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt6.QInputMethodEvent), param1 *qt6.QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQInputMethodEvent(unsafe.Pointer(param1))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_InputMethodQuery(param1 qt6.InputMethodQuery) *qt6.QVariant {

	_goptr := qt6.UnsafeNewQVariant(unsafe.Pointer(C.QWebEngineView_virtualbase_inputMethodQuery(unsafe.Pointer(this.h), (C.int)(param1))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWebEngineView) OnInputMethodQuery(slot func(super func(param1 qt6.InputMethodQuery) *qt6.QVariant, param1 qt6.InputMethodQuery) *qt6.QVariant) {
	ok := C.QWebEngineView_override_virtual_inputMethodQuery(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_inputMethodQuery
func miqt_exec_callback_QWebEngineView_inputMethodQuery(self *C.QWebEngineView, cb C.intptr_t, param1 C.int) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt6.InputMethodQuery) *qt6.QVariant, param1 qt6.InputMethodQuery) *qt6.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt6.InputMethodQuery)(param1)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return (*C.QVariant)(virtualReturn.UnsafePointer())

}

func (this *QWebEngineView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(C.QWebEngineView_virtualbase_focusNextPrevChild(unsafe.Pointer(this.h), (C.bool)(next)))

}
func (this *QWebEngineView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	ok := C.QWebEngineView_override_virtual_focusNextPrevChild(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_focusNextPrevChild
func miqt_exec_callback_QWebEngineView_focusNextPrevChild(self *C.QWebEngineView, cb C.intptr_t, next C.bool) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_EventFilter(watched *qt6.QObject, event *qt6.QEvent) bool {

	return (bool)(C.QWebEngineView_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QWebEngineView) OnEventFilter(slot func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool) {
	ok := C.QWebEngineView_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_eventFilter
func miqt_exec_callback_QWebEngineView_eventFilter(self *C.QWebEngineView, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QWebEngineView{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QWebEngineView) callVirtualBase_TimerEvent(event *qt6.QTimerEvent) {

	C.QWebEngineView_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnTimerEvent(slot func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent)) {
	ok := C.QWebEngineView_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_timerEvent
func miqt_exec_callback_QWebEngineView_timerEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_ChildEvent(event *qt6.QChildEvent) {

	C.QWebEngineView_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnChildEvent(slot func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent)) {
	ok := C.QWebEngineView_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_childEvent
func miqt_exec_callback_QWebEngineView_childEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_CustomEvent(event *qt6.QEvent) {

	C.QWebEngineView_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QWebEngineView) OnCustomEvent(slot func(super func(event *qt6.QEvent), event *qt6.QEvent)) {
	ok := C.QWebEngineView_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_customEvent
func miqt_exec_callback_QWebEngineView_customEvent(self *C.QWebEngineView, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent), event *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QWebEngineView) callVirtualBase_ConnectNotify(signal *qt6.QMetaMethod) {

	C.QWebEngineView_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebEngineView) OnConnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QWebEngineView_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_connectNotify
func miqt_exec_callback_QWebEngineView_connectNotify(self *C.QWebEngineView, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QWebEngineView) callVirtualBase_DisconnectNotify(signal *qt6.QMetaMethod) {

	C.QWebEngineView_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QWebEngineView) OnDisconnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QWebEngineView_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QWebEngineView_disconnectNotify
func miqt_exec_callback_QWebEngineView_disconnectNotify(self *C.QWebEngineView, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QWebEngineView{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QWebEngineView) Delete() {
	C.QWebEngineView_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineView) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineView) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
