package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var (
	nextSys   system
	nextSysUI widget.Box
)

func main() {
	a := app.New()
	w := a.NewWindow("Taveller SRD System Generator")

	nextSys.init()

	w.SetContent(&nextSysUI)
	w.ShowAndRun()
}
