package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qmenu.h
// dst-file: /src/widgets/qmenu.go
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
  // proto:  bool QMenu::isTearOffEnabled();
extern void _ZNK5QMenu16isTearOffEnabledEv(void* qthis);
  // proto:  bool QMenu::toolTipsVisible();
extern void _ZNK5QMenu15toolTipsVisibleEv(void* qthis);
  // proto:  QAction * QMenu::menuAction();
extern void _ZNK5QMenu10menuActionEv(void* qthis);
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text);
extern void _ZN5QMenu9addActionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QMenu::setTearOffEnabled(bool );
extern void _ZN5QMenu17setTearOffEnabledEb(void* qthis, bool arg0);
  // proto:  QAction * QMenu::addSection(const QString & text);
extern void _ZN5QMenu10addSectionERK7QString(void* qthis, void* arg0);
  // proto:  const QMetaObject * QMenu::metaObject();
extern void _ZNK5QMenu10metaObjectEv(void* qthis);
  // proto:  void QMenu::clear();
extern void _ZN5QMenu5clearEv(void* qthis);
  // proto:  QAction * QMenu::insertMenu(QAction * before, QMenu * menu);
extern void _ZN5QMenu10insertMenuEP7QActionPS_(void* qthis, void* arg0, void* arg1);
  // proto:  QIcon QMenu::icon();
extern void _ZNK5QMenu4iconEv(void* qthis);
  // proto:  QAction * QMenu::insertSection(QAction * before, const QString & text);
extern void _ZN5QMenu13insertSectionEP7QActionRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QPlatformMenu * QMenu::platformMenu();
extern void _ZN5QMenu12platformMenuEv(void* qthis);
  // proto:  void QMenu::setNoReplayFor(QWidget * widget);
extern void _ZN5QMenu14setNoReplayForEP7QWidget(void* qthis, void* arg0);
  // proto:  void QMenu::setIcon(const QIcon & icon);
extern void _ZN5QMenu7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  QAction * QMenu::exec(const QPoint & pos, QAction * at);
extern void _ZN5QMenu4execERK6QPointP7QAction(void* qthis, void* arg0, void* arg1);
  // proto:  bool QMenu::separatorsCollapsible();
extern void _ZNK5QMenu21separatorsCollapsibleEv(void* qthis);
  // proto:  QMenu * QMenu::addMenu(const QString & title);
extern void _ZN5QMenu7addMenuERK7QString(void* qthis, void* arg0);
  // proto:  QAction * QMenu::addSeparator();
extern void _ZN5QMenu12addSeparatorEv(void* qthis);
  // proto:  void QMenu::hideTearOffMenu();
extern void _ZN5QMenu15hideTearOffMenuEv(void* qthis);
  // proto:  QAction * QMenu::addAction(const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void _ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, char* arg2, void* arg3);
  // proto:  void QMenu::QMenu(QWidget * parent);
extern void* dector_ZN5QMenuC1EP7QWidget(void* arg0);
extern void _ZN5QMenuC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QMenu::setActiveAction(QAction * act);
extern void _ZN5QMenu15setActiveActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QMenu::setSeparatorsCollapsible(bool collapse);
extern void _ZN5QMenu24setSeparatorsCollapsibleEb(void* qthis, bool arg0);
  // proto:  void QMenu::QMenu(const QMenu & );
extern void* dector_ZN5QMenuC1ERKS_(void* arg0);
extern void _ZN5QMenuC1ERKS_(void* qthis, void* arg0);
  // proto:  QAction * QMenu::addAction(const QString & text);
extern void _ZN5QMenu9addActionERK7QString(void* qthis, void* arg0);
  // proto:  QAction * QMenu::activeAction();
extern void _ZNK5QMenu12activeActionEv(void* qthis);
  // proto:  bool QMenu::isEmpty();
extern void _ZNK5QMenu7isEmptyEv(void* qthis);
  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text, const QObject * receiver, const char * member, const QKeySequence & shortcut);
extern void _ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence(void* qthis, void* arg0, void* arg1, void* arg2, char* arg3, void* arg4);
  // proto:  QRect QMenu::actionGeometry(QAction * );
extern void _ZNK5QMenu14actionGeometryEP7QAction(void* qthis, void* arg0);
  // proto:  void QMenu::QMenu(const QString & title, QWidget * parent);
extern void* dector_ZN5QMenuC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN5QMenuC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  QAction * QMenu::insertSeparator(QAction * before);
extern void _ZN5QMenu15insertSeparatorEP7QAction(void* qthis, void* arg0);
  // proto:  QAction * QMenu::addSection(const QIcon & icon, const QString & text);
