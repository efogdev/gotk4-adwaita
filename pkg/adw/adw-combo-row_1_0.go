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
import "C"

// GType values.
var (
	GTypeComboRow = coreglib.Type(C.adw_combo_row_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeComboRow, F: marshalComboRow},
	})
}

// ComboRowOverrides contains methods that are overridable.
type ComboRowOverrides struct {
}

func defaultComboRowOverrides(v *ComboRow) ComboRowOverrides {
	return ComboRowOverrides{}
}

// ComboRow: gtk.ListBoxRow used to choose from a list of items.
//
// <picture> <source srcset="combo-row-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="combo-row.png" alt="combo-row"> </picture>
//
// The AdwComboRow widget allows the user to choose from a list of valid
// choices. The row displays the selected choice. When activated, the row
// displays a popover which allows the user to make a new choice.
//
// Example of an AdwComboRow UI definition:
//
//    <object class="AdwComboRow">
//      <property name="title" translatable="yes">Combo Row</property>
//      <property name="model">
//        <object class="GtkStringList">
//          <items>
//            <item translatable="yes">Foo</item>
//            <item translatable="yes">Bar</item>
//            <item translatable="yes">Baz</item>
//          </items>
//        </object>
//      </property>
//    </object>
//
// The comborow:selected and comborow:selected-item properties can be used to
// keep track of the selected item and react to their changes.
//
// AdwComboRow mirrors gtk.DropDown, see that widget for details.
//
// AdwComboRow is gtk.ListBoxRow:activatable if a model is set.
//
// # CSS nodes
//
// AdwComboRow has a main CSS node with name row and the .combo style class.
//
// Its popover has the node named popover with the .menu style class,
// it contains a gtk.ScrolledWindow, which in turn contains a gtk.ListView,
// both are accessible via their regular nodes.
//
// # Accessibility
//
// AdwComboRow uses the GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type ComboRow struct {
	_ [0]func() // equal guard
	ActionRow
}

var (
	_ gtk.Widgetter     = (*ComboRow)(nil)
	_ coreglib.Objector = (*ComboRow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ComboRow, *ComboRowClass, ComboRowOverrides](
		GTypeComboRow,
		initComboRowClass,
		wrapComboRow,
		defaultComboRowOverrides,
	)
}

func initComboRowClass(gclass unsafe.Pointer, overrides ComboRowOverrides, classInitFunc func(*ComboRowClass)) {
	if classInitFunc != nil {
		class := (*ComboRowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapComboRow(obj *coreglib.Object) *ComboRow {
	return &ComboRow{
		ActionRow: ActionRow{
			PreferencesRow: PreferencesRow{
				ListBoxRow: gtk.ListBoxRow{
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
				},
			},
		},
	}
}

func marshalComboRow(p uintptr) (interface{}, error) {
	return wrapComboRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewComboRow creates a new AdwComboRow.
//
// The function returns the following values:
//
//   - comboRow: newly created AdwComboRow.
//
func NewComboRow() *ComboRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_combo_row_new()

	var _comboRow *ComboRow // out

	_comboRow = wrapComboRow(coreglib.Take(unsafe.Pointer(_cret)))

	return _comboRow
}

// Expression gets the expression used to obtain strings from items.
//
// The function returns the following values:
//
//   - expression (optional) used to obtain strings from items.
//
func (self *ComboRow) Expression() gtk.Expressioner {
	var _arg0 *C.AdwComboRow   // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_combo_row_get_expression(_arg0)
	runtime.KeepAlive(self)

	var _expression gtk.Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Expressioner)
				return ok
			})
			rv, ok := casted.(gtk.Expressioner)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// Factory gets the factory for populating list items.
//
// The function returns the following values:
//
//   - listItemFactory (optional): factory in use.
//
func (self *ComboRow) Factory() *gtk.ListItemFactory {
	var _arg0 *C.AdwComboRow        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_combo_row_get_factory(_arg0)
	runtime.KeepAlive(self)

	var _listItemFactory *gtk.ListItemFactory // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listItemFactory = &gtk.ListItemFactory{
				Object: obj,
			}
		}
	}

	return _listItemFactory
}

// ListFactory gets the factory for populating list items in the popup.
//
// The function returns the following values:
//
//   - listItemFactory (optional): factory in use.
//
func (self *ComboRow) ListFactory() *gtk.ListItemFactory {
	var _arg0 *C.AdwComboRow        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_combo_row_get_list_factory(_arg0)
	runtime.KeepAlive(self)

	var _listItemFactory *gtk.ListItemFactory // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listItemFactory = &gtk.ListItemFactory{
				Object: obj,
			}
		}
	}

	return _listItemFactory
}

