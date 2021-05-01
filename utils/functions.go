package utils

import (
	"reflect"
	"sort"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/controllers"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/domain"
)

func CreateInitialGrid() *ui.Grid {
	w, h := ui.TerminalDimensions()

	grid := ui.NewGrid()
	grid.SetRect(0, 0, w, h)

	return grid
}

func CreateInitialWidgets() (*widgets.Paragraph, *widgets.Table, *widgets.Paragraph) {
	global := widgets.NewParagraph()
	table := widgets.NewTable()
	instructions := widgets.NewParagraph()

	return global, table, instructions
}

func ClearGrid(grid *ui.Grid) {
	w, h := ui.TerminalDimensions()
	grid = ui.NewGrid()
	grid.SetRect(0, 0, w, h)
	ShowLoading(grid)
	ui.Clear()
	ui.Render(grid)
	time.Sleep(150 * time.Millisecond)
}

func ShowLoading(grid *ui.Grid) {
	loading := widgets.NewParagraph()
	loading.Text = `
                                           d8b      888
                                           Y8P      888
                                                    888
.d88b.   .d88b.   .d8888b .d88b.  888  888 888  .d88888
d88P"88b d88""88b d88P"   d88""88b 888  888 888 d88" 888
888  888 888  888 888     888  888 Y88  88P 888 888  888
Y88b 888 Y88..88P Y88b.   Y88..88P  Y8bd8P  888 Y88b 888
"Y88888  "Y88P"   "Y8888P "Y88P"    Y88P   888  "Y88888
    888
Y8b d88P
"Y88P"


Worldwide Coronavirus (COVID-19) Statistics for your terminal

[Please wait until information is loading](fg:black,bg:yellow)`
	grid.Set(ui.NewRow(1, loading))
	ui.Clear()
	ui.Render(grid)
}

func ShowGlobalWidgets(grid *ui.Grid, global *widgets.Paragraph, table *widgets.Table, instructions *widgets.Paragraph) {
	globalWidget := ui.NewRow(0.15, ui.NewCol(1.0, global))
	countriesTable := ui.NewRow(0.70, ui.NewCol(1.0, table))
	instructionsWidget := ui.NewRow(0.15, ui.NewCol(1.0, instructions))
	grid.Set(globalWidget, countriesTable, instructionsWidget)
	ui.Clear()
	ui.Render(grid)
}

func PrintInstructions(instructions *widgets.Paragraph) {
	instructions.Title = "👉 Commands"
	instructions.Text = `
* Press W to order by country ASC    |   Press E to order by total cases ASC    | Press T to order by total deaths ASC
* Press S to order by country DESC   |   Press D to order by total cases DESC   | Press G to order by total deaths DESC
* Press R to refresh data.`
	instructions.Border = true
	instructions.BorderStyle.Fg = ui.ColorYellow
}

func SortByColumn(countries []domain.Country, column string, orderType string) {
	propertyName := strings.Title(column)
	sort.SliceStable(countries, func(i, j int) bool {
		a := getFieldValue(&countries[i], propertyName)
		switch a.Kind().String() {
		case "int64":
			a := getFieldValue(&countries[i], propertyName).Int()
			b := getFieldValue(&countries[j], propertyName).Int()
			if orderType == "ASC" {
				return a > b
			} else {
				return a < b
			}
		case "string":
			a := getFieldValue(&countries[i], propertyName).String()
			b := getFieldValue(&countries[j], propertyName).String()
			if orderType == "ASC" {
				return a > b
			} else {
				return a < b
			}
		default:
			return false
		}
	})
}

func getFieldValue(e *domain.Country, field string) reflect.Value {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f
}

func ListenCommands(countries []domain.Country, table *widgets.Table, all domain.All, global *widgets.Paragraph, grid *ui.Grid) {

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<Escape>":
			return
		case "w":
			ClearGrid(grid)
			SortByColumn(countries, "country", "ASC")
			controllers.PrintCountryTable(&countries, table)
		case "s":
			ClearGrid(grid)
			SortByColumn(countries, "country", "DESC")
			controllers.PrintCountryTable(&countries, table)
		case "e":
			ClearGrid(grid)
			SortByColumn(countries, "cases", "ASC")
			controllers.PrintCountryTable(&countries, table)
		case "d":
			ClearGrid(grid)
			SortByColumn(countries, "cases", "DESC")
			controllers.PrintCountryTable(&countries, table)
		case "t":
			ClearGrid(grid)
			SortByColumn(countries, "deaths", "ASC")
			controllers.PrintCountryTable(&countries, table)
		case "g":
			ClearGrid(grid)
			SortByColumn(countries, "deaths", "DESC")
			controllers.PrintCountryTable(&countries, table)
		case "r":
			ClearGrid(grid)
			controllers.GetAllData(&all, global)
			controllers.GetCountriesData(&countries, table)
		}
		ui.Render(grid)
	}
}
