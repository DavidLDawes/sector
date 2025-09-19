package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/widget"
)

type starBase struct {
	bType        string
	bDescription string
}

type system struct {
	Size             int8
	Atmosphere       int8
	Hydrology        int8
	Population       int8
	Government       int8
	Law_Level        int8
	Technology_Level int8
	Starport         int8
	bases            []starBase
}

type sizeDetail struct {
	size    string
	gravity string
	g       float32
}

type atmosphereDetail struct {
	description string
	pressure    string
	gear        string
	pMin        float32
	pMax        float32
}

type hydrologyDetail struct {
	hydroRange  string
	hydroMax    uint16
	description string
}

var (
	sizeDetails = []sizeDetail{
		{"800 km", "Negligible", 0.0},
		{"1,600 km", "0.05G", 0.05},
		{"3,200 km", "0.15G", 0.15},
		{"4,800 km", "0.25G", 0.25},
		{"6,400 km", "0.35G", 0.35},
		{"8,000 km", "0.45G", 0.45},
		{"9,600 km", "0.7G", 0.7},
		{"11,200 km", "0.9G", 0.9},
		{"12,800 km", "1.0G", 1.0},
		{"14,400 km", "1.25G", 1.25},
		{"16,000 km", "1.4G", 1.4},
	}

	atmosphereDetails = []atmosphereDetail{
		{"None", "0.00", "Vacc Suit", 0.0, 0.0},
		{"Trace", "0.001 to 0.09", "Vacc Suit", 0.001, 0.09},
		{"Very Thin, Tainted", "0.1 to 0.42", "Respirator, Filter", 0.1, 0.42},
		{"Very Thin", "0.1 to 0.42", "Respirator", 0.1, 0.42},
		{"Thin, Tainted", "0.43 to 0.7", "Respirator, Filter", 0.43, 0.7},
		{"Thin", "0.43 to 0.7", "Respirator", 0.43, 0.7},
		{"Standard", "0.71 to 1.49", "", 0.71, 1.49},
		{"Standard, Tainted", "0.71 to 1.49", "Filter", 0.71, 1.49},
		{"Dense", "1.5 to 2.5", "", 1.5, 2.5},
		{"Dense, Tainted", "1.5 to 2.5", "Filter", 1.5, 2.5},
		{"Exotic", "Varies", "Air Supply", 0.43, 2.5},
		{"Corrosive", "Varies", "Vacc Suit", 0.43, 2.5},
		{"Insidious", "Varies", "Vacc Suit", 0.43, 2.5},
		{"Dense", "High", "Pressure Suit", 2.51, 25},
		{"Thin", "Low", "Vacc Suit", 0.0, 0.5},
		{"Unusual", "Varies", "Varies", 0.1, 25},
	}

	hydroDetails = []hydrologyDetail{
		{"0% - 5%", 5, "Desert World"},
		{"6% - 15%", 15, "Dry World"},
		{"16% - 25%", 25, "A few small seas"},
		{"26% - 35%", 35, ""},
		{"36% - 45%", 45, "Wet World"},
		{"46% - 55%", 55, "Large Oceans"},
		{"56% - 65%", 65, "Small seas and oceans."},
		{"66% - 75%", 75, "Earth-like world"},
		{"76% - 85%", 85, "Water World"},
		{"86% - 95%", 95, "A few small islands and archipelagos"},
		{"96% - 100%", 100, "Almost Entirely Water"},
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
)

func (s *system) init(box *widget.Box) {

	// Filling out the System structure, Size first
	s.Size = int8(zero_to_ten())
	fmt.Print(s.Size)
	size.SetText(sz + " " + string(trv_int[s.Size]) + ", " +
		sizeDetails[s.Size].size + ", gravity " + sizeDetails[s.Size].gravity)

	// Continue filling out the System structure, Atmosphere next
	randomAtmosphere := zero_to_ten() + s.Size - 5
	if randomAtmosphere < 0 {
		randomAtmosphere = 0
	}
	if randomAtmosphere > 15 {
		randomAtmosphere = 15
	}
	s.Atmosphere = randomAtmosphere
	atmosphere.SetText(atmos + " " +
		string(trv_int[s.Atmosphere]) + ", " + atmosphereDetails[s.Atmosphere].description +
		", pressure " + atmosphereDetails[s.Atmosphere].pressure +
		", gear required: " + atmosphereDetails[s.Atmosphere].gear,
	)
	// Continue w/System structure, Hydrology next
	if s.Size < 2 {
		s.Hydrology = 0
	} else {
		s.Hydrology = zero_to_ten() - 5 + s.Size
		switch s.Atmosphere {
		case 0, 1, 10, 11, 12:
			s.Hydrology = s.Hydrology - 4
		}
		if s.Hydrology < 0 {
			s.Hydrology = 0
		}
		if s.Hydrology > 10 {
			s.Hydrology = 10
		}
	}
	hydrology.SetText(hydro + " " +
		string(trv_int[s.Hydrology]) + ", " +
		hydroDetails[s.Hydrology].description + ", range of " +
		hydroDetails[s.Hydrology].hydroRange,
	)
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

func zero_to_ten() int8 {
	return int8(6*rand.Float32()) + int8(6*rand.Float32())
}

func two_to_twelve() int8 {
	return zero_to_ten() + 2
}
