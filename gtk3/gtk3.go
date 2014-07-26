package gtk3

// #include "gtk3.go.h"
// #cgo pkg-config: gtk+-3.0
import "C"
import (
	"unsafe"
        "github.com/mattn/go-gtk/glib"
)

// Glib Stuff
//  yuck
func gstring(s *C.char) *C.gchar { return C.toGstr(s) }

// General Functions
// https://developer.gnome.org/gtk3/stable/gtk3-General.html
func Init() {
	C.gtk_init(nil, nil)
}

func Main() {
	C.gtk_main()
}

func MainQuit() {
	C.gtk_main_quit()
}

// Widget Functions
// https://developer.gnome.org/gtk3/stable/GtkWidget.html

type Widget struct {
	GWidget *C.GtkWidget
}

func WidgetFromNative(p unsafe.Pointer) *Widget {
	return &Widget{C.toGWidget(p)}
}

func (v *Widget) Show() {
	C.gtk_widget_show(v.GWidget)
}

func (v *Widget) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GWidget)).Connect(s, f, datas...)
}

func (v *Widget) Add(w *Widget) {
	C.gtk_container_add(C.toGContainer(v.GWidget), w.GWidget)
}

// Window Functions
// https://developer.gnome.org/gtk3/stable/GtkWindow.html

type Window struct {
	GWindow *C.GtkWindow
}

func WINDOW(p *Widget) *C.GtkWindow                 { return C.toGWindow(p.GWidget) }

type WindowType int

const (
	WINDOW_TOPLEVEL WindowType = C.GTK_WINDOW_TOPLEVEL
)

func NewWindow(t WindowType) *Widget {
	return &Widget{C.gtk_window_new(C.GtkWindowType(t))}
}

func (v *Widget) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free(unsafe.Pointer(ptr))
	C.gtk_window_set_title(WINDOW(v), gstring(ptr))
}
