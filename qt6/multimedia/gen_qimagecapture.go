package multimedia

/*

#include "gen_qimagecapture.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QImageCapture__Error int

const (
	QImageCapture__NoError                  QImageCapture__Error = 0
	QImageCapture__NotReadyError            QImageCapture__Error = 1
	QImageCapture__ResourceError            QImageCapture__Error = 2
	QImageCapture__OutOfSpaceError          QImageCapture__Error = 3
	QImageCapture__NotSupportedFeatureError QImageCapture__Error = 4
	QImageCapture__FormatError              QImageCapture__Error = 5
)

type QImageCapture__Quality int

const (
	QImageCapture__VeryLowQuality  QImageCapture__Quality = 0
	QImageCapture__LowQuality      QImageCapture__Quality = 1
	QImageCapture__NormalQuality   QImageCapture__Quality = 2
	QImageCapture__HighQuality     QImageCapture__Quality = 3
	QImageCapture__VeryHighQuality QImageCapture__Quality = 4
)

type QImageCapture__FileFormat int

const (
	QImageCapture__UnspecifiedFormat QImageCapture__FileFormat = 0
	QImageCapture__JPEG              QImageCapture__FileFormat = 1
	QImageCapture__PNG               QImageCapture__FileFormat = 2
	QImageCapture__WebP              QImageCapture__FileFormat = 3
	QImageCapture__Tiff              QImageCapture__FileFormat = 4
	QImageCapture__LastFileFormat    QImageCapture__FileFormat = 4
)

type QImageCapture struct {
	h *C.QImageCapture
	*qt6.QObject
}

func (this *QImageCapture) cPointer() *C.QImageCapture {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QImageCapture) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQImageCapture constructs the type using only CGO pointers.
func newQImageCapture(h *C.QImageCapture) *QImageCapture {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QImageCapture_virtbase(h, &outptr_QObject)

	return &QImageCapture{h: h,
		QObject: qt6.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQImageCapture constructs the type using only unsafe pointers.
func UnsafeNewQImageCapture(h unsafe.Pointer) *QImageCapture {
	return newQImageCapture((*C.QImageCapture)(h))
}

// NewQImageCapture constructs a new QImageCapture object.
func NewQImageCapture() *QImageCapture {

	return newQImageCapture(C.QImageCapture_new())
}

// NewQImageCapture2 constructs a new QImageCapture object.
func NewQImageCapture2(parent *qt6.QObject) *QImageCapture {

	return newQImageCapture(C.QImageCapture_new2((*C.QObject)(parent.UnsafePointer())))
}

func (this *QImageCapture) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QImageCapture_metaObject(this.h)))
}

func (this *QImageCapture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QImageCapture_metacast(this.h, param1_Cstring))
}

func QImageCapture_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QImageCapture_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) IsAvailable() bool {
	return (bool)(C.QImageCapture_isAvailable(this.h))
}

func (this *QImageCapture) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(C.QImageCapture_captureSession(this.h))
}

func (this *QImageCapture) Error() QImageCapture__Error {
	return (QImageCapture__Error)(C.QImageCapture_error(this.h))
}

func (this *QImageCapture) ErrorString() string {
	var _ms C.struct_miqt_string = C.QImageCapture_errorString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) IsReadyForCapture() bool {
	return (bool)(C.QImageCapture_isReadyForCapture(this.h))
}

func (this *QImageCapture) FileFormat() QImageCapture__FileFormat {
	return (QImageCapture__FileFormat)(C.QImageCapture_fileFormat(this.h))
}

func (this *QImageCapture) SetFileFormat(format QImageCapture__FileFormat) {
	C.QImageCapture_setFileFormat(this.h, (C.int)(format))
}

func QImageCapture_SupportedFormats() []QImageCapture__FileFormat {
	var _ma C.struct_miqt_array = C.QImageCapture_supportedFormats()
	_ret := make([]QImageCapture__FileFormat, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QImageCapture__FileFormat)(_outCast[i])
	}
	return _ret
}

func QImageCapture_FileFormatName(c QImageCapture__FileFormat) string {
	var _ms C.struct_miqt_string = C.QImageCapture_fileFormatName((C.int)(c))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QImageCapture_FileFormatDescription(c QImageCapture__FileFormat) string {
	var _ms C.struct_miqt_string = C.QImageCapture_fileFormatDescription((C.int)(c))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) Resolution() *qt6.QSize {
	_goptr := qt6.UnsafeNewQSize(unsafe.Pointer(C.QImageCapture_resolution(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageCapture) SetResolution(resolution *qt6.QSize) {
	C.QImageCapture_setResolution(this.h, (*C.QSize)(resolution.UnsafePointer()))
}

func (this *QImageCapture) SetResolution2(width int, height int) {
	C.QImageCapture_setResolution2(this.h, (C.int)(width), (C.int)(height))
}

func (this *QImageCapture) Quality() QImageCapture__Quality {
	return (QImageCapture__Quality)(C.QImageCapture_quality(this.h))
}

func (this *QImageCapture) SetQuality(quality QImageCapture__Quality) {
	C.QImageCapture_setQuality(this.h, (C.int)(quality))
}

func (this *QImageCapture) MetaData() *QMediaMetaData {
	_goptr := newQMediaMetaData(C.QImageCapture_metaData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageCapture) SetMetaData(metaData *QMediaMetaData) {
	C.QImageCapture_setMetaData(this.h, metaData.cPointer())
}

func (this *QImageCapture) AddMetaData(metaData *QMediaMetaData) {
	C.QImageCapture_addMetaData(this.h, metaData.cPointer())
}

func (this *QImageCapture) CaptureToFile() int {
	return (int)(C.QImageCapture_captureToFile(this.h))
}

func (this *QImageCapture) Capture() int {
	return (int)(C.QImageCapture_capture(this.h))
}

func (this *QImageCapture) ErrorChanged() {
	C.QImageCapture_errorChanged(this.h)
}
func (this *QImageCapture) OnErrorChanged(slot func()) {
	C.QImageCapture_connect_errorChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_errorChanged
func miqt_exec_callback_QImageCapture_errorChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) ErrorOccurred(id int, error QImageCapture__Error, errorString string) {
	errorString_ms := C.struct_miqt_string{}
	errorString_ms.data = C.CString(errorString)
	errorString_ms.len = C.size_t(len(errorString))
	defer C.free(unsafe.Pointer(errorString_ms.data))
	C.QImageCapture_errorOccurred(this.h, (C.int)(id), (C.int)(error), errorString_ms)
}
func (this *QImageCapture) OnErrorOccurred(slot func(id int, error QImageCapture__Error, errorString string)) {
	C.QImageCapture_connect_errorOccurred(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_errorOccurred
func miqt_exec_callback_QImageCapture_errorOccurred(cb C.intptr_t, id C.int, error C.int, errorString C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, error QImageCapture__Error, errorString string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := (QImageCapture__Error)(error)

	var errorString_ms C.struct_miqt_string = errorString
	errorString_ret := C.GoStringN(errorString_ms.data, C.int(int64(errorString_ms.len)))
	C.free(unsafe.Pointer(errorString_ms.data))
	slotval3 := errorString_ret

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QImageCapture) ReadyForCaptureChanged(ready bool) {
	C.QImageCapture_readyForCaptureChanged(this.h, (C.bool)(ready))
}
func (this *QImageCapture) OnReadyForCaptureChanged(slot func(ready bool)) {
	C.QImageCapture_connect_readyForCaptureChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_readyForCaptureChanged
func miqt_exec_callback_QImageCapture_readyForCaptureChanged(cb C.intptr_t, ready C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(ready bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(ready)

	gofunc(slotval1)
}

func (this *QImageCapture) MetaDataChanged() {
	C.QImageCapture_metaDataChanged(this.h)
}
func (this *QImageCapture) OnMetaDataChanged(slot func()) {
	C.QImageCapture_connect_metaDataChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_metaDataChanged
func miqt_exec_callback_QImageCapture_metaDataChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) FileFormatChanged() {
	C.QImageCapture_fileFormatChanged(this.h)
}
func (this *QImageCapture) OnFileFormatChanged(slot func()) {
	C.QImageCapture_connect_fileFormatChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_fileFormatChanged
func miqt_exec_callback_QImageCapture_fileFormatChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) QualityChanged() {
	C.QImageCapture_qualityChanged(this.h)
}
func (this *QImageCapture) OnQualityChanged(slot func()) {
	C.QImageCapture_connect_qualityChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_qualityChanged
func miqt_exec_callback_QImageCapture_qualityChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) ResolutionChanged() {
	C.QImageCapture_resolutionChanged(this.h)
}
func (this *QImageCapture) OnResolutionChanged(slot func()) {
	C.QImageCapture_connect_resolutionChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_resolutionChanged
func miqt_exec_callback_QImageCapture_resolutionChanged(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QImageCapture) ImageExposed(id int) {
	C.QImageCapture_imageExposed(this.h, (C.int)(id))
}
func (this *QImageCapture) OnImageExposed(slot func(id int)) {
	C.QImageCapture_connect_imageExposed(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_imageExposed
func miqt_exec_callback_QImageCapture_imageExposed(cb C.intptr_t, id C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	gofunc(slotval1)
}

func (this *QImageCapture) ImageCaptured(id int, preview *qt6.QImage) {
	C.QImageCapture_imageCaptured(this.h, (C.int)(id), (*C.QImage)(preview.UnsafePointer()))
}
func (this *QImageCapture) OnImageCaptured(slot func(id int, preview *qt6.QImage)) {
	C.QImageCapture_connect_imageCaptured(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_imageCaptured
func miqt_exec_callback_QImageCapture_imageCaptured(cb C.intptr_t, id C.int, preview *C.QImage) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, preview *qt6.QImage))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := qt6.UnsafeNewQImage(unsafe.Pointer(preview))

	gofunc(slotval1, slotval2)
}

func (this *QImageCapture) ImageMetadataAvailable(id int, metaData *QMediaMetaData) {
	C.QImageCapture_imageMetadataAvailable(this.h, (C.int)(id), metaData.cPointer())
}
func (this *QImageCapture) OnImageMetadataAvailable(slot func(id int, metaData *QMediaMetaData)) {
	C.QImageCapture_connect_imageMetadataAvailable(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_imageMetadataAvailable
func miqt_exec_callback_QImageCapture_imageMetadataAvailable(cb C.intptr_t, id C.int, metaData *C.QMediaMetaData) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, metaData *QMediaMetaData))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := newQMediaMetaData(metaData)

	gofunc(slotval1, slotval2)
}

func (this *QImageCapture) ImageAvailable(id int, frame *QVideoFrame) {
	C.QImageCapture_imageAvailable(this.h, (C.int)(id), frame.cPointer())
}
func (this *QImageCapture) OnImageAvailable(slot func(id int, frame *QVideoFrame)) {
	C.QImageCapture_connect_imageAvailable(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_imageAvailable
func miqt_exec_callback_QImageCapture_imageAvailable(cb C.intptr_t, id C.int, frame *C.QVideoFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, frame *QVideoFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := newQVideoFrame(frame)

	gofunc(slotval1, slotval2)
}

func (this *QImageCapture) ImageSaved(id int, fileName string) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	C.QImageCapture_imageSaved(this.h, (C.int)(id), fileName_ms)
}
func (this *QImageCapture) OnImageSaved(slot func(id int, fileName string)) {
	C.QImageCapture_connect_imageSaved(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageCapture_imageSaved
func miqt_exec_callback_QImageCapture_imageSaved(cb C.intptr_t, id C.int, fileName C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int, fileName string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	var fileName_ms C.struct_miqt_string = fileName
	fileName_ret := C.GoStringN(fileName_ms.data, C.int(int64(fileName_ms.len)))
	C.free(unsafe.Pointer(fileName_ms.data))
	slotval2 := fileName_ret

	gofunc(slotval1, slotval2)
}

func QImageCapture_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QImageCapture_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QImageCapture_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QImageCapture_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageCapture) CaptureToFileWithLocation(location string) int {
	location_ms := C.struct_miqt_string{}
	location_ms.data = C.CString(location)
	location_ms.len = C.size_t(len(location))
	defer C.free(unsafe.Pointer(location_ms.data))
	return (int)(C.QImageCapture_captureToFileWithLocation(this.h, location_ms))
}

// Sender can only be called from a QImageCapture that was directly constructed.
func (this *QImageCapture) Sender() *qt6.QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := qt6.UnsafeNewQObject(unsafe.Pointer(C.QImageCapture_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h))))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QImageCapture that was directly constructed.
func (this *QImageCapture) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QImageCapture_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QImageCapture that was directly constructed.
func (this *QImageCapture) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QImageCapture_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QImageCapture that was directly constructed.
func (this *QImageCapture) IsSignalConnected(signal *qt6.QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QImageCapture_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer())))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QImageCapture) callVirtualBase_Event(event *qt6.QEvent) bool {

	return (bool)(C.QImageCapture_virtualbase_event(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QImageCapture) OnEvent(slot func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool) {
	ok := C.QImageCapture_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QImageCapture_event
func miqt_exec_callback_QImageCapture_event(self *C.QImageCapture, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent) bool, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QImageCapture{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QImageCapture) callVirtualBase_EventFilter(watched *qt6.QObject, event *qt6.QEvent) bool {

	return (bool)(C.QImageCapture_virtualbase_eventFilter(unsafe.Pointer(this.h), (*C.QObject)(watched.UnsafePointer()), (*C.QEvent)(event.UnsafePointer())))

}
func (this *QImageCapture) OnEventFilter(slot func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool) {
	ok := C.QImageCapture_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QImageCapture_eventFilter
func miqt_exec_callback_QImageCapture_eventFilter(self *C.QImageCapture, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt6.QObject, event *qt6.QEvent) bool, watched *qt6.QObject, event *qt6.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QImageCapture{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QImageCapture) callVirtualBase_TimerEvent(event *qt6.QTimerEvent) {

	C.QImageCapture_virtualbase_timerEvent(unsafe.Pointer(this.h), (*C.QTimerEvent)(event.UnsafePointer()))

}
func (this *QImageCapture) OnTimerEvent(slot func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent)) {
	ok := C.QImageCapture_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QImageCapture_timerEvent
func miqt_exec_callback_QImageCapture_timerEvent(self *C.QImageCapture, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QTimerEvent), event *qt6.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QImageCapture{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QImageCapture) callVirtualBase_ChildEvent(event *qt6.QChildEvent) {

	C.QImageCapture_virtualbase_childEvent(unsafe.Pointer(this.h), (*C.QChildEvent)(event.UnsafePointer()))

}
func (this *QImageCapture) OnChildEvent(slot func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent)) {
	ok := C.QImageCapture_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QImageCapture_childEvent
func miqt_exec_callback_QImageCapture_childEvent(self *C.QImageCapture, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QChildEvent), event *qt6.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QImageCapture{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QImageCapture) callVirtualBase_CustomEvent(event *qt6.QEvent) {

	C.QImageCapture_virtualbase_customEvent(unsafe.Pointer(this.h), (*C.QEvent)(event.UnsafePointer()))

}
func (this *QImageCapture) OnCustomEvent(slot func(super func(event *qt6.QEvent), event *qt6.QEvent)) {
	ok := C.QImageCapture_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QImageCapture_customEvent
func miqt_exec_callback_QImageCapture_customEvent(self *C.QImageCapture, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt6.QEvent), event *qt6.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QImageCapture{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QImageCapture) callVirtualBase_ConnectNotify(signal *qt6.QMetaMethod) {

	C.QImageCapture_virtualbase_connectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QImageCapture) OnConnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QImageCapture_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QImageCapture_connectNotify
func miqt_exec_callback_QImageCapture_connectNotify(self *C.QImageCapture, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QImageCapture{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QImageCapture) callVirtualBase_DisconnectNotify(signal *qt6.QMetaMethod) {

	C.QImageCapture_virtualbase_disconnectNotify(unsafe.Pointer(this.h), (*C.QMetaMethod)(signal.UnsafePointer()))

}
func (this *QImageCapture) OnDisconnectNotify(slot func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod)) {
	ok := C.QImageCapture_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QImageCapture_disconnectNotify
func miqt_exec_callback_QImageCapture_disconnectNotify(self *C.QImageCapture, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt6.QMetaMethod), signal *qt6.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt6.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QImageCapture{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QImageCapture) Delete() {
	C.QImageCapture_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QImageCapture) GoGC() {
	runtime.SetFinalizer(this, func(this *QImageCapture) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
