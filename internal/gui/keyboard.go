package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type navigation struct {
	next, prev tview.Primitive
}

func (g *Gui) setupKeyboard() {
	_ = map[tview.Primitive]navigation{}

	// Setup app level keyboard shortcuts.
	g.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

	// Setup Tables element level keyboard shortcuts.
	g.Tables.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return nil
	})
}
