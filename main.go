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
	nextSysUI := widget.NewVBox()

	a := app.New()
	w := a.NewWindow("Taveller SRD System Generator")

	nextSys.init(*nextSysUI)

	w.SetContent(nextSysUI)
	w.ShowAndRun()
}

func getBox() *widget.Box {
	return &nextSysUI
}
