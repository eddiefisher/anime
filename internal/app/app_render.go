package app

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const updateInterval = 30 * time.Minute

type appRender struct {
	app          *tview.Application
	list         *tview.List
	errText      *tview.TextView
	collectAnime func(*tview.List)
}

func (s *appRender) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		s.app.Stop()
	case tcell.KeyRune:
		if event.Rune() == 'r' {
			s.list.Clear()
			s.collectAnime(s.list)
		}
	}

	return event
}

func (s *appRender) updateFunc() {
	for {
		time.Sleep(updateInterval)
		s.app.QueueUpdate(func() {
			s.list.Clear()
			s.errText.Clear()
			s.collectAnime(s.list)
			s.app.ForceDraw()
		})
	}
}
