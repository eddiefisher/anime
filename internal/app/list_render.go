package app

import (
	"strings"

	"github.com/eddiefisher/anime/internal/actions"
	"github.com/eddiefisher/anime/internal/browser"
	"github.com/eddiefisher/anime/internal/entity"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type listRender struct {
	animes  []entity.Anime
	anime   entity.Anime
	errText *tview.TextView
	list    *tview.List
}

func (s *listRender) Render(list *tview.List) {
	animes, err := actions.Load()
	if err != nil {
		s.errText.SetText(err.Error())

		return
	}
	s.animes = animes

	s.list.SetTitle(" Anime ").SetBorder(true)
	for _, anime := range s.animes {
		s.anime = anime
		s.renderItem()
	}
}

func (s *listRender) renderItem() {
	s.list.AddItem(s.mainText(), s.secondaryText(), 0, nil).
		SetSecondaryTextColor(tcell.Color111).
		SetHighlightFullLine(true).
		SetWrapAround(true).
		SetInputCapture(s.inputCapture).
		SetMouseCapture(s.mouseCapture)
}

func (s *listRender) mainText() string {
	data := []string{
		s.anime.CurrentEpisode,
		s.anime.Title,
		s.anime.Description,
	}

	return strings.Join(data, " ")
}

func (s *listRender) secondaryText() string {
	data := []string{
		"\t", s.anime.Date,
		"[" + string(s.anime.Status.Color) + "]",
		s.anime.Status.Name.String(),
	}

	return strings.Join(data, " ")
}

func (s *listRender) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
		s.openInBrowser()
	}

	return event
}

func (s *listRender) mouseCapture(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
	if action == tview.MouseLeftDoubleClick {
		s.openInBrowser()
	}

	return action, event
}

func (s *listRender) openInBrowser() {
	if err := browser.Open(s.animes[s.list.GetCurrentItem()].Source); err != nil {
		s.errText.SetText(err.Error())
	}
}
