package webengine

/*

#include "gen_qwebenginequotarequest.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"unsafe"
)

type QWebEngineQuotaRequest struct {
	h *C.QWebEngineQuotaRequest
}

func (this *QWebEngineQuotaRequest) cPointer() *C.QWebEngineQuotaRequest {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineQuotaRequest) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineQuotaRequest constructs the type using only CGO pointers.
func newQWebEngineQuotaRequest(h *C.QWebEngineQuotaRequest) *QWebEngineQuotaRequest {
	if h == nil {
		return nil
	}

	return &QWebEngineQuotaRequest{h: h}
}

// UnsafeNewQWebEngineQuotaRequest constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineQuotaRequest(h unsafe.Pointer) *QWebEngineQuotaRequest {
	return newQWebEngineQuotaRequest((*C.QWebEngineQuotaRequest)(h))
}

// NewQWebEngineQuotaRequest constructs a new QWebEngineQuotaRequest object.
func NewQWebEngineQuotaRequest() *QWebEngineQuotaRequest {

	return newQWebEngineQuotaRequest(C.QWebEngineQuotaRequest_new())
}

// NewQWebEngineQuotaRequest2 constructs a new QWebEngineQuotaRequest object.
func NewQWebEngineQuotaRequest2(param1 *QWebEngineQuotaRequest) *QWebEngineQuotaRequest {

	return newQWebEngineQuotaRequest(C.QWebEngineQuotaRequest_new2(param1.cPointer()))
}

func (this *QWebEngineQuotaRequest) Accept() {
	C.QWebEngineQuotaRequest_accept(this.h)
}

func (this *QWebEngineQuotaRequest) Reject() {
	C.QWebEngineQuotaRequest_reject(this.h)
}

func (this *QWebEngineQuotaRequest) Origin() *qt6.QUrl {
	_goptr := qt6.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineQuotaRequest_origin(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineQuotaRequest) RequestedSize() int64 {
	return (int64)(C.QWebEngineQuotaRequest_requestedSize(this.h))
}

func (this *QWebEngineQuotaRequest) OperatorEqual(that *QWebEngineQuotaRequest) bool {
	return (bool)(C.QWebEngineQuotaRequest_operatorEqual(this.h, that.cPointer()))
}

func (this *QWebEngineQuotaRequest) OperatorNotEqual(that *QWebEngineQuotaRequest) bool {
	return (bool)(C.QWebEngineQuotaRequest_operatorNotEqual(this.h, that.cPointer()))
}

// Delete this object from C++ memory.
func (this *QWebEngineQuotaRequest) Delete() {
	C.QWebEngineQuotaRequest_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineQuotaRequest) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineQuotaRequest) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
