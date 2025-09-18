package main

import (
	"math/rand"

	"fyne.io/fyne/widget"
)

type starBase struct {
	bType        string
	bDescription string
}

type system struct {
	Size             byte
	Atmosphere       byte
	Hydrology        byte
	Population       byte
	Government       byte
	Law_Level        byte
	Technology_Level byte
	Starport         byte
	bases            []starBase
}

type sizeDetail struct {
	size    string
	gravity string
	g       float32
}

var (
	sizeDetails = []sizeDetail{
		sizeDetail{"800 km", "Negligible", 0.0},
		sizeDetail{"1,600 km", "0.05G", 0.05},
		sizeDetail{"3,200 km", "0.15G", 0.15},
		sizeDetail{"4,800 km", "0.25G", 0.25},
		sizeDetail{"6,400 km", "0.35G", 0.35},
		sizeDetail{"8,000 km", "0.45G", 0.45},
		sizeDetail{"9,600 km", "0.7G", 0.7},
		sizeDetail{"11,200 km", "0.9G", 0.9},
		sizeDetail{"12,800 km", "1.0G", 1.0},
		sizeDetail{"14,400 km", "1.25G", 1.25},
	}

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
		widget.NewLabel("System Details"),
		size, atmosphere, hydrology,
		population, government,
		law_level, technology_level,
		starport, military_base,
		scout_base, reserach_base,
		pirate_base,
	)

	sizeOffset int
)

func (s *system) init(box *widget.Box) {

	s.Size = byte(zero_to_ten())
	sizeOffset = trv_to_int(s.Size)
	size.SetText(sz + " " +
		string(s.Size) + ", " + sizeDetails[sizeOffset].size + ", " +
		"gravity " + sizeDetails[sizeOffset].gravity)
	atmosphere = widget.NewLabel(atmos)
	hydrology = widget.NewLabel(hydro)
	population = widget.NewLabel(pop)
	government = widget.NewLabel(gov)
	law_level = widget.NewLabel(law)
	technology_level = widget.NewLabel(tech)
	starport = widget.NewLabel(sp)
	military_base = widget.NewLabel(mil)
	scout_base = widget.NewLabel(scout)
	reserach_base = widget.NewLabel(res)
	pirate_base = widget.NewLabel(pir)

	box.Children = append(box.Children, systemDetailsBox)
}

func (s *system) generateAtmos(sys_size byte) (randomAtmorsphere byte) {
	randomAtmosphere := zero_to_ten() + sys_size - 5
	if randomAtmosphere < 0 {
		randomAtmosphere = 0
	}
	return
}

func zero_to_ten() byte {
	return byte(6*rand.Float32()) + byte(6*rand.Float32())
}

func two_to_twelve() byte {
	return zero_to_ten() + 2
}

func trv_to_int(tint byte) int {
	return int(tint - '0')
}
