package webengine

/*

#include "gen_qwebengineclientcertificateselection.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/network"
	"runtime"
	"unsafe"
)

type QWebEngineClientCertificateSelection struct {
	h *C.QWebEngineClientCertificateSelection
}

func (this *QWebEngineClientCertificateSelection) cPointer() *C.QWebEngineClientCertificateSelection {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineClientCertificateSelection) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineClientCertificateSelection constructs the type using only CGO pointers.
func newQWebEngineClientCertificateSelection(h *C.QWebEngineClientCertificateSelection) *QWebEngineClientCertificateSelection {
	if h == nil {
		return nil
	}

	return &QWebEngineClientCertificateSelection{h: h}
}

// UnsafeNewQWebEngineClientCertificateSelection constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineClientCertificateSelection(h unsafe.Pointer) *QWebEngineClientCertificateSelection {
	return newQWebEngineClientCertificateSelection((*C.QWebEngineClientCertificateSelection)(h))
}

// NewQWebEngineClientCertificateSelection constructs a new QWebEngineClientCertificateSelection object.
func NewQWebEngineClientCertificateSelection(param1 *QWebEngineClientCertificateSelection) *QWebEngineClientCertificateSelection {

	return newQWebEngineClientCertificateSelection(C.QWebEngineClientCertificateSelection_new(param1.cPointer()))
}

func (this *QWebEngineClientCertificateSelection) OperatorAssign(param1 *QWebEngineClientCertificateSelection) {
	C.QWebEngineClientCertificateSelection_operatorAssign(this.h, param1.cPointer())
}

func (this *QWebEngineClientCertificateSelection) Host() *qt6.QUrl {
	_goptr := qt6.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineClientCertificateSelection_host(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineClientCertificateSelection) Select(certificate *network.QSslCertificate) {
	C.QWebEngineClientCertificateSelection_select(this.h, (*C.QSslCertificate)(certificate.UnsafePointer()))
}

func (this *QWebEngineClientCertificateSelection) SelectNone() {
	C.QWebEngineClientCertificateSelection_selectNone(this.h)
}

func (this *QWebEngineClientCertificateSelection) Certificates() []network.QSslCertificate {
	var _ma C.struct_miqt_array = C.QWebEngineClientCertificateSelection_certificates(this.h)
	_ret := make([]network.QSslCertificate, int(_ma.len))
	_outCast := (*[0xffff]*C.QSslCertificate)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := network.UnsafeNewQSslCertificate(unsafe.Pointer(_outCast[i]))
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

// Delete this object from C++ memory.
func (this *QWebEngineClientCertificateSelection) Delete() {
	C.QWebEngineClientCertificateSelection_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineClientCertificateSelection) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineClientCertificateSelection) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
