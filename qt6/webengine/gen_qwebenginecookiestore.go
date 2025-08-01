package webengine

/*

#include "gen_qwebenginecookiestore.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/e1z0/miqt/qt6"
	"github.com/e1z0/miqt/qt6/network"
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QWebEngineCookieStore struct {
	h *C.QWebEngineCookieStore
	*qt6.QObject
}

func (this *QWebEngineCookieStore) cPointer() *C.QWebEngineCookieStore {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineCookieStore) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineCookieStore constructs the type using only CGO pointers.
func newQWebEngineCookieStore(h *C.QWebEngineCookieStore) *QWebEngineCookieStore {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QWebEngineCookieStore_virtbase(h, &outptr_QObject)

	return &QWebEngineCookieStore{h: h,
		QObject: qt6.UnsafeNewQObject(unsafe.Pointer(outptr_QObject))}
}

// UnsafeNewQWebEngineCookieStore constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineCookieStore(h unsafe.Pointer) *QWebEngineCookieStore {
	return newQWebEngineCookieStore((*C.QWebEngineCookieStore)(h))
}

func (this *QWebEngineCookieStore) MetaObject() *qt6.QMetaObject {
	return qt6.UnsafeNewQMetaObject(unsafe.Pointer(C.QWebEngineCookieStore_metaObject(this.h)))
}

func (this *QWebEngineCookieStore) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QWebEngineCookieStore_metacast(this.h, param1_Cstring))
}

func QWebEngineCookieStore_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineCookieStore_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineCookieStore) SetCookie(cookie *network.QNetworkCookie) {
	C.QWebEngineCookieStore_setCookie(this.h, (*C.QNetworkCookie)(cookie.UnsafePointer()))
}

func (this *QWebEngineCookieStore) DeleteCookie(cookie *network.QNetworkCookie) {
	C.QWebEngineCookieStore_deleteCookie(this.h, (*C.QNetworkCookie)(cookie.UnsafePointer()))
}

func (this *QWebEngineCookieStore) DeleteSessionCookies() {
	C.QWebEngineCookieStore_deleteSessionCookies(this.h)
}

func (this *QWebEngineCookieStore) DeleteAllCookies() {
	C.QWebEngineCookieStore_deleteAllCookies(this.h)
}

func (this *QWebEngineCookieStore) LoadAllCookies() {
	C.QWebEngineCookieStore_loadAllCookies(this.h)
}

func (this *QWebEngineCookieStore) CookieAdded(cookie *network.QNetworkCookie) {
	C.QWebEngineCookieStore_cookieAdded(this.h, (*C.QNetworkCookie)(cookie.UnsafePointer()))
}
func (this *QWebEngineCookieStore) OnCookieAdded(slot func(cookie *network.QNetworkCookie)) {
	C.QWebEngineCookieStore_connect_cookieAdded(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineCookieStore_cookieAdded
func miqt_exec_callback_QWebEngineCookieStore_cookieAdded(cb C.intptr_t, cookie *C.QNetworkCookie) {
	gofunc, ok := cgo.Handle(cb).Value().(func(cookie *network.QNetworkCookie))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := network.UnsafeNewQNetworkCookie(unsafe.Pointer(cookie))

	gofunc(slotval1)
}

func (this *QWebEngineCookieStore) CookieRemoved(cookie *network.QNetworkCookie) {
	C.QWebEngineCookieStore_cookieRemoved(this.h, (*C.QNetworkCookie)(cookie.UnsafePointer()))
}
func (this *QWebEngineCookieStore) OnCookieRemoved(slot func(cookie *network.QNetworkCookie)) {
	C.QWebEngineCookieStore_connect_cookieRemoved(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWebEngineCookieStore_cookieRemoved
func miqt_exec_callback_QWebEngineCookieStore_cookieRemoved(cb C.intptr_t, cookie *C.QNetworkCookie) {
	gofunc, ok := cgo.Handle(cb).Value().(func(cookie *network.QNetworkCookie))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := network.UnsafeNewQNetworkCookie(unsafe.Pointer(cookie))

	gofunc(slotval1)
}

func QWebEngineCookieStore_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineCookieStore_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWebEngineCookieStore_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QWebEngineCookieStore_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineCookieStore) SetCookie2(cookie *network.QNetworkCookie, origin *qt6.QUrl) {
	C.QWebEngineCookieStore_setCookie2(this.h, (*C.QNetworkCookie)(cookie.UnsafePointer()), (*C.QUrl)(origin.UnsafePointer()))
}

func (this *QWebEngineCookieStore) DeleteCookie2(cookie *network.QNetworkCookie, origin *qt6.QUrl) {
	C.QWebEngineCookieStore_deleteCookie2(this.h, (*C.QNetworkCookie)(cookie.UnsafePointer()), (*C.QUrl)(origin.UnsafePointer()))
}

// Delete this object from C++ memory.
func (this *QWebEngineCookieStore) Delete() {
	C.QWebEngineCookieStore_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineCookieStore) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineCookieStore) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QWebEngineCookieStore__FilterRequest struct {
	h *C.QWebEngineCookieStore__FilterRequest
}

func (this *QWebEngineCookieStore__FilterRequest) cPointer() *C.QWebEngineCookieStore__FilterRequest {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineCookieStore__FilterRequest) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineCookieStore__FilterRequest constructs the type using only CGO pointers.
func newQWebEngineCookieStore__FilterRequest(h *C.QWebEngineCookieStore__FilterRequest) *QWebEngineCookieStore__FilterRequest {
	if h == nil {
		return nil
	}

	return &QWebEngineCookieStore__FilterRequest{h: h}
}

// UnsafeNewQWebEngineCookieStore__FilterRequest constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineCookieStore__FilterRequest(h unsafe.Pointer) *QWebEngineCookieStore__FilterRequest {
	return newQWebEngineCookieStore__FilterRequest((*C.QWebEngineCookieStore__FilterRequest)(h))
}

// NewQWebEngineCookieStore__FilterRequest constructs a new QWebEngineCookieStore::FilterRequest object.
func NewQWebEngineCookieStore__FilterRequest(param1 *QWebEngineCookieStore__FilterRequest) *QWebEngineCookieStore__FilterRequest {

	return newQWebEngineCookieStore__FilterRequest(C.QWebEngineCookieStore__FilterRequest_new(param1.cPointer()))
}

func (this *QWebEngineCookieStore__FilterRequest) OperatorAssign(param1 *QWebEngineCookieStore__FilterRequest) {
	C.QWebEngineCookieStore__FilterRequest_operatorAssign(this.h, param1.cPointer())
}

// Delete this object from C++ memory.
func (this *QWebEngineCookieStore__FilterRequest) Delete() {
	C.QWebEngineCookieStore__FilterRequest_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineCookieStore__FilterRequest) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineCookieStore__FilterRequest) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
