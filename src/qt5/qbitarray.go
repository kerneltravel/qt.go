package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qbitarray.h
// dst-file: /src/core/qbitarray.go
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
  // proto:  void QBitRef::QBitRef(QBitArray & array, int idx);
extern void* dector_ZN7QBitRefC1ER9QBitArrayi(void* arg0, int arg1);
extern void demth_ZN7QBitRefC1ER9QBitArrayi(void* qthis, void* arg0, int arg1);
  // proto:  void QBitArray::QBitArray(int size, bool val);
extern void* dector_ZN9QBitArrayC1Eib(int arg0, bool arg1);
extern void _ZN9QBitArrayC1Eib(void* qthis, int arg0, bool arg1);
  // proto:  bool QBitArray::isEmpty();
extern void demth_ZNK9QBitArray7isEmptyEv(void* qthis);
  // proto:  void QBitArray::setBit(int i);
extern void _ZN9QBitArray6setBitEi(void* qthis, int arg0);
  // proto:  int QBitArray::size();
extern void demth_ZNK9QBitArray4sizeEv(void* qthis);
  // proto:  void QBitArray::swap(QBitArray & other);
extern void demth_ZN9QBitArray4swapERS_(void* qthis, void* arg0);
  // proto:  int QBitArray::count();
extern void demth_ZNK9QBitArray5countEv(void* qthis);
  // proto:  int QBitArray::count(bool on);
extern void _ZNK9QBitArray5countEb(void* qthis, bool arg0);
  // proto:  void QBitArray::detach();
extern void demth_ZN9QBitArray6detachEv(void* qthis);
  // proto:  void QBitArray::QBitArray();
extern void* dector_ZN9QBitArrayC1Ev();
extern void demth_ZN9QBitArrayC1Ev(void* qthis);
  // proto:  bool QBitArray::at(int i);
extern void _ZNK9QBitArray2atEi(void* qthis, int arg0);
  // proto:  void QBitArray::clear();
extern void demth_ZN9QBitArray5clearEv(void* qthis);
  // proto:  void QBitArray::clearBit(int i);
extern void _ZN9QBitArray8clearBitEi(void* qthis, int arg0);
  // proto:  bool QBitArray::testBit(int i);
extern void _ZNK9QBitArray7testBitEi(void* qthis, int arg0);
  // proto:  void QBitArray::truncate(int pos);
extern void demth_ZN9QBitArray8truncateEi(void* qthis, int arg0);
  // proto:  bool QBitArray::toggleBit(int i);
extern void _ZN9QBitArray9toggleBitEi(void* qthis, int arg0);
  // proto:  void QBitArray::QBitArray(const QBitArray & other);
extern void* dector_ZN9QBitArrayC1ERKS_(void* arg0);
extern void demth_ZN9QBitArrayC1ERKS_(void* qthis, void* arg0);
  // proto:  void QBitArray::fill(bool val, int first, int last);
extern void _ZN9QBitArray4fillEbii(void* qthis, bool arg0, int arg1, int arg2);
  // proto:  bool QBitArray::isNull();
extern void demth_ZNK9QBitArray6isNullEv(void* qthis);
  // proto:  void QBitArray::setBit(int i, bool val);
extern void _ZN9QBitArray6setBitEib(void* qthis, int arg0, bool arg1);
  // proto:  void QBitArray::resize(int size);
extern void _ZN9QBitArray6resizeEi(void* qthis, int arg0);
  // proto:  bool QBitArray::isDetached();
extern void demth_ZNK9QBitArray10isDetachedEv(void* qthis);
  // proto:  bool QBitArray::fill(bool val, int size);
extern void demth_ZN9QBitArray4fillEbi(void* qthis, bool arg0, int arg1);
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

// class sizeof(QBitRef)=16
type QBitRef struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QBitArray)=8
type QBitArray struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QBitRef::QBitRef(QBitArray & array, int idx);
func NewQBitRef(args ...interface{}) QBitRef {
  return QBitRef{}
}

  // proto:  void QBitArray::QBitArray(int size, bool val);
