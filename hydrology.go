package main

type hydrologyDetail struct {
	hydroRange  string
	hydroMax    uint16
	description string
}

var (
	hydroDetails = []hydrologyDetail{
		{"0% - 5%", 5, "Desert World"},
		{"6% - 15%", 15, "Dry World"},
		{"16% - 25%", 25, "A few small seas"},
		{"26% - 35%", 35, "Small to medium seas"},
		{"36% - 45%", 45, "Wet World"},
		{"46% - 55%", 55, "Large Oceans"},
		{"56% - 65%", 65, "Small seas and oceans."},
		{"66% - 75%", 75, "Earth-like world"},
		{"76% - 85%", 85, "Water World"},
		{"86% - 95%", 95, "A few small islands and archipelagos"},
		{"96% - 100%", 100, "Almost Entirely Water"},
	}
)

func (s *system) getHydrology() {
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
}
