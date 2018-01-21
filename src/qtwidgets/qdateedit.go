package qtwidgets

// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QDateEdit struct {
	*QDateTimeEdit
}

func (this *QDateEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDateTimeEdit.GetCthis()
	}
}
func NewQDateEditFromPointer(cthis unsafe.Pointer) *QDateEdit {
	bcthis0 := NewQDateTimeEditFromPointer(cthis)
	return &QDateEdit{bcthis0}
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:217
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDateEdit) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDateEdit10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:220
// index:0
// Public
// void QDateEdit(class QWidget *)
func NewQDateEdit(parent *QWidget /*444 QWidget **/) *QDateEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:221
// index:1
// Public
// void QDateEdit(const class QDate &, class QWidget *)
func NewQDateEdit_1(date *qtcore.QDate, parent *QWidget /*444 QWidget **/) *QDateEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = date.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateEditC2ERK5QDateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:222
// index:0
// Public virtual
// void ~QDateEdit()
func DeleteQDateEdit(*QDateEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:225
// index:0
// Public
// void userDateChanged(const class QDate &)
func (this *QDateEdit) UserDateChanged(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDateEdit15userDateChangedERK5QDate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end