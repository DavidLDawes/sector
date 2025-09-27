package main

var (
	Scout = starBase{
		sct,
		" Offers customized docking, pickup for X-Boats, refined\nfuel and supplies to scout ships and couriers.\n",
	}
	Research = starBase{
		rsrch,
		" Large organization dedicated to a particular field\nof research.\n",
	}
	Consulate = starBase{
		cons,
		" Administration office for various departments such as\ncommerce, justice and foreign affairs. Characters wishing to report significant\ncrimes or obtain various permits will need to visit a consulate.\n",
	}
	Hidden = starBase{
		hid,
		" The presence of a pirate base hidden in a system indicates\nthat a group of thieves is active in the area. Pirates will be operating out of\na remote base away from the starport itself; agents working for the pirates and\nsympathizers (their fences and suppliers too) at the starpot and they will be\non the lookout for likely prey.\n",
	}
	Pirate = starBase{
		pir,
		" The presence of a pirate base in a system indicates that\na group of thieves who raid shipping is active in the area. On this Law Level 0\nworld Pirates may operate directly out of the starport itself, with agents at\nthe port selling loot and ransoming prisoners.\n",
	}
	Ancient = starBase{
		anc,
		" Archeological treasure troves from a long vanished \nrace/society with occasionally very high tech items on a random planet or moon.\n",
	}
	Merchant = starBase{
		merch,
		" Trading and import/export infrastructure with warehouses,\nmarkets, cargo docks, and regular starship visits.\n",
	}
	Yard = starBase{
		yard,
		" Ship manufacturing, refit, repair, and servicing facilities,\nusually in orbit.\n",
	}

	Megacorp = starBase{
		mega,
		" Huge corporate entity with vast resources and capital,\ndrives tech improvement and industries with plenty of ancilliary business\nopportunities.\n",
	}

	Naval = starBase{
		nav,
		" Naval warships and service vessels are stationed or based here,\nalong with the facilities, administrative and personnel resources these fleets\nrequire. Plenty of business and indusrty driven by the Navy.\n",
	}

	University = starBase{
		uni,
		" Medical Center and Research Hospital, drives medical and\nbiological advances, jobs, and commerce.\n",
	}

	Hostel = starBase{
		host,
		" Hostel for Scout personnel, mostly pilots waiting to\nrelieve X-Boat pilots.\n",
	}

	Travellers = starBase{
		trv,
		" Low end facility, reasonable rooms well secured,\nnot fancy.\n",
	}

	Chapter = starBase{
		chap,
		" Full time staffed office, legal, cultural, commercial and\npolitical contacts available, reasonable budgets and connections.\n",
	}

	Firstclass = starBase{
		first,
		" All the comforts, features and luxuries you could ask\nfor (if you can afford them). Free room and board for TAS members.\n",
	}
)

func (s system) getBases() {
	if !s.Detailed {
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
		consulate := false
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
				consulate = true
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
				consulate = true
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
				consulate = true
				numBases++
			}
		}

		// Pirate Base: Starport B: Throw 12+, Starport C: Throw 10+, Starport D or E: Throw 12+.
		pirate := false
		hidden := false
		switch s.Starport {
		case 4, 2, 1:
			// B, D or E
			if zero_to_ten() == 9 {

				if s.Law_Level > 0 {
					hidden = true
				} else {
					pirate = true
				}
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
				numBases++
			}
		case 4, 3:
			// B or C
			roll := zero_to_ten()
			if roll > 5 {
				scout = true
				numBases++
			}
			if roll > 8 {
				hostel = true
				numBases++
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
				numBases++
			} else if roll > 1 {
				travellers = true
				numBases++
			}
			if roll > 7 {
				chapter = true
				numBases++
			}
		case 4:
			// B
			roll := zero_to_ten()
			if roll > 6 {
				firstclass = true
				numBases++
			} else if roll > 3 {
				travellers = true
				numBases++
			}
		case 3:
			// C
			if zero_to_ten() > 7 {
				travellers = true
				numBases++
			}

		}

		s.Bases = make([]starBase, numBases)
		i := 0
		// Build a string with all the Bases
		// Ancient Artifact?
		if ancient {
			s.Bases[i] = Ancient
			i++
		}

		if merchant {
			s.Bases[i] = Merchant
			i++
		}
		if yard {
			s.Bases[i] = Yard
			i++

		}

		if megacorp {
			s.Bases[i] = Megacorp
			i++
		}

		if consulate {
			s.Bases[i] = Consulate
			i++

		}

		if naval {
			s.Bases[i] = Naval
			i++

		}

		if hidden {
			s.Bases[i] = Hidden
			i++
		}

		if pirate {
			s.Bases[i] = Pirate
			i++
		}

		if research {
			s.Bases[i] = Research
			i++
		}

		if university {
			s.Bases[i] = University
			i++
		}

		if scout {
			s.Bases[i] = Scout
			i++
		}

		if hostel {
			s.Bases[i] = Hostel
			i++
		}

		if travellers {
			s.Bases[i] = Travellers
			i++
		}

		if firstclass {
			s.Bases[i] = Firstclass
			i++
		}

		if chapter {
			s.Bases[i] = Chapter
			i++
		}
	}

	allBases := ""
	for i := 0; i < len(s.Bases); i++ {
		switch s.Bases[i].base {
		case sct:
			allBases += Scout.base + " " + Scout.description
		case rsrch:
			allBases += Research.base + " " + Research.description
		case cons:
			allBases += Consulate.base + " " + Consulate.description
		case pir:
			allBases += Pirate.base + " " + Pirate.description
		case hid:
			allBases += Hidden.base + " " + Hidden.description
		case anc:
			allBases += Ancient.base + " " + Ancient.description
		case merch:
			allBases += Merchant.base + " " + Merchant.description
		case yard:
			allBases += Yard.base + " " + Yard.description
		case mega:
			allBases += Megacorp.base + " " + Megacorp.description
		case nav:
			allBases += Naval.base + " " + Naval.description
		case uni:
			allBases += University.base + " " + University.description
		case host:
			allBases += Hostel.base + " " + Hostel.description
		case trv:
			allBases += Travellers.base + " " + Travellers.description
		case chap:
			allBases += Chapter.base + " " + Chapter.description
		case first:
			allBases += Firstclass.base + " " + Firstclass.description
		default:
		}
	}
	Bases.SetText(allBases)
}
