// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeToastOverlay = coreglib.Type(C.adw_toast_overlay_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeToastOverlay, F: marshalToastOverlay},
	})
}

// ToastOverlayOverrides contains methods that are overridable.
type ToastOverlayOverrides struct {
}

func defaultToastOverlayOverrides(v *ToastOverlay) ToastOverlayOverrides {
	return ToastOverlayOverrides{}
}

// ToastOverlay: widget showing toasts above its content.
//
// <picture> <source srcset="toast-overlay-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="toast-overlay.png"
// alt="toast-overlay"> </picture>
//
// Toasts can be shown with toastoverlay.AddToast.
//
// See toast for details.
//
// CSS nodes
//
//    toastoverlay
//    ├── [child]
//    ├── toast
//    ┊   ├── widget
//    ┊   │   ├── [label.heading]
//        │   ╰── [custom title]
//        ├── [button]
//        ╰── button.circular.flat
//
// AdwToastOverlay's CSS node is called toastoverlay. It contains the child,
// as well as zero or more toast subnodes.
//
// Each of the toast nodes contains a widget subnode, optionally a button
// subnode, and another button subnode with .circular and .flat style classes.
//
// The widget subnode contains a label subnode with the .heading style class,
// or a custom widget provided by the application.
//
// # Accessibility
//
// AdwToastOverlay uses the GTK_ACCESSIBLE_ROLE_TAB_GROUP role.
type ToastOverlay struct {
	_ [0]func() // equal guard
	gtk.Widget
}

var (
	_ gtk.Widgetter = (*ToastOverlay)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ToastOverlay, *ToastOverlayClass, ToastOverlayOverrides](
		GTypeToastOverlay,
		initToastOverlayClass,
		wrapToastOverlay,
		defaultToastOverlayOverrides,
	)
}

func initToastOverlayClass(gclass unsafe.Pointer, overrides ToastOverlayOverrides, classInitFunc func(*ToastOverlayClass)) {
	if classInitFunc != nil {
		class := (*ToastOverlayClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapToastOverlay(obj *coreglib.Object) *ToastOverlay {
	return &ToastOverlay{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalToastOverlay(p uintptr) (interface{}, error) {
	return wrapToastOverlay(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewToastOverlay creates a new AdwToastOverlay.
//
// The function returns the following values:
//
//   - toastOverlay: new created AdwToastOverlay.
//
func NewToastOverlay() *ToastOverlay {
	var _cret *C.GtkWidget // in

	_cret = C.adw_toast_overlay_new()

	var _toastOverlay *ToastOverlay // out

	_toastOverlay = wrapToastOverlay(coreglib.Take(unsafe.Pointer(_cret)))

	return _toastOverlay
}

// AddToast displays toast.
//
// Only one toast can be shown at a time; if a toast is already being displayed,
// either toast or the original toast will be placed in a queue, depending on
// the priority of toast. See toast:priority.
//
// If called on a toast that's already displayed, its timeout will be reset.
//
// If called on a toast currently in the queue, the toast will be bumped forward
// to be shown as soon as possible.
//
// The function takes the following parameters:
//
//   - toast: toast.
//
func (self *ToastOverlay) AddToast(toast *Toast) {
	var _arg0 *C.AdwToastOverlay // out
	var _arg1 *C.AdwToast        // out

	_arg0 = (*C.AdwToastOverlay)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.AdwToast)(unsafe.Pointer(coreglib.InternObject(toast).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(toast).Native()))

	C.adw_toast_overlay_add_toast(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(toast)
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//   - widget (optional): child widget of self.
//
func (self *ToastOverlay) Child() gtk.Widgetter {
	var _arg0 *C.AdwToastOverlay // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.AdwToastOverlay)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_toast_overlay_get_child(_arg0)
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

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//   - child (optional) widget.
//
func (self *ToastOverlay) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwToastOverlay // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.AdwToastOverlay)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.adw_toast_overlay_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}
