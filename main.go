package main

import (
	"os"
	"time"

	"github.com/pablosproject/pomogo/app"
	"github.com/pablosproject/pomogo/timer"
)

func main() {

	// workLenght := flag.Int("l", 25, "lenght of work in minutes")
	// shortBreak := flag.Int("s", 5, "short pause (in minute)")
	// flag.Parse()

	config := timer.PomodoroConfig{
		WorkLenght:       10 * time.Second,
		ShortPauseLenght: 5 * time.Second,
		LongPauseLenght:  1 * time.Second,
	}

	app := app.NewApp(config)

	if err := app.Start(); err != nil {
		os.Exit(-1)
		// TODO: handle exit state
	}

	os.Exit(0)
}
