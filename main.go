package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/controllers"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/domain"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/utils"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	grid := utils.CreateInitialGrid()
	utils.ClearGrid(grid)

	global, table, instructions := utils.CreateInitialWidgets()

	var all domain.All
	controllers.GetAllData(&all, global)

	var countries []domain.Country
	controllers.GetCountriesData(&countries, table)

	utils.PrintInstructions(instructions)
	utils.ShowGlobalWidgets(grid, global, table, instructions)
	utils.ListenCommands(countries, table, all, global, grid)
}
