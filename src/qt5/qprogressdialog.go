package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qprogressdialog.h
// dst-file: /src/widgets/qprogressdialog.go
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
  // proto:  void QProgressDialog::setAutoClose(bool close);
extern void _ZN15QProgressDialog12setAutoCloseEb(void* qthis, bool arg0);
  // proto:  void QProgressDialog::open(QObject * receiver, const char * member);
extern void _ZN15QProgressDialog4openEP7QObjectPKc(void* qthis, void* arg0, char* arg1);
  // proto:  void QProgressDialog::setMaximum(int maximum);
extern void _ZN15QProgressDialog10setMaximumEi(void* qthis, int arg0);
  // proto:  void QProgressDialog::setMinimum(int minimum);
extern void _ZN15QProgressDialog10setMinimumEi(void* qthis, int arg0);
  // proto:  void QProgressDialog::setLabelText(const QString & text);
extern void _ZN15QProgressDialog12setLabelTextERK7QString(void* qthis, void* arg0);
  // proto:  bool QProgressDialog::wasCanceled();
extern void _ZNK15QProgressDialog11wasCanceledEv(void* qthis);
  // proto:  void QProgressDialog::~QProgressDialog();
extern void _ZN15QProgressDialogD0Ev(void* qthis);
  // proto:  int QProgressDialog::minimumDuration();
extern void _ZNK15QProgressDialog15minimumDurationEv(void* qthis);
  // proto:  void QProgressDialog::setMinimumDuration(int ms);
extern void _ZN15QProgressDialog18setMinimumDurationEi(void* qthis, int arg0);
  // proto:  int QProgressDialog::maximum();
extern void _ZNK15QProgressDialog7maximumEv(void* qthis);
  // proto:  void QProgressDialog::setBar(QProgressBar * bar);
extern void _ZN15QProgressDialog6setBarEP12QProgressBar(void* qthis, void* arg0);
  // proto:  void QProgressDialog::cancel();
extern void _ZN15QProgressDialog6cancelEv(void* qthis);
  // proto:  bool QProgressDialog::autoClose();
extern void _ZNK15QProgressDialog9autoCloseEv(void* qthis);
  // proto:  int QProgressDialog::minimum();
extern void _ZNK15QProgressDialog7minimumEv(void* qthis);
  // proto:  bool QProgressDialog::autoReset();
extern void _ZNK15QProgressDialog9autoResetEv(void* qthis);
  // proto:  void QProgressDialog::reset();
extern void _ZN15QProgressDialog5resetEv(void* qthis);
  // proto:  void QProgressDialog::QProgressDialog(const QProgressDialog & );
extern void* dector_ZN15QProgressDialogC1ERKS_(void* arg0);
extern void _ZN15QProgressDialogC1ERKS_(void* qthis, void* arg0);
  // proto:  void QProgressDialog::setRange(int minimum, int maximum);
extern void _ZN15QProgressDialog8setRangeEii(void* qthis, int arg0, int arg1);
  // proto:  void QProgressDialog::setCancelButtonText(const QString & text);
extern void _ZN15QProgressDialog19setCancelButtonTextERK7QString(void* qthis, void* arg0);
  // proto:  QSize QProgressDialog::sizeHint();
extern void _ZNK15QProgressDialog8sizeHintEv(void* qthis);
  // proto:  QString QProgressDialog::labelText();
extern void _ZNK15QProgressDialog9labelTextEv(void* qthis);
  // proto:  void QProgressDialog::setLabel(QLabel * label);
extern void _ZN15QProgressDialog8setLabelEP6QLabel(void* qthis, void* arg0);
  // proto:  const QMetaObject * QProgressDialog::metaObject();
extern void _ZNK15QProgressDialog10metaObjectEv(void* qthis);
  // proto:  void QProgressDialog::setAutoReset(bool reset);
extern void _ZN15QProgressDialog12setAutoResetEb(void* qthis, bool arg0);
  // proto:  int QProgressDialog::value();
extern void _ZNK15QProgressDialog5valueEv(void* qthis);
  // proto:  void QProgressDialog::setCancelButton(QPushButton * button);
extern void _ZN15QProgressDialog15setCancelButtonEP11QPushButton(void* qthis, void* arg0);
  // proto:  void QProgressDialog::setValue(int progress);
extern void _ZN15QProgressDialog8setValueEi(void* qthis, int arg0);
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

// class sizeof(QProgressDialog)=1
type QProgressDialog struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
//  _canceled QProgressDialog_canceled_signal;
}

  // proto:  void QProgressDialog::setAutoClose(bool close);
func (this *QProgressDialog) setAutoClose(args ...interface{}) () {
  // setAutoClose(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setAutoCloseEb
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoClose", args)
  }

}

  // proto:  void QProgressDialog::open(QObject * receiver, const char * member);
func (this *QProgressDialog) open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog4openEP7QObjectPKc
  default:
    qtrt.ErrorResolve("QProgressDialog", "open", args)
  }

}

  // proto:  void QProgressDialog::setMaximum(int maximum);
