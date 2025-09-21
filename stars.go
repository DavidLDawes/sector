package main

import "strconv"

type starSystem struct {
	description string
	count       int
}

var (
	starCounts = []starSystem{
		{"Solitary star, primary is ", 1},
		{"Binary stars, primary is ", 2},
		{"Trinary stars, primary is ", 3},
	}

	starTypes = []string{
		"Type A ",
		"Type F ",
		"Type G ",
		"Type K ",
		"Type M ",
		"Type L ",
	}

	orbitTypes = []string{
		"Tight orbit",
		"Close orbit",
		"Moderate orbit",
		"Distant orbit",
	}
	planetTypes = []string{
		"Asteroid Belt",
		"Dwarf Planet",
		"Terrestrial Planet",
		"Helian Planet",
		"Jovian Planet",
		"Companion Star",
	}
)

// Filling out the System structure's Size, adding the text to
// the label that's in the Box
func (s *system) getStars() {
	s.Size = int8(zero_to_ten())

	starDetails := s.getStarsDetails()

	stars.SetText(starCounts[s.numStars-1].description +
		starDetails)
}

func (s *system) getStarsDetails() string {
	switch zero_to_fifteen() {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		s.numStars = 1
	case 8, 9, 10, 11, 12:
		s.numStars = 2
	case 13, 14, 15:
		s.numStars = 3
	default:
		s.numStars = 1
	}
	s.starTypes = make([]string, s.numStars)
	return s.getStarTypes() + s.getPlanetDetails()
}

