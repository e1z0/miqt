package multimedia

/*

#include "gen_qaudiosystem.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QAbstractAudioDeviceInfo struct {
	h *C.QAbstractAudioDeviceInfo
	*qt.QObject
}

func (this *QAbstractAudioDeviceInfo) cPointer() *C.QAbstractAudioDeviceInfo {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QAbstractAudioDeviceInfo) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQAbstractAudioDeviceInfo constructs the type using only CGO pointers.
func newQAbstractAudioDeviceInfo(h *C.QAbstractAudioDeviceInfo) *QAbstractAudioDeviceInfo {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QAbstractAudioDeviceInfo_virtbase(h, &outptr_QObject)

	return &QAbstractAudioDeviceInfo{h: h,
		QObject: qt.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQAbstractAudioDeviceInfo constructs the type using only unsafe pointers.
func UnsafeNewQAbstractAudioDeviceInfo(h unsafe.Pointer) *QAbstractAudioDeviceInfo {
	return newQAbstractAudioDeviceInfo((*C.QAbstractAudioDeviceInfo)(h))
}

func (this *QAbstractAudioDeviceInfo) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QAbstractAudioDeviceInfo_metaObject(this.h)))
}

func (this *QAbstractAudioDeviceInfo) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QAbstractAudioDeviceInfo_metacast(this.h, param1_Cstring))
}

func QAbstractAudioDeviceInfo_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioDeviceInfo_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioDeviceInfo_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioDeviceInfo_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractAudioDeviceInfo) PreferredFormat() *QAudioFormat {
	_goptr := newQAudioFormat(C.QAbstractAudioDeviceInfo_preferredFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractAudioDeviceInfo) IsFormatSupported(format *QAudioFormat) bool {
	return (bool)(C.QAbstractAudioDeviceInfo_isFormatSupported(this.h, format.cPointer()))
}

func (this *QAbstractAudioDeviceInfo) DeviceName() string {
	var _ms C.struct_miqt_string = C.QAbstractAudioDeviceInfo_deviceName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractAudioDeviceInfo) SupportedCodecs() []string {
	var _ma C.struct_miqt_array = C.QAbstractAudioDeviceInfo_supportedCodecs(this.h)
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

func (this *QAbstractAudioDeviceInfo) SupportedSampleRates() []int {
	var _ma C.struct_miqt_array = C.QAbstractAudioDeviceInfo_supportedSampleRates(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QAbstractAudioDeviceInfo) SupportedChannelCounts() []int {
	var _ma C.struct_miqt_array = C.QAbstractAudioDeviceInfo_supportedChannelCounts(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QAbstractAudioDeviceInfo) SupportedSampleSizes() []int {
	var _ma C.struct_miqt_array = C.QAbstractAudioDeviceInfo_supportedSampleSizes(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QAbstractAudioDeviceInfo) SupportedByteOrders() []QAudioFormat__Endian {
	var _ma C.struct_miqt_array = C.QAbstractAudioDeviceInfo_supportedByteOrders(this.h)
	_ret := make([]QAudioFormat__Endian, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QAudioFormat__Endian)(_outCast[i])
	}
	return _ret
}

func (this *QAbstractAudioDeviceInfo) SupportedSampleTypes() []QAudioFormat__SampleType {
	var _ma C.struct_miqt_array = C.QAbstractAudioDeviceInfo_supportedSampleTypes(this.h)
	_ret := make([]QAudioFormat__SampleType, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QAudioFormat__SampleType)(_outCast[i])
	}
	return _ret
}

func QAbstractAudioDeviceInfo_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioDeviceInfo_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioDeviceInfo_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioDeviceInfo_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioDeviceInfo_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioDeviceInfo_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioDeviceInfo_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioDeviceInfo_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QAbstractAudioDeviceInfo) Delete() {
	C.QAbstractAudioDeviceInfo_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAbstractAudioDeviceInfo) GoGC() {
	runtime.SetFinalizer(this, func(this *QAbstractAudioDeviceInfo) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QAbstractAudioOutput struct {
	h *C.QAbstractAudioOutput
	*qt.QObject
}

func (this *QAbstractAudioOutput) cPointer() *C.QAbstractAudioOutput {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QAbstractAudioOutput) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQAbstractAudioOutput constructs the type using only CGO pointers.
func newQAbstractAudioOutput(h *C.QAbstractAudioOutput) *QAbstractAudioOutput {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QAbstractAudioOutput_virtbase(h, &outptr_QObject)

	return &QAbstractAudioOutput{h: h,
		QObject: qt.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQAbstractAudioOutput constructs the type using only unsafe pointers.
func UnsafeNewQAbstractAudioOutput(h unsafe.Pointer) *QAbstractAudioOutput {
	return newQAbstractAudioOutput((*C.QAbstractAudioOutput)(h))
}

func (this *QAbstractAudioOutput) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QAbstractAudioOutput_metaObject(this.h)))
}

func (this *QAbstractAudioOutput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QAbstractAudioOutput_metacast(this.h, param1_Cstring))
}

func QAbstractAudioOutput_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioOutput_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioOutput_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioOutput_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractAudioOutput) Start(device *qt.QIODevice) {
	C.QAbstractAudioOutput_start(this.h, (*C.QIODevice)(device.UnsafePointer()))
}

func (this *QAbstractAudioOutput) Start2() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(C.QAbstractAudioOutput_start2(this.h)))
}

func (this *QAbstractAudioOutput) Stop() {
	C.QAbstractAudioOutput_stop(this.h)
}

func (this *QAbstractAudioOutput) Reset() {
	C.QAbstractAudioOutput_reset(this.h)
}

func (this *QAbstractAudioOutput) Suspend() {
	C.QAbstractAudioOutput_suspend(this.h)
}

func (this *QAbstractAudioOutput) Resume() {
	C.QAbstractAudioOutput_resume(this.h)
}

func (this *QAbstractAudioOutput) BytesFree() int {
	return (int)(C.QAbstractAudioOutput_bytesFree(this.h))
}

func (this *QAbstractAudioOutput) PeriodSize() int {
	return (int)(C.QAbstractAudioOutput_periodSize(this.h))
}

func (this *QAbstractAudioOutput) SetBufferSize(value int) {
	C.QAbstractAudioOutput_setBufferSize(this.h, (C.int)(value))
}

func (this *QAbstractAudioOutput) BufferSize() int {
	return (int)(C.QAbstractAudioOutput_bufferSize(this.h))
}

func (this *QAbstractAudioOutput) SetNotifyInterval(milliSeconds int) {
	C.QAbstractAudioOutput_setNotifyInterval(this.h, (C.int)(milliSeconds))
}

func (this *QAbstractAudioOutput) NotifyInterval() int {
	return (int)(C.QAbstractAudioOutput_notifyInterval(this.h))
}

func (this *QAbstractAudioOutput) ProcessedUSecs() int64 {
	return (int64)(C.QAbstractAudioOutput_processedUSecs(this.h))
}

func (this *QAbstractAudioOutput) ElapsedUSecs() int64 {
	return (int64)(C.QAbstractAudioOutput_elapsedUSecs(this.h))
}

func (this *QAbstractAudioOutput) Error() QAudio__Error {
	return (QAudio__Error)(C.QAbstractAudioOutput_error(this.h))
}

func (this *QAbstractAudioOutput) State() QAudio__State {
	return (QAudio__State)(C.QAbstractAudioOutput_state(this.h))
}

func (this *QAbstractAudioOutput) SetFormat(fmt *QAudioFormat) {
	C.QAbstractAudioOutput_setFormat(this.h, fmt.cPointer())
}

func (this *QAbstractAudioOutput) Format() *QAudioFormat {
	_goptr := newQAudioFormat(C.QAbstractAudioOutput_format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractAudioOutput) SetVolume(volume float64) {
	C.QAbstractAudioOutput_setVolume(this.h, (C.double)(volume))
}

func (this *QAbstractAudioOutput) Volume() float64 {
	return (float64)(C.QAbstractAudioOutput_volume(this.h))
}

func (this *QAbstractAudioOutput) Category() string {
	var _ms C.struct_miqt_string = C.QAbstractAudioOutput_category(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractAudioOutput) SetCategory(category string) {
	category_ms := C.struct_miqt_string{}
	category_ms.data = C.CString(category)
	category_ms.len = C.size_t(len(category))
	defer C.free(unsafe.Pointer(category_ms.data))
	C.QAbstractAudioOutput_setCategory(this.h, category_ms)
}

func (this *QAbstractAudioOutput) ErrorChanged(error QAudio__Error) {
	C.QAbstractAudioOutput_errorChanged(this.h, (C.int)(error))
}
func (this *QAbstractAudioOutput) OnErrorChanged(slot func(error QAudio__Error)) {
	C.QAbstractAudioOutput_connect_errorChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAudioOutput_errorChanged
func miqt_exec_callback_QAbstractAudioOutput_errorChanged(cb C.intptr_t, error C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QAudio__Error))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudio__Error)(error)

	gofunc(slotval1)
}

func (this *QAbstractAudioOutput) StateChanged(state QAudio__State) {
	C.QAbstractAudioOutput_stateChanged(this.h, (C.int)(state))
}
func (this *QAbstractAudioOutput) OnStateChanged(slot func(state QAudio__State)) {
	C.QAbstractAudioOutput_connect_stateChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAudioOutput_stateChanged
func miqt_exec_callback_QAbstractAudioOutput_stateChanged(cb C.intptr_t, state C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state QAudio__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudio__State)(state)

	gofunc(slotval1)
}

func (this *QAbstractAudioOutput) Notify() {
	C.QAbstractAudioOutput_notify(this.h)
}
func (this *QAbstractAudioOutput) OnNotify(slot func()) {
	C.QAbstractAudioOutput_connect_notify(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAudioOutput_notify
func miqt_exec_callback_QAbstractAudioOutput_notify(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QAbstractAudioOutput_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioOutput_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioOutput_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioOutput_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioOutput_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioOutput_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioOutput_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioOutput_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QAbstractAudioOutput) Delete() {
	C.QAbstractAudioOutput_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAbstractAudioOutput) GoGC() {
	runtime.SetFinalizer(this, func(this *QAbstractAudioOutput) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QAbstractAudioInput struct {
	h *C.QAbstractAudioInput
	*qt.QObject
}

func (this *QAbstractAudioInput) cPointer() *C.QAbstractAudioInput {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QAbstractAudioInput) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQAbstractAudioInput constructs the type using only CGO pointers.
func newQAbstractAudioInput(h *C.QAbstractAudioInput) *QAbstractAudioInput {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QAbstractAudioInput_virtbase(h, &outptr_QObject)

	return &QAbstractAudioInput{h: h,
		QObject: qt.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQAbstractAudioInput constructs the type using only unsafe pointers.
func UnsafeNewQAbstractAudioInput(h unsafe.Pointer) *QAbstractAudioInput {
	return newQAbstractAudioInput((*C.QAbstractAudioInput)(h))
}

func (this *QAbstractAudioInput) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QAbstractAudioInput_metaObject(this.h)))
}

func (this *QAbstractAudioInput) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QAbstractAudioInput_metacast(this.h, param1_Cstring))
}

func QAbstractAudioInput_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioInput_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioInput_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioInput_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractAudioInput) Start(device *qt.QIODevice) {
	C.QAbstractAudioInput_start(this.h, (*C.QIODevice)(device.UnsafePointer()))
}

func (this *QAbstractAudioInput) Start2() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(C.QAbstractAudioInput_start2(this.h)))
}

func (this *QAbstractAudioInput) Stop() {
	C.QAbstractAudioInput_stop(this.h)
}

func (this *QAbstractAudioInput) Reset() {
	C.QAbstractAudioInput_reset(this.h)
}

func (this *QAbstractAudioInput) Suspend() {
	C.QAbstractAudioInput_suspend(this.h)
}

func (this *QAbstractAudioInput) Resume() {
	C.QAbstractAudioInput_resume(this.h)
}

func (this *QAbstractAudioInput) BytesReady() int {
	return (int)(C.QAbstractAudioInput_bytesReady(this.h))
}

func (this *QAbstractAudioInput) PeriodSize() int {
	return (int)(C.QAbstractAudioInput_periodSize(this.h))
}

func (this *QAbstractAudioInput) SetBufferSize(value int) {
	C.QAbstractAudioInput_setBufferSize(this.h, (C.int)(value))
}

func (this *QAbstractAudioInput) BufferSize() int {
	return (int)(C.QAbstractAudioInput_bufferSize(this.h))
}

func (this *QAbstractAudioInput) SetNotifyInterval(milliSeconds int) {
	C.QAbstractAudioInput_setNotifyInterval(this.h, (C.int)(milliSeconds))
}

func (this *QAbstractAudioInput) NotifyInterval() int {
	return (int)(C.QAbstractAudioInput_notifyInterval(this.h))
}

func (this *QAbstractAudioInput) ProcessedUSecs() int64 {
	return (int64)(C.QAbstractAudioInput_processedUSecs(this.h))
}

func (this *QAbstractAudioInput) ElapsedUSecs() int64 {
	return (int64)(C.QAbstractAudioInput_elapsedUSecs(this.h))
}

func (this *QAbstractAudioInput) Error() QAudio__Error {
	return (QAudio__Error)(C.QAbstractAudioInput_error(this.h))
}

func (this *QAbstractAudioInput) State() QAudio__State {
	return (QAudio__State)(C.QAbstractAudioInput_state(this.h))
}

func (this *QAbstractAudioInput) SetFormat(fmt *QAudioFormat) {
	C.QAbstractAudioInput_setFormat(this.h, fmt.cPointer())
}

func (this *QAbstractAudioInput) Format() *QAudioFormat {
	_goptr := newQAudioFormat(C.QAbstractAudioInput_format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractAudioInput) SetVolume(volume float64) {
	C.QAbstractAudioInput_setVolume(this.h, (C.double)(volume))
}

func (this *QAbstractAudioInput) Volume() float64 {
	return (float64)(C.QAbstractAudioInput_volume(this.h))
}

func (this *QAbstractAudioInput) ErrorChanged(error QAudio__Error) {
	C.QAbstractAudioInput_errorChanged(this.h, (C.int)(error))
}
func (this *QAbstractAudioInput) OnErrorChanged(slot func(error QAudio__Error)) {
	C.QAbstractAudioInput_connect_errorChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAudioInput_errorChanged
func miqt_exec_callback_QAbstractAudioInput_errorChanged(cb C.intptr_t, error C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QAudio__Error))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudio__Error)(error)

	gofunc(slotval1)
}

func (this *QAbstractAudioInput) StateChanged(state QAudio__State) {
	C.QAbstractAudioInput_stateChanged(this.h, (C.int)(state))
}
func (this *QAbstractAudioInput) OnStateChanged(slot func(state QAudio__State)) {
	C.QAbstractAudioInput_connect_stateChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAudioInput_stateChanged
func miqt_exec_callback_QAbstractAudioInput_stateChanged(cb C.intptr_t, state C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(state QAudio__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QAudio__State)(state)

	gofunc(slotval1)
}

func (this *QAbstractAudioInput) Notify() {
	C.QAbstractAudioInput_notify(this.h)
}
func (this *QAbstractAudioInput) OnNotify(slot func()) {
	C.QAbstractAudioInput_connect_notify(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractAudioInput_notify
func miqt_exec_callback_QAbstractAudioInput_notify(cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QAbstractAudioInput_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioInput_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioInput_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioInput_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioInput_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioInput_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractAudioInput_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QAbstractAudioInput_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QAbstractAudioInput) Delete() {
	C.QAbstractAudioInput_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QAbstractAudioInput) GoGC() {
	runtime.SetFinalizer(this, func(this *QAbstractAudioInput) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
