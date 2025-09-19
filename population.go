package main

import (
	"strconv"
)

func (s *system) getPopulation() {
	// Continue w/System structure, Population next
	s.Population = zero_to_ten()
	popTotal := int(1)
	if s.Population == 0 {
		popTotal = 0
	}
	for i := 0; i < int(s.Population); i++ {
		popTotal = popTotal * 10
	}
	factor := int(1)
	if popTotal > 1000 {
		factor = int(one_to_nine())
	} else {
		// bias tiny settlements to a bit higher numbers
		factor = int(7 + zero_to_two())
	}
	popTotal = popTotal * factor

	population.SetText(pop + " " +
		string(trv_int[s.Population]) + ", " +
		strconv.Itoa(popTotal),
	)
}
