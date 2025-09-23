package main

var trv_int = []byte("0123456789ABCDEFGHJKLMNPQRSTUVWXYZ")

// trv_int skips I and O, otherwise it is pretty much like Hex except
//  it continues through the rest of the alphabet

const (
	sz    = string("Planet Size ")
	atmos = string("Atmosphere ")
	hydro = string("Hydrology")
	pop   = string("Population ")
	gov   = string("Government ")
	law   = string("Law Level ")
	tech  = string("Technology Level ")
	sp    = string("Starport class ")
	anc   = string("Ancient Artifact")
	mil   = string("Military Base")
	nav   = string("Naval Base")
	merch = string("Merchant Base")
	mega  = string("Megacorp Home or Major Office")
	sct   = string("Scout Base")
	host  = string("Scout Hostel")
	pir   = string("Pirate Base")
	rsrch = string("Research Canter")
	hosp  = string("University Hospital and Medical Center")
	cons  = string("Imperial Consulate")
	uwp   = string("Universal World Profile ")
	uni   = string("University and Research Hospital")
	chap  = string("Traveller's Aid Society Chapter House")
	trv   = string("Traveller's Aid Society Hostel")
	first = string("Traveller's Aid Society First Class Hostel")
	yard  = "Starship Yard"

	starportNone = 'X'
	starportE    = 'E'
	starportD    = 'D'
	starportC    = 'C'
	starportB    = 'B'
	starportA    = 'A'
)