extern void _ZN5QMenu10addSectionERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  bool QMenu::isTearOffMenuVisible();
extern void _ZNK5QMenu20isTearOffMenuVisibleEv(void* qthis);
  // proto:  void QMenu::~QMenu();
extern void _ZN5QMenuD0Ev(void* qthis);
  // proto:  QString QMenu::title();
extern void _ZNK5QMenu5titleEv(void* qthis);
  // proto:  QAction * QMenu::defaultAction();
extern void _ZNK5QMenu13defaultActionEv(void* qthis);
  // proto:  QAction * QMenu::addMenu(QMenu * menu);
extern void _ZN5QMenu7addMenuEPS_(void* qthis, void* arg0);
  // proto:  QSize QMenu::sizeHint();
extern void _ZNK5QMenu8sizeHintEv(void* qthis);
  // proto:  void QMenu::setDefaultAction(QAction * );
extern void _ZN5QMenu16setDefaultActionEP7QAction(void* qthis, void* arg0);
  // proto:  QAction * QMenu::actionAt(const QPoint & );
extern void _ZNK5QMenu8actionAtERK6QPoint(void* qthis, void* arg0);
  // proto:  QAction * QMenu::insertSection(QAction * before, const QIcon & icon, const QString & text);
extern void _ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QMenu::popup(const QPoint & pos, QAction * at);
extern void _ZN5QMenu5popupERK6QPointP7QAction(void* qthis, void* arg0, void* arg1);
  // proto:  void QMenu::setToolTipsVisible(bool visible);
extern void _ZN5QMenu18setToolTipsVisibleEb(void* qthis, bool arg0);
  // proto:  void QMenu::setTitle(const QString & title);
extern void _ZN5QMenu8setTitleERK7QString(void* qthis, void* arg0);
  // proto:  QMenu * QMenu::addMenu(const QIcon & icon, const QString & title);
extern void _ZN5QMenu7addMenuERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QAction * QMenu::exec();
extern void _ZN5QMenu4execEv(void* qthis);
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

// class sizeof(QMenu)=1
type QMenu struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToShow QMenu_aboutToShow_signal;
//  _triggered QMenu_triggered_signal;
//  _hovered QMenu_hovered_signal;
//  _aboutToHide QMenu_aboutToHide_signal;
}

  // proto:  bool QMenu::isTearOffEnabled();
func (this *QMenu) isTearOffEnabled(args ...interface{}) () {
  // isTearOffEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu16isTearOffEnabledEv
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffEnabled", args)
  }

}

  // proto:  bool QMenu::toolTipsVisible();
func (this *QMenu) toolTipsVisible(args ...interface{}) () {
  // toolTipsVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu15toolTipsVisibleEv
  default:
    qtrt.ErrorResolve("QMenu", "toolTipsVisible", args)
  }

}

  // proto:  QAction * QMenu::menuAction();
func (this *QMenu) menuAction(args ...interface{}) () {
  // menuAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu10menuActionEv
  default:
    qtrt.ErrorResolve("QMenu", "menuAction", args)
  }

}

  // proto:  QAction * QMenu::addAction(const QIcon & icon, const QString & text);
func (this *QMenu) addAction(args ...interface{}) () {
  // addAction(const class QIcon &, const class QString &)
  // addAction(const class QString &, const class QObject *, const char *, const class QKeySequence &)
  // addAction(const class QString &)
  // addAction(const class QIcon &, const class QString &, const class QObject *, const char *, const class QKeySequence &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1][3] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[3][3] = qtrt.ByteTy(true) // "const char *"
  vtys[3][4] = reflect.TypeOf(QKeySequence{}) // "const QKeySequence &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QString
  case 1:
    // invoke: _ZN5QMenu9addActionERK7QStringPK7QObjectPKcRK12QKeySequence
  case 2:
    // invoke: _ZN5QMenu9addActionERK7QString
  case 3:
    // invoke: _ZN5QMenu9addActionERK5QIconRK7QStringPK7QObjectPKcRK12QKeySequence
  default:
    qtrt.ErrorResolve("QMenu", "addAction", args)
  }

}

  // proto:  void QMenu::setTearOffEnabled(bool );
func (this *QMenu) setTearOffEnabled(args ...interface{}) () {
  // setTearOffEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu17setTearOffEnabledEb
  default:
    qtrt.ErrorResolve("QMenu", "setTearOffEnabled", args)
  }

}

  // proto:  QAction * QMenu::addSection(const QString & text);
