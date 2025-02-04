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
	GTypeSpringAnimation = coreglib.Type(C.adw_spring_animation_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpringAnimation, F: marshalSpringAnimation},
	})
}

// SpringAnimation: spring-based animation.
//
// AdwSpringAnimation implements an animation driven by a physical model
// of a spring described by springparams, with a resting position in
// springanimation:value-to, stretched to springanimation:value-from.
//
// Since the animation is physically simulated, spring animations don't
// have a fixed duration. The animation will stop when the simulated
// spring comes to a rest - when the amplitude of the oscillations becomes
// smaller than springanimation:epsilon, or immediately when it reaches
// springanimation:value-to if springanimation:clamp is set to TRUE. The
// estimated duration can be obtained with springanimation:estimated-duration.
//
// Due to the nature of spring-driven motion the animation can overshoot
// springanimation:value-to before coming to a rest. Whether the animation will
// overshoot or not depends on the damping ratio of the spring. See springparams
// for more information about specific damping ratio values.
//
// If springanimation:clamp is TRUE, the animation will abruptly end as soon as
// it reaches the final value, preventing overshooting.
//
// Animations can have an initial velocity value, set via
// springanimation:initial-velocity, which adjusts the curve without changing
// the duration. This makes spring animations useful for deceleration at the end
// of gestures.
//
// If the initial and final values are equal, and the initial velocity is not 0,
// the animation value will bounce and return to its resting position.
type SpringAnimation struct {
	_ [0]func() // equal guard
	Animation
}

var (
	_ Animationer = (*SpringAnimation)(nil)
)

func wrapSpringAnimation(obj *coreglib.Object) *SpringAnimation {
	return &SpringAnimation{
		Animation: Animation{
			Object: obj,
		},
	}
}

func marshalSpringAnimation(p uintptr) (interface{}, error) {
	return wrapSpringAnimation(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSpringAnimation creates a new AdwSpringAnimation on widget.
//
// The animation will animate target from from to to with the dynamics of a
// spring described by spring_params.
//
// The function takes the following parameters:
//
//   - widget to create animation on.
//   - from: value to animate from.
//   - to: value to animate to.
//   - springParams: physical parameters of the spring.
//   - target value to animate.
//
// The function returns the following values:
//
//   - springAnimation: newly created animation.
func NewSpringAnimation(widget gtk.Widgetter, from, to float64, springParams *SpringParams, target AnimationTargetter) *SpringAnimation {
	var _arg1 *C.GtkWidget          // out
	var _arg2 C.double              // out
	var _arg3 C.double              // out
	var _arg4 *C.AdwSpringParams    // out
	var _arg5 *C.AdwAnimationTarget // out
	var _cret *C.AdwAnimation       // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	_arg2 = C.double(from)
	_arg3 = C.double(to)
	_arg4 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(springParams)))
	_arg5 = (*C.AdwAnimationTarget)(unsafe.Pointer(coreglib.InternObject(target).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(target).Native()))

	_cret = C.adw_spring_animation_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	runtime.KeepAlive(springParams)
	runtime.KeepAlive(target)

	var _springAnimation *SpringAnimation // out

	_springAnimation = wrapSpringAnimation(coreglib.Take(unsafe.Pointer(_cret)))

	return _springAnimation
}

