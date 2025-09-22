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
