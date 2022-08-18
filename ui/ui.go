package ui

import (
	"log"
	"os"
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
	if err = u.registerKeybindings(); err != nil {
		//TODO: error surface to top level
		log.Panicln(err)
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
