package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qlayoutitem.h
// dst-file: /src/widgets/qlayoutitem.go
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
  // proto:  QSpacerItem * QLayoutItem::spacerItem();
extern void _ZN11QLayoutItem10spacerItemEv(void* qthis);
  // proto:  QSize QLayoutItem::minimumSize();
extern void _ZNK11QLayoutItem11minimumSizeEv(void* qthis);
  // proto:  QWidget * QLayoutItem::widget();
extern void _ZN11QLayoutItem6widgetEv(void* qthis);
  // proto:  void QLayoutItem::invalidate();
extern void _ZN11QLayoutItem10invalidateEv(void* qthis);
  // proto:  void QLayoutItem::setGeometry(const QRect & );
extern void _ZN11QLayoutItem11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  QLayout * QLayoutItem::layout();
extern void _ZN11QLayoutItem6layoutEv(void* qthis);
  // proto:  bool QLayoutItem::isEmpty();
extern void _ZNK11QLayoutItem7isEmptyEv(void* qthis);
  // proto:  QSize QLayoutItem::sizeHint();
extern void _ZNK11QLayoutItem8sizeHintEv(void* qthis);
  // proto:  void QLayoutItem::~QLayoutItem();
extern void _ZN11QLayoutItemD0Ev(void* qthis);
  // proto:  bool QLayoutItem::hasHeightForWidth();
extern void _ZNK11QLayoutItem17hasHeightForWidthEv(void* qthis);
  // proto:  int QLayoutItem::heightForWidth(int );
extern void _ZNK11QLayoutItem14heightForWidthEi(void* qthis, int arg0);
  // proto:  QRect QLayoutItem::geometry();
extern void _ZNK11QLayoutItem8geometryEv(void* qthis);
  // proto:  QSize QLayoutItem::maximumSize();
extern void _ZNK11QLayoutItem11maximumSizeEv(void* qthis);
  // proto:  int QLayoutItem::minimumHeightForWidth(int );
extern void _ZNK11QLayoutItem21minimumHeightForWidthEi(void* qthis, int arg0);
  // proto:  QSize QSpacerItem::minimumSize();
extern void _ZNK11QSpacerItem11minimumSizeEv(void* qthis);
  // proto:  QSizePolicy QSpacerItem::sizePolicy();
extern void _ZNK11QSpacerItem10sizePolicyEv(void* qthis);
  // proto:  void QSpacerItem::~QSpacerItem();
extern void _ZN11QSpacerItemD0Ev(void* qthis);
  // proto:  QSize QSpacerItem::sizeHint();
extern void _ZNK11QSpacerItem8sizeHintEv(void* qthis);
  // proto:  QSize QSpacerItem::maximumSize();
extern void _ZNK11QSpacerItem11maximumSizeEv(void* qthis);
  // proto:  bool QSpacerItem::isEmpty();
extern void _ZNK11QSpacerItem7isEmptyEv(void* qthis);
  // proto:  QRect QSpacerItem::geometry();
extern void _ZNK11QSpacerItem8geometryEv(void* qthis);
  // proto:  void QSpacerItem::setGeometry(const QRect & );
extern void _ZN11QSpacerItem11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  QSpacerItem * QSpacerItem::spacerItem();
extern void _ZN11QSpacerItem10spacerItemEv(void* qthis);
  // proto:  QSize QWidgetItem::sizeHint();
extern void _ZNK11QWidgetItem8sizeHintEv(void* qthis);
  // proto:  QSize QWidgetItem::minimumSize();
extern void _ZNK11QWidgetItem11minimumSizeEv(void* qthis);
  // proto:  bool QWidgetItem::hasHeightForWidth();
extern void _ZNK11QWidgetItem17hasHeightForWidthEv(void* qthis);
  // proto:  void QWidgetItem::~QWidgetItem();
extern void _ZN11QWidgetItemD0Ev(void* qthis);
  // proto:  void QWidgetItem::QWidgetItem(QWidget * w);
extern void* dector_ZN11QWidgetItemC1EP7QWidget(void* arg0);
extern void _ZN11QWidgetItemC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QWidget * QWidgetItem::widget();
extern void _ZN11QWidgetItem6widgetEv(void* qthis);
  // proto:  void QWidgetItem::setGeometry(const QRect & );
extern void _ZN11QWidgetItem11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  int QWidgetItem::heightForWidth(int );
extern void _ZNK11QWidgetItem14heightForWidthEi(void* qthis, int arg0);
  // proto:  void QWidgetItem::QWidgetItem(const QWidgetItem & );
extern void* dector_ZN11QWidgetItemC1ERKS_(void* arg0);
extern void _ZN11QWidgetItemC1ERKS_(void* qthis, void* arg0);
  // proto:  QSize QWidgetItem::maximumSize();
extern void _ZNK11QWidgetItem11maximumSizeEv(void* qthis);
  // proto:  bool QWidgetItem::isEmpty();
extern void _ZNK11QWidgetItem7isEmptyEv(void* qthis);
  // proto:  QRect QWidgetItem::geometry();
