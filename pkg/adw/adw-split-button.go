// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
// extern void _gotk4_adw1_SplitButton_ConnectClicked(gpointer, guintptr);
// extern void _gotk4_adw1_SplitButton_ConnectActivate(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeSplitButton = coreglib.Type(C.adw_split_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSplitButton, F: marshalSplitButton},
	})
}

// SplitButtonOverrides contains methods that are overridable.
type SplitButtonOverrides struct {
}

func defaultSplitButtonOverrides(v *SplitButton) SplitButtonOverrides {
	return SplitButtonOverrides{}
}

// SplitButton: combined button and dropdown widget.
//
// <picture> <source srcset="split-button-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="split-button.png"
// alt="split-button"> </picture>
//
// AdwSplitButton is typically used to present a set of actions in a menu,
// but allow access to one of them with a single click.
//
// The API is very similar to gtk.Button and gtk.MenuButton, see their
// documentation for details.
//
// CSS nodes
//
//	splitbutton[.image-button][.text-button]
//	├── button
//	│   ╰── <content>
//	├── separator
//	╰── menubutton
//	    ╰── button.toggle
//	        ╰── arrow
//
// AdwSplitButton's CSS node is called splitbutton. It contains the css nodes:
// button, separator, menubutton. See gtk.MenuButton documentation for the
// menubutton contents.
//
// The main CSS node will contain the .image-button or .text-button style
// classes matching the button contents. The nested button nodes will never
// contain them.
//
// # Style classes
//
// AdwSplitButton can use some of the same style classes as gtk.Button:
//
// - .suggested-action (style-classes.html#suggested-action)
//
// - .destructive-action (style-classes.html#destructive-action)
//
// - .flat (style-classes.html#flat)
//
// - .raised (style-classes.html#raised)
//
// Other style classes, like .pill, cannot be used.
//
// # Accessibility
//
// AdwSplitButton uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type SplitButton struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	gtk.Actionable
}

var (
	_ gtk.Widgetter     = (*SplitButton)(nil)
	_ coreglib.Objector = (*SplitButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*SplitButton, *SplitButtonClass, SplitButtonOverrides](
		GTypeSplitButton,
		initSplitButtonClass,
		wrapSplitButton,
		defaultSplitButtonOverrides,
	)
}

