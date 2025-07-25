package webengine

/*

#include "gen_qwebengineurlrequestjob.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"unsafe"
)

type QWebEngineUrlRequestJob__Error int

const (
	QWebEngineUrlRequestJob__NoError        QWebEngineUrlRequestJob__Error = 0
	QWebEngineUrlRequestJob__UrlNotFound    QWebEngineUrlRequestJob__Error = 1
	QWebEngineUrlRequestJob__UrlInvalid     QWebEngineUrlRequestJob__Error = 2
	QWebEngineUrlRequestJob__RequestAborted QWebEngineUrlRequestJob__Error = 3
	QWebEngineUrlRequestJob__RequestDenied  QWebEngineUrlRequestJob__Error = 4
	QWebEngineUrlRequestJob__RequestFailed  QWebEngineUrlRequestJob__Error = 5
)

type QWebEngineUrlRequestJob struct {
	h *C.QWebEngineUrlRequestJob
	*qt.QObject
}

func (this *QWebEngineUrlRequestJob) cPointer() *C.QWebEngineUrlRequestJob {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineUrlRequestJob) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineUrlRequestJob constructs the type using only CGO pointers.
func newQWebEngineUrlRequestJob(h *C.QWebEngineUrlRequestJob) *QWebEngineUrlRequestJob {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QWebEngineUrlRequestJob_virtbase(h, &outptr_QObject)

	return &QWebEngineUrlRequestJob{h: h,
		QObject: qt.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQWebEngineUrlRequestJob constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineUrlRequestJob(h unsafe.Pointer) *QWebEngineUrlRequestJob {
	return newQWebEngineUrlRequestJob((*C.QWebEngineUrlRequestJob)(h))
}

func (this *QWebEngineUrlRequestJob) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(C.QWebEngineUrlRequestJob_metaObject(this.h)))
}

func (this *QWebEngineUrlRequestJob) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QWebEngineUrlRequestJob_metacast(this.h, param1_Cstring))
}

func QWebEngineUrlRequestJob_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineUrlRequestJob_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineUrlRequestJob_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineUrlRequestJob_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineUrlRequestJob) RequestUrl() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineUrlRequestJob_requestUrl(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineUrlRequestJob) RequestMethod() []byte {
	var _bytearray C.struct_miqt_string = C.QWebEngineUrlRequestJob_requestMethod(this.h)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QWebEngineUrlRequestJob) Initiator() *qt.QUrl {
	_goptr := qt.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineUrlRequestJob_initiator(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineUrlRequestJob) Reply(contentType []byte, device *qt.QIODevice) {
	contentType_alias := C.struct_miqt_string{}
	if len(contentType) > 0 {
		contentType_alias.data = (*C.char)(unsafe.Pointer(&contentType[0]))
	} else {
		contentType_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	contentType_alias.len = C.size_t(len(contentType))
	C.QWebEngineUrlRequestJob_reply(this.h, contentType_alias, (*C.QIODevice)(device.UnsafePointer()))
}

func (this *QWebEngineUrlRequestJob) Fail(error QWebEngineUrlRequestJob__Error) {
	C.QWebEngineUrlRequestJob_fail(this.h, (C.int)(error))
}

func (this *QWebEngineUrlRequestJob) Redirect(url *qt.QUrl) {
	C.QWebEngineUrlRequestJob_redirect(this.h, (*C.QUrl)(url.UnsafePointer()))
}

func QWebEngineUrlRequestJob_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineUrlRequestJob_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineUrlRequestJob_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineUrlRequestJob_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineUrlRequestJob_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineUrlRequestJob_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineUrlRequestJob_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineUrlRequestJob_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

// Delete this object from C++ memory.
func (this *QWebEngineUrlRequestJob) Delete() {
	C.QWebEngineUrlRequestJob_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineUrlRequestJob) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineUrlRequestJob) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
