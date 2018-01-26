package qtgui

// /usr/include/qt/QtGui/qaccessiblebridge.h
// #include <qaccessiblebridge.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QAccessibleBridge struct {
	*qtrt.CObject
}

func (this *QAccessibleBridge) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleBridge) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAccessibleBridgeFromPointer(cthis unsafe.Pointer) *QAccessibleBridge {
	return &QAccessibleBridge{&qtrt.CObject{cthis}}
}
func (*QAccessibleBridge) NewFromPointer(cthis unsafe.Pointer) *QAccessibleBridge {
	return NewQAccessibleBridgeFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:58
// index:0
// Public inline virtual
// void ~QAccessibleBridge()
func DeleteQAccessibleBridge(*QAccessibleBridge) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleBridgeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:59
// index:0
// Public pure virtual
// void setRootObject(class QAccessibleInterface *)
func (this *QAccessibleBridge) SetRootObject(arg0 *QAccessibleInterface /*777 QAccessibleInterface **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleBridge13setRootObjectEP20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:60
// index:0
// Public pure virtual
// void notifyAccessibilityUpdate(class QAccessibleEvent *)
func (this *QAccessibleBridge) NotifyAccessibilityUpdate(event *QAccessibleEvent /*777 QAccessibleEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleBridge25notifyAccessibilityUpdateEP16QAccessibleEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end