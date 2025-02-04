// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

// GetEnableAnimations checks whether animations are enabled for widget.
//
// This should be used when implementing an animated widget to know whether to
// animate it or not.
//
// The function takes the following parameters:
//
//   - widget: GtkWidget.
//
// The function returns the following values:
//
//   - ok: whether animations are enabled for widget.
func GetEnableAnimations(widget gtk.Widgetter) bool {
	var _arg1 *C.GtkWidget // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	_cret = C.adw_get_enable_animations(_arg1)
	runtime.KeepAlive(widget)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Lerp computes the linear interpolation between a and b for t.
//
// The function takes the following parameters:
//
//   - a: start.
//   - b: end.
//   - t: interpolation rate.
//
// The function returns the following values:
//
//   - gdouble: computed value.
func Lerp(a, b, t float64) float64 {
	var _arg1 C.double // out
	var _arg2 C.double // out
	var _arg3 C.double // out
	var _cret C.double // in

	_arg1 = C.double(a)
	_arg2 = C.double(b)
	_arg3 = C.double(t)

	_cret = C.adw_lerp(_arg1, _arg2, _arg3)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	runtime.KeepAlive(t)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
