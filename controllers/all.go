package controllers

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/clients"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/domain"
)

func GetAllData(all *domain.All, global *widgets.Paragraph) {
	err := clients.GetAll(all)
	if err != nil {
		log.Fatal(err)
	}
	PrintAllParagraph(all, global)
}

func PrintAllParagraph(all *domain.All, global *widgets.Paragraph) {
	global.Title = "🌐 Global statistics"
	global.Text = fmt.Sprintf("[Infections](fg:blue): %d (%d today)\n", all.Cases, all.TodayCases)
	global.Text += fmt.Sprintf("[Deaths](fg:red): %d (%d today)\n", all.Deaths, all.TodayDeaths)
	global.Text += fmt.Sprintf("[Recoveries](fg:green): %d (%d remaining)\n", all.Recovered, all.Active)
	if all.Critical > 0 {
		global.Text += fmt.Sprintf("[Critical](fg:yellow): %d (%.2f%% of cases)\n", all.Critical, float64(all.Critical)/float64(all.Cases)*100)
	}
	global.Text += fmt.Sprintf("[Mortality rate (IFR)](fg:cyan): %.2f%%\n", float64(all.Deaths)/float64(all.Cases)*100)
	global.Text += fmt.Sprintf("[Mortality rate (CFR)](fg:cyan): %.2f%%\n", float64(all.Deaths)/(float64(all.Recovered)+float64(all.Deaths))*100)
	if all.AffectedCountries > 0 {
		global.Text += fmt.Sprintf("[Affected Countries](fg:magenta): %d\n", all.AffectedCountries)
	}
	global.SetRect(0, 0, 50, 10)
	global.BorderStyle.Fg = ui.ColorYellow
	global.TitleStyle = ui.NewStyle(ui.ColorClear)
	global.TextStyle = ui.NewStyle(ui.ColorClear)
}
