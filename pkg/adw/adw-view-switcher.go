// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_view_switcher_get_type()), F: marshalViewSwitcherer},
	})
}

type ViewSwitcher struct {
	gtk.Widget
}

var _ gextras.Nativer = (*ViewSwitcher)(nil)

func wrapViewSwitcher(obj *externglib.Object) *ViewSwitcher {
	return &ViewSwitcher{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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

func marshalViewSwitcherer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapViewSwitcher(obj), nil
}

// NewViewSwitcher creates a new ViewSwitcher widget.
func NewViewSwitcher() *ViewSwitcher {
	var _cret *C.GtkWidget // in

	_cret = C.adw_view_switcher_new()

	var _viewSwitcher *ViewSwitcher // out

	_viewSwitcher = wrapViewSwitcher(externglib.Take(unsafe.Pointer(_cret)))

	return _viewSwitcher
}

// NarrowEllipsize: get the ellipsizing position of the narrow mode label. See
// adw_view_switcher_set_narrow_ellipsize().
func (self *ViewSwitcher) NarrowEllipsize() pango.EllipsizeMode {
	var _arg0 *C.AdwViewSwitcher   // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_get_narrow_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// Policy gets the policy of self.
func (self *ViewSwitcher) Policy() ViewSwitcherPolicy {
	var _arg0 *C.AdwViewSwitcher      // out
	var _cret C.AdwViewSwitcherPolicy // in

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_get_policy(_arg0)

	var _viewSwitcherPolicy ViewSwitcherPolicy // out

	_viewSwitcherPolicy = ViewSwitcherPolicy(_cret)

	return _viewSwitcherPolicy
}

// Stack: get the Stack being controlled by the ViewSwitcher.
//
// See: adw_view_switcher_set_stack()
func (self *ViewSwitcher) Stack() *gtk.Stack {
	var _arg0 *C.AdwViewSwitcher // out
	var _cret *C.GtkStack        // in

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_switcher_get_stack(_arg0)

	var _stack *gtk.Stack // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_stack = &gtk.Stack{
			Widget: gtk.Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
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

	return _stack
}

// SetNarrowEllipsize: set the mode used to ellipsize the text in narrow mode if
// there is not enough space to render the entire string.
func (self *ViewSwitcher) SetNarrowEllipsize(mode pango.EllipsizeMode) {
	var _arg0 *C.AdwViewSwitcher   // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.adw_view_switcher_set_narrow_ellipsize(_arg0, _arg1)
}

// SetPolicy sets the policy of self.
func (self *ViewSwitcher) SetPolicy(policy ViewSwitcherPolicy) {
	var _arg0 *C.AdwViewSwitcher      // out
	var _arg1 C.AdwViewSwitcherPolicy // out

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwViewSwitcherPolicy(policy)

	C.adw_view_switcher_set_policy(_arg0, _arg1)
}

// SetStack sets the Stack to control.
func (self *ViewSwitcher) SetStack(stack *gtk.Stack) {
	var _arg0 *C.AdwViewSwitcher // out
	var _arg1 *C.GtkStack        // out

	_arg0 = (*C.AdwViewSwitcher)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.adw_view_switcher_set_stack(_arg0, _arg1)
}
