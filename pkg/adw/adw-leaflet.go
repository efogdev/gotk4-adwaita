// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_leaflet_transition_type_get_type()), F: marshalLeafletTransitionType},
		{T: externglib.Type(C.adw_leaflet_get_type()), F: marshalLeafletter},
		{T: externglib.Type(C.adw_leaflet_page_get_type()), F: marshalLeafletPager},
	})
}

// LeafletTransitionType: this enumeration value describes the possible
// transitions between modes and children in a Leaflet widget.
//
// New values may be added to this enumeration over time.
type LeafletTransitionType int

const (
	// LeafletTransitionTypeOver the old page or uncover the new page, sliding
	// from or towards the end according to orientation, text direction and
	// children order
	LeafletTransitionTypeOver LeafletTransitionType = iota
	// LeafletTransitionTypeUnder: uncover the new page or cover the old page,
	// sliding from or towards the start according to orientation, text
	// direction and children order
	LeafletTransitionTypeUnder
	// LeafletTransitionTypeSlide: slide from left, right, up or down according
	// to the orientation, text direction and the children order
	LeafletTransitionTypeSlide
)

func marshalLeafletTransitionType(p uintptr) (interface{}, error) {
	return LeafletTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for LeafletTransitionType.
func (l LeafletTransitionType) String() string {
	switch l {
	case LeafletTransitionTypeOver:
		return "Over"
	case LeafletTransitionTypeUnder:
		return "Under"
	case LeafletTransitionTypeSlide:
		return "Slide"
	default:
		return fmt.Sprintf("LeafletTransitionType(%d)", l)
	}
}

type Leaflet struct {
	gtk.Widget

	Swipeable
	gtk.Orientable
}

var _ gextras.Nativer = (*Leaflet)(nil)

func wrapLeaflet(obj *externglib.Object) *Leaflet {
	return &Leaflet{
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
		Swipeable: Swipeable{
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
		},
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalLeafletter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLeaflet(obj), nil
}

func NewLeaflet() *Leaflet {
	var _cret *C.GtkWidget // in

	_cret = C.adw_leaflet_new()

	var _leaflet *Leaflet // out

	_leaflet = wrapLeaflet(externglib.Take(unsafe.Pointer(_cret)))

	return _leaflet
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *Leaflet) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// Append adds a child to self.
func (self *Leaflet) Append(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.adw_leaflet_append(_arg0, _arg1)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// AdjacentChild gets the previous or next child that doesn't have 'navigatable'
// child property set to FALSE, or NULL if it doesn't exist. This will be the
// same widget adw_leaflet_navigate() will navigate to.
func (self *Leaflet) AdjacentChild(direction NavigationDirection) gtk.Widgetter {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwNavigationDirection // out
	var _cret *C.GtkWidget             // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwNavigationDirection(direction)

	_cret = C.adw_leaflet_get_adjacent_child(_arg0, _arg1)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// CanSwipeBack returns whether the Leaflet allows swiping to the previous
// child.
func (self *Leaflet) CanSwipeBack() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_can_swipe_back(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanSwipeForward returns whether the Leaflet allows swiping to the next child.
func (self *Leaflet) CanSwipeForward() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_can_swipe_forward(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (self *Leaflet) CanUnfold() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_can_unfold(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildByName finds the child of self with the name given as the argument.
// Returns NULL if there is no child with this name.
func (self *Leaflet) ChildByName(name string) gtk.Widgetter {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.char       // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))

	_cret = C.adw_leaflet_get_child_by_name(_arg0, _arg1)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// ChildTransitionDuration returns the amount of time (in milliseconds) that
// transitions between children in self will take.
func (self *Leaflet) ChildTransitionDuration() uint {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.guint       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_child_transition_duration(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ChildTransitionRunning returns whether self is currently in a transition from
// one page to another.
func (self *Leaflet) ChildTransitionRunning() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_child_transition_running(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Folded gets whether self is folded.
func (self *Leaflet) Folded() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_folded(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous gets whether self is homogeneous for the given fold and
// orientation. See adw_leaflet_set_homogeneous().
func (self *Leaflet) Homogeneous(folded bool, orientation gtk.Orientation) bool {
	var _arg0 *C.AdwLeaflet    // out
	var _arg1 C.gboolean       // out
	var _arg2 C.GtkOrientation // out
	var _cret C.gboolean       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if folded {
		_arg1 = C.TRUE
	}
	_arg2 = C.GtkOrientation(orientation)

	_cret = C.adw_leaflet_get_homogeneous(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize returns whether the Leaflet is set up to interpolate between
// the sizes of children on page switch.
func (self *Leaflet) InterpolateSize() bool {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_interpolate_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ModeTransitionDuration returns the amount of time (in milliseconds) that
// transitions between modes in self will take.
func (self *Leaflet) ModeTransitionDuration() uint {
	var _arg0 *C.AdwLeaflet // out
	var _cret C.guint       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_mode_transition_duration(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Page returns the LeafletPage object for child.
func (self *Leaflet) Page(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.adw_leaflet_get_page(_arg0, _arg1)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Pages returns a Model that contains the pages of the leaflet, and can be used
// to keep an up-to-date view. The model also implements SelectionModel and can
// be used to track the visible page.
func (self *Leaflet) Pages() gtk.SelectionModeller {
	var _arg0 *C.AdwLeaflet        // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_pages(_arg0)

	var _selectionModel gtk.SelectionModeller // out

	_selectionModel = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gtk.SelectionModeller)

	return _selectionModel
}

// TransitionType gets the type of animation that will be used for transitions
// between modes and children in self.
func (self *Leaflet) TransitionType() LeafletTransitionType {
	var _arg0 *C.AdwLeaflet              // out
	var _cret C.AdwLeafletTransitionType // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_transition_type(_arg0)

	var _leafletTransitionType LeafletTransitionType // out

	_leafletTransitionType = LeafletTransitionType(_cret)

	return _leafletTransitionType
}

// VisibleChild gets the visible child widget.
func (self *Leaflet) VisibleChild() gtk.Widgetter {
	var _arg0 *C.AdwLeaflet // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_visible_child(_arg0)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// VisibleChildName gets the name of the currently visible child widget.
func (self *Leaflet) VisibleChildName() string {
	var _arg0 *C.AdwLeaflet // out
	var _cret *C.char       // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_get_visible_child_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// InsertChildAfter inserts child in the position after sibling in the list of
// children. If sibling is NULL, insert child at the first position.
func (self *Leaflet) InsertChildAfter(child gtk.Widgetter, sibling gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _arg2 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((sibling).(gextras.Nativer).Native()))

	_cret = C.adw_leaflet_insert_child_after(_arg0, _arg1, _arg2)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Navigate switches to the previous or next child that doesn't have
// 'navigatable' child property set to FALSE, similar to performing a swipe
// gesture to go in direction.
func (self *Leaflet) Navigate(direction NavigationDirection) bool {
	var _arg0 *C.AdwLeaflet            // out
	var _arg1 C.AdwNavigationDirection // out
	var _cret C.gboolean               // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwNavigationDirection(direction)

	_cret = C.adw_leaflet_navigate(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Prepend inserts child at the first position in self.
func (self *Leaflet) Prepend(child gtk.Widgetter) *LeafletPage {
	var _arg0 *C.AdwLeaflet     // out
	var _arg1 *C.GtkWidget      // out
	var _cret *C.AdwLeafletPage // in

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	_cret = C.adw_leaflet_prepend(_arg0, _arg1)

	var _leafletPage *LeafletPage // out

	_leafletPage = wrapLeafletPage(externglib.Take(unsafe.Pointer(_cret)))

	return _leafletPage
}

// Remove removes a child widget from self.
func (self *Leaflet) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.adw_leaflet_remove(_arg0, _arg1)
}

// ReorderChildAfter moves child to the position after sibling in the list of
// children. If sibling is NULL, move child to the first position.
func (self *Leaflet) ReorderChildAfter(child gtk.Widgetter, sibling gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer((sibling).(gextras.Nativer).Native()))

	C.adw_leaflet_reorder_child_after(_arg0, _arg1, _arg2)
}

// SetCanSwipeBack sets whether or not self allows switching to the previous
// child that has 'navigatable' child property set to TRUE via a swipe gesture
func (self *Leaflet) SetCanSwipeBack(canSwipeBack bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if canSwipeBack {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_swipe_back(_arg0, _arg1)
}

// SetCanSwipeForward sets whether or not self allows switching to the next
// child that has 'navigatable' child property set to TRUE via a swipe gesture.
func (self *Leaflet) SetCanSwipeForward(canSwipeForward bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if canSwipeForward {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_swipe_forward(_arg0, _arg1)
}

func (self *Leaflet) SetCanUnfold(canUnfold bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if canUnfold {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_can_unfold(_arg0, _arg1)
}

// SetChildTransitionDuration sets the duration that transitions between
// children in self will take.
func (self *Leaflet) SetChildTransitionDuration(duration uint) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.guint       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(duration)

	C.adw_leaflet_set_child_transition_duration(_arg0, _arg1)
}

// SetHomogeneous sets the Leaflet to be homogeneous or not for the given fold
// and orientation. If it is homogeneous, the Leaflet will request the same
// width or height for all its children depending on the orientation. If it
// isn't and it is folded, the leaflet may change width or height when a
// different child becomes visible.
func (self *Leaflet) SetHomogeneous(folded bool, orientation gtk.Orientation, homogeneous bool) {
	var _arg0 *C.AdwLeaflet    // out
	var _arg1 C.gboolean       // out
	var _arg2 C.GtkOrientation // out
	var _arg3 C.gboolean       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if folded {
		_arg1 = C.TRUE
	}
	_arg2 = C.GtkOrientation(orientation)
	if homogeneous {
		_arg3 = C.TRUE
	}

	C.adw_leaflet_set_homogeneous(_arg0, _arg1, _arg2, _arg3)
}

// SetInterpolateSize sets whether or not self will interpolate its size when
// changing the visible child. If the Leaflet:interpolate-size property is set
// to TRUE, self will interpolate its size between the current one and the one
// it'll take after changing the visible child, according to the set transition
// duration.
func (self *Leaflet) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_set_interpolate_size(_arg0, _arg1)
}

// SetModeTransitionDuration sets the duration that transitions between modes in
// self will take.
func (self *Leaflet) SetModeTransitionDuration(duration uint) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 C.guint       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(duration)

	C.adw_leaflet_set_mode_transition_duration(_arg0, _arg1)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between modes and children in self.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the mode or child that is about to
// become current.
func (self *Leaflet) SetTransitionType(transition LeafletTransitionType) {
	var _arg0 *C.AdwLeaflet              // out
	var _arg1 C.AdwLeafletTransitionType // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = C.AdwLeafletTransitionType(transition)

	C.adw_leaflet_set_transition_type(_arg0, _arg1)
}

// SetVisibleChild makes visible_child visible using a transition determined by
// AdwLeaflet:transition-type and AdwLeaflet:child-transition-duration. The
// transition can be cancelled by the user, in which case visible child will
// change back to the previously visible child.
func (self *Leaflet) SetVisibleChild(visibleChild gtk.Widgetter) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((visibleChild).(gextras.Nativer).Native()))

	C.adw_leaflet_set_visible_child(_arg0, _arg1)
}

// SetVisibleChildName makes the child with the name name visible.
//
// See adw_leaflet_set_visible_child() for more details.
func (self *Leaflet) SetVisibleChildName(name string) {
	var _arg0 *C.AdwLeaflet // out
	var _arg1 *C.char       // out

	_arg0 = (*C.AdwLeaflet)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))

	C.adw_leaflet_set_visible_child_name(_arg0, _arg1)
}

type LeafletPage struct {
	*externglib.Object
}

var _ gextras.Nativer = (*LeafletPage)(nil)

func wrapLeafletPage(obj *externglib.Object) *LeafletPage {
	return &LeafletPage{
		Object: obj,
	}
}

func marshalLeafletPager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLeafletPage(obj), nil
}

// Child returns the leaflet child to which self belongs.
func (self *LeafletPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwLeafletPage // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_page_get_child(_arg0)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// Name returns the current value of the LeafletPage:name property.
func (self *LeafletPage) Name() string {
	var _arg0 *C.AdwLeafletPage // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_page_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Navigatable gets whether the child can be navigated to when folded.
//
// See adw_leaflet_page_set_navigatable() and LeafletPage:navigatable.
func (self *LeafletPage) Navigatable() bool {
	var _arg0 *C.AdwLeafletPage // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_leaflet_page_get_navigatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetName sets the new value of the LeafletPage:name property. See also
// adw_leaflet_page_get_name()
func (self *LeafletPage) SetName(name string) {
	var _arg0 *C.AdwLeafletPage // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))

	C.adw_leaflet_page_set_name(_arg0, _arg1)
}

// SetNavigatable sets whether the child can be navigated to when folded. If
// FALSE, the child will be ignored by adw_leaflet_get_adjacent_child(),
// adw_leaflet_navigate(), and swipe gestures.
//
// This can be used used to prevent switching to widgets like separators.
//
// Sets the new value of the LeafletPage:navigatable property to navigatable.
func (self *LeafletPage) SetNavigatable(navigatable bool) {
	var _arg0 *C.AdwLeafletPage // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwLeafletPage)(unsafe.Pointer(self.Native()))
	if navigatable {
		_arg1 = C.TRUE
	}

	C.adw_leaflet_page_set_navigatable(_arg0, _arg1)
}
