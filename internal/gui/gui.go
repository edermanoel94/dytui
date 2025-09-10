package gui

import (
	"dytui/internal/controller"

	"github.com/rivo/tview"
)

var (
	app   *tview.Application
	pages *tview.Pages
)

type Gui struct {
	ctrl *controller.Controller

	// View commponents
	app   *tview.Application
	pages *tview.Pages

	profiles *tview.List
	tables   *tview.List
	result   *tview.Table
}

func New(ctrl *controller.Controller) *Gui {
	return &Gui{
		ctrl: ctrl,
	}
}

func (g *Gui) Run() error {

	app = tview.NewApplication()

	profiles := tview.NewList()
	profiles.SetBorder(true)

	tables := tview.NewList()
	tables.SetBorder(true)

	result := tview.NewTable()
	result.SetBorder(true)

	for _, table := range []string{} {
		tables.AddItem(table, "", 0, nil)
	}

	flex := tview.NewFlex().
		AddItem(tables, 0, 1, false).
		AddItem(profiles, 0, 1, false).
		AddItem(result, 0, 3, true)

	pages := tview.NewPages().
		AddPage("DYTUI", flex, true, true)

	app.SetRoot(pages, true)

	if err := app.Run(); err != nil {
		return err
	}
	return nil
}
