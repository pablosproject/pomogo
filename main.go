package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

func main() {

	workLenght := flag.Int("l", 25, "lenght of work in minutes")
	shortBreak := flag.Int("s", 5, "short pause (in minute)")
	flag.Parse()

	done := make(chan bool)
	timer := NewTimer(*workLenght, *shortBreak)

	go func() {
		for {
			select {
			case currentTime := <-timer.TimeC:
				fmt.Println(formatDuration(currentTime))
			case newState := <-timer.StateC:
				switch newState {
				case IDLE:
					fmt.Println("Idle state")
				case WORK:
					fmt.Println("work state")
				case SHORTBREAK:
					fmt.Println("shortbreak state")
				case LONGBREAK:
					fmt.Println("longbreak state")
				}
			}
		}
	}()

	timer.Start()

	<-done
}

func formatDuration(d time.Duration) string {
	m := int64(d.Minutes())
	s := int64(math.Mod(d.Seconds(), 60))
	return fmt.Sprintf("%02d:%02d", m, s)
}
