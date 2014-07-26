package vte3

// #include "vte3.go.h"
// #cgo pkg-config: vte-2.90
import "C"

import (
	"github.com/willhaines/willterm/gtk3"
	"unsafe"
)

func (t Terminal) VteToGtk() *gtk3.Widget {
        return gtk3.WidgetFromNative(unsafe.Pointer(t.Widget))
}

type Terminal struct {
	Widget *C.GtkWidget
}

func NewTerminal() Terminal {
	//new_term := unsafe.Pointer(C.vte_terminal_new())
	//widget := gtk3.WidgetFromNative(new_term)
	return Terminal{C.vte_terminal_new()}
}

func (t Terminal) Fork(args []string) {
	argv := C.make_strings(C.int(len(args) + 1))
	defer C.free(unsafe.Pointer(argv))
	for i, arg := range args {
		ptr := C.CString(arg)
		defer C.free(unsafe.Pointer(ptr))
		C.set_string(argv, C.int(i), ptr)
	}
	C.set_string(argv, C.int(len(args)), nil)
	C.vte_terminal_fork_command_full(C.toVTerminal(t.Widget), C.VTE_PTY_DEFAULT,
		nil, argv, nil, C.G_SPAWN_SEARCH_PATH, nil, nil, nil, nil )
}
