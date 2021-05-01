package controllers

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/clients"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/domain"
)

func GetCountriesData(countries *[]domain.Country, table *widgets.Table) {
	err := clients.GetCountries(countries)
	if err != nil {
		log.Fatal(err)
	}

	PrintCountryTable(countries, table)
}

func PrintCountryTable(countries *[]domain.Country, table *widgets.Table) {
	tableHeader := []string{"#", "Country", "Total Cases", "Cases (today)", "Total Deaths", "Deaths (today)", "Recoveries", "Active", "Critical", "Mortality"}
	table.Rows = [][]string{tableHeader}

	for i, v := range *countries {
		table.Rows = append(table.Rows, []string{
			fmt.Sprintf("%d", i+1),
			v.Country,
			fmt.Sprintf("%d", v.Cases),
			fmt.Sprintf("%d", v.TodayCases),
			fmt.Sprintf("%d", v.Deaths),
			fmt.Sprintf("%d", v.TodayDeaths),
			fmt.Sprintf("%d", v.Recovered),
			fmt.Sprintf("%d", v.Active),
			fmt.Sprintf("%d", v.Critical),
			fmt.Sprintf("%.2f%s", float64(v.Deaths)/float64(v.Cases)*100, "%"),
		})
	}

	table.ColumnWidths = []int{5, 22, 20, 20, 18, 18, 15, 15, 15, 15}
	table.TextAlignment = ui.AlignCenter
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.FillRow = true
	table.RowSeparator = false
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table.BorderLeft = false
	table.BorderRight = false
}
