package main

type sizeDetail struct {
	size    string
	gravity string
	g       float32
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
)

// Filling out the System structure's Size, adding the text to
// the label that's in the Box
func (s *system) getSize() {
	s.Size = int8(zero_to_ten())

	size.SetText(sz + " " + string(trv_int[s.Size]) + ", " +
		sizeDetails[s.Size].size + ", gravity " +
		sizeDetails[s.Size].gravity)
}
