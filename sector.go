package main

import (
	"strconv"
	"strings"

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
	sysPerLine := 0
	numPopMax := 0
	numTechMax := 0
	for row := 0; row < 10; row++ {
		for column := 0; column < 8; column++ {
			if zero_to_one() == 1 {
				systemGrid[column][row].init(&nxtUI)
				b.Children = make([]fyne.CanvasObject, 0)
				systems++

				if sysPerLine%3 == 0 && sysPerLine > 0 {
					s.system_list += "\n"
				}
				sysPerLine++

				s.system_list += strconv.Itoa(column) + strconv.Itoa(row) +
					" " + uw_profile.Text +
					strings.Repeat(" ", 32-len(systemGrid[column][row].Codes))

				nxtPop := systemGrid[column][row].Population
				totalPop += int(systemGrid[column][row].PopulationCount)
				if nxtPop > maxPop {
					numPopMax = 1
					maxPop = nxtPop
					s.population_report = "Maximum population " + strconv.Itoa(int(maxPop)) +
						" in " + strconv.Itoa(column) + strconv.Itoa(row) +
						" " + uw_profile.Text
				} else if nxtPop == maxPop {
					numPopMax++
					if numPopMax&1 == 0 {
						s.population_report += ", "
					}
					s.population_report += strconv.Itoa(column) + strconv.Itoa(row) +
						" " + uw_profile.Text
					if numPopMax&1 == 0 {
						s.population_report += "\n"
					}
				}

				nxtTech := systemGrid[column][row].Technology_Level
				if nxtTech > maxTech {
					numPopMax = 1
					maxTech = nxtTech
					s.technology_report = "Maximum technolgy level " + strconv.Itoa(int(maxTech)) +
						" in " + strconv.Itoa(column) + strconv.Itoa(row) +
						" " + uw_profile.Text
				} else if nxtTech == maxTech {
					numTechMax++
					if numTechMax&1 == 0 {
						s.technology_report += ", "
					}
					s.technology_report += strconv.Itoa(column) + strconv.Itoa(row) +
						" " + uw_profile.Text
					if numTechMax&1 == 0 {
						s.technology_report += "\n"
					}
				}
			}
		}
	}
	s.population_report += "Total sector population " +
		strconv.Itoa(totalPop) + "\n"
	system_list.SetText(s.system_list)
	pop_report.SetText(s.population_report)
	tech_report.SetText(s.technology_report)
	b.Children = append(b.Children, sectorDetailsBox)
}
