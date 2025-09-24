package main

import (
	"fyne.io/fyne/widget"
)

type starBase struct {
	base        string
	description string
}

type planet struct {
	content    string
	satellites string
	rings      string
}

type system struct {
	Size               int8
	numStars           int8
	starTypes          []string
	Close_Companion    bool
	Epistellar_Planets int8
	EpiPlanets         []planet
	Inner_Planets      int8
	InPlanets          []planet
	Moderate_Companion bool
	Outer_Planets      int8
	OutPlanets         []planet
	Distant_Companion  bool
	numGasGiants       int8
	numAsteroids       int8
	UWP                string
	Atmosphere         int8
	Hydrology          int8
	Population         int8
	Government         int8
	Law_Level          int8
	NumFactions        int8
	Factions           []int8
	Technology_Level   int8
	Starport           int8
	bases              []starBase
}

var (
	uw_profile       = widget.NewLabel("")
	stars            = widget.NewLabel("")
	size             = widget.NewLabel("")
	atmosphere       = widget.NewLabel("")
	hydrology        = widget.NewLabel("")
	population       = widget.NewLabel("")
	government       = widget.NewLabel("")
	law_level        = widget.NewLabel("")
	technology_level = widget.NewLabel("")
	starport         = widget.NewLabel("")
	bases            = widget.NewLabel("")
	trade_codes      = widget.NewLabel("")

	systemDetailsBox = widget.NewVBox(
		uw_profile, widget.NewLabel("System Details"),
		stars, size, atmosphere, hydrology,
		population, government,
		law_level, technology_level,
		starport, bases, trade_codes,
	)
)

func (s *system) init(box *widget.Box) {
	s.getStars()
	s.getSize()
	s.getAtmosphere()
	s.getHydrology()
	s.getPopulation()
	s.getGovernment()
	s.getLaw()
	s.getStarport()
	s.getTechLevel()
	s.getBases()
	s.getCodes()

	box.Children = append(box.Children, systemDetailsBox)
}
