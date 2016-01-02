package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qopenglpaintdevice.h
// dst-file: /src/gui/qopenglpaintdevice.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QPaintEngine * QOpenGLPaintDevice::paintEngine();
extern void _ZNK18QOpenGLPaintDevice11paintEngineEv(void* qthis);
  // proto:  QSize QOpenGLPaintDevice::size();
extern void _ZNK18QOpenGLPaintDevice4sizeEv(void* qthis);
  // proto:  void QOpenGLPaintDevice::setPaintFlipped(bool flipped);
extern void _ZN18QOpenGLPaintDevice15setPaintFlippedEb(void* qthis, bool arg0);
  // proto:  void QOpenGLPaintDevice::~QOpenGLPaintDevice();
extern void _ZN18QOpenGLPaintDeviceD0Ev(void* qthis);
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(int width, int height);
extern void* dector_ZN18QOpenGLPaintDeviceC1Eii(int arg0, int arg1);
extern void _ZN18QOpenGLPaintDeviceC1Eii(void* qthis, int arg0, int arg1);
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(const QOpenGLPaintDevice & );
extern void* dector_ZN18QOpenGLPaintDeviceC1ERKS_(void* arg0);
extern void _ZN18QOpenGLPaintDeviceC1ERKS_(void* qthis, void* arg0);
  // proto:  QOpenGLContext * QOpenGLPaintDevice::context();
extern void _ZNK18QOpenGLPaintDevice7contextEv(void* qthis);
  // proto:  void QOpenGLPaintDevice::setDevicePixelRatio(qreal devicePixelRatio);
extern void _ZN18QOpenGLPaintDevice19setDevicePixelRatioEd(void* qthis, double arg0);
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice();
extern void* dector_ZN18QOpenGLPaintDeviceC1Ev();
extern void _ZN18QOpenGLPaintDeviceC1Ev(void* qthis);
  // proto:  int QOpenGLPaintDevice::devType();
extern void _ZNK18QOpenGLPaintDevice7devTypeEv(void* qthis);
  // proto:  qreal QOpenGLPaintDevice::dotsPerMeterX();
extern void _ZNK18QOpenGLPaintDevice13dotsPerMeterXEv(void* qthis);
  // proto:  void QOpenGLPaintDevice::setDotsPerMeterX(qreal );
extern void _ZN18QOpenGLPaintDevice16setDotsPerMeterXEd(void* qthis, double arg0);
  // proto:  qreal QOpenGLPaintDevice::dotsPerMeterY();
extern void _ZNK18QOpenGLPaintDevice13dotsPerMeterYEv(void* qthis);
  // proto:  void QOpenGLPaintDevice::setDotsPerMeterY(qreal );
extern void _ZN18QOpenGLPaintDevice16setDotsPerMeterYEd(void* qthis, double arg0);
  // proto:  bool QOpenGLPaintDevice::paintFlipped();
extern void _ZNK18QOpenGLPaintDevice12paintFlippedEv(void* qthis);
  // proto:  void QOpenGLPaintDevice::setSize(const QSize & size);
extern void _ZN18QOpenGLPaintDevice7setSizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QOpenGLPaintDevice::ensureActiveTarget();
extern void _ZN18QOpenGLPaintDevice18ensureActiveTargetEv(void* qthis);
  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(const QSize & size);
extern void* dector_ZN18QOpenGLPaintDeviceC1ERK5QSize(void* arg0);
extern void _ZN18QOpenGLPaintDeviceC1ERK5QSize(void* qthis, void* arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QOpenGLPaintDevice)=1
type QOpenGLPaintDevice struct {
  /*qbase*/ QPaintDevice;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QPaintEngine * QOpenGLPaintDevice::paintEngine();
func (this *QOpenGLPaintDevice) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice11paintEngineEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintEngine", args)
  }

}

  // proto:  QSize QOpenGLPaintDevice::size();
func (this *QOpenGLPaintDevice) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice4sizeEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "size", args)
  }

}

  // proto:  void QOpenGLPaintDevice::setPaintFlipped(bool flipped);
func (this *QOpenGLPaintDevice) setPaintFlipped(args ...interface{}) () {
  // setPaintFlipped(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice15setPaintFlippedEb
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setPaintFlipped", args)
  }

}

  // proto:  void QOpenGLPaintDevice::~QOpenGLPaintDevice();
func (this *QOpenGLPaintDevice) FreeQOpenGLPaintDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "~QOpenGLPaintDevice", args)
  }

}

  // proto:  void QOpenGLPaintDevice::QOpenGLPaintDevice(int width, int height);
func NewQOpenGLPaintDevice(args ...interface{}) QOpenGLPaintDevice {
  return QOpenGLPaintDevice{}
}

  // proto:  QOpenGLContext * QOpenGLPaintDevice::context();
func (this *QOpenGLPaintDevice) context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice7contextEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "context", args)
  }

}

  // proto:  void QOpenGLPaintDevice::setDevicePixelRatio(qreal devicePixelRatio);
func (this *QOpenGLPaintDevice) setDevicePixelRatio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice19setDevicePixelRatioEd
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDevicePixelRatio", args)
  }

}

  // proto:  int QOpenGLPaintDevice::devType();
func (this *QOpenGLPaintDevice) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice7devTypeEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "devType", args)
  }

}

  // proto:  qreal QOpenGLPaintDevice::dotsPerMeterX();
func (this *QOpenGLPaintDevice) dotsPerMeterX(args ...interface{}) () {
  // dotsPerMeterX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice13dotsPerMeterXEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterX", args)
  }

}

  // proto:  void QOpenGLPaintDevice::setDotsPerMeterX(qreal );
func (this *QOpenGLPaintDevice) setDotsPerMeterX(args ...interface{}) () {
  // setDotsPerMeterX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice16setDotsPerMeterXEd
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterX", args)
  }

}

  // proto:  qreal QOpenGLPaintDevice::dotsPerMeterY();
func (this *QOpenGLPaintDevice) dotsPerMeterY(args ...interface{}) () {
  // dotsPerMeterY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice13dotsPerMeterYEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterY", args)
  }

}

  // proto:  void QOpenGLPaintDevice::setDotsPerMeterY(qreal );
func (this *QOpenGLPaintDevice) setDotsPerMeterY(args ...interface{}) () {
  // setDotsPerMeterY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice16setDotsPerMeterYEd
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterY", args)
  }

}

  // proto:  bool QOpenGLPaintDevice::paintFlipped();
func (this *QOpenGLPaintDevice) paintFlipped(args ...interface{}) () {
  // paintFlipped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice12paintFlippedEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintFlipped", args)
  }

}

  // proto:  void QOpenGLPaintDevice::setSize(const QSize & size);
func (this *QOpenGLPaintDevice) setSize(args ...interface{}) () {
  // setSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice7setSizeERK5QSize
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setSize", args)
  }

}

  // proto:  void QOpenGLPaintDevice::ensureActiveTarget();
func (this *QOpenGLPaintDevice) ensureActiveTarget(args ...interface{}) () {
  // ensureActiveTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice18ensureActiveTargetEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "ensureActiveTarget", args)
  }

}

// <= body block end

