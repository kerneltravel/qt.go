package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 100
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
type QVariantComparisonHelper struct {
	*qtrt.CObject
}

func (this *QVariantComparisonHelper) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQVariantComparisonHelperFromPointer(cthis unsafe.Pointer) *QVariantComparisonHelper {
	return &QVariantComparisonHelper{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qvariant.h:560
// index:0
// Public inline
// void QVariantComparisonHelper(const class QVariant &)
func NewQVariantComparisonHelper(var_ *QVariant) *QVariantComparisonHelper {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = var_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QVariantComparisonHelperC2ERK8QVariant", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQVariantComparisonHelperFromPointer(cthis)
	return gothis
}

//  body block end