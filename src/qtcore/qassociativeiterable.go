package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

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
}

//  ext block end

//  body block begin
type QAssociativeIterable struct {
	*qtrt.CObject
}

func (this *QAssociativeIterable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQAssociativeIterableFromPointer(cthis unsafe.Pointer) *QAssociativeIterable {
	return &QAssociativeIterable{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qvariant.h:680
// index:0
// Public
// QAssociativeIterable::const_iterator begin()
func (this *QAssociativeIterable) Begin() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:681
// index:0
// Public
// QAssociativeIterable::const_iterator end()
func (this *QAssociativeIterable) End() unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:682
// index:0
// Public
// QAssociativeIterable::const_iterator find(const class QVariant &)
func (this *QAssociativeIterable) Find(key *QVariant) unsafe.Pointer /*444*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable4findERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qvariant.h:684
// index:0
// Public
// QVariant value(const class QVariant &)
func (this *QAssociativeIterable) Value(key *QVariant) *QVariant /*123*/ {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable5valueERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:686
// index:0
// Public
// int size()
func (this *QAssociativeIterable) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QAssociativeIterable4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end