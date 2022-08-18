package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/pablosproject/pomogo/timer"
)

func (u *UI) layoutFooter() error {
	width, height := u.gui.Size()
	footerWidth := 40
	footerHeight := 3
	footerPadding := +5
	footerTop, footerBottom := footerViewPosition(width, height, footerWidth, footerHeight, -footerPadding)
	if _, err := u.gui.SetView("footer", footerTop.x, footerTop.y, footerBottom.x, footerBottom.y); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}

func (u *UI) renderFooter() {
	v, _ := u.gui.View("footer")

	footerMessage := "ctrt+c/q: quit    "
	if u.timer.State() == timer.IDLE {
		footerMessage += "S: Start Pomodoro "
	}
	if u.timer.State() == timer.LONGBREAK || u.timer.State() == timer.SHORTBREAK {
		footerMessage += "S: Stop Current Break "
	}
	if u.timer.State() == timer.WORK {
		footerMessage += "S: Stop Current Work "
	}

	u.gui.Update(func(gui *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, centeredString(v, footerMessage))
		return nil
	})
}
