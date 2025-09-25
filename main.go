package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var (
	nextSys    system
	mySector   sector
	nxtUI      widget.Box
	systemGrid [8][10]system
)

func main() {
	mySector.init(&nxtUI)

	a := app.New()
	w := a.NewWindow("Taveller SRD Sector Generator")

	//nextSys.init(&nxtUI)
	w.SetContent(&nxtUI)

	w.ShowAndRun()
}
