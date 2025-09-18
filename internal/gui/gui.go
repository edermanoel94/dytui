package gui

import (
	"context"
	"dytui/internal/controller"
	"fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type KeyOp int16

const (
	KeyTablesOp KeyOp = iota
	KeyResultOp
	KeyPreviewOp
	KeyQueryOp
)

var (
	KeyMapping = map[KeyOp]tcell.Key{
		KeyTablesOp:  tcell.KeyCtrlD,
		KeyResultOp:  tcell.KeyCtrlA,
		KeyPreviewOp: tcell.KeyCtrlE,
		KeyQueryOp:   tcell.KeyCtrlQ,
	}
)

var (
	TitleTablesView  = fmt.Sprintf("Sources [ %s ]", tcell.KeyNames[KeyMapping[KeyTablesOp]])
	TitleResultView  = fmt.Sprintf("Schemas [ %s ]", tcell.KeyNames[KeyMapping[KeyResultOp]])
	TitlePreviewView = fmt.Sprintf("Tables [ %s ]", tcell.KeyNames[KeyMapping[KeyPreviewOp]])
	TitleQueryView   = fmt.Sprintf("Preview [ %s ]", tcell.KeyNames[KeyMapping[KeyQueryOp]])

	TitleFooterView = "Navigate [ Tab / Shift-Tab ] · Focus [ Ctrl-F ] · Exit [ Ctrl-C ] \n Tables specific: Describe [ e ] · Preview [ p ]"
)

type Gui struct {
	ctrl *controller.Controller

	focusMode bool

	// View commponents
	App   *tview.Application
	Pages *tview.Pages

	MainFlexView   *tview.Grid
	ResultGridView *tview.Grid

	Tables       *tview.List
	ResultView   *tview.List
	PreviewTable *tview.Table
	QueryInput   *tview.InputField
	FooterText   *tview.TextView
}

func New(ctrl *controller.Controller) *Gui {

	g := &Gui{
		ctrl: ctrl,
	}

	g.App = tview.NewApplication()

	g.Tables = tview.NewList()
	g.Tables.
		SetBorder(true).
		SetTitle("Tables")

	g.Tables.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {

		fmt.Println("Mensagem da tabela que foi selecionda: ", index, mainText, shortcut)
	})

	g.ResultView = tview.NewList()
	g.ResultView.
		SetBorder(true).
		SetTitle("Result")

	g.PreviewTable = tview.NewTable()

	g.FooterText = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(TitleFooterView).SetTextColor(tcell.ColorGray)

	g.QueryInput = tview.NewInputField()
	g.QueryInput.
		SetBorder(true).
		SetTitle("Query")

	navigate := tview.NewGrid().
		AddItem(g.Tables, 0, 0, 1, 1, 0, 0, true)

	previewAndQuery := tview.NewGrid().
		SetRows(0, 3).
		AddItem(g.ResultView, 0, 0, 1, 1, 0, 0, false).
		AddItem(g.QueryInput, 1, 0, 1, 1, 0, 0, false)

	g.MainFlexView = tview.NewGrid().
		SetRows(0, 2).
		SetColumns(40, 0).
		SetBorders(false).
		AddItem(navigate, 0, 0, 1, 1, 0, 0, false).
		AddItem(previewAndQuery, 0, 1, 1, 1, 0, 0, false).
		AddItem(g.FooterText, 1, 0, 1, 2, 0, 0, false)

	g.loadData()

	return g
}

func (g *Gui) Run() error {
	return g.App.SetRoot(g.MainFlexView, true).EnableMouse(true).Run()
}

func (g *Gui) loadData() {

	g.Tables.Clear()
	g.ResultView.Clear()

	tables, err := g.ctrl.Current().ListTables(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	g.queueUpdateDraw(func() {
		for _, t := range tables {
			g.Tables.AddItem(t, "", 0, nil)
		}

		g.App.SetFocus(g.ResultView)
	})

}

func (g *Gui) showData() {
	g.queueUpdateDraw(func() {

		g.PreviewTable.Clear()

		g.PreviewTable.SetTitle(fmt.Sprintf("%s: %s", TitlePreviewView, "Result"))
		g.PreviewTable.SetFixed(1, 1)
		g.PreviewTable.SetSelectable(true, false)
		g.PreviewTable.ScrollToBeginning()
	})
}

func (g *Gui) queueUpdateDraw(f func()) {
	go func() {
		g.App.QueueUpdateDraw(f)
	}()
}

func (g *Gui) showMessage(msg string) {
	g.queueUpdateDraw(func() {
		g.FooterText.SetText(msg).SetTextColor(tcell.ColorGreen)
	})
	go time.AfterFunc(3*time.Second, g.resetMessage)
}

func (g *Gui) resetMessage() {
	g.queueUpdateDraw(func() {
		g.FooterText.SetText(TitleFooterView).SetTextColor(tcell.ColorGray)
	})
}
