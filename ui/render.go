package ui

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/pablosproject/pomogo/timer"
)

type UI struct {
	timer *timer.PomodoroTimer
	gui   *gocui.Gui
}

func NewUI(timer *timer.PomodoroTimer) *UI {
	return &UI{
		timer: timer,
	}
}

func (u *UI) Run() error {
	gui, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		// TODO: surface the error on top level and listen to error up
		os.Exit(-1)
	}
	u.gui = gui
	defer gui.Close()

	gui.SetManagerFunc(u.layout)
	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := gui.SetKeybinding("", 's', gocui.ModNone, u.start); err != nil {
		return err
	}
	go u.startPolling()
	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}

func (u *UI) startPolling() {
	ticker := time.NewTicker(30 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		u.render()
	}
}

func (u *UI) render() {
	u.renderState()

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
		fmt.Fprint(v, centeredString(v, fmt.Sprintf("%d", u.timer.State())))
		return nil
	})
}

func (u *UI) layout(g *gocui.Gui) error {
	width, height := g.Size()

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

	return nil
}

func (u *UI) start(gui *gocui.Gui, v *gocui.View) error {
	u.timer.Start()
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

