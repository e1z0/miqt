package printsupport

/*

#include "gen_qprinterinfo.h"
#include <stdlib.h>

*/
import "C"

import (
	"github.com/mappu/miqt/qt"
	"runtime"
	"unsafe"
)

type QPrinterInfo struct {
	h *C.QPrinterInfo
}

func (this *QPrinterInfo) cPointer() *C.QPrinterInfo {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QPrinterInfo) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQPrinterInfo constructs the type using only CGO pointers.
func newQPrinterInfo(h *C.QPrinterInfo) *QPrinterInfo {
	if h == nil {
		return nil
	}

	return &QPrinterInfo{h: h}
}

// UnsafeNewQPrinterInfo constructs the type using only unsafe pointers.
func UnsafeNewQPrinterInfo(h unsafe.Pointer) *QPrinterInfo {
	return newQPrinterInfo((*C.QPrinterInfo)(h))
}

// NewQPrinterInfo constructs a new QPrinterInfo object.
func NewQPrinterInfo() *QPrinterInfo {

	return newQPrinterInfo(C.QPrinterInfo_new())
}

// NewQPrinterInfo2 constructs a new QPrinterInfo object.
func NewQPrinterInfo2(other *QPrinterInfo) *QPrinterInfo {

	return newQPrinterInfo(C.QPrinterInfo_new2(other.cPointer()))
}

// NewQPrinterInfo3 constructs a new QPrinterInfo object.
func NewQPrinterInfo3(printer *QPrinter) *QPrinterInfo {

	return newQPrinterInfo(C.QPrinterInfo_new3(printer.cPointer()))
}

func (this *QPrinterInfo) OperatorAssign(other *QPrinterInfo) {
	C.QPrinterInfo_operatorAssign(this.h, other.cPointer())
}