func initSplitButtonClass(gclass unsafe.Pointer, overrides SplitButtonOverrides, classInitFunc func(*SplitButtonClass)) {
	if classInitFunc != nil {
		class := (*SplitButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSplitButton(obj *coreglib.Object) *SplitButton {
	return &SplitButton{
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
		Actionable: gtk.Actionable{
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
	}
}

func marshalSplitButton(p uintptr) (interface{}, error) {
	return wrapSplitButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate is emitted to animate press then release.
//
// This is an action signal. Applications should never connect to this signal,
// but use the splitbutton::clicked signal.
func (self *SplitButton) ConnectActivate(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "activate", false, unsafe.Pointer(C._gotk4_adw1_SplitButton_ConnectActivate), f)
}

// ConnectClicked is emitted when the button has been activated (pressed and
// released).
func (self *SplitButton) ConnectClicked(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "clicked", false, unsafe.Pointer(C._gotk4_adw1_SplitButton_ConnectClicked), f)
}

// NewSplitButton creates a new AdwSplitButton.
//
// The function returns the following values:
//
//   - splitButton: newly created AdwSplitButton.
func NewSplitButton() *SplitButton {
	var _cret *C.GtkWidget // in

	_cret = C.adw_split_button_new()

	var _splitButton *SplitButton // out

	_splitButton = wrapSplitButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _splitButton
}

// CanShrink gets whether the button can be smaller than the natural size of its
// contents.
//
// The function returns the following values:
//
//   - ok: whether the button can shrink.
func (self *SplitButton) CanShrink() bool {
	var _arg0 *C.AdwSplitButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_can_shrink(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child widget.
//
// The function returns the following values:
//
//   - widget (optional): child widget.
func (self *SplitButton) Child() gtk.Widgetter {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_child(_arg0)
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

// Direction gets the direction in which the popup will be popped up.
//
// The function returns the following values:
//
//   - arrowType: direction.
func (self *SplitButton) Direction() gtk.ArrowType {
	var _arg0 *C.AdwSplitButton // out
	var _cret C.GtkArrowType    // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_direction(_arg0)
	runtime.KeepAlive(self)

	var _arrowType gtk.ArrowType // out

	_arrowType = gtk.ArrowType(_cret)

	return _arrowType
}

// DropdownTooltip gets the tooltip of the dropdown button of self.
//
// The function returns the following values:
//
//   - utf8: dropdown tooltip of self.
func (self *SplitButton) DropdownTooltip() string {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_dropdown_tooltip(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IconName gets the name of the icon used to automatically populate the button.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name.
func (self *SplitButton) IconName() string {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Label gets the label for self.
//
// The function returns the following values:
//
//   - utf8 (optional): label for self.
func (self *SplitButton) Label() string {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_label(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// MenuModel gets the menu model from which the popup will be created.
//
// The function returns the following values:
//
//   - menuModel (optional): menu model.
func (self *SplitButton) MenuModel() gio.MenuModeller {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.GMenuModel     // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_menu_model(_arg0)
	runtime.KeepAlive(self)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gio.MenuModeller)
				return ok
			})
			rv, ok := casted.(gio.MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// Popover gets the popover that will be popped up when the dropdown is clicked.
//
// The function returns the following values:
//
//   - popover (optional): popover.
func (self *SplitButton) Popover() *gtk.Popover {
	var _arg0 *C.AdwSplitButton // out
	var _cret *C.GtkPopover     // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_popover(_arg0)
	runtime.KeepAlive(self)

	var _popover *gtk.Popover // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_popover = &gtk.Popover{
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
				NativeSurface: gtk.NativeSurface{
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
				ShortcutManager: gtk.ShortcutManager{
					Object: obj,
				},
			}
		}
	}

	return _popover
}

// UseUnderline gets whether an underline in the text indicates a mnemonic.
//
// The function returns the following values:
//
//   - ok: whether an underline in the text indicates a mnemonic.
func (self *SplitButton) UseUnderline() bool {
	var _arg0 *C.AdwSplitButton // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_split_button_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown dismisses the menu.
func (self *SplitButton) Popdown() {
	var _arg0 *C.AdwSplitButton // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_split_button_popdown(_arg0)
	runtime.KeepAlive(self)
}

// Popup pops up the menu.
func (self *SplitButton) Popup() {
	var _arg0 *C.AdwSplitButton // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	C.adw_split_button_popup(_arg0)
	runtime.KeepAlive(self)
}

// SetCanShrink sets whether the button can be smaller than the natural size of
// its contents.
//
// If set to TRUE, the label will ellipsize.
//
// See gtk.Button.SetCanShrink() and gtk.MenuButton.SetCanShrink().
//
// The function takes the following parameters:
//
//   - canShrink: whether the button can shrink.
func (self *SplitButton) SetCanShrink(canShrink bool) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if canShrink {
		_arg1 = C.TRUE
	}

	C.adw_split_button_set_can_shrink(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(canShrink)
}

// SetChild sets the child widget.
//
// Setting the child widget will set splitbutton:label and splitbutton:icon-name
// to NULL.
//
// The function takes the following parameters:
//
//   - child (optional): new child widget.
func (self *SplitButton) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.adw_split_button_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetDirection sets the direction in which the popup will be popped up.
//
// The dropdown arrow icon will point at the same direction.
//
// If the does not fit in the available space in the given direction, GTK will
// try its best to keep it inside the screen and fully visible.
//
// If you pass GTK_ARROW_NONE, it's equivalent to GTK_ARROW_DOWN.
//
// The function takes the following parameters:
//
//   - direction: direction.
func (self *SplitButton) SetDirection(direction gtk.ArrowType) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 C.GtkArrowType    // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkArrowType(direction)

	C.adw_split_button_set_direction(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(direction)
}

// SetDropdownTooltip sets the tooltip of the dropdown button of self.
//
// The tooltip can be marked up with the Pango text markup language.
//
// The function takes the following parameters:
//
//   - tooltip: dropdown tooltip of self.
func (self *SplitButton) SetDropdownTooltip(tooltip string) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(tooltip)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_split_button_set_dropdown_tooltip(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(tooltip)
}

// SetIconName sets the name of the icon used to automatically populate the
// button.
//
// Setting the icon name will set splitbutton:label and splitbutton:child to
// NULL.
//
// The function takes the following parameters:
//
//   - iconName: icon name to set.
func (self *SplitButton) SetIconName(iconName string) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_split_button_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetLabel sets the label for self.
//
// Setting the label will set splitbutton:icon-name and splitbutton:child to
// NULL.
//
// The function takes the following parameters:
//
//   - label to set.
func (self *SplitButton) SetLabel(label string) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_split_button_set_label(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(label)
}

// SetMenuModel sets the menu model from which the popup will be created.
//
// If the menu model is NULL, the dropdown is disabled.
//
// A gtk.Popover will be created from the menu model with
// gtk.PopoverMenu.NewFromModel. Actions will be connected as documented for
// this function.
//
// If splitbutton:popover is already set, it will be dissociated from the
// button, and the property is set to NULL.
//
// The function takes the following parameters:
//
//   - menuModel (optional): menu model.
func (self *SplitButton) SetMenuModel(menuModel gio.MenuModeller) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.GMenuModel     // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if menuModel != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(menuModel).Native()))
	}

	C.adw_split_button_set_menu_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(menuModel)
}

// SetPopover sets the popover that will be popped up when the dropdown is
// clicked.
//
// If the popover is NULL, the dropdown is disabled.
//
// If splitbutton:menu-model is set, the menu model is dissociated from the
// button, and the property is set to NULL.
//
// The function takes the following parameters:
//
//   - popover (optional): popover.
func (self *SplitButton) SetPopover(popover *gtk.Popover) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 *C.GtkPopover     // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if popover != nil {
		_arg1 = (*C.GtkPopover)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	}

	C.adw_split_button_set_popover(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(popover)
}

// SetUseUnderline sets whether an underline in the text indicates a mnemonic.
//
// See splitbutton:label.
//
// The function takes the following parameters:
//
//   - useUnderline: whether an underline in the text indicates a mnemonic.
func (self *SplitButton) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwSplitButton // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwSplitButton)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_split_button_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}

// SplitButtonClass: instance of this type is always passed by reference.
type SplitButtonClass struct {
	*splitButtonClass
}

// splitButtonClass is the struct that's finalized.
type splitButtonClass struct {
	native *C.AdwSplitButtonClass
}

func (s *SplitButtonClass) ParentClass() *gtk.WidgetClass {
	valptr := &s.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
