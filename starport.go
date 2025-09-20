package main

type starportDetail struct {
	description string
	code        int8
}

var (
	starPortDetails = []starportDetail{
		{"No Starport. No provision is made for any ship landings. ", starportNone},
		{"Frontier Installation. Essentially a marked spot of bedrock with no fuel, facilities, or bases present.",
			starportE},
		{"Poor Quality. Only unrefined fuel available. No repair facilities present. Scout base may be present.", starportD},
		{"Routine Quality. Only unrefined fuel available. Reasonable repair facilities present. Scout base may be present.", starportC},
		{"Good Quality. Refined fuel available. Annual maintenance overhaul available. Shipyard capable of constructing non-starships present. Naval base and/or scout base may be present.", starportB},
		{"Excellent Quality. Refined fuel available. Annual maintenance overhaul available. Shipyard capable of constructing starships and non-starships present. Naval base and/or scout base may be present.", starportA},
	}
)

func (s *system) getStarport() {
	// Continue filling out the System structure, Atmosphere next
	switch zero_to_ten() {
	case 0:
		s.Starport = 0
	case 1, 2:
		s.Starport = 1
	case 3, 4:
		s.Starport = 2
	case 5, 6:
		s.Starport = 3
	case 7, 8:
		s.Starport = 4
	case 9, 10, 11:
		s.Starport = 5
	default:
		s.Starport = 3
	}

	starport.SetText(sp + string(starPortDetails[s.Starport].code) +
		", " + starPortDetails[s.Starport].description,
	)

}
