package main

var (
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
)

func (s system) getBases() {
	numBases := 0
	// Ancient Site
	ancient := false
	if zero_to_ten() == 10 {
		// 1 in 36
		ancient = true
		numBases++
	}

	// Imperial Consulate, Merchant Base & Naval Base all depend on starport
	merchant := false
	yard := false
	megacorp := false
	imperial := false
	naval := false
	switch s.Starport {
	case 5:
		// Merchant Base: Starport A, roll 6+ in 2D6 or 4+ on 2D0-10, above 3
		// Imperial Consulate: Starport A, roll 6+ in 2D6 or 4+ on 2D0-10, "
		roll := zero_to_ten()
		if roll > 3 {
			merchant = true
			numBases++
		}
		if roll > 6 {
			yard = true
			numBases++
		}
		if roll > 9 {
			megacorp = true
			numBases++
		}

		if zero_to_ten() > 3 {
			imperial = true
			numBases++
		}
		if zero_to_ten() > 5 {
			naval = true
			numBases++
		}
	case 4:
		// Starport B, roll 8+ in 2D6 or 6+ on 2D0-10
		roll := zero_to_ten()
		if roll > 5 {
			merchant = true
			numBases++
		}
		if roll > 8 {
			yard = true
			numBases++
		}
		if zero_to_ten() > 5 {
			imperial = true
			numBases++
		}
		if zero_to_ten() > 5 {
			naval = true
			numBases++
		}
	case 3:
		// Starport C, roll 10+ in 2D6 or 8+ on 2D0-10
		if zero_to_ten() > 5 {
			naval = true
			numBases++
		}
		if zero_to_ten() > 7 {
			merchant = true
			numBases++
		}
		if zero_to_ten() > 7 {
			imperial = true
			numBases++
		}
	}

	// Pirate Base: Starport B: Throw 12+, Starport C: Throw 10+, Starport D or E: Throw 12+.
	pirate := false
	switch s.Starport {
	case 4, 2, 1:
		// B, D or E
		if zero_to_ten() == 9 {
			pirate = true
			numBases++
		}
	case 3:
		// C
		if zero_to_ten() > 7 {
			pirate = true
			numBases++
		}
	}

	// Research Base
	research := false
	university := false
	switch s.Starport {
	case 5:
		// A
		roll := zero_to_ten()
		if roll > 5 {
			research = true
			numBases++
		}
		if roll > 8 {
			university = true
			numBases++
		}
	case 4, 3:
		// B or C
		if zero_to_ten() > 8 {
			research = true
			numBases++
		}
	}

	// Scout Base
	// Scout Hostel
	scout := false
	hostel := false
	switch s.Starport {
	case 5:
		// A
		if zero_to_ten() > 7 {
			scout = true
		}
	case 4, 3:
		// B or C
		roll := zero_to_ten()
		if roll > 5 {
			scout = true
		}
		if roll > 8 {
			hostel = true
		}
	}

	travellers := false
	firstclass := false
	chapter := false
	switch s.Starport {
	case 5:
		// A
		roll := zero_to_ten()
		if roll > 4 {
			firstclass = true
		} else if roll > 1 {
			travellers = true
		}
		if roll > 7 {
			chapter = true
		}
	case 4:
		// B
		roll := zero_to_ten()
		if roll > 6 {
			firstclass = true
		} else if roll > 3 {
			travellers = true
		}
	case 3:
		// C
		if zero_to_ten() > 7 {
			travellers = true
		}
	}

	// Build a string with all the bases
	allBases := ""
	// Ancient Artifact?
	if ancient {
		allBases += "Ancient Artifact (random planet)\n"
	}

	if merchant {
		allBases += "Merchant Base\n"
	}
	if yard {
		allBases += "Starship Yard\n"
	}

	if megacorp {
		allBases += "Megacorp Home Office\n"
	}

	if imperial {
		allBases += "Imperial Consulate\n"
	}

	if naval {
		allBases += "Naval Base\n"
	}

	if pirate {
		allBases += "pirate Base\n"
	}

	if research {
		allBases += "Research Base\n"
	}

	if university {
		allBases += "University and Research Hospital\n"
	}

	if scout {
		allBases += "Scout Base\n"
	}

	if hostel {
		allBases += "Scout Hostel\n"
	}

	if travellers {
		allBases += "Traveller's Aid Society Hostel\n"
	}

	if firstclass {
		allBases += "Traveller's Aid Society First Class Hotel\n"
	}

	if chapter {
		allBases += "Traveller's Aid Society Chapter House\n"
	}

	bases.SetText(allBases)
}
