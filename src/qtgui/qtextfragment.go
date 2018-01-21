package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QTextFragment struct {
	*qtrt.CObject
}

func (this *QTextFragment) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQTextFragmentFromPointer(cthis unsafe.Pointer) *QTextFragment {
	return &QTextFragment{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextobject.h:307
// index:0
// Public inline
// void QTextFragment()
func NewQTextFragment() *QTextFragment {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QTextFragmentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFragmentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:311
// index:0
// Public inline
// bool isValid()
func (this *QTextFragment) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:317
// index:0
// Public
// int position()
func (this *QTextFragment) Position() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:318
// index:0
// Public
// int length()
func (this *QTextFragment) Length() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment6lengthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:319
// index:0
// Public
// bool contains(int)
func (this *QTextFragment) Contains(position int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment8containsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextobject.h:321
// index:0
// Public
// QTextCharFormat charFormat()
func (this *QTextFragment) CharFormat() *QTextCharFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment10charFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:322
// index:0
// Public
// int charFormatIndex()
func (this *QTextFragment) CharFormatIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment15charFormatIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextobject.h:323
// index:0
// Public
// QString text()
func (this *QTextFragment) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QTextFragment4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end