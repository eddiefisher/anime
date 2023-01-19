package app

import (
	"github.com/eddiefisher/anime/internal/actions"
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
	ar := appRender{
		app:          app,
		list:         list,
		errText:      errText,
		collectAnime: collectAnime,
	}

	collectAnime(list)
	list.SetTitle(" Anime ").SetBorder(true)

	flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(list, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(menu, 0, 1, false).AddItem(errText, 0, 1, false), 1, 1, false),
		0, 2, true)

	flex.SetInputCapture(ar.inputCapture)

	go ar.updateFunc()

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func collectAnime(list *tview.List) {
	animes, err := actions.Load()
	if err != nil {
		errText.SetText(err.Error())
		return
	}
	lr := listRender{
		animes:  animes,
		errText: errText,
		list:    list,
	}
	for _, anime := range animes {
		lr.anime = anime
		list.AddItem(lr.mainText(), lr.secondaryText(), 0, nil).
			SetSecondaryTextColor(tcell.Color111).
			SetHighlightFullLine(true).
			SetWrapAround(true).
			SetInputCapture(lr.inputCapture).
			SetMouseCapture(lr.mouseCapture)
	}
}
