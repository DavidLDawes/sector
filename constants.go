package main

var trv_int = []byte("0123456789ABCDEFGHJKLMNPQRSTUVWXYZ")

// trv_int skips I and O, otherwise it is pretty much like Hex except
//  it continues through the rest of the alphabet

const (
	sz    = "Planet Size "
	atmos = "Atmosphere "
	hydro = "Hydrology"
	pop   = "Population "
	gov   = "Government "
	law   = "Law Level "
	tech  = "Technology Level "
	sp    = "Starport class "
	anc   = "Ancient Artifact"
	mil   = "Military Base"
	scout = "Scout Base"
	res   = "Research Base"
	nav   = "Naval Base"
	merch = "Merchant Base"
	mega  = "Megacorp Home or Major Office"
	sct   = "Scout Base"
	host  = "Scout Hostel"
	pir   = "Pirate Base"
	rsrch = "Research Canter"
	hosp  = "University Hospital and Medical Center"
	cons  = "Imperial Consulate"
	uwp   = "UWP "
	uni   = "University and Research Hospital"
	chap  = "Traveller's Aid Society Chapter House"
	trv   = "Traveller's Aid Society Hostel"
	first = "Traveller's Aid Society First Class Hostel"
	yard  = "Starship Yard"

	starportNone = 'X'
	starportE    = 'E'
	starportD    = 'D'
	starportC    = 'C'
	starportB    = 'B'
	starportA    = 'A'
)
