// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

//export _gotk4_adw1_AlertDialog_ConnectResponse
func _gotk4_adw1_AlertDialog_ConnectResponse(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(response string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(response string))
	}

	var _response string // out

	_response = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_response)
}