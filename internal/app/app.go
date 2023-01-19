package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app  = tview.NewApplication()
	flex = tview.NewFlex()
	list = tview.NewList()
	menu = tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(Enter) open in browser | (r) reload | (Esc) to quit")
	errText = tview.NewTextView().
		SetTextColor(tcell.ColorRed)
)

func New() {
	lr := listRender{
		errText: errText,
		list:    list,
	}

	ar := appRender{
		app:        app,
		list:       list,
		flex:       flex,
		menu:       menu,
		errText:    errText,
		listRender: lr,
	}

	lr.Render(list)
	ar.Render()
}
