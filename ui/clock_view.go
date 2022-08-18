package ui

import (
	"log"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"github.com/pablosproject/pomogo/timer"
)

func (u *UI) layoutClock() error {
	width, height := u.gui.Size()
	timerwidth := 20
	timerHeight := 5
	timerTop, timerBottom := centeredView(width, height, timerwidth, timerHeight, 0, 0)

	// Layout timer box
	if _, err := u.gui.SetView("timer", timerTop.x, timerTop.y, timerBottom.x, timerBottom.y); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
	}
	return nil
}

func (u *UI) renderClock() {
	v, _ := u.gui.View("timer")

	u.gui.Update(func(gui *gocui.Gui) error {
		v.Clear()
		if u.timer.State() != timer.IDLE {
			s := centeredString(v, formatDuration(u.timer.RemainingTime()))
			color.New(color.FgRed).Fprint(v, s)
		}
		return nil
	})
}
