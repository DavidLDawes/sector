package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type sector struct {
	system_list       string
	population_report string
	technology_report string
}

var (
	system_list = widget.NewLabel("")
	pop_report  = widget.NewLabel("")
	tech_report = widget.NewLabel("")

	sectorDetailsBox = widget.NewVBox(
		system_list,
		pop_report,
		tech_report,
	)
)

func (s sector) init(b *widget.Box) {
	systems := 0
	totalPop := 0
	maxPop := int8(-1)
	s.system_list = ""
	maxTech := int8(-1)
	numMax := 0
	for row := 0; row < 10; row++ {
		for column := 0; column < 8; column++ {
			if zero_to_one() == 1 {
				nxtSys := systemGrid[column][row]
				nxtSys.init(&nxtUI)
				b.Children = make([]fyne.CanvasObject, 0)
				systems++

				s.system_list += strconv.Itoa(column) + strconv.Itoa(row) +
					" " + uw_profile.Text + "\n"

				nxtPop := nxtSys.Population
				totalPop += int(nxtPop)
				if nxtPop > maxPop {
					numMax = 1
					maxPop = nxtPop
					s.population_report = "Maximum population " + strconv.Itoa(int(maxPop)) +
						" in " + strconv.Itoa(column) + strconv.Itoa(row) +
						" " + uw_profile.Text
				} else if nxtPop == maxPop {
					numMax++
					if numMax&1 == 0 {
						s.population_report += ", "
					}
					s.population_report += strconv.Itoa(column) + strconv.Itoa(row) +
						" " + uw_profile.Text
					if numMax&1 == 0 {
						s.population_report += "\n"
					}
				}

				nxtTech := nxtSys.Technology_Level
				if nxtTech > maxTech {
					maxTech = nxtTech
					s.technology_report = "Maximum technolgy level " + strconv.Itoa(int(maxTech)) +
						" in " + strconv.Itoa(column) + strconv.Itoa(row) +
						" " + uw_profile.Text
				}
			}
		}
	}
	system_list.SetText(s.system_list)
	pop_report.SetText(s.population_report)
	tech_report.SetText(s.technology_report)
	b.Children = append(b.Children, sectorDetailsBox)
}
