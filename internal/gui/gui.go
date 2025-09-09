package gui

import (
	"context"
	"dytui/internal/awsutil"
	"dytui/internal/controller"
	"dytui/internal/dynamo"
	"log"

	"github.com/rivo/tview"
)

var (
	app   *tview.Application
	pages *tview.Pages
)

type Gui struct {
	ctrl controller.Controller

	// View commponents
	app   *tview.Application
	pages *tview.Pages

	profiles *tview.List
	tables   *tview.List
	result   *tview.Table
}

func Start() {

	ctx := context.Background()

	app = tview.NewApplication()

	profiles := tview.NewList()
	profiles.SetBorder(true)

	tables := tview.NewList()
	tables.SetBorder(true)

	result := tview.NewTable()
	result.SetBorder(true)

	credentials, err := awsutil.LoadAWSCredentials()

	if err != nil {
		log.Fatal(err)
	}

	for _, cred := range credentials {
		profiles.AddItem(cred.Name, "", 0, nil)
	}

	tableNames, err := dynamo.ListTables(ctx)

	if err != nil {
		log.Fatal(err)
	}

	for _, table := range tableNames {
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
		log.Fatal(err)
	}
}
