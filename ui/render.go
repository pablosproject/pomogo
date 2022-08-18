package ui

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/pablosproject/pomogo/timer"
)

func (u *UI) render() {
	u.renderState()
	u.renderFooter()

	switch u.timer.State() {
	case timer.WORK, timer.SHORTBREAK, timer.LONGBREAK:
		u.renderTimer()
	case timer.IDLE:
		u.renderIdle()
	}
}

func (u *UI) renderTimer() {
	v, _ := u.gui.View("timer")

	u.gui.Update(func(gui *gocui.Gui) error {
		s := centeredString(v, formatDuration(u.timer.RemainingTime()))
		v.Clear()
		fmt.Fprint(v, s)
		return nil
	})
}

func (u *UI) renderIdle() {
	v, _ := u.gui.View("timer")

	u.gui.Update(func(gui *gocui.Gui) error {
		v.Clear()
		return nil
	})
}

func (u *UI) renderState() {
	v, _ := u.gui.View("state")

	u.gui.Update(func(gui *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, centeredString(v, formatState(u.timer.State())))
		return nil
	})
}

func (u *UI) renderFooter() {
	v, _ := u.gui.View("state")

	u.gui.Update(func(gui *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, centeredString(v, formatState(u.timer.State())))
		return nil
	})
}

func (u *UI) layout(g *gocui.Gui) error {
	width, height := g.Size()

	// TODO: create an abstraction for this windows and move initialization there
	// Layout main view
	timerwidth := 20
	timerHeight := 5
	timerTop, timerBottom := centeredView(width, height, timerwidth, timerHeight, 0, 0)

	// Layout timer box
	if _, err := g.SetView("timer", timerTop.x, timerTop.y, timerBottom.x, timerBottom.y); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
	}

	// Layout state View
	stateWidth := 40
	stateHeight := 5
	paddingTop := 10
	stateTop, stateBottom := centeredView(width, height, stateWidth, stateHeight, 0, -paddingTop)
	if _, err := g.SetView("state", stateTop.x, stateTop.y, stateBottom.x, stateBottom.y); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
	}

	// Layout footer view
	footerWidth := 40
	footerHeight := 3
	footerPadding := 20
	footerTop, footerBottom := footerView(width, height, footerWidth, footerHeight, -footerPadding)
	if _, err := g.SetView("footer", footerTop.x, footerTop.y, footerBottom.x, footerBottom.y); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
			return err
		}
	}

	return nil
}
