// Code generated by girgen. DO NOT EDIT.

package adw

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

//export _gotk4_adw1_EntryRow_ConnectApply
func _gotk4_adw1_EntryRow_ConnectApply(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_adw1_EntryRow_ConnectEntryActivated
func _gotk4_adw1_EntryRow_ConnectEntryActivated(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
