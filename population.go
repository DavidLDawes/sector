package main

import (
	"strconv"
)

func (s *system) getPopulation() {
	if !s.Detailed {
		// Continue w/System structure, Population next
		s.Population = zero_to_ten()
		if s.Population == 0 {
			s.PopulationCount = 0
		} else {
			s.PopulationCount = 1
			for i := 0; i < int(s.Population); i++ {
				s.PopulationCount = s.PopulationCount * 10
			}
			if s.PopulationCount > 1000 {

				s.PopulationCount = s.PopulationCount * int(one_to_nine())
			} else {
				// bias tiny settlements to a the higher end of the range
				s.PopulationCount = s.PopulationCount * int((7 + zero_to_two()))
			}
		}
	}

	population.SetText(pop + " " +
		string(trv_int[s.Population]) + ", " +
		strconv.Itoa(s.PopulationCount),
	)
}
