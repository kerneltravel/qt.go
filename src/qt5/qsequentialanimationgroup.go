package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qsequentialanimationgroup.h
// dst-file: /src/core/qsequentialanimationgroup.go
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
  // proto:  QPauseAnimation * QSequentialAnimationGroup::insertPause(int index, int msecs);
extern void _ZN25QSequentialAnimationGroup11insertPauseEii(void* qthis, int arg0, int arg1);
  // proto:  void QSequentialAnimationGroup::QSequentialAnimationGroup(QObject * parent);
extern void* dector_ZN25QSequentialAnimationGroupC1EP7QObject(void* arg0);
extern void _ZN25QSequentialAnimationGroupC1EP7QObject(void* qthis, void* arg0);
  // proto:  QPauseAnimation * QSequentialAnimationGroup::addPause(int msecs);
extern void _ZN25QSequentialAnimationGroup8addPauseEi(void* qthis, int arg0);
  // proto:  void QSequentialAnimationGroup::~QSequentialAnimationGroup();
extern void _ZN25QSequentialAnimationGroupD0Ev(void* qthis);
  // proto:  void QSequentialAnimationGroup::QSequentialAnimationGroup(const QSequentialAnimationGroup & );
extern void* dector_ZN25QSequentialAnimationGroupC1ERKS_(void* arg0);
extern void _ZN25QSequentialAnimationGroupC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QSequentialAnimationGroup::metaObject();
extern void _ZNK25QSequentialAnimationGroup10metaObjectEv(void* qthis);
  // proto:  QAbstractAnimation * QSequentialAnimationGroup::currentAnimation();
extern void _ZNK25QSequentialAnimationGroup16currentAnimationEv(void* qthis);
  // proto:  int QSequentialAnimationGroup::duration();
extern void _ZNK25QSequentialAnimationGroup8durationEv(void* qthis);
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

// class sizeof(QSequentialAnimationGroup)=1
type QSequentialAnimationGroup struct {
  /*qbase*/ QAnimationGroup;
  qclsinst uint64 /* *mut c_void*/;
//  _currentAnimationChanged QSequentialAnimationGroup_currentAnimationChanged_signal;
}

  // proto:  QPauseAnimation * QSequentialAnimationGroup::insertPause(int index, int msecs);
func (this *QSequentialAnimationGroup) insertPause(args ...interface{}) () {
  // insertPause(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroup11insertPauseEii
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "insertPause", args)
  }

}

  // proto:  void QSequentialAnimationGroup::QSequentialAnimationGroup(QObject * parent);
func NewQSequentialAnimationGroup(args ...interface{}) QSequentialAnimationGroup {
  return QSequentialAnimationGroup{}
}

  // proto:  QPauseAnimation * QSequentialAnimationGroup::addPause(int msecs);
func (this *QSequentialAnimationGroup) addPause(args ...interface{}) () {
  // addPause(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QSequentialAnimationGroup8addPauseEi
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "addPause", args)
  }

}

  // proto:  void QSequentialAnimationGroup::~QSequentialAnimationGroup();
func (this *QSequentialAnimationGroup) FreeQSequentialAnimationGroup(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "~QSequentialAnimationGroup", args)
  }

}

  // proto:  const QMetaObject * QSequentialAnimationGroup::metaObject();
func (this *QSequentialAnimationGroup) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup10metaObjectEv
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "metaObject", args)
  }

}

  // proto:  QAbstractAnimation * QSequentialAnimationGroup::currentAnimation();
func (this *QSequentialAnimationGroup) currentAnimation(args ...interface{}) () {
  // currentAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup16currentAnimationEv
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "currentAnimation", args)
  }

}

  // proto:  int QSequentialAnimationGroup::duration();
func (this *QSequentialAnimationGroup) duration(args ...interface{}) () {
  // duration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QSequentialAnimationGroup8durationEv
  default:
    qtrt.ErrorResolve("QSequentialAnimationGroup", "duration", args)
  }

}

// <= body block end

