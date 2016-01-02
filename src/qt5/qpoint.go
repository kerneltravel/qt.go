package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qpoint.h
// dst-file: /src/core/qpoint.go
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
  // proto:  int & QPoint::ry();
extern void demth_ZN6QPoint2ryEv(void* qthis);
  // proto: static int QPoint::dotProduct(const QPoint & p1, const QPoint & p2);
extern void _ZN6QPoint10dotProductERKS_S1_(void* arg0, void* arg1);
  // proto:  int QPoint::x();
extern void _ZNK6QPoint1xEv(void* qthis);
  // proto:  void QPoint::QPoint(int xpos, int ypos);
extern void* dector_ZN6QPointC1Eii(int arg0, int arg1);
extern void _ZN6QPointC1Eii(void* qthis, int arg0, int arg1);
  // proto:  int QPoint::y();
extern void _ZNK6QPoint1yEv(void* qthis);
  // proto:  void QPoint::setX(int x);
extern void demth_ZN6QPoint4setXEi(void* qthis, int arg0);
  // proto:  bool QPoint::isNull();
extern void _ZNK6QPoint6isNullEv(void* qthis);
  // proto:  void QPoint::QPoint();
extern void* dector_ZN6QPointC1Ev();
extern void _ZN6QPointC1Ev(void* qthis);
  // proto:  void QPoint::setY(int y);
extern void demth_ZN6QPoint4setYEi(void* qthis, int arg0);
  // proto:  int & QPoint::rx();
extern void demth_ZN6QPoint2rxEv(void* qthis);
  // proto:  int QPoint::manhattanLength();
extern void _ZNK6QPoint15manhattanLengthEv(void* qthis);
  // proto:  void QPointF::QPointF(qreal xpos, qreal ypos);
extern void* dector_ZN7QPointFC1Edd(double arg0, double arg1);
extern void _ZN7QPointFC1Edd(void* qthis, double arg0, double arg1);
  // proto:  void QPointF::QPointF();
extern void* dector_ZN7QPointFC1Ev();
extern void _ZN7QPointFC1Ev(void* qthis);
  // proto:  qreal QPointF::manhattanLength();
extern void _ZNK7QPointF15manhattanLengthEv(void* qthis);
  // proto:  QPoint QPointF::toPoint();
extern void _ZNK7QPointF7toPointEv(void* qthis);
  // proto:  qreal & QPointF::rx();
extern void demth_ZN7QPointF2rxEv(void* qthis);
  // proto:  qreal QPointF::y();
extern void _ZNK7QPointF1yEv(void* qthis);
  // proto:  bool QPointF::isNull();
extern void demth_ZNK7QPointF6isNullEv(void* qthis);
  // proto:  qreal QPointF::x();
extern void _ZNK7QPointF1xEv(void* qthis);
  // proto:  void QPointF::QPointF(const QPoint & p);
extern void* dector_ZN7QPointFC1ERK6QPoint(void* arg0);
extern void _ZN7QPointFC1ERK6QPoint(void* qthis, void* arg0);
  // proto:  void QPointF::setX(qreal x);
extern void demth_ZN7QPointF4setXEd(void* qthis, double arg0);
  // proto:  qreal & QPointF::ry();
extern void demth_ZN7QPointF2ryEv(void* qthis);
  // proto: static qreal QPointF::dotProduct(const QPointF & p1, const QPointF & p2);
extern void _ZN7QPointF10dotProductERKS_S1_(void* arg0, void* arg1);
  // proto:  void QPointF::setY(qreal y);
extern void demth_ZN7QPointF4setYEd(void* qthis, double arg0);
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

// class sizeof(QPoint)=8
type QPoint struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPointF)=16
type QPointF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int & QPoint::ry();
func (this *QPoint) ry(args ...interface{}) () {
  // ry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint2ryEv
  default:
    qtrt.ErrorResolve("QPoint", "ry", args)
  }

}

  // proto: static int QPoint::dotProduct(const QPoint & p1, const QPoint & p2);