// Model gets the model that provides the displayed items.
//
// The function returns the following values:
//
//   - listModel (optional): model in use.
//
func (self *ComboRow) Model() *gio.ListModel {
	var _arg0 *C.AdwComboRow // out
	var _cret *C.GListModel  // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_combo_row_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel *gio.ListModel // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_listModel = &gio.ListModel{
				Object: obj,
			}
		}
	}

	return _listModel
}

// Selected gets the position of the selected item.
//
// The function returns the following values:
//
//   - guint: position of the selected item, or gtk.INVALIDLISTPOSITION if no
//     item is selected.
//
func (self *ComboRow) Selected() uint {
	var _arg0 *C.AdwComboRow // out
	var _cret C.guint        // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_combo_row_get_selected(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SelectedItem gets the selected item.
//
// The function returns the following values:
//
//   - object (optional): selected item.
//
func (self *ComboRow) SelectedItem() *coreglib.Object {
	var _arg0 *C.AdwComboRow // out
	var _cret C.gpointer     // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_combo_row_get_selected_item(_arg0)
	runtime.KeepAlive(self)

	var _object *coreglib.Object // out

	_object = coreglib.Take(unsafe.Pointer(_cret))

	return _object
}

// UseSubtitle gets whether to use the current value as the subtitle.
//
// The function returns the following values:
//
//   - ok: whether to use the current value as the subtitle.
//
func (self *ComboRow) UseSubtitle() bool {
	var _arg0 *C.AdwComboRow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_combo_row_get_use_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpression sets the expression used to obtain strings from items.
//
// The expression must have a value type of G_TYPE_STRING.
//
// It's used to bind strings to labels produced by the default factory if
// comborow:factory is not set, or when comborow:use-subtitle is set to TRUE.
//
// The function takes the following parameters:
//
//   - expression (optional): expression.
//
func (self *ComboRow) SetExpression(expression gtk.Expressioner) {
	var _arg0 *C.AdwComboRow   // out
	var _arg1 *C.GtkExpression // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(coreglib.InternObject(expression).Native()))
	}

	C.adw_combo_row_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetFactory sets the factory for populating list items.
//
// This factory is always used for the item in the row. It is also used for
// items in the popup unless comborow:list-factory is set.
//
// The function takes the following parameters:
//
//   - factory (optional) to use.
//
func (self *ComboRow) SetFactory(factory *gtk.ListItemFactory) {
	var _arg0 *C.AdwComboRow        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	}

	C.adw_combo_row_set_factory(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetListFactory sets the factory for populating list items in the popup.
//
// If this is not set, comborow:factory is used.
//
// The function takes the following parameters:
//
//   - factory (optional) to use.
//
func (self *ComboRow) SetListFactory(factory *gtk.ListItemFactory) {
	var _arg0 *C.AdwComboRow        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	}

	C.adw_combo_row_set_list_factory(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetModel sets the model that provides the displayed items.
//
// The function takes the following parameters:
//
//   - model (optional) to use.
//
func (self *ComboRow) SetModel(model gio.ListModeller) {
	var _arg0 *C.AdwComboRow // out
	var _arg1 *C.GListModel  // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	C.adw_combo_row_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetSelected selects the item at the given position.
//
// The function takes the following parameters:
//
//   - position of the item to select, or gtk.INVALIDLISTPOSITION.
//
func (self *ComboRow) SetSelected(position uint) {
	var _arg0 *C.AdwComboRow // out
	var _arg1 C.guint        // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint(position)

	C.adw_combo_row_set_selected(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}

// SetUseSubtitle sets whether to use the current value as the subtitle.
//
// If you use a custom list item factory, you will need to give the row a name
// conversion expression with comborow:expression.
//
// If set to TRUE, you should not access actionrow:subtitle.
//
// The subtitle is interpreted as Pango markup if preferencesrow:use-markup is
// set to TRUE.
//
// The function takes the following parameters:
//
//   - useSubtitle: whether to use the current value as the subtitle.
//
func (self *ComboRow) SetUseSubtitle(useSubtitle bool) {
	var _arg0 *C.AdwComboRow // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.AdwComboRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useSubtitle {
		_arg1 = C.TRUE
	}

	C.adw_combo_row_set_use_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useSubtitle)
}
