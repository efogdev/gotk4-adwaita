// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

// DialogClass: instance of this type is always passed by reference.
type DialogClass struct {
	*dialogClass
}

// dialogClass is the struct that's finalized.
type dialogClass struct {
	native *C.AdwDialogClass
}

func (d *DialogClass) ParentClass() *gtk.WidgetClass {
	valptr := &d.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}