func (this *QPoint) dotProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPoint", "dotProduct", args)
  }

}

  // proto:  int QPoint::x();
func (this *QPoint) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint1xEv
  default:
    qtrt.ErrorResolve("QPoint", "x", args)
  }

}

  // proto:  void QPoint::QPoint(int xpos, int ypos);
func NewQPoint(args ...interface{}) QPoint {
  return QPoint{}
}

  // proto:  int QPoint::y();
func (this *QPoint) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint1yEv
  default:
    qtrt.ErrorResolve("QPoint", "y", args)
  }

}

  // proto:  void QPoint::setX(int x);
func (this *QPoint) setX(args ...interface{}) () {
  // setX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint4setXEi
  default:
    qtrt.ErrorResolve("QPoint", "setX", args)
  }

}

  // proto:  bool QPoint::isNull();
func (this *QPoint) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint6isNullEv
  default:
    qtrt.ErrorResolve("QPoint", "isNull", args)
  }

}

  // proto:  void QPoint::setY(int y);
func (this *QPoint) setY(args ...interface{}) () {
  // setY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint4setYEi
  default:
    qtrt.ErrorResolve("QPoint", "setY", args)
  }

}

  // proto:  int & QPoint::rx();
func (this *QPoint) rx(args ...interface{}) () {
  // rx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QPoint2rxEv
  default:
    qtrt.ErrorResolve("QPoint", "rx", args)
  }

}

  // proto:  int QPoint::manhattanLength();
func (this *QPoint) manhattanLength(args ...interface{}) () {
  // manhattanLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QPoint15manhattanLengthEv
  default:
    qtrt.ErrorResolve("QPoint", "manhattanLength", args)
  }

}

  // proto:  void QPointF::QPointF(qreal xpos, qreal ypos);
func NewQPointF(args ...interface{}) QPointF {
  return QPointF{}
}

  // proto:  qreal QPointF::manhattanLength();
func (this *QPointF) manhattanLength(args ...interface{}) () {
  // manhattanLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF15manhattanLengthEv
  default:
    qtrt.ErrorResolve("QPointF", "manhattanLength", args)
  }

}

  // proto:  QPoint QPointF::toPoint();
func (this *QPointF) toPoint(args ...interface{}) () {
  // toPoint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF7toPointEv
  default:
    qtrt.ErrorResolve("QPointF", "toPoint", args)
  }

}

  // proto:  qreal & QPointF::rx();
func (this *QPointF) rx(args ...interface{}) () {
  // rx()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF2rxEv
  default:
    qtrt.ErrorResolve("QPointF", "rx", args)
  }

}

  // proto:  qreal QPointF::y();
func (this *QPointF) y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF1yEv
  default:
    qtrt.ErrorResolve("QPointF", "y", args)
  }

}

  // proto:  bool QPointF::isNull();
func (this *QPointF) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF6isNullEv
  default:
    qtrt.ErrorResolve("QPointF", "isNull", args)
  }

}

  // proto:  qreal QPointF::x();
func (this *QPointF) x(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPointF1xEv
  default:
    qtrt.ErrorResolve("QPointF", "x", args)
  }

}

  // proto:  void QPointF::setX(qreal x);
func (this *QPointF) setX(args ...interface{}) () {
  // setX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF4setXEd
  default:
    qtrt.ErrorResolve("QPointF", "setX", args)
  }

}

  // proto:  qreal & QPointF::ry();
func (this *QPointF) ry(args ...interface{}) () {
  // ry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF2ryEv
  default:
    qtrt.ErrorResolve("QPointF", "ry", args)
  }

}

  // proto: static qreal QPointF::dotProduct(const QPointF & p1, const QPointF & p2);
func (this *QPointF) dotProduct_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPointF", "dotProduct", args)
  }

}

  // proto:  void QPointF::setY(qreal y);
func (this *QPointF) setY(args ...interface{}) () {
  // setY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPointF4setYEd
  default:
    qtrt.ErrorResolve("QPointF", "setY", args)
  }

}

// <= body block end

