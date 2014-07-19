package main

// go version of demo at the top of this page:
// https://developer.gnome.org/gtk3/stable/gtk-getting-started.html

import (
	"github.com/willhaines/willterm/gtk3"
)

func main() {
	gtk3.Init()
	window := gtk3.NewWindow(gtk3.WINDOW_TOPLEVEL)
	window.SetTitle("This better fucking work")
	window.Show()
	window.Connect("destroy", gtk3.MainQuit)
	gtk3.Main()
}
