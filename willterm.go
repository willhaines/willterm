package main

import (
	"github.com/mattn/go-gtk/gtk"
	"github.com/str1ngs/vte"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Notebook")
	window.Connect("destroy", gtk.MainQuit)

	window2 := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window2.SetTitle("GTK Notebook")
	window2.Connect("destroy", gtk.MainQuit)

	notebook := gtk.NewNotebook()
	page := gtk.NewFrame("demo")
	notebook.AppendPage(page, gtk.NewLabel("demo"))
	notebook.SetTabDetachable(page, true)

	terminal := vte.NewTerminal()
	terminal.Fork([]string{"/bin/zsh", "--login"})

	page.Add(terminal)

	notebook.SetGroupName("foo")
	window.Add(notebook)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	notebook2 := gtk.NewNotebook()
	notebook2.SetGroupName("foo")
	window2.Add(notebook2)
	window2.SetSizeRequest(400, 200)
	window2.ShowAll()

	gtk.Main()
}