func NewQBitArray(args ...interface{}) QBitArray {
  return QBitArray{}
}

  // proto:  bool QBitArray::isEmpty();
func (this *QBitArray) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray7isEmptyEv
  default:
    qtrt.ErrorResolve("QBitArray", "isEmpty", args)
  }

}

  // proto:  void QBitArray::setBit(int i);
func (this *QBitArray) setBit(args ...interface{}) () {
  // setBit(int)
  // setBit(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6setBitEi
  case 1:
    // invoke: _ZN9QBitArray6setBitEib
  default:
    qtrt.ErrorResolve("QBitArray", "setBit", args)
  }

}

  // proto:  int QBitArray::size();
func (this *QBitArray) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray4sizeEv
  default:
    qtrt.ErrorResolve("QBitArray", "size", args)
  }

}

  // proto:  void QBitArray::swap(QBitArray & other);
func (this *QBitArray) swap(args ...interface{}) () {
  // swap(class QBitArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitArray{}) // "QBitArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray4swapERS_
  default:
    qtrt.ErrorResolve("QBitArray", "swap", args)
  }

}

  // proto:  int QBitArray::count();
func (this *QBitArray) count(args ...interface{}) () {
  // count()
  // count(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray5countEv
  case 1:
    // invoke: _ZNK9QBitArray5countEb
  default:
    qtrt.ErrorResolve("QBitArray", "count", args)
  }

}

  // proto:  void QBitArray::detach();
func (this *QBitArray) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6detachEv
  default:
    qtrt.ErrorResolve("QBitArray", "detach", args)
  }

}

  // proto:  bool QBitArray::at(int i);
func (this *QBitArray) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray2atEi
  default:
    qtrt.ErrorResolve("QBitArray", "at", args)
  }

}

  // proto:  void QBitArray::clear();
func (this *QBitArray) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray5clearEv
  default:
    qtrt.ErrorResolve("QBitArray", "clear", args)
  }

}

  // proto:  void QBitArray::clearBit(int i);
func (this *QBitArray) clearBit(args ...interface{}) () {
  // clearBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray8clearBitEi
  default:
    qtrt.ErrorResolve("QBitArray", "clearBit", args)
  }

}

  // proto:  bool QBitArray::testBit(int i);
func (this *QBitArray) testBit(args ...interface{}) () {
  // testBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray7testBitEi
  default:
    qtrt.ErrorResolve("QBitArray", "testBit", args)
  }

}

  // proto:  void QBitArray::truncate(int pos);
func (this *QBitArray) truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray8truncateEi
  default:
    qtrt.ErrorResolve("QBitArray", "truncate", args)
  }

}

  // proto:  bool QBitArray::toggleBit(int i);
func (this *QBitArray) toggleBit(args ...interface{}) () {
  // toggleBit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray9toggleBitEi
  default:
    qtrt.ErrorResolve("QBitArray", "toggleBit", args)
  }

}

  // proto:  void QBitArray::fill(bool val, int first, int last);
func (this *QBitArray) fill(args ...interface{}) () {
  // fill(_Bool, int, int)
  // fill(_Bool, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.BoolTy(false) // "bool"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray4fillEbii
  case 1:
    // invoke: _ZN9QBitArray4fillEbi
  default:
    qtrt.ErrorResolve("QBitArray", "fill", args)
  }

}

  // proto:  bool QBitArray::isNull();
func (this *QBitArray) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray6isNullEv
  default:
    qtrt.ErrorResolve("QBitArray", "isNull", args)
  }

}

  // proto:  void QBitArray::resize(int size);
func (this *QBitArray) resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QBitArray6resizeEi
  default:
    qtrt.ErrorResolve("QBitArray", "resize", args)
  }

}

  // proto:  bool QBitArray::isDetached();
func (this *QBitArray) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QBitArray10isDetachedEv
  default:
    qtrt.ErrorResolve("QBitArray", "isDetached", args)
  }

}

// <= body block end