func (this *QProgressDialog) setMaximum(args ...interface{}) () {
  // setMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog10setMaximumEi
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMaximum", args)
  }

}

  // proto:  void QProgressDialog::setMinimum(int minimum);
func (this *QProgressDialog) setMinimum(args ...interface{}) () {
  // setMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog10setMinimumEi
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimum", args)
  }

}

  // proto:  void QProgressDialog::setLabelText(const QString & text);
func (this *QProgressDialog) setLabelText(args ...interface{}) () {
  // setLabelText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setLabelTextERK7QString
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabelText", args)
  }

}

  // proto:  bool QProgressDialog::wasCanceled();
func (this *QProgressDialog) wasCanceled(args ...interface{}) () {
  // wasCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog11wasCanceledEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "wasCanceled", args)
  }

}

  // proto:  void QProgressDialog::~QProgressDialog();
func (this *QProgressDialog) FreeQProgressDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProgressDialog", "~QProgressDialog", args)
  }

}

  // proto:  int QProgressDialog::minimumDuration();
func (this *QProgressDialog) minimumDuration(args ...interface{}) () {
  // minimumDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog15minimumDurationEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimumDuration", args)
  }

}

  // proto:  void QProgressDialog::setMinimumDuration(int ms);
func (this *QProgressDialog) setMinimumDuration(args ...interface{}) () {
  // setMinimumDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog18setMinimumDurationEi
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimumDuration", args)
  }

}

  // proto:  int QProgressDialog::maximum();
func (this *QProgressDialog) maximum(args ...interface{}) () {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7maximumEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "maximum", args)
  }

}

  // proto:  void QProgressDialog::setBar(QProgressBar * bar);
func (this *QProgressDialog) setBar(args ...interface{}) () {
  // setBar(class QProgressBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProgressBar{}) // "QProgressBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog6setBarEP12QProgressBar
  default:
    qtrt.ErrorResolve("QProgressDialog", "setBar", args)
  }

}

  // proto:  void QProgressDialog::cancel();
func (this *QProgressDialog) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog6cancelEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "cancel", args)
  }

}

  // proto:  bool QProgressDialog::autoClose();
func (this *QProgressDialog) autoClose(args ...interface{}) () {
  // autoClose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoCloseEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoClose", args)
  }

}

  // proto:  int QProgressDialog::minimum();
func (this *QProgressDialog) minimum(args ...interface{}) () {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7minimumEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimum", args)
  }

}

  // proto:  bool QProgressDialog::autoReset();
func (this *QProgressDialog) autoReset(args ...interface{}) () {
  // autoReset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoResetEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoReset", args)
  }

}

  // proto:  void QProgressDialog::reset();
func (this *QProgressDialog) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog5resetEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "reset", args)
  }

}

  // proto:  void QProgressDialog::QProgressDialog(const QProgressDialog & );
func NewQProgressDialog(args ...interface{}) QProgressDialog {
  return QProgressDialog{}
}

  // proto:  void QProgressDialog::setRange(int minimum, int maximum);
func (this *QProgressDialog) setRange(args ...interface{}) () {
  // setRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setRangeEii
  default:
    qtrt.ErrorResolve("QProgressDialog", "setRange", args)
  }

}

  // proto:  void QProgressDialog::setCancelButtonText(const QString & text);
func (this *QProgressDialog) setCancelButtonText(args ...interface{}) () {
  // setCancelButtonText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog19setCancelButtonTextERK7QString
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButtonText", args)
  }

}

  // proto:  QSize QProgressDialog::sizeHint();
func (this *QProgressDialog) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog8sizeHintEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "sizeHint", args)
  }

}

  // proto:  QString QProgressDialog::labelText();
func (this *QProgressDialog) labelText(args ...interface{}) () {
  // labelText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9labelTextEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "labelText", args)
  }

}

  // proto:  void QProgressDialog::setLabel(QLabel * label);
func (this *QProgressDialog) setLabel(args ...interface{}) () {
  // setLabel(class QLabel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLabel{}) // "QLabel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setLabelEP6QLabel
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabel", args)
  }

}

  // proto:  const QMetaObject * QProgressDialog::metaObject();
func (this *QProgressDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog10metaObjectEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "metaObject", args)
  }

}

  // proto:  void QProgressDialog::setAutoReset(bool reset);
func (this *QProgressDialog) setAutoReset(args ...interface{}) () {
  // setAutoReset(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setAutoResetEb
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoReset", args)
  }

}

  // proto:  int QProgressDialog::value();
func (this *QProgressDialog) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog5valueEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "value", args)
  }

}

  // proto:  void QProgressDialog::setCancelButton(QPushButton * button);
func (this *QProgressDialog) setCancelButton(args ...interface{}) () {
  // setCancelButton(class QPushButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPushButton{}) // "QPushButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog15setCancelButtonEP11QPushButton
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButton", args)
  }

}

  // proto:  void QProgressDialog::setValue(int progress);
func (this *QProgressDialog) setValue(args ...interface{}) () {
  // setValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setValueEi
  default:
    qtrt.ErrorResolve("QProgressDialog", "setValue", args)
  }

}

// <= body block end

