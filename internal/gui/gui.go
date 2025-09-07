package gui

import (
	"log"

	"github.com/rivo/tview"
)

var (
	app   *tview.Application
	pages *tview.Pages
)

func Start() {

	app = tview.NewApplication()

	tables := tview.NewList()
	tables.SetBorder(true)

	result := tview.NewTable()
	result.SetBorder(true)

	flex := tview.NewFlex().
		AddItem(tables, 0, 1, true).
		AddItem(result, 0, 1, false)

	tables.AddItem("teste", "", 0, func() {

	})

	pages := tview.NewPages().
		AddPage("Teste", flex, true, true)

	app.SetRoot(pages, false)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
