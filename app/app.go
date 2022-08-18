package app

import (
	"os"

	"github.com/pablosproject/pomogo/timer"
	"github.com/pablosproject/pomogo/ui"
)

type App struct {
	timer *timer.PomodoroTimer
	ui    *ui.UI
}

func NewApp(config timer.PomodoroConfig) *App {
	timer := timer.NewTimer(config)
	return &App{
		timer: timer,
		ui:    ui.NewUI(timer),
	}
}

func (a *App) Start() error {
	if err := a.ui.Run(); err != nil {
		os.Exit(-1)
		// TODO: handle exit state
	}

	return nil
}
