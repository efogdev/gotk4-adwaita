// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeTabViewShortcuts = coreglib.Type(C.adw_tab_view_shortcuts_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTabViewShortcuts, F: marshalTabViewShortcuts},
	})
}

// TabViewShortcuts describes available shortcuts in an tabview.
//
// Shortcuts can be set with tabview:shortcuts, or added/removed individually
// with tabview.AddShortcuts and tabview.RemoveShortcuts.
//
// New values may be added to this enumeration over time.
type TabViewShortcuts C.guint

const (
	// TabViewShortcutNone: no shortcuts.
	TabViewShortcutNone TabViewShortcuts = 0b0
	// TabViewShortcutControlTab: <kbd>Ctrl</kbd>+<kbd>Tab</kbd> - switch to the
	// next page, with looping.
	TabViewShortcutControlTab TabViewShortcuts = 0b1
	// TabViewShortcutControlShiftTab:
	// <kbd>Shift</kbd>+<kbd>Ctrl</kbd>+<kbd>Tab</kbd> - switch to the previous
	// page, with looping.
	TabViewShortcutControlShiftTab TabViewShortcuts = 0b10
	// TabViewShortcutControlPageUp: <kbd>Ctrl</kbd>+<kbd>Page Up</kbd> - switch
	// to the previous page.
	TabViewShortcutControlPageUp TabViewShortcuts = 0b100
	// TabViewShortcutControlPageDown: <kbd>Ctrl</kbd>+<kbd>Page Down</kbd> -
	// switch to the next page.
	TabViewShortcutControlPageDown TabViewShortcuts = 0b1000
	// TabViewShortcutControlHome: <kbd>Ctrl</kbd>+<kbd>Home</kbd> - switch to
	// the first page.
	TabViewShortcutControlHome TabViewShortcuts = 0b10000
	// TabViewShortcutControlEnd: <kbd>Ctrl</kbd>+<kbd>End</kbd> - switch to the
	// last page.
	TabViewShortcutControlEnd TabViewShortcuts = 0b100000
	// TabViewShortcutControlShiftPageUp:
	// <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>Page Up</kbd> - move the selected
	// page backward.
	TabViewShortcutControlShiftPageUp TabViewShortcuts = 0b1000000
	// TabViewShortcutControlShiftPageDown:
	// <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>Page Down</kbd> - move the selected
	// page forward.
	TabViewShortcutControlShiftPageDown TabViewShortcuts = 0b10000000
	// TabViewShortcutControlShiftHome:
	// <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>Home</kbd> - move the selected page
	// at the start.
	TabViewShortcutControlShiftHome TabViewShortcuts = 0b100000000
	// TabViewShortcutControlShiftEnd:
	// <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>End</kbd> - move the current page
	// at the end.
	TabViewShortcutControlShiftEnd TabViewShortcuts = 0b1000000000
	// TabViewShortcutAltDigits: <kbd>Alt</kbd>+<kbd>1</kbd>⋯<kbd>9</kbd> -
	// switch to pages 1-9.
	TabViewShortcutAltDigits TabViewShortcuts = 0b10000000000
	// TabViewShortcutAltZero: <kbd>Alt</kbd>+<kbd>0</kbd> - switch to page 10.
	TabViewShortcutAltZero TabViewShortcuts = 0b100000000000
	// TabViewShortcutAllShortcuts: all of the shortcuts.
	TabViewShortcutAllShortcuts TabViewShortcuts = 0b111111111111
)

func marshalTabViewShortcuts(p uintptr) (interface{}, error) {
	return TabViewShortcuts(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for TabViewShortcuts.
func (t TabViewShortcuts) String() string {
	if t == 0 {
		return "TabViewShortcuts(0)"
	}

	var builder strings.Builder
	builder.Grow(256)

	for t != 0 {
		next := t & (t - 1)
		bit := t - next

		switch bit {
		case TabViewShortcutNone:
			builder.WriteString("None|")
		case TabViewShortcutControlTab:
			builder.WriteString("ControlTab|")
		case TabViewShortcutControlShiftTab:
			builder.WriteString("ControlShiftTab|")
		case TabViewShortcutControlPageUp:
			builder.WriteString("ControlPageUp|")
		case TabViewShortcutControlPageDown:
			builder.WriteString("ControlPageDown|")
		case TabViewShortcutControlHome:
			builder.WriteString("ControlHome|")
		case TabViewShortcutControlEnd:
			builder.WriteString("ControlEnd|")
		case TabViewShortcutControlShiftPageUp:
			builder.WriteString("ControlShiftPageUp|")
		case TabViewShortcutControlShiftPageDown:
			builder.WriteString("ControlShiftPageDown|")
		case TabViewShortcutControlShiftHome:
			builder.WriteString("ControlShiftHome|")
		case TabViewShortcutControlShiftEnd:
			builder.WriteString("ControlShiftEnd|")
		case TabViewShortcutAltDigits:
			builder.WriteString("AltDigits|")
		case TabViewShortcutAltZero:
			builder.WriteString("AltZero|")
		case TabViewShortcutAllShortcuts:
			builder.WriteString("AllShortcuts|")
		default:
			builder.WriteString(fmt.Sprintf("TabViewShortcuts(0b%b)|", bit))
		}

		t = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if t contains other.
func (t TabViewShortcuts) Has(other TabViewShortcuts) bool {
	return (t & other) == other
}