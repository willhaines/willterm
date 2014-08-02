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
func gboolean(b C.int) C.gboolean { return C.toGbool(b) }

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

// Application Functions
// https://developer.gnome.org/gtk3/stable/GtkApplication.html

type Application struct {
	GApplication *C.GtkApplication
}

func NewApplication() *Application {
	//ptr := C.CString("com.github.com.willhaines.willterm")
	//ptr := C.CString("")
	//defer C.free(unsafe.Pointer(ptr))
	//return &Application{C.gtk_application_new(gstring(ptr), 0)}
	return &Application{C.gtk_application_new(nil, 0)}
}

func (a *Application) Register() {
	C.g_application_register(C.toGApp(a.GApplication), nil, nil)
}

func (a *Application) Run() {
	C.g_application_run(C.toGApp(a.GApplication), 0, nil)
}

func (a *Application) AddWindow(w *Widget) {
	C.gtk_application_add_window(a.GApplication, WINDOW(w))
}

func (a *Application) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(a.GApplication)).Connect(s, f, datas...)
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

func (v *Widget) GetParent() *Widget {
	return WidgetFromNative(unsafe.Pointer(C.gtk_widget_get_parent(v.GWidget)))
}

func (v *Widget) Destroy() {
	C.gtk_widget_destroy(v.GWidget)
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

func (v *Widget) SetDecorated(setting int) {
	C.gtk_window_set_decorated(WINDOW(v), gboolean(C.int(setting)))
}

func (v *Widget) SetHideTitlebarWhenMaximized(setting int) {
	C.gtk_window_set_hide_titlebar_when_maximized(WINDOW(v), gboolean(C.int(setting)))
}

// Notebook Functions
// https://developer.gnome.org/gtk3/stable/GtkNotebook.html

type Notebook struct {
	GNotebook *C.GtkNotebook
}

func NOTEBOOK(p *Widget) *C.GtkNotebook	{ return C.toGNotebook(p.GWidget) }

func NewNotebook() *Widget {
	return &Widget{C.gtk_notebook_new()}
}

func (n *Widget) PageNum(child *Widget) C.gint {
	return C.gtk_notebook_page_num(NOTEBOOK(n), child.GWidget)
}

func (n *Widget) AppendPage(child *Widget) {
	C.gtk_notebook_append_page(NOTEBOOK(n), child.GWidget, nil)
}

func (n *Widget) removePage(pageNum C.gint) {
	C.gtk_notebook_remove_page(NOTEBOOK(n), pageNum)
}

func RemovePageCallback(child *glib.CallbackContext) {
	childWidget := child.Data().(*Widget)
	//notebook := childWidget.GetParent()
	//pageNum := notebook.PageNum(childWidget)
	childWidget.Destroy()
	//notebook.removePage(pageNum)
}

func (n *Widget) PopupEnable() {
	C.gtk_notebook_popup_enable(NOTEBOOK(n))
}

func (n *Widget) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(NOTEBOOK(n)))
}

func (n *Widget) SetGroupName(name string) {
	ptr := C.CString(name)
	defer C.free(unsafe.Pointer(ptr))
	C.gtk_notebook_set_group_name(NOTEBOOK(n), gstring(ptr))
}

func (n *Widget) SetTabDetachable(child *Widget) {
	C.gtk_notebook_set_tab_detachable(NOTEBOOK(n), child.GWidget, gboolean(C.int(1)))
}

func (n *Widget) SetTabLabelText(child *Widget, title string) {
	C.gtk_notebook_set_tab_label_text(
		NOTEBOOK(n),
		child.GWidget,
		gstring(C.CString(title)))
}

