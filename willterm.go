package main

import (
	"github.com/willhaines/willterm/gtk3"
	"github.com/willhaines/willterm/vte3"
	"github.com/mattn/go-gtk/glib"
	//"unsafe"
	//"fmt"
	//"reflect"
)

func numPagesCallback(context *glib.CallbackContext) {
	notebook := context.Data().(*gtk3.Widget)
	n := notebook.GetNPages()
	switch n {
	case 0:
		notebook.GetParent().Destroy()
	}
}

func addNewTerminal(notebook *gtk3.Widget) {
	argv := []string{"/bin/zsh"}

	terminal := vte3.NewTerminal()
	widget := terminal.VteToGtk()

	terminal.Fork(argv)
	widget.Show()
	notebook.AppendPage(widget)
	notebook.SetTabDetachable(widget)

	widget.Connect("child-exited", widget.Destroy)
	widget.Connect("icon-title-changed", func (child *glib.CallbackContext) {
		terminal = child.Data().(vte3.Terminal)
		page := terminal.VteToGtk().GetParent()
		title := terminal.GetIconTitle()
		page.SetTabLabelText(terminal.VteToGtk(), title)
	}, terminal)

}

func addNewWindow(app *gtk3.Application) *gtk3.Widget {
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("This better fucking work")
	window.SetHideTitlebarWhenMaximized(1)

	notebook := gtk3.NewNotebook()
	window.Add(notebook)
	notebook.PopupEnable()
	notebook.SetGroupName("willterm")

	addNewTerminal(notebook)
	notebook.Show()
	window.Show()

	notebook.Connect("page-removed", numPagesCallback, notebook)
	//window.Connect("destroy", gtk3.MainQuit)
	app.AddWindow(window)

	return notebook
}

func noop() {
}

func main() {
	//gtk3.Init()
	app := gtk3.NewApplication()
	app.Connect("activate", noop)
	app.Register()

	notebook := addNewWindow(app)
	addNewWindow(app)

	addNewTerminal(notebook)
	addNewTerminal(notebook)
	addNewTerminal(notebook)

	/*
	terminal.VteToGtk().Connect("child-exited",
		notebook.RemovePageCallback, terminal.VteToGtk())
	terminal2.VteToGtk().Connect("child-exited",
		notebook.RemovePageCallback, terminal2.VteToGtk())
	*/
	/*
	notebook.Connect("page-removed", func(ctx *glib.CallbackContext) {
		fmt.Println(reflect.TypeOf(ctx.Target()))

		fmt.Println(reflect.ValueOf(ctx.Target()))
		fmt.Println(ctx.Target().Elem())
		notebook := gtk3.WidgetFromNative(unsafe.Pointer(
			ctx.Target().Elem()))
		if notebook.GetNPages() == 0 {
			gtk3.MainQuit()
		}
	})
	*/

	app.Run()
	//gtk3.Main()
}