func (this *QMenu) addSection(args ...interface{}) () {
  // addSection(const class QString &)
  // addSection(const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu10addSectionERK7QString
  case 1:
    // invoke: _ZN5QMenu10addSectionERK5QIconRK7QString
  default:
    qtrt.ErrorResolve("QMenu", "addSection", args)
  }

}

  // proto:  const QMetaObject * QMenu::metaObject();
func (this *QMenu) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu10metaObjectEv
  default:
    qtrt.ErrorResolve("QMenu", "metaObject", args)
  }

}

  // proto:  void QMenu::clear();
func (this *QMenu) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu5clearEv
  default:
    qtrt.ErrorResolve("QMenu", "clear", args)
  }

}

  // proto:  QAction * QMenu::insertMenu(QAction * before, QMenu * menu);
func (this *QMenu) insertMenu(args ...interface{}) () {
  // insertMenu(class QAction *, class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu10insertMenuEP7QActionPS_
  default:
    qtrt.ErrorResolve("QMenu", "insertMenu", args)
  }

}

  // proto:  QIcon QMenu::icon();
func (this *QMenu) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu4iconEv
  default:
    qtrt.ErrorResolve("QMenu", "icon", args)
  }

}

  // proto:  QAction * QMenu::insertSection(QAction * before, const QString & text);
func (this *QMenu) insertSection(args ...interface{}) () {
  // insertSection(class QAction *, const class QString &)
  // insertSection(class QAction *, const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu13insertSectionEP7QActionRK7QString
  case 1:
    // invoke: _ZN5QMenu13insertSectionEP7QActionRK5QIconRK7QString
  default:
    qtrt.ErrorResolve("QMenu", "insertSection", args)
  }

}

  // proto:  QPlatformMenu * QMenu::platformMenu();
func (this *QMenu) platformMenu(args ...interface{}) () {
  // platformMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu12platformMenuEv
  default:
    qtrt.ErrorResolve("QMenu", "platformMenu", args)
  }

}

  // proto:  void QMenu::setNoReplayFor(QWidget * widget);
func (this *QMenu) setNoReplayFor(args ...interface{}) () {
  // setNoReplayFor(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu14setNoReplayForEP7QWidget
  default:
    qtrt.ErrorResolve("QMenu", "setNoReplayFor", args)
  }

}

  // proto:  void QMenu::setIcon(const QIcon & icon);
func (this *QMenu) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu7setIconERK5QIcon
  default:
    qtrt.ErrorResolve("QMenu", "setIcon", args)
  }

}

  // proto:  QAction * QMenu::exec(const QPoint & pos, QAction * at);
func (this *QMenu) exec(args ...interface{}) () {
  // exec(const class QPoint &, class QAction *)
  // exec(QList<class QAction *>, const class QPoint &, class QAction *, class QWidget *)
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "QList<QAction *>"
  vtys[1][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][2] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[1][3] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu4execERK6QPointP7QAction
  case 1:
    // invoke: _ZN5QMenu4execE5QListIP7QActionERK6QPointS2_P7QWidget
  case 2:
    // invoke: _ZN5QMenu4execEv
  default:
    qtrt.ErrorResolve("QMenu", "exec", args)
  }

}

  // proto:  bool QMenu::separatorsCollapsible();
func (this *QMenu) separatorsCollapsible(args ...interface{}) () {
  // separatorsCollapsible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu21separatorsCollapsibleEv
  default:
    qtrt.ErrorResolve("QMenu", "separatorsCollapsible", args)
  }

}

  // proto:  QMenu * QMenu::addMenu(const QString & title);
func (this *QMenu) addMenu(args ...interface{}) () {
  // addMenu(const class QString &)
  // addMenu(class QMenu *)
  // addMenu(const class QIcon &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMenu{}) // "QMenu *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu7addMenuERK7QString
  case 1:
    // invoke: _ZN5QMenu7addMenuEPS_
  case 2:
    // invoke: _ZN5QMenu7addMenuERK5QIconRK7QString
  default:
    qtrt.ErrorResolve("QMenu", "addMenu", args)
  }

}

  // proto:  QAction * QMenu::addSeparator();
func (this *QMenu) addSeparator(args ...interface{}) () {
  // addSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu12addSeparatorEv
  default:
    qtrt.ErrorResolve("QMenu", "addSeparator", args)
  }

}

  // proto:  void QMenu::hideTearOffMenu();
