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
type QStyleOptionFrame struct {
	*QStyleOption
}

func (this *QStyleOptionFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func NewQStyleOptionFrameFromPointer(cthis unsafe.Pointer) *QStyleOptionFrame {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionFrame{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:143
// index:0
// Public
// void QStyleOptionFrame()
func NewQStyleOptionFrame() *QStyleOptionFrame {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QStyleOptionFrameC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFrameFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:147
// index:1
// Protected
// void QStyleOptionFrame(int)
func NewQStyleOptionFrame_1(version int) *QStyleOptionFrame {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QStyleOptionFrameC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFrameFromPointer(cthis)
	return gothis
}

//  body block end