package webengine

/*

#include "gen_qwebenginescript.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt6"
	"runtime"
	"unsafe"
)

type QWebEngineScript__InjectionPoint int

const (
	QWebEngineScript__Deferred         QWebEngineScript__InjectionPoint = 0
	QWebEngineScript__DocumentReady    QWebEngineScript__InjectionPoint = 1
	QWebEngineScript__DocumentCreation QWebEngineScript__InjectionPoint = 2
)

type QWebEngineScript__ScriptWorldId int

const (
	QWebEngineScript__MainWorld        QWebEngineScript__ScriptWorldId = 0
	QWebEngineScript__ApplicationWorld QWebEngineScript__ScriptWorldId = 1
	QWebEngineScript__UserWorld        QWebEngineScript__ScriptWorldId = 2
)

type QWebEngineScript struct {
	h *C.QWebEngineScript
}

func (this *QWebEngineScript) cPointer() *C.QWebEngineScript {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineScript) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineScript constructs the type using only CGO pointers.
func newQWebEngineScript(h *C.QWebEngineScript) *QWebEngineScript {
	if h == nil {
		return nil
	}

	return &QWebEngineScript{h: h}
}

// UnsafeNewQWebEngineScript constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineScript(h unsafe.Pointer) *QWebEngineScript {
	return newQWebEngineScript((*C.QWebEngineScript)(h))
}

// NewQWebEngineScript constructs a new QWebEngineScript object.
func NewQWebEngineScript() *QWebEngineScript {

	return newQWebEngineScript(C.QWebEngineScript_new())
}

// NewQWebEngineScript2 constructs a new QWebEngineScript object.
func NewQWebEngineScript2(other *QWebEngineScript) *QWebEngineScript {

	return newQWebEngineScript(C.QWebEngineScript_new2(other.cPointer()))
}

func (this *QWebEngineScript) OperatorAssign(other *QWebEngineScript) {
	C.QWebEngineScript_operatorAssign(this.h, other.cPointer())
}

func (this *QWebEngineScript) Name() string {
	var _ms C.struct_miqt_string = C.QWebEngineScript_name(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineScript) SetName(name string) {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	C.QWebEngineScript_setName(this.h, name_ms)
}

func (this *QWebEngineScript) SourceUrl() *qt6.QUrl {
	_goptr := qt6.UnsafeNewQUrl(unsafe.Pointer(C.QWebEngineScript_sourceUrl(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWebEngineScript) SetSourceUrl(url *qt6.QUrl) {
	C.QWebEngineScript_setSourceUrl(this.h, (*C.QUrl)(url.UnsafePointer()))
}

func (this *QWebEngineScript) SourceCode() string {
	var _ms C.struct_miqt_string = C.QWebEngineScript_sourceCode(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWebEngineScript) SetSourceCode(sourceCode string) {
	sourceCode_ms := C.struct_miqt_string{}
	sourceCode_ms.data = C.CString(sourceCode)
	sourceCode_ms.len = C.size_t(len(sourceCode))
	defer C.free(unsafe.Pointer(sourceCode_ms.data))
	C.QWebEngineScript_setSourceCode(this.h, sourceCode_ms)
}

func (this *QWebEngineScript) InjectionPoint() QWebEngineScript__InjectionPoint {
	return (QWebEngineScript__InjectionPoint)(C.QWebEngineScript_injectionPoint(this.h))
}

func (this *QWebEngineScript) SetInjectionPoint(injectionPoint QWebEngineScript__InjectionPoint) {
	C.QWebEngineScript_setInjectionPoint(this.h, (C.int)(injectionPoint))
}

func (this *QWebEngineScript) WorldId() uint {
	return (uint)(C.QWebEngineScript_worldId(this.h))
}

func (this *QWebEngineScript) SetWorldId(worldId uint) {
	C.QWebEngineScript_setWorldId(this.h, (C.uint)(worldId))
}

func (this *QWebEngineScript) RunsOnSubFrames() bool {
	return (bool)(C.QWebEngineScript_runsOnSubFrames(this.h))
}

func (this *QWebEngineScript) SetRunsOnSubFrames(on bool) {
	C.QWebEngineScript_setRunsOnSubFrames(this.h, (C.bool)(on))
}

func (this *QWebEngineScript) OperatorEqual(other *QWebEngineScript) bool {
	return (bool)(C.QWebEngineScript_operatorEqual(this.h, other.cPointer()))
}

func (this *QWebEngineScript) OperatorNotEqual(other *QWebEngineScript) bool {
	return (bool)(C.QWebEngineScript_operatorNotEqual(this.h, other.cPointer()))
}

func (this *QWebEngineScript) Swap(other *QWebEngineScript) {
	C.QWebEngineScript_swap(this.h, other.cPointer())
}

// Delete this object from C++ memory.
func (this *QWebEngineScript) Delete() {
	C.QWebEngineScript_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineScript) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineScript) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
