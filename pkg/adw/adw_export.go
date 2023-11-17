// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

//export _gotk4_adw1_ActionRowClass_activate
func _gotk4_adw1_ActionRowClass_activate(arg0 *C.AdwActionRow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ActionRowOverrides](instance0)
	if overrides.Activate == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ActionRowOverrides.Activate, got none")
	}

	overrides.Activate()
}

//export _gotk4_adw1_MessageDialogClass_response
func _gotk4_adw1_MessageDialogClass_response(arg0 *C.AdwMessageDialog, arg1 *C.char) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[MessageDialogOverrides](instance0)
	if overrides.Response == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected MessageDialogOverrides.Response, got none")
	}

	var _response string // out

	_response = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	overrides.Response(_response)
}
