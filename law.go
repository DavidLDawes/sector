package main

var lawDetails = []string{
	"No Restrictions",
	"No poison gas, explosives, undetectable weapons, or WMD",
	"No portable energy weapons (except ship-mounted weapons)",
	"No heavy weapons",
	"No light assault weapons or submachine guns",
	"No personal concealable weapons",
	"No firearms except shotguns and stunners; carrying weapons discouraged",
	"Only stunners allowed; carrying weapons discouraged",
	"All bladed weapons, no firearms at all",
	"No weapons of any sort",
}

// Filling out the System structure's Law Level, adding the text to
// the label that's in the Box
func (s *system) getLaw() {
	s.Law_Level = zero_to_ten() - 7 + s.Government
	if s.Law_Level < 0 {
		s.Law_Level = 0
	}
	if int(s.Law_Level) > len(lawDetails)-1 {
		s.Law_Level = int8(len(lawDetails) - 1)
	}

	law_level.SetText(law + " " + string(trv_int[s.Law_Level]) + ", " +
		lawDetails[s.Law_Level])
}
