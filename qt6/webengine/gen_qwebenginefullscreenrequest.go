package webengine

/*

#include "gen_qwebenginefullscreenrequest.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"unsafe"
)

type QWebEngineFullScreenRequest struct {
	h *C.QWebEngineFullScreenRequest
}

func (this *QWebEngineFullScreenRequest) cPointer() *C.QWebEngineFullScreenRequest {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineFullScreenRequest) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineFullScreenRequest constructs the type using only CGO pointers.
func newQWebEngineFullScreenRequest(h *C.QWebEngineFullScreenRequest) *QWebEngineFullScreenRequest {
	if h == nil {
		return nil
	}

	return &QWebEngineFullScreenRequest{h: h}
}

// UnsafeNewQWebEngineFullScreenRequest constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineFullScreenRequest(h unsafe.Pointer) *QWebEngineFullScreenRequest {
	return newQWebEngineFullScreenRequest((*C.QWebEngineFullScreenRequest)(h))
}

// NewQWebEngineFullScreenRequest constructs a new QWebEngineFullScreenRequest object.
func NewQWebEngineFullScreenRequest(other *QWebEngineFullScreenRequest) *QWebEngineFullScreenRequest {

	return newQWebEngineFullScreenRequest(C.QWebEngineFullScreenRequest_new(other.cPointer()))
}

func (this *QWebEngineFullScreenRequest) OperatorAssign(other *QWebEngineFullScreenRequest) {
	C.QWebEngineFullScreenRequest_operatorAssign(this.h, other.cPointer())
}

func (this *QWebEngineFullScreenRequest) Reject() {
	C.QWebEngineFullScreenRequest_reject(this.h)
}

func (this *QWebEngineFullScreenRequest) Accept() {
	C.QWebEngineFullScreenRequest_accept(this.h)
}

func (this *QWebEngineFullScreenRequest) ToggleOn() bool {
	return (bool)(C.QWebEngineFullScreenRequest_toggleOn(this.h))
}

func (this *QWebEngineFullScreenRequest) Origin() *qt6.QUrl {
	_goptr := qt6.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineFullScreenRequest_origin(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QWebEngineFullScreenRequest) Delete() {
	C.QWebEngineFullScreenRequest_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineFullScreenRequest) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineFullScreenRequest) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
