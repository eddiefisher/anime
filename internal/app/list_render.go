package app

import (
	"strings"

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

func (s *listRender) mainText() string {
	return strings.Join([]string{s.anime.CurrentEpisode, s.anime.Title, s.anime.Description}, " ")
}

func (s *listRender) secondaryText() string {
	return strings.Join([]string{"\t", s.anime.Date, "[" + string(s.anime.Status.Color) + "]", s.anime.Status.Name.String()}, " ")
}

func (s *listRender) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyEnter {
		if err := browser.Open(s.currentItem().Source); err != nil {
			s.errText.SetText(err.Error())
		}
		return event
	}

	return event
}

func (s *listRender) mouseCapture(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
	if action == tview.MouseLeftDoubleClick {
		if err := browser.Open(s.currentItem().Source); err != nil {
			s.errText.SetText(err.Error())
		}
	}
	return action, event
}

func (s listRender) currentItem() entity.Anime {
	return s.animes[s.list.GetCurrentItem()]
}
