package main

type trade struct {
	description string
	code        string
}

var (
	TradeCodes = []trade{
		{"Agricultural", "Ag"},     // : Atmos 4-9, Hydro 4-8, Pop 5-7.
		{"Asteroid Belt", "As"},    //: Asteroid Belt.
		{"Desert", "De"},           //: Atmos 2-D, Hydro 0.
		{"Fluid Oceans", "Fl"},     //: Atmos A+ or non-water biochemistry, Hydro 1-B.
		{"Garden", "Ga"},           //: Size 5-A, Atmos 4-9, Hydro 4-8.
		{"High Population", "Hi"},  //: Pop 9+.
		{"High Technology", "Ht"},  //: Industry", "TL-3"}, //+.
		{"Ice-Capped", "Ic"},       //: Atmos 0-1, Hydro 1+.
		{"Industrial", "In"},       //: Pop 9+, Industry 6+.
		{"Low Population", "Lo"},   //: Pop 1-3.
		{"Low Technology", "Lt"},   //: Industry 5-.
		{"Non-Agricultural", "Na"}, //: Atmos 0-3 or B+, Hydro 0-3 or B+, Pop 6+.
		{"Non-Industrial", "Ni"},   //: Pop 4-6.
		{"Poor", "Po"},             //: Atmos 2-5, Hydro 0-3.
		{"Rich", "Ri"},             //: Atmos 6 or 8, Pop 6-8.
		{"Sterile", "St"},          //: Bio 0.
		{"Water World", "Wa"},      //: Atmos 2+, Hydro A-B.
		{"Vacuum", "Va"},           //: Atmos 0.
	}
)

// Filling out the System structure's Trade Codes, adding the text to
// the label that's in the Box
//
// Also fill-in the ShortCodes. adding them to the UWP profile
func (s *system) getCodes() {
	codes := ""
	compressed := "  "
	// Agricultural: Atmos 4-9, Hydro 4-8, Pop 5-7
	if s.Atmosphere > 3 && s.Atmosphere < 10 &&
		s.Hydrology > 3 && s.Hydrology < 9 &&
		s.Population > 4 && s.Population < 8 {
		compressed += " " + TradeCodes[0].code
		codes += TradeCodes[0].code + " " + TradeCodes[0].description + "\n"
	}
	// Asteroids
	if s.numAsteroids > 0 {
		compressed += " " + TradeCodes[1].code
		codes += TradeCodes[1].code + " " + TradeCodes[1].description + "\n"
	}
	// Desert
	if s.Atmosphere > 2 && s.Atmosphere < 14 &&
		s.Hydrology == 0 {
		compressed += " " + TradeCodes[2].code
		codes += TradeCodes[2].code + " " + TradeCodes[2].description + "\n"
	}
	// Fluid Oceans
	if s.Atmosphere > 9 &&
		s.Hydrology > 0 && s.Hydrology < 12 {
		compressed += " " + TradeCodes[3].code
		codes += TradeCodes[3].code + " " + TradeCodes[3].description + "\n"
	}
	// Garden
	if s.Size > 4 && s.Size < 11 &&
		s.Atmosphere > 3 && s.Atmosphere < 10 &&
		s.Hydrology > 3 && s.Hydrology < 9 {
		compressed += " " + TradeCodes[4].code
		codes += TradeCodes[4].code + " " + TradeCodes[4].description + "\n"
	}
	// High Populatuion
	if s.Population > 8 {
		compressed += " " + TradeCodes[5].code
		codes += TradeCodes[5].code + " " + TradeCodes[5].description + "\n"
	}
	// High Technology
	if s.Technology_Level > 9 {
		compressed += " " + TradeCodes[6].code
		codes += TradeCodes[6].code + " " + TradeCodes[6].description + "\n"
	}
	// Ice Capped
	if s.Atmosphere < 2 && s.Hydrology > 0 {
		compressed += " " + TradeCodes[7].code
		codes += TradeCodes[7].code + " " + TradeCodes[7].description + "\n"
	}
	// Industrial
	if s.Population > 9 && s.Technology_Level > 5 {
		compressed += " " + TradeCodes[8].code
		codes += TradeCodes[8].code + " " + TradeCodes[8].description + "\n"
	}
	// Low Population
	if s.Population < 4 {
		compressed += " " + TradeCodes[9].code
		codes += TradeCodes[9].code + " " + TradeCodes[9].description + "\n"
	}
	// Low Technology
	if s.Technology_Level < 6 {
		compressed += " " + TradeCodes[10].code
		codes += TradeCodes[10].code + " " + TradeCodes[10].description + "\n"
	}
	// Non-Agricultural
	if (s.Atmosphere < 4 || s.Atmosphere > 10) ||
		(s.Hydrology < 4 || s.Hydrology > 10) &&
			s.Population > 5 {
		compressed += " " + TradeCodes[11].code
		codes += TradeCodes[11].code + " " + TradeCodes[11].description + "\n"
	}
	// Non-Industrial
	if s.Population > 3 && s.Population < 6 {
		compressed += " " + TradeCodes[12].code
		codes += TradeCodes[12].code + " " + TradeCodes[12].description + "\n"
	}
	// Poor
	if s.Atmosphere > 1 && s.Atmosphere < 6 &&
		s.Hydrology < 4 {
		compressed += " " + TradeCodes[13].code
		codes += TradeCodes[13].code + " " + TradeCodes[13].description + "\n"
	}
	// Rich
	if (s.Atmosphere == 6 || s.Atmosphere == 8) &&
		s.Population > 5 && s.Population < 9 {
		compressed += " " + TradeCodes[14].code
		codes += TradeCodes[14].code + " " + TradeCodes[14].description + "\n"
	}
	// Sterile
	if s.Size == 0 || s.Atmosphere == 0 {
		compressed += " " + TradeCodes[15].code
		codes += TradeCodes[15].code + " " + TradeCodes[15].description + "\n"
	}
	// Water World
	if s.Atmosphere > 1 &&
		s.Hydrology > 8 && s.Hydrology < 12 {
		compressed += " " + TradeCodes[16].code
		codes += TradeCodes[16].code + " " + TradeCodes[16].description + "\n"
	}
	if s.Atmosphere == 0 {
		compressed += " " + TradeCodes[17].code
		codes += TradeCodes[17].code + " " + TradeCodes[17].description + "\n"
	}
	trade_codes.SetText(codes)
	uw_profile.SetText(uw_profile.Text + compressed)
	if !s.Detailed {
		s.Codes = compressed
	}
}
