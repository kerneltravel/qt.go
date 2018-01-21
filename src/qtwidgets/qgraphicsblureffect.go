package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QGraphicsBlurEffect struct {
	*QGraphicsEffect
}

func (this *QGraphicsBlurEffect) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsEffect.GetCthis()
	}
}
func NewQGraphicsBlurEffectFromPointer(cthis unsafe.Pointer) *QGraphicsBlurEffect {
	bcthis0 := NewQGraphicsEffectFromPointer(cthis)
	return &QGraphicsBlurEffect{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:157
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsBlurEffect) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:170
// index:0
// Public
// void QGraphicsBlurEffect(class QObject *)
func NewQGraphicsBlurEffect(parent *qtcore.QObject /*444 QObject **/) *QGraphicsBlurEffect {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsBlurEffectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:171
// index:0
// Public virtual
// void ~QGraphicsBlurEffect()
func DeleteQGraphicsBlurEffect(*QGraphicsBlurEffect) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:173
// index:0
// Public virtual
// QRectF boundingRectFor(const class QRectF &)
func (this *QGraphicsBlurEffect) BoundingRectFor(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect15boundingRectForERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:174
// index:0
// Public
// qreal blurRadius()
func (this *QGraphicsBlurEffect) BlurRadius() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect10blurRadiusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:175
// index:0
// Public
// QGraphicsBlurEffect::BlurHints blurHints()
func (this *QGraphicsBlurEffect) BlurHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QGraphicsBlurEffect9blurHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:178
// index:0
// Public
// void setBlurRadius(qreal)
func (this *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect13setBlurRadiusEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:182
// index:0
// Public
// void blurRadiusChanged(qreal)
func (this *QGraphicsBlurEffect) BlurRadiusChanged(blurRadius float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect17blurRadiusChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &blurRadius)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicseffect.h:186
// index:0
// Protected virtual
// void draw(class QPainter *)
func (this *QGraphicsBlurEffect) Draw(painter *qtgui.QPainter /*444 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QGraphicsBlurEffect4drawEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end