func (this *QMenu) hideTearOffMenu(args ...interface{}) () {
  // hideTearOffMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu15hideTearOffMenuEv
  default:
    qtrt.ErrorResolve("QMenu", "hideTearOffMenu", args)
  }

}

  // proto:  void QMenu::QMenu(QWidget * parent);
func NewQMenu(args ...interface{}) QMenu {
  return QMenu{}
}

  // proto:  void QMenu::setActiveAction(QAction * act);
func (this *QMenu) setActiveAction(args ...interface{}) () {
  // setActiveAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu15setActiveActionEP7QAction
  default:
    qtrt.ErrorResolve("QMenu", "setActiveAction", args)
  }

}

  // proto:  void QMenu::setSeparatorsCollapsible(bool collapse);
func (this *QMenu) setSeparatorsCollapsible(args ...interface{}) () {
  // setSeparatorsCollapsible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu24setSeparatorsCollapsibleEb
  default:
    qtrt.ErrorResolve("QMenu", "setSeparatorsCollapsible", args)
  }

}

  // proto:  QAction * QMenu::activeAction();
func (this *QMenu) activeAction(args ...interface{}) () {
  // activeAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu12activeActionEv
  default:
    qtrt.ErrorResolve("QMenu", "activeAction", args)
  }

}

  // proto:  bool QMenu::isEmpty();
func (this *QMenu) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu7isEmptyEv
  default:
    qtrt.ErrorResolve("QMenu", "isEmpty", args)
  }

}

  // proto:  QRect QMenu::actionGeometry(QAction * );
func (this *QMenu) actionGeometry(args ...interface{}) () {
  // actionGeometry(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu14actionGeometryEP7QAction
  default:
    qtrt.ErrorResolve("QMenu", "actionGeometry", args)
  }

}

  // proto:  QAction * QMenu::insertSeparator(QAction * before);
func (this *QMenu) insertSeparator(args ...interface{}) () {
  // insertSeparator(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu15insertSeparatorEP7QAction
  default:
    qtrt.ErrorResolve("QMenu", "insertSeparator", args)
  }

}

  // proto:  bool QMenu::isTearOffMenuVisible();
func (this *QMenu) isTearOffMenuVisible(args ...interface{}) () {
  // isTearOffMenuVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu20isTearOffMenuVisibleEv
  default:
    qtrt.ErrorResolve("QMenu", "isTearOffMenuVisible", args)
  }

}

  // proto:  void QMenu::~QMenu();
func (this *QMenu) FreeQMenu(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMenu", "~QMenu", args)
  }

}

  // proto:  QString QMenu::title();
func (this *QMenu) title(args ...interface{}) () {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu5titleEv
  default:
    qtrt.ErrorResolve("QMenu", "title", args)
  }

}

  // proto:  QAction * QMenu::defaultAction();
func (this *QMenu) defaultAction(args ...interface{}) () {
  // defaultAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu13defaultActionEv
  default:
    qtrt.ErrorResolve("QMenu", "defaultAction", args)
  }

}

  // proto:  QSize QMenu::sizeHint();
func (this *QMenu) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu8sizeHintEv
  default:
    qtrt.ErrorResolve("QMenu", "sizeHint", args)
  }

}

  // proto:  void QMenu::setDefaultAction(QAction * );
func (this *QMenu) setDefaultAction(args ...interface{}) () {
  // setDefaultAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu16setDefaultActionEP7QAction
  default:
    qtrt.ErrorResolve("QMenu", "setDefaultAction", args)
  }

}

  // proto:  QAction * QMenu::actionAt(const QPoint & );
func (this *QMenu) actionAt(args ...interface{}) () {
  // actionAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QMenu8actionAtERK6QPoint
  default:
    qtrt.ErrorResolve("QMenu", "actionAt", args)
  }

}

  // proto:  void QMenu::popup(const QPoint & pos, QAction * at);
func (this *QMenu) popup(args ...interface{}) () {
  // popup(const class QPoint &, class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu5popupERK6QPointP7QAction
  default:
    qtrt.ErrorResolve("QMenu", "popup", args)
  }

}

  // proto:  void QMenu::setToolTipsVisible(bool visible);
func (this *QMenu) setToolTipsVisible(args ...interface{}) () {
  // setToolTipsVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu18setToolTipsVisibleEb
  default:
    qtrt.ErrorResolve("QMenu", "setToolTipsVisible", args)
  }

}

  // proto:  void QMenu::setTitle(const QString & title);
func (this *QMenu) setTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QMenu8setTitleERK7QString
  default:
    qtrt.ErrorResolve("QMenu", "setTitle", args)
  }

}

// <= body block end

