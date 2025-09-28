package main

type techDetail struct {
	description string
	details     string
}

var (
	techDetails = []techDetail{
		{"Primitive", "No technology "},
		{"Primitive", "Roughly on a par with Bronze or Iron age technology."},
		{"Primitive", "Renaissance technology."},
		{"Primitive", "The germ of industrial revolution and steam power."},
		{"Industrial", "The transition to industrial revolution is complete, bringing plastics, radio\nand other such inventions."},
		{"Industrial", "Widespread electrification, telecommunications and internal combustion."},
		{"Industrial", "Fission power and more advanced computing."},
		{"Pre-Stellar", "A pre-stellar society can reach orbit reliably and has telecommunications\nsatellites."},
		{"Pre-Stellar", "At TL 8, it is possible to reach other worlds in the same system, although\nterraforming or full colonisation are not within reach."},
		{"Pre-Stellar", "Gravity manipulation, which makes space travel vastly safer and faster."},
		{"Early Stellar", "With the advent of Jump, nearby systems are opened up."},
		{"Early Stellar", "The first true artificial intelligences become possible, as computers are\nable to model synaptic networks."},
		{"Average Stellar", "Weather control revolutionises terraforming and agriculture."},
		{"Average Stellar", "The battle dress appears on the battlefield in response to the new\nweapons."},
		{"Average Stellar", "Fusion weapons become man-portable."},
		{"High Stellar", "Black globe generators suggest a new direction for defensive technologies,\nwhile the development of synthetic anagathics means that the human lifespan is\nnow vastly increased. Higher Technology Levels exist and\nmay appear in other settings or be discovered by pioneering scientists."},
		{"High Stellar", "Antimatter engines allow much longer jumps (J-8) with reasonable cargo,\ncounter weapons and shields improve, powered/active armor increases durability,\ncommunications equivalent to personal telepathy and good\nshielding"},
		{"Low Sector", "All engines are compact, take fewer engineers, and perform extremely well,\nup to J-12 and M-16, P-18, human multiplicity physically and virtually is\ncommon, super-intelligent AI, low berths are 0 risk and cure\nalmost all cancers and diseases, forms are mutable."},
		{"Average Sector", "Speculative, J-16, M-20, P-24, extremely good armor, black globe with\ncomplete variability control and huge storage, cvapital ships can use black\nglobes to power outrageous arsenals."},
		{"High Sector", "J-20, J-30 when not near massive objects, J-40 with obscure psionics\nintegrated into production and operation of engines and ships. Resurections,\nbackups, robo/physical/virtual multi-embodiement with merged\nconsciousness, meta-human groupings with others for even larger merged consciousnesses.\nImmunity to pretty much every form of death."},
	}
)

func (s system) getTechLevel() int8 {
	if !s.Detailed {
		s.Technology_Level = s.getTech()
	}
	technology_level.SetText(tech + string(trv_int[s.Technology_Level]) +
		", " + techDetails[s.Technology_Level].description + "\n" +
		techDetails[s.Technology_Level].details,
	)

	s.UWP = string(starPortDetails[s.Starport].code) + string(trv_int[s.Size]) +
		string(trv_int[s.Size]) +
		string(trv_int[s.Atmosphere]) +
		string(trv_int[s.Hydrology]) +
		string(trv_int[s.Population]) +
		string(trv_int[s.Government]) +
		string(trv_int[s.Law_Level]) + "-" +
		string(trv_int[s.Technology_Level])
	uw_profile.SetText(uwp + s.UWP)

	return s.Technology_Level
}

// Because of all the dependencies, getTechLevel is the last of the top level get*
// functions to be called from system.go
func (s system) getTech() (tech int8) {
	tech = zero_to_five() + 1 // in other words a 1D6
	switch s.Starport {
	case 0:
		tech -= 4
	case 3:
		tech += 2
	case 4:
		tech += 4
	case 5:
		tech += 6
	}

	if s.Size < 2 {
		tech += 2
	} else if s.Size < 5 {
		tech += 1
	}

	if s.Atmosphere < 4 {
		tech += 1
	} else if s.Atmosphere > 9 {
		tech += 1
	}

	if s.Hydrology == 0 || s.Hydrology == 9 {
		tech++
	} else if s.Hydrology == 10 {
		tech += 2
	}

	if s.Population < 6 || s.Population == 9 {
		tech++
	} else if s.Population == 10 {
		tech += 2
	} else if s.Population == 11 {
		tech += 3
	} else if s.Population == 12 {
		tech += 4
	} else if s.Population == 13 {
		tech += 5
	} else if s.Population == 14 {
		tech += 6
	}

	if s.Government == 0 || s.Government == 5 {
		tech++
	} else if s.Government == 7 {
		tech += 2
	} else if s.Government == 13 || s.Government == 14 {
		tech -= 2
	}

	if tech < 0 {
		tech = 0
	}

	return
}
