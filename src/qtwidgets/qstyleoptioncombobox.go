package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QStyleOptionComboBox struct {
	*QStyleOptionComplex
}

func (this *QStyleOptionComboBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func NewQStyleOptionComboBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionComboBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionComboBox{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:601
// index:0
// Public
// void QStyleOptionComboBox()
func NewQStyleOptionComboBox() *QStyleOptionComboBox {
	cthis := qtrt.Calloc(1, 256) // 120
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionComboBoxFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:605
// index:1
// Protected
// void QStyleOptionComboBox(int)
func NewQStyleOptionComboBox_1(version int) *QStyleOptionComboBox {
	cthis := qtrt.Calloc(1, 256) // 120
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionComboBoxC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionComboBoxFromPointer(cthis)
	return gothis
}

//  body block end