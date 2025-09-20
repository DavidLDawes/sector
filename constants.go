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
	mil   = string("Military Base")
	scout = string("Scout Base")
	res   = string("Research Base")
	pir   = string("Pirate Base")
	uwp   = string("Universal World Profile ")

	starportNone = 'X'
	starportE    = 'E'
	starportD    = 'D'
	starportC    = 'C'
	starportB    = 'B'
	starportA    = 'A'
)
