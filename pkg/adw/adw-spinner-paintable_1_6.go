// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeSpinnerPaintable = coreglib.Type(C.adw_spinner_paintable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpinnerPaintable, F: marshalSpinnerPaintable},
	})
}

// SpinnerPaintableOverrides contains methods that are overridable.
type SpinnerPaintableOverrides struct {
}

func defaultSpinnerPaintableOverrides(v *SpinnerPaintable) SpinnerPaintableOverrides {
	return SpinnerPaintableOverrides{}
}

// SpinnerPaintable: paintable showing a loading spinner.
//
// <picture> <source srcset="spinner-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="spinner.png" alt="spinner"> </picture>
//
// AdwSpinnerPaintable size varies depending on the available space, but is
// capped at 64×64 pixels.
//
// To be able to animate, AdwSpinnerPaintable needs a widget. It will be
// animated according to that widget's frame clock, and only if that widget is
// mapped. Ideally it should be the same widget the paintable is displayed in,
// but that's not a requirement.
//
// Most applications should be using spinner instead. AdwSpinnerPaintable is
// provided for the cases where using a widget is impractical or impossible,
// such as statuspage:paintable:
//
//	<object class="AdwStatusPage" id="status_page">
//	  <property name="paintable">
//	    <object class="AdwSpinnerPaintable">
//	      <property name="widget">status_page</property>
//	    </object>
//	  </property>
//	  <!-- ... -->
//	</object>.
type SpinnerPaintable struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gtk.SymbolicPaintable
}

var (
	_ coreglib.Objector = (*SpinnerPaintable)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SpinnerPaintable, *SpinnerPaintableClass, SpinnerPaintableOverrides](
		GTypeSpinnerPaintable,
		initSpinnerPaintableClass,
		wrapSpinnerPaintable,
		defaultSpinnerPaintableOverrides,
	)
}

func initSpinnerPaintableClass(gclass unsafe.Pointer, overrides SpinnerPaintableOverrides, classInitFunc func(*SpinnerPaintableClass)) {
	if classInitFunc != nil {
		class := (*SpinnerPaintableClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSpinnerPaintable(obj *coreglib.Object) *SpinnerPaintable {
	return &SpinnerPaintable{
		Object: obj,
		SymbolicPaintable: gtk.SymbolicPaintable{
			Paintable: gdk.Paintable{
				Object: obj,
			},
		},
	}
}

func marshalSpinnerPaintable(p uintptr) (interface{}, error) {
	return wrapSpinnerPaintable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSpinnerPaintable creates a new AdwSpinnerPaintable for widget.
//
// The function takes the following parameters:
//
//   - widget (optional) used for frame clock.
//
// The function returns the following values:
//
//   - spinnerPaintable: newly created AdwSpinnerPaintable.
func NewSpinnerPaintable(widget gtk.Widgetter) *SpinnerPaintable {
	var _arg1 *C.GtkWidget           // out
	var _cret *C.AdwSpinnerPaintable // in

	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	_cret = C.adw_spinner_paintable_new(_arg1)
	runtime.KeepAlive(widget)

	var _spinnerPaintable *SpinnerPaintable // out

	_spinnerPaintable = wrapSpinnerPaintable(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _spinnerPaintable
}

// Widget gets the widget used for frame clock.
//
// The function returns the following values:
//
//   - widget (optional): widget.
func (self *SpinnerPaintable) Widget() gtk.Widgetter {
	var _arg0 *C.AdwSpinnerPaintable // out
	var _cret *C.GtkWidget           // in

	_arg0 = (*C.AdwSpinnerPaintable)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spinner_paintable_get_widget(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetWidget sets the widget used for frame clock.
//
// The function takes the following parameters:
//
//   - widget (optional) to use for frame clock.
func (self *SpinnerPaintable) SetWidget(widget gtk.Widgetter) {
	var _arg0 *C.AdwSpinnerPaintable // out
	var _arg1 *C.GtkWidget           // out

	_arg0 = (*C.AdwSpinnerPaintable)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	C.adw_spinner_paintable_set_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}
