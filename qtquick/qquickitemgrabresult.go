package qtquick

// /usr/include/qt/QtQuick/qquickitemgrabresult.h
// #include <qquickitemgrabresult.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QQuickItemGrabResult) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QQuickItemGrabResult struct {
	*qtcore.QObject
}
type QQuickItemGrabResult_ITF interface {
	qtcore.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult { return ptr }

func (this *QQuickItemGrabResult) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickItemGrabResult) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickItemGrabResultFromPointer(cthis unsafe.Pointer) *QQuickItemGrabResult {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickItemGrabResult{bcthis0}
}
func (*QQuickItemGrabResult) NewFromPointer(cthis unsafe.Pointer) *QQuickItemGrabResult {
	return NewQQuickItemGrabResultFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQuickItemGrabResult) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:64
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage image()
func (this *QQuickItemGrabResult) Image() *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult5imageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QQuickItemGrabResult) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool saveToFile(const QString &)
func (this *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResult10saveToFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:70
// index:1
// Public Visibility=Default Availability=Available
// [1] bool saveToFile(const QString &)
func (this *QQuickItemGrabResult) SaveToFile_1(fileName string) bool {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult10saveToFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QQuickItemGrabResult) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResult5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ready()
func (this *QQuickItemGrabResult) Ready() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResult5readyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQQuickItemGrabResult(this *QQuickItemGrabResult) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResultD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end