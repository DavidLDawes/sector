package main

type sizeDetail struct {
	size    string
	gravity string
	g       float32
}

type starSystem struct {
	description string
	count       int
}

var (
	sizeDetails = []sizeDetail{
		{"800 km", "Negligible", 0.0},
		{"1,600 km", "0.05G", 0.05},
		{"3,200 km", "0.15G", 0.15},
		{"4,800 km", "0.25G", 0.25},
		{"6,400 km", "0.35G", 0.35},
		{"8,000 km", "0.45G", 0.45},
		{"9,600 km", "0.7G", 0.7},
		{"11,200 km", "0.9G", 0.9},
		{"12,800 km", "1.0G", 1.0},
		{"14,400 km", "1.25G", 1.25},
		{"16,000 km", "1.4G", 1.4},
	}

	starCounts = []starSystem{
		{"Solitary star, primary is ", 1},
		{"Binary stars, primary is ", 2},
		{"Trinary stars, primary is ", 3},
	}
	starTypes = []string{
		"Type A",
		"Type F",
		"Type G",
		"Type K",
		"Type M",
		"Type L",
	}
)

// Filling out the System structure's Size, adding the text to
// the label that's in the Box
func (s *system) getSize() {
	s.Size = int8(zero_to_ten())
	s.getStarsDetails()

	size.SetText(starCounts[s.numStars].description +
		sz + " " + string(trv_int[s.Size]) + ", " +
		sizeDetails[s.Size].size + ", gravity " +
		sizeDetails[s.Size].gravity)
}

func (s *system) getStarsDetails() {
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
	for i := int8(0); i < s.numStars; i++ {

	}
}

func (s *system) getStarTypes() {
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
	if s.numStars > 1 {
		for i := 1; i < int(s.numStars); i++ {
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
		}
	}
}
