package main

import (
	"strconv"
)

var (
	govDetails = []string{
		"None",
		"Company/Corporation",
		"Participating Democracy",
		"Self-Perpetuating Oligarchy",
		"Representative Democracy",
		"Feudal Technocracy",
		"Captive Government",
		"Balkanisation",
		"Civil Service Bureaucracy",
		"Impersonal Bureaucracy",
		"Charismatic Dictator",
		"Non-Charistmatic Dictator",
		"Charismatic Oligarchy",
		"Religious Dictatorship",
	}

	importance = []string{
		" Importance: Obscure",
		" Importance: Fringe",
		" Importance: Minor",
		" Importance: Notable",
		" Importance: Significant",
		" Importance: Overwhelming",
	}
)

func (s *system) getGovernment() {
	s.Government = s.getGov()

	// Factions are an indepenendent variable, and also part
	// of the Government info, so we use a separate int8 for
	// each in the system structure, but combine the details
	// into the single Government text label we display.

	s.NumFactions = zero_to_two() + 1
	if s.Government == 0 || s.Government == 7 {
		s.NumFactions += 1
	} else if s.Government > 9 {
		s.NumFactions -= 1
	}
	if s.NumFactions < 0 {
		s.NumFactions = 0
	}

	govText := gov + " " + govDetails[s.Government]
	if s.NumFactions > 0 {
		s.Factions = make([]int8, s.NumFactions)
		for i := int8(0); i < s.NumFactions; i++ {
			s.Factions[i] = s.getGov()
			if s.Factions[i] == s.Government {
				govText += ";\n Splinter Faction #" +
					strconv.Itoa(int(i+1)) + " " +
					govDetails[s.Factions[i]] + " " +
					getImportance()
			} else {
				govText += ";\n Faction #" + strconv.Itoa(int(i+1)) +
					" " + govDetails[s.Factions[i]] +
					getImportance()
			}
		}
		government.SetText(govText)
	} else {
		s.Factions = make([]int8, 0)
		government.SetText(govText)
	}

}

func (s system) getGov() (gov int8) {
	gov = zero_to_ten() - 5 + s.Population
	if gov < 0 {
		gov = 0
	} else if gov > 13 {
		gov = 13
	}
	return
}

func getImportance() string {
	switch two_to_twelve() {
	case 2, 3:
		return importance[0]
	case 4, 5:
		return importance[1]
	case 6, 7:
		return importance[2]
	case 8, 9:
		return importance[3]
	case 10, 11:
		return importance[4]
	case 12:
		return importance[5]
	default:
		return importance[0]
	}
}
