package main

import (
	"flag"
	"os"
	"time"

	"github.com/pablosproject/pomogo/timer"
	"github.com/pablosproject/pomogo/ui"
)

func main() {

	// workLenght := flag.Int("l", 25, "lenght of work in minutes")
	// shortBreak := flag.Int("s", 5, "short pause (in minute)")
	flag.Parse()

	timer := timer.NewTimer(
		// PomodoroConfig{
		// 	workLenght:       time.Duration(*workLenght) * time.Minute,
		// 	shortPauseLenght: time.Duration(*shortBreak) * time.Minute,
		// 	longPauseLenght:  0,
		// },
		timer.PomodoroConfig{
			WorkLenght:       10 * time.Second,
			ShortPauseLenght: 5 * time.Second,
			LongPauseLenght:  1 * time.Second,
		},
	)
	render := ui.NewUI(timer)

	if err := render.Run(); err != nil {
		os.Exit(-1)
		// TODO: handle exit state
	}

	os.Exit(0)
}
