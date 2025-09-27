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
	PopulationCount    int
	Government         int8
	Law_Level          int8
	NumFactions        int8
	Factions           []int8
	Technology_Level   int8
	Starport           int8
	Bases              []starBase
	Codes              string
	Detailed           bool
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
	Bases            = widget.NewLabel("")
	trade_codes      = widget.NewLabel("")

	systemDetailsBox = widget.NewVBox(
		uw_profile, widget.NewLabel("System Details"),
		stars, size, atmosphere, hydrology,
		population, government,
		law_level, technology_level,
		starport, Bases, trade_codes,
	)
)

func (s *system) init(box *widget.Box) {
	s.getDetails()
	// All the details filled in, from here on out we can
	// just get each one out of the text as needed.
	s.Detailed = true

	box.Children = append(box.Children, systemDetailsBox)
}

func (s *system) getDetails() {
	s.getStars()
	s.getSize()
	s.getAtmosphere()
	s.getHydrology()
	s.getPopulation()
	s.getGovernment()
	s.getLaw()
	s.getStarport()
	s.Technology_Level = s.getTechLevel()
	s.getBases()
	s.getCodes()
}
