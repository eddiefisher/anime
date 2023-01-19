package app

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const updateInterval = 30 * time.Minute

type appRender struct {
	app        *tview.Application
	list       *tview.List
	errText    *tview.TextView
	listRender listRender
}

func (s *appRender) Render() {
	flex.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(s.list, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(menu, 0, 1, false).AddItem(s.errText, 0, 1, false), 1, 1, false),
		0, 2, true)

	flex.SetInputCapture(s.inputCapture)

	go s.updateFunc()

	if err := s.app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func (s *appRender) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		s.app.Stop()
	case tcell.KeyRune:
		if event.Rune() == 'r' {
			s.list.Clear()
			s.listRender.Render(s.list)
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
			s.listRender.Render(s.list)
			s.app.ForceDraw()
		})
	}
}