func (s *system) getPlanetDetails() (result string) {
	s.numGasGiants = int8(0)
	if s.starTypes[0] != starTypes[5] {
		s.numGasGiants = zero_to_five() - 2
		if s.numGasGiants < 0 {
			s.numGasGiants = 0
		}
	}

	remainingGasGiants := s.numGasGiants
	countGasGiants := 0
	if s.Close_Companion {
		s.Epistellar_Planets = 0
	} else {
		s.Epistellar_Planets = zero_to_five() - 2
		if s.Epistellar_Planets < 0 {
			s.Epistellar_Planets = 0
		} else if s.Epistellar_Planets > 2 {
			s.Epistellar_Planets = 2
		}
	}

	epiList := ""
	s.EpiPlanets = make([]planet, s.Epistellar_Planets)
	if s.Epistellar_Planets > 0 && remainingGasGiants > 0 {
		if s.Epistellar_Planets > remainingGasGiants {
			for i := 0; i < int(remainingGasGiants); i++ {
				numSat := zero_to_five() + 1
				if zero_to_five() < 5 {
					// All satelites are dwarfs
					if numSat == 1 {
						s.EpiPlanets[i] = planet{planetTypes[4], "Sattelite " +
							planetTypes[1], ""}
						countGasGiants++
						epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
					} else {
						s.EpiPlanets[i] = planet{planetTypes[4], "Satellite " +
							strconv.Itoa(int(numSat)) + "x " + planetTypes[1], ""}
						countGasGiants++
						epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
					}
				} else {
					// One satellite is terrestrial
					if numSat == 1 {
						s.EpiPlanets[i] = planet{planetTypes[4], "Satellite " +
							planetTypes[2], ""}
						countGasGiants++
						epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
					} else if numSat == 2 {
						s.EpiPlanets[i] = planet{planetTypes[4], "Satellites " +
							planetTypes[2] + " & " + planetTypes[1], ""}
						countGasGiants++
						epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
					} else {
						s.EpiPlanets[i] = planet{planetTypes[4], "Sattelites " +
							planetTypes[2] + strconv.Itoa(int(numSat)) +
							"x " + planetTypes[1], ""}
						countGasGiants++
						epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
					}
				}
			}
			for i := remainingGasGiants; i < s.Epistellar_Planets; i++ {
				s.EpiPlanets[i] = s.getPlanet()
				if s.EpiPlanets[i].content[0:3] == planetTypes[4][0:3] {
					// Another gas giant
					countGasGiants++
				}
				epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
			}
			remainingGasGiants = 0
		} else {
			// more gas giants than epistellar planet slots to put them in
			for i := 0; i < int(s.Epistellar_Planets); i++ {
				s.EpiPlanets[i] = planet{planetTypes[4], "", ""}
				countGasGiants++
				epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
			}
			remainingGasGiants -= s.Epistellar_Planets
		}
	} else if s.Epistellar_Planets > 0 {
		for i := 0; i < int(s.Epistellar_Planets); i++ {
			s.EpiPlanets[i] = s.getPlanet()
			if s.EpiPlanets[i].content[0:3] == planetTypes[4][0:3] {
				// Another gas giant
				countGasGiants++
			}
			epiList += s.EpiPlanets[i].content + ", " + s.EpiPlanets[i].satellites + "\n"
		}

	}

	inList := ""
	s.Inner_Planets = zero_to_five()
	if s.starTypes[0] == starTypes[4] {
		// Type M star is -1
		s.Inner_Planets--
	}
	if s.Inner_Planets < 0 {
		s.Inner_Planets = 0
	}

	s.InPlanets = make([]planet, s.Inner_Planets)
	if s.Inner_Planets > 0 && remainingGasGiants > 0 {
		if s.Inner_Planets > remainingGasGiants {
			for i := 0; i < int(remainingGasGiants); i++ {
				s.InPlanets[i] = planet{planetTypes[4], "", ""}
				countGasGiants++
				inList += s.InPlanets[i].content + ", " + s.InPlanets[i].satellites + "\n"
			}
			for i := remainingGasGiants; i < s.Inner_Planets; i++ {
				s.InPlanets[i] = s.getPlanet()
				if s.InPlanets[i].content[0:3] == planetTypes[4][0:3] {
					// Another gas giant
					countGasGiants++
				}
				inList += s.InPlanets[i].content + ", " + s.InPlanets[i].satellites + "\n"
			}
			remainingGasGiants = 0
		} else {
			// more gas giants than inner planet slots to put them in
			for i := 0; i < int(s.Inner_Planets); i++ {
				s.InPlanets[i] = planet{planetTypes[4], "", ""}
				inList += s.InPlanets[i].content + ", " + s.InPlanets[i].satellites + "\n"
			}
			remainingGasGiants -= s.Inner_Planets
		}
	} else if s.Inner_Planets > 0 {
		for i := 0; i < int(s.Inner_Planets); i++ {
			s.InPlanets[i] = s.getPlanet()
			if s.InPlanets[i].content[0:3] == planetTypes[4][0:3] {
				// Another gas giant
				countGasGiants++
			}
			inList += s.InPlanets[i].content + ", " + s.InPlanets[i].satellites + "\n"
		}
	}

	outList := ""
	s.Outer_Planets = zero_to_five()
	if s.starTypes[0] == starTypes[4] {
		// Type M star is -1
		s.Outer_Planets--
	}
	if s.Moderate_Companion {
		// Moderate companion star clears out outer planets some...
		s.Outer_Planets--
	}
	if s.Outer_Planets < 0 {
		s.Outer_Planets = 0
	}

	s.OutPlanets = make([]planet, s.Outer_Planets)
	if s.Outer_Planets > 0 && remainingGasGiants > 0 {
		if s.Outer_Planets > remainingGasGiants {
			for i := 0; i < int(remainingGasGiants); i++ {
				s.OutPlanets[i] = planet{planetTypes[4], "", ""}
				countGasGiants++
				outList += s.OutPlanets[i].content + ", " + s.OutPlanets[i].satellites + "\n"
			}
			for i := remainingGasGiants; i < s.Outer_Planets; i++ {
				s.OutPlanets[i] = s.getPlanet()
				if s.OutPlanets[i].content[0:3] == planetTypes[4][0:3] {
					// Another gas giant
					countGasGiants++
				}
				outList += s.OutPlanets[i].content + ", " + s.OutPlanets[i].satellites + "\n"
			}
			remainingGasGiants = 0
		} else {
			// more gas giants than outer planet slots to put them in
			for i := 0; i < int(s.Outer_Planets); i++ {
				s.OutPlanets[i] = planet{planetTypes[4], "", ""}
				countGasGiants++
				outList += s.OutPlanets[i].content + ", " + s.OutPlanets[i].satellites + "\n"
			}
			remainingGasGiants -= s.Outer_Planets
		}
	} else if s.Outer_Planets > 0 {
		for i := 0; i < int(s.Outer_Planets); i++ {
			s.OutPlanets[i] = s.getPlanet()
			outList += s.OutPlanets[i].content + ", " + s.OutPlanets[i].satellites + "\n"
		}
	}

	result = ""
	if s.Epistellar_Planets > 0 {
		result = "\nEpistellar planets: " + strconv.Itoa(int(s.Epistellar_Planets)) +
			"\n" + epiList
	}
	if s.Inner_Planets > 0 {
		result += "\nInner planets: " + strconv.Itoa(int(s.Inner_Planets)) +
			"\n" + inList
	}
	if s.Outer_Planets > 0 {
		result += "\nOuter planets: " + strconv.Itoa(int(s.Outer_Planets)) +
			"\n" + outList
	}
	result += "\nGas Giants: " + strconv.Itoa(int(countGasGiants))
	return
}

