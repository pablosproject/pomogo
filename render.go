package main

import (
	"fmt"
	"math"
	"time"

	"github.com/nsf/termbox-go"
)

type RenderState struct {
	state         WorkState
	remainingTime time.Duration
}

type Renderer interface {
	Render(state RenderState) error
}

type Render struct {
}

func (r *Render) Render(state RenderState) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	switch state.state {
	case IDLE:
		r.renderTitle("State: IDLE")
	case WORK:
		r.renderTitle("State: WORK")
		r.renderTime(state.remainingTime)
	case SHORTBREAK:
		r.renderTitle("State: SHORTBREAK")
		r.renderTime(state.remainingTime)
	case LONGBREAK:
		r.renderTitle("State: LONGBREAK")
		r.renderTime(state.remainingTime)
	}

	termbox.Flush()
}

func (r *Render) renderTitle(title string) {
	x, y := midPosition(title, 0, -10)
	tbprint(x, y, termbox.ColorRed, termbox.ColorDefault, title)
}
func (r *Render) renderTime(time time.Duration) {
	timeString := formatDuration(time)
	x, y := midPosition(timeString, 0, 0)
	tbprint(x, y, termbox.ColorCyan|termbox.AttrBlink, termbox.ColorDefault, timeString)
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func formatDuration(d time.Duration) string {
	m := int64(d.Minutes())
	s := int64(math.Mod(d.Seconds(), 60))
	return fmt.Sprintf("%02d:%02d", m, s)
}

func midPosition(message string, x_offset, y_offset int) (x int, y int) {
	w, h := termbox.Size()
	midy := h / 2
	midx := (w - len(message)) / 2
	return midx + x_offset, midy + y_offset
}