extern void _ZNK11QWidgetItem8geometryEv(void* qthis);
  // proto:  QSize QWidgetItemV2::sizeHint();
extern void _ZNK13QWidgetItemV28sizeHintEv(void* qthis);
  // proto:  QSize QWidgetItemV2::minimumSize();
extern void _ZNK13QWidgetItemV211minimumSizeEv(void* qthis);
  // proto:  int QWidgetItemV2::heightForWidth(int width);
extern void _ZNK13QWidgetItemV214heightForWidthEi(void* qthis, int arg0);
  // proto:  void QWidgetItemV2::~QWidgetItemV2();
extern void _ZN13QWidgetItemV2D0Ev(void* qthis);
  // proto:  void QWidgetItemV2::QWidgetItemV2(QWidget * widget);
extern void* dector_ZN13QWidgetItemV2C1EP7QWidget(void* arg0);
extern void _ZN13QWidgetItemV2C1EP7QWidget(void* qthis, void* arg0);
  // proto:  QSize QWidgetItemV2::maximumSize();
extern void _ZNK13QWidgetItemV211maximumSizeEv(void* qthis);
  // proto:  void QWidgetItemV2::QWidgetItemV2(const QWidgetItemV2 & );
extern void* dector_ZN13QWidgetItemV2C1ERKS_(void* arg0);
extern void _ZN13QWidgetItemV2C1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QLayoutItem)=1
type QLayoutItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QSpacerItem)=1
type QSpacerItem struct {
  /*qbase*/ QLayoutItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QWidgetItem)=1
type QWidgetItem struct {
  /*qbase*/ QLayoutItem;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QWidgetItemV2)=1
type QWidgetItemV2 struct {
  /*qbase*/ QWidgetItem;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QSpacerItem * QLayoutItem::spacerItem();
func (this *QLayoutItem) spacerItem(args ...interface{}) () {
  // spacerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem10spacerItemEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "spacerItem", args)
  }

}

  // proto:  QSize QLayoutItem::minimumSize();
func (this *QLayoutItem) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem11minimumSizeEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "minimumSize", args)
  }

}

  // proto:  QWidget * QLayoutItem::widget();
func (this *QLayoutItem) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem6widgetEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "widget", args)
  }

}

  // proto:  void QLayoutItem::invalidate();
func (this *QLayoutItem) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem10invalidateEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "invalidate", args)
  }

}

  // proto:  void QLayoutItem::setGeometry(const QRect & );
func (this *QLayoutItem) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem11setGeometryERK5QRect
  default:
    qtrt.ErrorResolve("QLayoutItem", "setGeometry", args)
  }

}

  // proto:  QLayout * QLayoutItem::layout();
func (this *QLayoutItem) layout(args ...interface{}) () {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QLayoutItem6layoutEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "layout", args)
  }

}

  // proto:  bool QLayoutItem::isEmpty();
func (this *QLayoutItem) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem7isEmptyEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "isEmpty", args)
  }

}

  // proto:  QSize QLayoutItem::sizeHint();
func (this *QLayoutItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem8sizeHintEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "sizeHint", args)
  }

}

  // proto:  void QLayoutItem::~QLayoutItem();
func (this *QLayoutItem) FreeQLayoutItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLayoutItem", "~QLayoutItem", args)
  }

}

  // proto:  bool QLayoutItem::hasHeightForWidth();
func (this *QLayoutItem) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem17hasHeightForWidthEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "hasHeightForWidth", args)
  }

}

  // proto:  int QLayoutItem::heightForWidth(int );
func (this *QLayoutItem) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem14heightForWidthEi
  default:
    qtrt.ErrorResolve("QLayoutItem", "heightForWidth", args)
  }

}

  // proto:  QRect QLayoutItem::geometry();
func (this *QLayoutItem) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem8geometryEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "geometry", args)
  }

}

  // proto:  QSize QLayoutItem::maximumSize();
func (this *QLayoutItem) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem11maximumSizeEv
  default:
    qtrt.ErrorResolve("QLayoutItem", "maximumSize", args)
  }

}

  // proto:  int QLayoutItem::minimumHeightForWidth(int );
func (this *QLayoutItem) minimumHeightForWidth(args ...interface{}) () {
  // minimumHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QLayoutItem21minimumHeightForWidthEi
  default:
    qtrt.ErrorResolve("QLayoutItem", "minimumHeightForWidth", args)
  }

}

  // proto:  QSize QSpacerItem::minimumSize();
func (this *QSpacerItem) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem11minimumSizeEv
  default:
    qtrt.ErrorResolve("QSpacerItem", "minimumSize", args)
  }

}

  // proto:  QSizePolicy QSpacerItem::sizePolicy();
func (this *QSpacerItem) sizePolicy(args ...interface{}) () {
  // sizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem10sizePolicyEv
  default:
    qtrt.ErrorResolve("QSpacerItem", "sizePolicy", args)
  }

}

  // proto:  void QSpacerItem::~QSpacerItem();
func (this *QSpacerItem) FreeQSpacerItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSpacerItem", "~QSpacerItem", args)
  }

}

  // proto:  QSize QSpacerItem::sizeHint();
