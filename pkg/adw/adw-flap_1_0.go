// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
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
	GTypeFlapFoldPolicy     = coreglib.Type(C.adw_flap_fold_policy_get_type())
	GTypeFlapTransitionType = coreglib.Type(C.adw_flap_transition_type_get_type())
	GTypeFlap               = coreglib.Type(C.adw_flap_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFlapFoldPolicy, F: marshalFlapFoldPolicy},
		coreglib.TypeMarshaler{T: GTypeFlapTransitionType, F: marshalFlapTransitionType},
		coreglib.TypeMarshaler{T: GTypeFlap, F: marshalFlap},
	})
}

// FlapFoldPolicy describes the possible folding behavior of a flap widget.
type FlapFoldPolicy C.gint

const (
	// FlapFoldPolicyNever: disable folding, the flap cannot reach narrow sizes.
	FlapFoldPolicyNever FlapFoldPolicy = iota
	// FlapFoldPolicyAlways: keep the flap always folded.
	FlapFoldPolicyAlways
	// FlapFoldPolicyAuto: fold and unfold the flap based on available space.
	FlapFoldPolicyAuto
)

func marshalFlapFoldPolicy(p uintptr) (interface{}, error) {
	return FlapFoldPolicy(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FlapFoldPolicy.
func (f FlapFoldPolicy) String() string {
	switch f {
	case FlapFoldPolicyNever:
		return "Never"
	case FlapFoldPolicyAlways:
		return "Always"
	case FlapFoldPolicyAuto:
		return "Auto"
	default:
		return fmt.Sprintf("FlapFoldPolicy(%d)", f)
	}
}

// FlapTransitionType describes transitions types of a flap widget.
//
// It determines the type of animation when transitioning between children in a
// flap widget, as well as which areas can be swiped via flap:swipe-to-open and
// flap:swipe-to-close.
//
// New values may be added to this enum over time.
type FlapTransitionType C.gint

const (
	// FlapTransitionTypeOver: flap slides over the content, which is dimmed.
	// When folded, only the flap can be swiped.
	FlapTransitionTypeOver FlapTransitionType = iota
	// FlapTransitionTypeUnder: content slides over the flap. Only the content
	// can be swiped.
	FlapTransitionTypeUnder
	// FlapTransitionTypeSlide: flap slides offscreen when hidden, neither the
	// flap nor content overlap each other. Both widgets can be swiped.
	FlapTransitionTypeSlide
)

func marshalFlapTransitionType(p uintptr) (interface{}, error) {
	return FlapTransitionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for FlapTransitionType.
func (f FlapTransitionType) String() string {
	switch f {
	case FlapTransitionTypeOver:
		return "Over"
	case FlapTransitionTypeUnder:
		return "Under"
	case FlapTransitionTypeSlide:
		return "Slide"
	default:
		return fmt.Sprintf("FlapTransitionType(%d)", f)
	}
}

// FlapOverrides contains methods that are overridable.
type FlapOverrides struct {
}

func defaultFlapOverrides(v *Flap) FlapOverrides {
	return FlapOverrides{}
}

// Flap: adaptive container acting like a box or an overlay.
//
// <picture> <source srcset="flap-wide-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="flap-wide.png" alt="flap-wide"> </picture> <picture>
// <source srcset="flap-narrow-dark.png" media="(prefers-color-scheme: dark)">
// <img src="flap-narrow.png" alt="flap-narrow"> </picture>
//
// The AdwFlap widget can display its children like a gtk.Box does or like a
// gtk.Overlay does, according to the flap:fold-policy value.
//
// AdwFlap has at most three children: flap:content, flap:flap and
// flap:separator. Content is the primary child, flap is displayed next to it
// when unfolded, or overlays it when folded. Flap can be shown or hidden by
// changing the flap:reveal-flap value, as well as via swipe gestures if
// flap:swipe-to-open and/or flap:swipe-to-close are set to TRUE.
//
// Optionally, a separator can be provided, which would be displayed between the
// content and the flap when there's no shadow to separate them, depending on
// the transition type.
//
// flap:flap is transparent by default; add the .background
// (style-classes.html#background) style class to it if this is unwanted.
//
// If flap:modal is set to TRUE, content becomes completely inaccessible when
// the flap is revealed while folded.
//
// The position of the flap and separator children relative to the content is
// determined by orientation, as well as the flap:flap-position value.
//
// Folding the flap will automatically hide the flap widget, and unfolding it
// will automatically reveal it. If this behavior is not desired, the
// flap:locked property can be used to override it.
//
// Common use cases include sidebars, header bars that need to be able to
// overlap the window content (for example, in fullscreen mode) and bottom
// sheets.
//
//
// AdwFlap as GtkBuildable
//
// The AdwFlap implementation of the gtk.Buildable interface supports setting
// the flap child by specifying “flap” as the “type” attribute of a <child>
// element, and separator by specifying “separator”. Specifying “content” child
// type or omitting it results in setting the content child.
//
//
// CSS nodes
//
// AdwFlap has a single CSS node with name flap. The node will get the style
// classes .folded when it is folded, and .unfolded when it's not.
type Flap struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	Swipeable
	gtk.Orientable
}

var (
	_ gtk.Widgetter     = (*Flap)(nil)
	_ coreglib.Objector = (*Flap)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Flap, *FlapClass, FlapOverrides](
		GTypeFlap,
		initFlapClass,
		wrapFlap,
		defaultFlapOverrides,
	)
}

func initFlapClass(gclass unsafe.Pointer, overrides FlapOverrides, classInitFunc func(*FlapClass)) {
	if classInitFunc != nil {
		class := (*FlapClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapFlap(obj *coreglib.Object) *Flap {
	return &Flap{
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
		Object: obj,
		Swipeable: Swipeable{
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
		},
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalFlap(p uintptr) (interface{}, error) {
	return wrapFlap(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFlap creates a new AdwFlap.
//
// The function returns the following values:
//
//    - flap: newly created AdwFlap.
//
func NewFlap() *Flap {
	var _cret *C.GtkWidget // in

	_cret = C.adw_flap_new()

	var _flap *Flap // out

	_flap = wrapFlap(coreglib.Take(unsafe.Pointer(_cret)))

	return _flap
}

// Content gets the content widget for self.
//
// The function returns the following values:
//
//    - widget (optional): content widget for self.
//
func (self *Flap) Content() gtk.Widgetter {
	var _arg0 *C.AdwFlap   // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_content(_arg0)
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

// Flap gets the flap widget for self.
//
// The function returns the following values:
//
//    - widget (optional): flap widget for self.
//
func (self *Flap) Flap() gtk.Widgetter {
	var _arg0 *C.AdwFlap   // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_flap(_arg0)
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

// FlapPosition gets the flap position for self.
//
// The function returns the following values:
//
//    - packType: flap position for self.
//
func (self *Flap) FlapPosition() gtk.PackType {
	var _arg0 *C.AdwFlap    // out
	var _cret C.GtkPackType // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_flap_position(_arg0)
	runtime.KeepAlive(self)

	var _packType gtk.PackType // out

	_packType = gtk.PackType(_cret)

	return _packType
}

// FoldDuration gets the duration that fold transitions in self will take.
//
// The function returns the following values:
//
//    - guint: fold transition duration.
//
func (self *Flap) FoldDuration() uint {
	var _arg0 *C.AdwFlap // out
	var _cret C.guint    // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_fold_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// FoldPolicy gets the fold policy for self.
//
// The function returns the following values:
//
//    - flapFoldPolicy: fold policy for self.
//
func (self *Flap) FoldPolicy() FlapFoldPolicy {
	var _arg0 *C.AdwFlap          // out
	var _cret C.AdwFlapFoldPolicy // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_fold_policy(_arg0)
	runtime.KeepAlive(self)

	var _flapFoldPolicy FlapFoldPolicy // out

	_flapFoldPolicy = FlapFoldPolicy(_cret)

	return _flapFoldPolicy
}

// FoldThresholdPolicy gets the fold threshold policy for self.
//
// The function returns the following values:
//
func (self *Flap) FoldThresholdPolicy() FoldThresholdPolicy {
	var _arg0 *C.AdwFlap               // out
	var _cret C.AdwFoldThresholdPolicy // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_fold_threshold_policy(_arg0)
	runtime.KeepAlive(self)

	var _foldThresholdPolicy FoldThresholdPolicy // out

	_foldThresholdPolicy = FoldThresholdPolicy(_cret)

	return _foldThresholdPolicy
}

// Folded gets whether self is currently folded.
//
// The function returns the following values:
//
//    - ok: TRUE if self is currently folded.
//
func (self *Flap) Folded() bool {
	var _arg0 *C.AdwFlap // out
	var _cret C.gboolean // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_folded(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Locked gets whether self is locked.
//
// The function returns the following values:
//
//    - ok: TRUE if self is locked.
//
func (self *Flap) Locked() bool {
	var _arg0 *C.AdwFlap // out
	var _cret C.gboolean // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_locked(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Modal gets whether self is modal.
//
// The function returns the following values:
//
//    - ok: TRUE if self is modal.
//
func (self *Flap) Modal() bool {
	var _arg0 *C.AdwFlap // out
	var _cret C.gboolean // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_modal(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RevealFlap gets whether the flap widget is revealed for self.
//
// The function returns the following values:
//
//    - ok: TRUE if the flap widget is revealed.
//
func (self *Flap) RevealFlap() bool {
	var _arg0 *C.AdwFlap // out
	var _cret C.gboolean // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_reveal_flap(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RevealParams gets the reveal animation spring parameters for self.
//
// The function returns the following values:
//
//    - springParams: reveal animation parameters.
//
func (self *Flap) RevealParams() *SpringParams {
	var _arg0 *C.AdwFlap         // out
	var _cret *C.AdwSpringParams // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_reveal_params(_arg0)
	runtime.KeepAlive(self)

	var _springParams *SpringParams // out

	_springParams = (*SpringParams)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_springParams)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_spring_params_unref((*C.AdwSpringParams)(intern.C))
		},
	)

	return _springParams
}

// RevealProgress gets the current reveal progress for self.
//
// The function returns the following values:
//
//    - gdouble: current reveal progress for self.
//
func (self *Flap) RevealProgress() float64 {
	var _arg0 *C.AdwFlap // out
	var _cret C.double   // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_reveal_progress(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Separator gets the separator widget for self.
//
// The function returns the following values:
//
//    - widget (optional): separator widget for self.
//
func (self *Flap) Separator() gtk.Widgetter {
	var _arg0 *C.AdwFlap   // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_separator(_arg0)
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

// SwipeToClose gets whether self can be closed with a swipe gesture.
//
// The function returns the following values:
//
//    - ok: TRUE if self can be closed with a swipe gesture.
//
func (self *Flap) SwipeToClose() bool {
	var _arg0 *C.AdwFlap // out
	var _cret C.gboolean // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_swipe_to_close(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SwipeToOpen gets whether self can be opened with a swipe gesture.
//
// The function returns the following values:
//
//    - ok: TRUE if self can be opened with a swipe gesture.
//
func (self *Flap) SwipeToOpen() bool {
	var _arg0 *C.AdwFlap // out
	var _cret C.gboolean // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_swipe_to_open(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionType gets the type of animation used for reveal and fold
// transitions in self.
//
// The function returns the following values:
//
//    - flapTransitionType: current transition type of self.
//
func (self *Flap) TransitionType() FlapTransitionType {
	var _arg0 *C.AdwFlap              // out
	var _cret C.AdwFlapTransitionType // in

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_flap_get_transition_type(_arg0)
	runtime.KeepAlive(self)

	var _flapTransitionType FlapTransitionType // out

	_flapTransitionType = FlapTransitionType(_cret)

	return _flapTransitionType
}

// SetContent sets the content widget for self.
//
// The function takes the following parameters:
//
//    - content (optional) widget.
//
func (self *Flap) SetContent(content gtk.Widgetter) {
	var _arg0 *C.AdwFlap   // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if content != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(content).Native()))
	}

	C.adw_flap_set_content(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(content)
}

// SetFlap sets the flap widget for self.
//
// The function takes the following parameters:
//
//    - flap (optional) widget.
//
func (self *Flap) SetFlap(flap gtk.Widgetter) {
	var _arg0 *C.AdwFlap   // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if flap != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(flap).Native()))
	}

	C.adw_flap_set_flap(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(flap)
}

// SetFlapPosition sets the flap position for self.
//
// The function takes the following parameters:
//
//    - position: new value.
//
func (self *Flap) SetFlapPosition(position gtk.PackType) {
	var _arg0 *C.AdwFlap    // out
	var _arg1 C.GtkPackType // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkPackType(position)

	C.adw_flap_set_flap_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}

// SetFoldDuration sets the duration that fold transitions in self will take.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (self *Flap) SetFoldDuration(duration uint) {
	var _arg0 *C.AdwFlap // out
	var _arg1 C.guint    // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(duration)

	C.adw_flap_set_fold_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetFoldPolicy sets the fold policy for self.
//
// The function takes the following parameters:
//
//    - policy: fold policy.
//
func (self *Flap) SetFoldPolicy(policy FlapFoldPolicy) {
	var _arg0 *C.AdwFlap          // out
	var _arg1 C.AdwFlapFoldPolicy // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwFlapFoldPolicy(policy)

	C.adw_flap_set_fold_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetFoldThresholdPolicy sets the fold threshold policy for self.
//
// The function takes the following parameters:
//
//    - policy to use.
//
func (self *Flap) SetFoldThresholdPolicy(policy FoldThresholdPolicy) {
	var _arg0 *C.AdwFlap               // out
	var _arg1 C.AdwFoldThresholdPolicy // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwFoldThresholdPolicy(policy)

	C.adw_flap_set_fold_threshold_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetLocked sets whether self is locked.
//
// The function takes the following parameters:
//
//    - locked: new value.
//
func (self *Flap) SetLocked(locked bool) {
	var _arg0 *C.AdwFlap // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if locked {
		_arg1 = C.TRUE
	}

	C.adw_flap_set_locked(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(locked)
}

// SetModal sets whether self is modal.
//
// The function takes the following parameters:
//
//    - modal: whether self is modal.
//
func (self *Flap) SetModal(modal bool) {
	var _arg0 *C.AdwFlap // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.adw_flap_set_modal(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(modal)
}

// SetRevealFlap sets whether the flap widget is revealed for self.
//
// The function takes the following parameters:
//
//    - revealFlap: whether to reveal the flap widget.
//
func (self *Flap) SetRevealFlap(revealFlap bool) {
	var _arg0 *C.AdwFlap // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if revealFlap {
		_arg1 = C.TRUE
	}

	C.adw_flap_set_reveal_flap(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(revealFlap)
}

// SetRevealParams sets the reveal animation spring parameters for self.
//
// The function takes the following parameters:
//
//    - params: new parameters.
//
func (self *Flap) SetRevealParams(params *SpringParams) {
	var _arg0 *C.AdwFlap         // out
	var _arg1 *C.AdwSpringParams // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.AdwSpringParams)(gextras.StructNative(unsafe.Pointer(params)))

	C.adw_flap_set_reveal_params(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(params)
}

// SetSeparator sets the separator widget for self.
//
// The function takes the following parameters:
//
//    - separator (optional) widget.
//
func (self *Flap) SetSeparator(separator gtk.Widgetter) {
	var _arg0 *C.AdwFlap   // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if separator != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(separator).Native()))
	}

	C.adw_flap_set_separator(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(separator)
}

// SetSwipeToClose sets whether self can be closed with a swipe gesture.
//
// The function takes the following parameters:
//
//    - swipeToClose: whether self can be closed with a swipe gesture.
//
func (self *Flap) SetSwipeToClose(swipeToClose bool) {
	var _arg0 *C.AdwFlap // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if swipeToClose {
		_arg1 = C.TRUE
	}

	C.adw_flap_set_swipe_to_close(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(swipeToClose)
}

// SetSwipeToOpen sets whether self can be opened with a swipe gesture.
//
// The function takes the following parameters:
//
//    - swipeToOpen: whether self can be opened with a swipe gesture.
//
func (self *Flap) SetSwipeToOpen(swipeToOpen bool) {
	var _arg0 *C.AdwFlap // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if swipeToOpen {
		_arg1 = C.TRUE
	}

	C.adw_flap_set_swipe_to_open(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(swipeToOpen)
}

// SetTransitionType sets the type of animation used for reveal and fold
// transitions in self.
//
// The function takes the following parameters:
//
//    - transitionType: new transition type.
//
func (self *Flap) SetTransitionType(transitionType FlapTransitionType) {
	var _arg0 *C.AdwFlap              // out
	var _arg1 C.AdwFlapTransitionType // out

	_arg0 = (*C.AdwFlap)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwFlapTransitionType(transitionType)

	C.adw_flap_set_transition_type(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(transitionType)
}