func (s *system) getStarTypes() (result string) {
	switch two_to_twelve() {
	case 2:
		s.starTypes[0] = starTypes[0]
	case 3:
		s.starTypes[0] = starTypes[1]
	case 4:
		s.starTypes[0] = starTypes[2]
	case 5:
		s.starTypes[0] = starTypes[3]
	case 6, 7, 8, 9, 10, 11, 12, 13:
		s.starTypes[0] = starTypes[4]
	default:
		s.starTypes[0] = starTypes[3]
	}
	result = s.starTypes[0]
	if s.numStars > 1 {
		for i := 1; i < int(s.numStars); i++ {
			nextOrbit := s.getCompanionStarOrbit()
			switch zero_to_five() {
			case 0:
				s.starTypes[i] = starTypes[0]
			case 1:
				s.starTypes[i] = starTypes[1]
			case 2:
				s.starTypes[i] = starTypes[2]
			case 3:
				s.starTypes[i] = starTypes[3]
			case 4, 5, 6:
				s.starTypes[i] = starTypes[4]
			default:
				s.starTypes[i] = starTypes[3]
			}
			result += ", secondary " + s.starTypes[i] + " in " + nextOrbit
		}
		result += "\n"

	}
	return
}

func (s *system) getCompanionStarOrbit() string {
	switch zero_to_five() {
	case 0, 1:
		s.Close_Companion = true
		return orbitTypes[0]
	case 2, 3:
		s.Close_Companion = true
		return orbitTypes[1]
	case 4:
		s.Moderate_Companion = true
		return orbitTypes[2]
	case 5:
		s.Distant_Companion = true
		return orbitTypes[3]
	default:
		s.Close_Companion = true
		return orbitTypes[0]
	}
}

func (s *system) getPlanet() planet {
	delta := int8(0)
	if s.starTypes[0] == starTypes[5] {
		delta = -1
	}
	switch zero_to_five() + delta {
	case -1, 0:
		if zero_to_five() > 3 {
			return planet{planetTypes[0], "Mostly Dwarf planetodis, one " + planetTypes[1], ""}
		} else {
			return planet{planetTypes[0], "", ""}
		}
	case 1:
		if zero_to_five() > 4 {
			return planet{planetTypes[1], "Satellite " + planetTypes[1], ""}
		} else {
			return planet{planetTypes[1], "", ""}
		}
	case 2:
		if zero_to_five() > 3 {
			return planet{planetTypes[2], "Satellite " + planetTypes[1], ""}
		} else {
			return planet{planetTypes[2], "", ""}
		}
	case 3:
		numSat := zero_to_five() - 2
		if numSat < 0 {
			numSat = 0
		}

		if numSat > 0 {
			if zero_to_five() < 5 {
				// All satelites are dwarfs
				if numSat == 1 {
					return planet{planetTypes[3], "Sattelite " +
						planetTypes[1], ""}
				} else {
					return planet{planetTypes[3], "Satellite " +
						strconv.Itoa(int(numSat)) + "x " + planetTypes[1], ""}
				}
			} else {
				// One satellite is terrestrial
				if numSat == 1 {
					return planet{planetTypes[3], "Satellite " +
						planetTypes[2], ""}
				} else if numSat == 2 {
					return planet{planetTypes[3], "Satellite " +
						planetTypes[2] + " & " + planetTypes[1], ""}
				} else {
					return planet{planetTypes[3], "Sattelite " +
						planetTypes[2] + strconv.Itoa(int(numSat)) +
						" " + planetTypes[1], ""}
				}
			}
		} else {
			return planet{planetTypes[3], "", ""}
		}
	case 4, 5, 6:
		numSat := zero_to_five() + 1
		if zero_to_five() < 5 {
			// All satelites are dwarfs
			if numSat == 1 {
				return planet{planetTypes[4], "Sattelite " +
					planetTypes[1], ""}
			} else {
				return planet{planetTypes[4], "Satellite " +
					strconv.Itoa(int(numSat)) + "x " + planetTypes[1], ""}
			}
		} else {
			// One satellite is terrestrial
			if numSat == 1 {
				return planet{planetTypes[4], "Satellite " +
					planetTypes[2], ""}
			} else if numSat == 2 {
				return planet{planetTypes[4], "Satellites " +
					planetTypes[2] + " & " + planetTypes[1], ""}
			} else {
				return planet{planetTypes[4], "Sattelites " +
					planetTypes[2] + strconv.Itoa(int(numSat)) +
					"x " + planetTypes[1], ""}
			}
		}
	default:
		if zero_to_five() > 3 {
			return planet{planetTypes[2], "Satellite " + planetTypes[1], ""}
		} else {
			return planet{planetTypes[2], "", ""}
		}
	}
}