func (this *QSpacerItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem8sizeHintEv
  default:
    qtrt.ErrorResolve("QSpacerItem", "sizeHint", args)
  }

}

  // proto:  QSize QSpacerItem::maximumSize();
func (this *QSpacerItem) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem11maximumSizeEv
  default:
    qtrt.ErrorResolve("QSpacerItem", "maximumSize", args)
  }

}

  // proto:  bool QSpacerItem::isEmpty();
func (this *QSpacerItem) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem7isEmptyEv
  default:
    qtrt.ErrorResolve("QSpacerItem", "isEmpty", args)
  }

}

  // proto:  QRect QSpacerItem::geometry();
func (this *QSpacerItem) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QSpacerItem8geometryEv
  default:
    qtrt.ErrorResolve("QSpacerItem", "geometry", args)
  }

}

  // proto:  void QSpacerItem::setGeometry(const QRect & );
func (this *QSpacerItem) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSpacerItem11setGeometryERK5QRect
  default:
    qtrt.ErrorResolve("QSpacerItem", "setGeometry", args)
  }

}

  // proto:  QSpacerItem * QSpacerItem::spacerItem();
func (this *QSpacerItem) spacerItem(args ...interface{}) () {
  // spacerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QSpacerItem10spacerItemEv
  default:
    qtrt.ErrorResolve("QSpacerItem", "spacerItem", args)
  }

}

  // proto:  QSize QWidgetItem::sizeHint();
func (this *QWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem8sizeHintEv
  default:
    qtrt.ErrorResolve("QWidgetItem", "sizeHint", args)
  }

}

  // proto:  QSize QWidgetItem::minimumSize();
func (this *QWidgetItem) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem11minimumSizeEv
  default:
    qtrt.ErrorResolve("QWidgetItem", "minimumSize", args)
  }

}

  // proto:  bool QWidgetItem::hasHeightForWidth();
func (this *QWidgetItem) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem17hasHeightForWidthEv
  default:
    qtrt.ErrorResolve("QWidgetItem", "hasHeightForWidth", args)
  }

}

  // proto:  void QWidgetItem::~QWidgetItem();
func (this *QWidgetItem) FreeQWidgetItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidgetItem", "~QWidgetItem", args)
  }

}

  // proto:  void QWidgetItem::QWidgetItem(QWidget * w);
func NewQWidgetItem(args ...interface{}) QWidgetItem {
  return QWidgetItem{}
}

  // proto:  QWidget * QWidgetItem::widget();
func (this *QWidgetItem) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWidgetItem6widgetEv
  default:
    qtrt.ErrorResolve("QWidgetItem", "widget", args)
  }

}

  // proto:  void QWidgetItem::setGeometry(const QRect & );
func (this *QWidgetItem) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QWidgetItem11setGeometryERK5QRect
  default:
    qtrt.ErrorResolve("QWidgetItem", "setGeometry", args)
  }

}

  // proto:  int QWidgetItem::heightForWidth(int );
func (this *QWidgetItem) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem14heightForWidthEi
  default:
    qtrt.ErrorResolve("QWidgetItem", "heightForWidth", args)
  }

}

  // proto:  QSize QWidgetItem::maximumSize();
func (this *QWidgetItem) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem11maximumSizeEv
  default:
    qtrt.ErrorResolve("QWidgetItem", "maximumSize", args)
  }

}

  // proto:  bool QWidgetItem::isEmpty();
func (this *QWidgetItem) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem7isEmptyEv
  default:
    qtrt.ErrorResolve("QWidgetItem", "isEmpty", args)
  }

}

  // proto:  QRect QWidgetItem::geometry();
func (this *QWidgetItem) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QWidgetItem8geometryEv
  default:
    qtrt.ErrorResolve("QWidgetItem", "geometry", args)
  }

}

  // proto:  QSize QWidgetItemV2::sizeHint();
func (this *QWidgetItemV2) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV28sizeHintEv
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "sizeHint", args)
  }

}

  // proto:  QSize QWidgetItemV2::minimumSize();
func (this *QWidgetItemV2) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV211minimumSizeEv
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "minimumSize", args)
  }

}

  // proto:  int QWidgetItemV2::heightForWidth(int width);
func (this *QWidgetItemV2) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV214heightForWidthEi
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "heightForWidth", args)
  }

}

  // proto:  void QWidgetItemV2::~QWidgetItemV2();
func (this *QWidgetItemV2) FreeQWidgetItemV2(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "~QWidgetItemV2", args)
  }

}

  // proto:  void QWidgetItemV2::QWidgetItemV2(QWidget * widget);
func NewQWidgetItemV2(args ...interface{}) QWidgetItemV2 {
  return QWidgetItemV2{}
}

  // proto:  QSize QWidgetItemV2::maximumSize();
func (this *QWidgetItemV2) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetItemV211maximumSizeEv
  default:
    qtrt.ErrorResolve("QWidgetItemV2", "maximumSize", args)
  }

}

// <= body block end

