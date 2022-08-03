package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	// defining flags
	// No flag is also possible

	workLenght := flag.Int("l", 25, "lenght of work in minutes")
	shortBreak := flag.Int("s", 5, "short pause (in minute)")
	flag.Parse()

	timer := PomodoroTimer{
		workLenght:       time.Duration(*workLenght) * time.Minute,
		shortPauseLenght: time.Duration(*shortBreak) * time.Minute,
		longPauseLenght:  0,
		ticker:           &time.Ticker{},
		startTime:        time.Time{},
	}

	timer.start()
	for {
		select {
		case <-timer.ticker.C:
			if timer.remainingTime() <= 0 {
				timer.stop()
				os.Exit(2)
			}
			fmt.Println(formatDuration(timer.remainingTime()))
		}
	}
}

func formatDuration(d time.Duration) string {
	m := int64(d.Minutes())
	s := int64(math.Mod(d.Seconds(), 60))
	return fmt.Sprintf("%02d:%02d", m, s)
}
