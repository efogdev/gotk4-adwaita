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
// extern void _gotk4_adw1_SwipeTracker_ConnectUpdateSwipe(gpointer, gdouble, guintptr);
// extern void _gotk4_adw1_SwipeTracker_ConnectPrepare(gpointer, AdwNavigationDirection, guintptr);
// extern void _gotk4_adw1_SwipeTracker_ConnectEndSwipe(gpointer, gdouble, gdouble, guintptr);
// extern void _gotk4_adw1_SwipeTracker_ConnectBeginSwipe(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeSwipeTracker = coreglib.Type(C.adw_swipe_tracker_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSwipeTracker, F: marshalSwipeTracker},
	})
}

// SwipeTrackerOverrides contains methods that are overridable.
type SwipeTrackerOverrides struct {
}

func defaultSwipeTrackerOverrides(v *SwipeTracker) SwipeTrackerOverrides {
	return SwipeTrackerOverrides{}
}

// SwipeTracker: swipe tracker used in carousel, flap and leaflet.
//
// The AdwSwipeTracker object can be used for implementing widgets with swipe
// gestures. It supports touch-based swipes, pointer dragging, and touchpad
// scrolling.
//
// The widgets will probably want to expose the swipetracker:enabled property.
// If they expect to use horizontal orientation, swipetracker:reversed can be
// used for supporting RTL text direction.
type SwipeTracker struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gtk.Orientable
}

var (
	_ coreglib.Objector = (*SwipeTracker)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SwipeTracker, *SwipeTrackerClass, SwipeTrackerOverrides](
		GTypeSwipeTracker,
		initSwipeTrackerClass,
		wrapSwipeTracker,
		defaultSwipeTrackerOverrides,
	)
}

func initSwipeTrackerClass(gclass unsafe.Pointer, overrides SwipeTrackerOverrides, classInitFunc func(*SwipeTrackerClass)) {
	if classInitFunc != nil {
		class := (*SwipeTrackerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSwipeTracker(obj *coreglib.Object) *SwipeTracker {
	return &SwipeTracker{
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalSwipeTracker(p uintptr) (interface{}, error) {
	return wrapSwipeTracker(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectBeginSwipe: this signal is emitted right before a swipe will be
// started, after the drag threshold has been passed.
func (self *SwipeTracker) ConnectBeginSwipe(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "begin-swipe", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectBeginSwipe), f)
}

// ConnectEndSwipe: this signal is emitted as soon as the gesture has stopped.
//
// The user is expected to animate the deceleration from the current progress
// value to to with an animation using velocity as the initial velocity,
// provided in pixels per second. springanimation is usually a good fit for
// this.
func (self *SwipeTracker) ConnectEndSwipe(f func(velocity, to float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "end-swipe", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectEndSwipe), f)
}

// ConnectPrepare: this signal is emitted when a possible swipe is detected.
//
// The direction value can be used to restrict the swipe to a certain direction.
func (self *SwipeTracker) ConnectPrepare(f func(direction NavigationDirection)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "prepare", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectPrepare), f)
}

// ConnectUpdateSwipe: this signal is emitted every time the progress value
// changes.
func (self *SwipeTracker) ConnectUpdateSwipe(f func(progress float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "update-swipe", false, unsafe.Pointer(C._gotk4_adw1_SwipeTracker_ConnectUpdateSwipe), f)
}

// NewSwipeTracker creates a new AdwSwipeTracker for widget.
//
// The function takes the following parameters:
//
//   - swipeable: widget to add the tracker on.
//
// The function returns the following values:
//
//   - swipeTracker: newly created AdwSwipeTracker.
//
func NewSwipeTracker(swipeable Swipeabler) *SwipeTracker {
	var _arg1 *C.AdwSwipeable    // out
	var _cret *C.AdwSwipeTracker // in

	_arg1 = (*C.AdwSwipeable)(unsafe.Pointer(coreglib.InternObject(swipeable).Native()))

	_cret = C.adw_swipe_tracker_new(_arg1)
	runtime.KeepAlive(swipeable)

	var _swipeTracker *SwipeTracker // out

	_swipeTracker = wrapSwipeTracker(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _swipeTracker
}

// AllowLongSwipes gets whether to allow swiping for more than one snap point at
// a time.
//
// The function returns the following values:
//
//   - ok: whether long swipes are allowed.
//
func (self *SwipeTracker) AllowLongSwipes() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_allow_long_swipes(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AllowMouseDrag gets whether self can be dragged with mouse pointer.
//
// The function returns the following values:
//
//   - ok: whether mouse dragging is allowed.
//
func (self *SwipeTracker) AllowMouseDrag() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_allow_mouse_drag(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Enabled gets whether self is enabled.
//
// The function returns the following values:
//
//   - ok: whether self is enabled.
//
func (self *SwipeTracker) Enabled() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_enabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Reversed gets whether self is reversing the swipe direction.
//
// The function returns the following values:
//
//   - ok: whether the direction is reversed.
//
func (self *SwipeTracker) Reversed() bool {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret C.gboolean         // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_reversed(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Swipeable: get the widget self is attached to.
//
// The function returns the following values:
//
//   - swipeable widget.
//
func (self *SwipeTracker) Swipeable() *Swipeable {
	var _arg0 *C.AdwSwipeTracker // out
	var _cret *C.AdwSwipeable    // in

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_swipe_tracker_get_swipeable(_arg0)
	runtime.KeepAlive(self)

	var _swipeable *Swipeable // out

	_swipeable = wrapSwipeable(coreglib.Take(unsafe.Pointer(_cret)))

	return _swipeable
}

// SetAllowLongSwipes sets whether to allow swiping for more than one snap point
// at a time.
//
// If the value is FALSE, each swipe can only move to the adjacent snap points.
//
// The function takes the following parameters:
//
//   - allowLongSwipes: whether to allow long swipes.
//
func (self *SwipeTracker) SetAllowLongSwipes(allowLongSwipes bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if allowLongSwipes {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_allow_long_swipes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowLongSwipes)
}

// SetAllowMouseDrag sets whether self can be dragged with mouse pointer.
//
// The function takes the following parameters:
//
//   - allowMouseDrag: whether to allow mouse dragging.
//
func (self *SwipeTracker) SetAllowMouseDrag(allowMouseDrag bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if allowMouseDrag {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_allow_mouse_drag(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowMouseDrag)
}

// SetEnabled sets whether self is enabled.
//
// When it's not enabled, no events will be processed. Usually widgets will want
// to expose this via a property.
//
// The function takes the following parameters:
//
//   - enabled: whether self is enabled.
//
func (self *SwipeTracker) SetEnabled(enabled bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enabled)
}

// SetReversed sets whether to reverse the swipe direction.
//
// If the swipe tracker is horizontal, it can be used for supporting RTL text
// direction.
//
// The function takes the following parameters:
//
//   - reversed: whether to reverse the swipe direction.
//
func (self *SwipeTracker) SetReversed(reversed bool) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if reversed {
		_arg1 = C.TRUE
	}

	C.adw_swipe_tracker_set_reversed(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(reversed)
}

// ShiftPosition moves the current progress value by delta.
//
// This can be used to adjust the current position if snap points move during
// the gesture.
//
// The function takes the following parameters:
//
//   - delta: position delta.
//
func (self *SwipeTracker) ShiftPosition(delta float64) {
	var _arg0 *C.AdwSwipeTracker // out
	var _arg1 C.double           // out

	_arg0 = (*C.AdwSwipeTracker)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(delta)

	C.adw_swipe_tracker_shift_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(delta)
}
