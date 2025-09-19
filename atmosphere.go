package main

type atmosphereDetail struct {
	description string
	pressure    string
	gear        string
	pMin        float32
	pMax        float32
}

var (
	atmosphereDetails = []atmosphereDetail{
		{"None", "0.00", "Vacc Suit", 0.0, 0.0},
		{"Trace", "0.001 to 0.09", "Vacc Suit", 0.001, 0.09},
		{"Very Thin, Tainted", "0.1 to 0.42", "Respirator, Filter", 0.1, 0.42},
		{"Very Thin", "0.1 to 0.42", "Respirator", 0.1, 0.42},
		{"Thin, Tainted", "0.43 to 0.7", "Respirator, Filter", 0.43, 0.7},
		{"Thin", "0.43 to 0.7", "Respirator", 0.43, 0.7},
		{"Standard", "0.71 to 1.49", "", 0.71, 1.49},
		{"Standard, Tainted", "0.71 to 1.49", "Filter", 0.71, 1.49},
		{"Dense", "1.5 to 2.5", "", 1.5, 2.5},
		{"Dense, Tainted", "1.5 to 2.5", "Filter", 1.5, 2.5},
		{"Exotic", "Varies", "Air Supply", 0.43, 2.5},
		{"Corrosive", "Varies", "Vacc Suit", 0.43, 2.5},
		{"Insidious", "Varies", "Vacc Suit", 0.43, 2.5},
		{"Dense", "High", "Pressure Suit", 2.51, 25},
		{"Thin", "Low", "Vacc Suit", 0.0, 0.5},
		{"Unusual", "Varies", "Varies", 0.1, 25},
	}
)

func (s *system) getAtmosphere() {
	// Continue filling out the System structure, Atmosphere next
	randomAtmosphere := zero_to_ten() + s.Size - 5
	if randomAtmosphere < 0 {
		randomAtmosphere = 0
	}
	if randomAtmosphere > 15 {
		randomAtmosphere = 15
	}
	s.Atmosphere = randomAtmosphere
	atmosphere.SetText(atmos + " " +
		string(trv_int[s.Atmosphere]) + ", " + atmosphereDetails[s.Atmosphere].description +
		", pressure " + atmosphereDetails[s.Atmosphere].pressure +
		", gear required: " + atmosphereDetails[s.Atmosphere].gear,
	)

}
