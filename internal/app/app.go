package app

import (
	"strings"
	"time"

	"github.com/eddiefisher/anime/internal/browser"
	"github.com/eddiefisher/anime/internal/data"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const updateInterval = 30 * time.Minute

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
	collectAnime(list)

	list.SetTitle(" Anime ").SetBorder(true)

	flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(list, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(menu, 0, 1, false).AddItem(errText, 0, 1, false), 1, 1, false),
		0, 2, true)

	inputCapture := func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			app.Stop()
		case tcell.KeyRune:
			switch event.Rune() {
			case 'r':
				list.Clear()
				collectAnime(list)
			}
		}

		return event
	}

	updateFunc := func() {
		for {
			time.Sleep(updateInterval)
			app.QueueUpdate(func() {
				list.Clear()
				collectAnime(list)
				app.ForceDraw()
			})
		}
	}

	flex.SetInputCapture(inputCapture)

	go updateFunc()

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func collectAnime(list *tview.List) {
	animes, err := data.Load()
	if err != nil {
		errText.SetText(err.Error())
		return
	}
	for _, anime := range animes {

		mainText := []string{
			anime.CurrentEpisode, anime.Title, anime.Description,
		}
		secondaryText := []string{
			"\t", anime.Date, "[" + string(anime.Status.Color) + "]", anime.Status.Name.String(),
		}
		inputCapture := func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEnter:
				if err := browser.Open(animes[list.GetCurrentItem()].Source); err != nil {
					errText.SetText(err.Error())
				}
				return event
			}

			return event
		}
		mouseCapture := func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
			switch action {
			case tview.MouseLeftDoubleClick:
				if err := browser.Open(animes[list.GetCurrentItem()].Source); err != nil {
					errText.SetText(err.Error())
				}
			}
			return action, event
		}
		list.AddItem(strings.Join(mainText, " "), strings.Join(secondaryText, " "), 0, nil).
			SetSecondaryTextColor(tcell.Color111).
			SetHighlightFullLine(true).
			SetWrapAround(true).
			SetInputCapture(inputCapture).
			SetMouseCapture(mouseCapture)
	}
}
