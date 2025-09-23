package main

import "fyne.io/fyne/widget"

func (s *system) initButton() {
	generateAgain = *widget.NewButton("Generate Another", s.again)
}

func (s *system) again() {
	s.createSystem()
}
