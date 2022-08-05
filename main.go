package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

func main() {

	err := termbox.Init()
	if err != nil {
		//TODO: better error handkung
		panic(err)
	}
	defer termbox.Close()

	workLenght := flag.Int("l", 25, "lenght of work in minutes")
	shortBreak := flag.Int("s", 5, "short pause (in minute)")
	flag.Parse()

	done := make(chan bool)
	timer := NewTimer(*workLenght, *shortBreak)
	render := Render{}

	go func() {
		for {
			select {
			case currentTime := <-timer.TimeC:
				render.Render(RenderState{
					state:         timer.State(),
					remainingTime: currentTime,
				})
			case newState := <-timer.StateC:
				render.Render(RenderState{
					state:         newState,
					remainingTime: 0,
				})
			}
		}
	}()

	go func() {
		for {
			event := termbox.PollEvent()
			fmt.Println(event)
			termbox.Close()
			os.Exit(0)
		}
	}()

	timer.Start()

	<-done
}
