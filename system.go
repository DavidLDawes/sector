package main

import (
	"fyne.io/fyne/widget"
)

type starBase struct {
	bType        string
	bDescription string
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
	Epistellar_planets int8
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
	//bases            []starBase
}

var (
	Scout = starBase{
		"Scout",
		"A scout base offers refined fuel and supplies to scout ships.",
	}
	Research = starBase{
		"Research",
		"A Research base is dedicated to a particular field of research.",
	}
	Consulate = starBase{
		"Consulate",
		"A consulate is an administration office for various departments such as commerce, justice and foreign affairs. Characters wishing to report significant crimes or obtain various permits will need to visit a consulate.",
	}
	Pirate = starBase{
		"Pirate",
		"The presence of a pirate base in a system indicates that a group of thieves is active in the area. Pirates are unlikely to be operating out of the starport itself (except on a Law Level 0 world), but no doubt have agents at the port on the lookout for likely prey.",
	}

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
	military_base    = widget.NewLabel("")
	scout_base       = widget.NewLabel("")
	reserach_base    = widget.NewLabel("")
	pirate_base      = widget.NewLabel("")

	systemDetailsBox = widget.NewVBox(
		uw_profile, widget.NewLabel("System Details"),
		stars, size, atmosphere, hydrology,
		population, government,
		law_level, technology_level,
		starport, military_base,
		scout_base, reserach_base,
		pirate_base,
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
	military_base.SetText(mil)
	scout_base.SetText(scout)
	reserach_base.SetText(res)
	pirate_base.SetText(pir)

	box.Children = append(box.Children, systemDetailsBox)
}
