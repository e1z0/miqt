package network

/*

#include "gen_qsslcertificateextension.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"unsafe"
)

type QSslCertificateExtension struct {
	h *C.QSslCertificateExtension
}

func (this *QSslCertificateExtension) cPointer() *C.QSslCertificateExtension {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QSslCertificateExtension) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQSslCertificateExtension constructs the type using only CGO pointers.
func newQSslCertificateExtension(h *C.QSslCertificateExtension) *QSslCertificateExtension {
	if h == nil {
		return nil
	}

	return &QSslCertificateExtension{h: h}
}

// UnsafeNewQSslCertificateExtension constructs the type using only unsafe pointers.
func UnsafeNewQSslCertificateExtension(h unsafe.Pointer) *QSslCertificateExtension {
	return newQSslCertificateExtension((*C.QSslCertificateExtension)(h))
}

// NewQSslCertificateExtension constructs a new QSslCertificateExtension object.
func NewQSslCertificateExtension() *QSslCertificateExtension {

	return newQSslCertificateExtension(C.QSslCertificateExtension_new())
}

// NewQSslCertificateExtension2 constructs a new QSslCertificateExtension object.
func NewQSslCertificateExtension2(other *QSslCertificateExtension) *QSslCertificateExtension {

	return newQSslCertificateExtension(C.QSslCertificateExtension_new2(other.cPointer()))
}

func (this *QSslCertificateExtension) OperatorAssign(other *QSslCertificateExtension) {
	C.QSslCertificateExtension_operatorAssign(this.h, other.cPointer())
}

func (this *QSslCertificateExtension) Swap(other *QSslCertificateExtension) {
	C.QSslCertificateExtension_swap(this.h, other.cPointer())
}

func (this *QSslCertificateExtension) Oid() string {
	var _ms C.struct_miqt_string = C.QSslCertificateExtension_oid(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCertificateExtension) Name() string {
	var _ms C.struct_miqt_string = C.QSslCertificateExtension_name(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSslCertificateExtension) Value() *qt6.QVariant {
	_goptr := qt6.UnsafeNewQVariant(unsafe.Pointer(C.QSslCertificateExtension_value(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSslCertificateExtension) IsCritical() bool {
	return (bool)(C.QSslCertificateExtension_isCritical(this.h))
}

func (this *QSslCertificateExtension) IsSupported() bool {
	return (bool)(C.QSslCertificateExtension_isSupported(this.h))
}

// Delete this object from C++ memory.
func (this *QSslCertificateExtension) Delete() {
	C.QSslCertificateExtension_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QSslCertificateExtension) GoGC() {
	runtime.SetFinalizer(this, func(this *QSslCertificateExtension) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