// CalculateValue calculates the value self will have at time.
//
// The time starts at 0 and ends at springanimation:estimated_duration.
//
// See also springanimation.CalculateVelocity.
//
// The function takes the following parameters:
//
//   - time: elapsed time, in milliseconds.
//
// The function returns the following values:
//
//   - gdouble: value at time.
func (self *SpringAnimation) CalculateValue(time uint) float64 {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 C.guint               // out
	var _cret C.double              // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(time)

	_cret = C.adw_spring_animation_calculate_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(time)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// CalculateVelocity calculates the velocity self will have at time.
//
// The time starts at 0 and ends at springanimation:estimated_duration.
//
// See also springanimation.CalculateValue.
//
// The function takes the following parameters:
//
//   - time: elapsed time, in milliseconds.
//
// The function returns the following values:
//
//   - gdouble: velocity at time.
func (self *SpringAnimation) CalculateVelocity(time uint) float64 {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 C.guint               // out
	var _cret C.double              // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(time)

	_cret = C.adw_spring_animation_calculate_velocity(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(time)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Clamp gets whether self should be clamped.
//
// The function returns the following values:
//
//   - ok: whether self is clamped.
func (self *SpringAnimation) Clamp() bool {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret C.gboolean            // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_clamp(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Epsilon gets the precision of the spring.
//
// The function returns the following values:
//
//   - gdouble: epsilon value.
func (self *SpringAnimation) Epsilon() float64 {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret C.double              // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_epsilon(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// EstimatedDuration gets the estimated duration of self, in milliseconds.
//
// Can be duration_infinite if the spring damping is set to 0.
//
// The function returns the following values:
//
//   - guint: estimated duration.
func (self *SpringAnimation) EstimatedDuration() uint {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret C.guint               // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_estimated_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// InitialVelocity gets the initial velocity of self.
//
// The function returns the following values:
//
//   - gdouble: initial velocity.
func (self *SpringAnimation) InitialVelocity() float64 {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret C.double              // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_initial_velocity(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SpringParams gets the physical parameters of the spring of self.
//
// The function returns the following values:
//
//   - springParams: spring parameters.
func (self *SpringAnimation) SpringParams() *SpringParams {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret *C.AdwSpringParams    // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_spring_params(_arg0)
	runtime.KeepAlive(self)

	var _springParams *SpringParams // out

	_springParams = (*SpringParams)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.adw_spring_params_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_springParams)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_spring_params_unref((*C.AdwSpringParams)(intern.C))
		},
	)

	return _springParams
}

// ValueFrom gets the value self will animate from.
//
// The function returns the following values:
//
//   - gdouble: value to animate from.
func (self *SpringAnimation) ValueFrom() float64 {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret C.double              // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_value_from(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ValueTo gets the value self will animate to.
//
// The function returns the following values:
//
//   - gdouble: value to animate to.
func (self *SpringAnimation) ValueTo() float64 {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret C.double              // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_value_to(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Velocity gets the current velocity of self.
//
// The function returns the following values:
//
//   - gdouble: current velocity.
func (self *SpringAnimation) Velocity() float64 {
	var _arg0 *C.AdwSpringAnimation // out
	var _cret C.double              // in

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_spring_animation_get_velocity(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetClamp sets whether self should be clamped.
//
// If set to TRUE, the animation will abruptly end as soon as it reaches the
// final value, preventing overshooting.
//
// It won't prevent overshooting springanimation:value-from if a relative
// negative springanimation:initial-velocity is set.
//
// The function takes the following parameters:
//
//   - clamp: new value.
func (self *SpringAnimation) SetClamp(clamp bool) {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if clamp {
		_arg1 = C.TRUE
	}

	C.adw_spring_animation_set_clamp(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(clamp)
}

// SetEpsilon sets the precision of the spring.
//
// The level of precision used to determine when the animation has come to a
// rest, that is, when the amplitude of the oscillations becomes smaller than
// this value.
//
// If the epsilon value is too small, the animation will take a long time to
// stop after the animated value has stopped visibly changing.
//
// If the epsilon value is too large, the animation will end prematurely.
//
// The default value is 0.001.
//
// The function takes the following parameters:
//
//   - epsilon: new value.
func (self *SpringAnimation) SetEpsilon(epsilon float64) {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 C.double              // out

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(epsilon)

	C.adw_spring_animation_set_epsilon(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(epsilon)
}

// SetInitialVelocity sets the initial velocity of self.
//
// Initial velocity affects only the animation curve, but not its duration.
//
// The function takes the following parameters:
//
//   - velocity: initial velocity.
func (self *SpringAnimation) SetInitialVelocity(velocity float64) {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 C.double              // out

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(velocity)

	C.adw_spring_animation_set_initial_velocity(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(velocity)
}

// SetSpringParams sets the physical parameters of the spring of self.
//
// The function takes the following parameters:
//
//   - springParams: new spring parameters.
func (self *SpringAnimation) SetSpringParams(springParams *SpringParams) {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 *C.AdwSpringParams    // out

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(springParams)))

	C.adw_spring_animation_set_spring_params(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(springParams)
}

// SetValueFrom sets the value self will animate from.
//
// The animation will start at this value and end at springanimation:value-to.
//
// The function takes the following parameters:
//
//   - value to animate from.
func (self *SpringAnimation) SetValueFrom(value float64) {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 C.double              // out

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(value)

	C.adw_spring_animation_set_value_from(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetValueTo sets the value self will animate to.
//
// The animation will start at springanimation:value-from and end at this value.
//
// The function takes the following parameters:
//
//   - value to animate to.
func (self *SpringAnimation) SetValueTo(value float64) {
	var _arg0 *C.AdwSpringAnimation // out
	var _arg1 C.double              // out

	_arg0 = (*C.AdwSpringAnimation)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.double(value)

	C.adw_spring_animation_set_value_to(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}