func (this *QPrinterInfo) PrinterName() string {
	var _ms C.struct_miqt_string = C.QPrinterInfo_printerName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) Description() string {
	var _ms C.struct_miqt_string = C.QPrinterInfo_description(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) Location() string {
	var _ms C.struct_miqt_string = C.QPrinterInfo_location(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) MakeAndModel() string {
	var _ms C.struct_miqt_string = C.QPrinterInfo_makeAndModel(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPrinterInfo) IsNull() bool {
	return (bool)(C.QPrinterInfo_isNull(this.h))
}

func (this *QPrinterInfo) IsDefault() bool {
	return (bool)(C.QPrinterInfo_isDefault(this.h))
}

func (this *QPrinterInfo) IsRemote() bool {
	return (bool)(C.QPrinterInfo_isRemote(this.h))
}

func (this *QPrinterInfo) State() QPrinter__PrinterState {
	return (QPrinter__PrinterState)(C.QPrinterInfo_state(this.h))
}

func (this *QPrinterInfo) SupportedPageSizes() []qt.QPageSize {
	var _ma C.struct_miqt_array = C.QPrinterInfo_supportedPageSizes(this.h)
	_ret := make([]qt.QPageSize, int(_ma.len))
	_outCast := (*[0xffff]*C.QPageSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(_outCast[i]))
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPrinterInfo) DefaultPageSize() *qt.QPageSize {
	_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(C.QPrinterInfo_defaultPageSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinterInfo) SupportsCustomPageSizes() bool {
	return (bool)(C.QPrinterInfo_supportsCustomPageSizes(this.h))
}

func (this *QPrinterInfo) MinimumPhysicalPageSize() *qt.QPageSize {
	_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(C.QPrinterInfo_minimumPhysicalPageSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinterInfo) MaximumPhysicalPageSize() *qt.QPageSize {
	_goptr := qt.UnsafeNewQPageSize(unsafe.Pointer(C.QPrinterInfo_maximumPhysicalPageSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPrinterInfo) SupportedPaperSizes() []qt.QPagedPaintDevice__PageSize {
	var _ma C.struct_miqt_array = C.QPrinterInfo_supportedPaperSizes(this.h)
	_ret := make([]qt.QPagedPaintDevice__PageSize, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (qt.QPagedPaintDevice__PageSize)(_outCast[i])
	}
	return _ret
}

func (this *QPrinterInfo) SupportedSizesWithNames() []struct {
	First  string
	Second qt.QSizeF
} {
	var _ma C.struct_miqt_array = C.QPrinterInfo_supportedSizesWithNames(this.h)
	_ret := make([]struct {
		First  string
		Second qt.QSizeF
	}, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_map)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_mm C.struct_miqt_map = _outCast[i]
		_lv_First_CArray := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_lv_mm.keys))
		_lv_Second_CArray := (*[0xffff]*C.QSizeF)(unsafe.Pointer(_lv_mm.values))
		var _lv_first_ms C.struct_miqt_string = _lv_First_CArray[0]
		_lv_first_ret := C.GoStringN(_lv_first_ms.data, C.int(int64(_lv_first_ms.len)))
		C.free(unsafe.Pointer(_lv_first_ms.data))
		_lv_entry_First := _lv_first_ret
		_lv_second_goptr := qt.UnsafeNewQSizeF(unsafe.Pointer(_lv_Second_CArray[0]))
		_lv_second_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_lv_entry_Second := *_lv_second_goptr

		_ret[i] = struct {
			First  string
			Second qt.QSizeF
		}{First: _lv_entry_First, Second: _lv_entry_Second}
	}
	return _ret
}

func (this *QPrinterInfo) SupportedResolutions() []int {
	var _ma C.struct_miqt_array = C.QPrinterInfo_supportedResolutions(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QPrinterInfo) DefaultDuplexMode() QPrinter__DuplexMode {
	return (QPrinter__DuplexMode)(C.QPrinterInfo_defaultDuplexMode(this.h))
}

func (this *QPrinterInfo) SupportedDuplexModes() []QPrinter__DuplexMode {
	var _ma C.struct_miqt_array = C.QPrinterInfo_supportedDuplexModes(this.h)
	_ret := make([]QPrinter__DuplexMode, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QPrinter__DuplexMode)(_outCast[i])
	}
	return _ret
}

func (this *QPrinterInfo) DefaultColorMode() QPrinter__ColorMode {
	return (QPrinter__ColorMode)(C.QPrinterInfo_defaultColorMode(this.h))
}

func (this *QPrinterInfo) SupportedColorModes() []QPrinter__ColorMode {
	var _ma C.struct_miqt_array = C.QPrinterInfo_supportedColorModes(this.h)
	_ret := make([]QPrinter__ColorMode, int(_ma.len))
	_outCast := (*[0xffff]C.int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QPrinter__ColorMode)(_outCast[i])
	}
	return _ret
}

func QPrinterInfo_AvailablePrinterNames() []string {
	var _ma C.struct_miqt_array = C.QPrinterInfo_availablePrinterNames()
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

func QPrinterInfo_AvailablePrinters() []QPrinterInfo {
	var _ma C.struct_miqt_array = C.QPrinterInfo_availablePrinters()
	_ret := make([]QPrinterInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QPrinterInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQPrinterInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QPrinterInfo_DefaultPrinterName() string {
	var _ms C.struct_miqt_string = C.QPrinterInfo_defaultPrinterName()
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPrinterInfo_DefaultPrinter() *QPrinterInfo {
	_goptr := newQPrinterInfo(C.QPrinterInfo_defaultPrinter())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPrinterInfo_PrinterInfo(printerName string) *QPrinterInfo {
	printerName_ms := C.struct_miqt_string{}
	printerName_ms.data = C.CString(printerName)
	printerName_ms.len = C.size_t(len(printerName))
	defer C.free(unsafe.Pointer(printerName_ms.data))
	_goptr := newQPrinterInfo(C.QPrinterInfo_printerInfo(printerName_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QPrinterInfo) Delete() {
	C.QPrinterInfo_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QPrinterInfo) GoGC() {
	runtime.SetFinalizer(this, func(this *QPrinterInfo) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
