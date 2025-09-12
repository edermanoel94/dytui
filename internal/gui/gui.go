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
	App   *tview.Application
	Pages *tview.Pages

	Profiles *tview.List
	Tables   *tview.List
	Result   *tview.Table
}

func New(ctrl *controller.Controller) *Gui {

	g := &Gui{
		ctrl: ctrl,
	}

	g.App = tview.NewApplication()

	g.Profiles = tview.NewList()
	g.Profiles.SetBorder(true)

	g.Tables = tview.NewList()
	g.Tables.SetBorder(true)

	g.Result = tview.NewTable()
	g.Result.SetBorder(true)

	for _, table := range []string{} {
		g.Tables.AddItem(table, "", 0, nil)
	}

	flex := tview.NewFlex().
		AddItem(g.Tables, 0, 1, false).
		AddItem(g.Profiles, 0, 1, false).
		AddItem(g.Result, 0, 3, true)

	g.Pages = tview.NewPages().
		AddPage("DYTUI", flex, true, true)

	g.App.SetRoot(g.Pages, true)

	return g
}

func (g *Gui) Run() error {

	if err := g.App.Run(); err != nil {
		return err
	}
	return nil
}
