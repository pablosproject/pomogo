package ui

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/pablosproject/pomogo/timer"
)

func formatDuration(d time.Duration) string {
	m := int64(d.Minutes())
	s := int64(math.Mod(d.Seconds(), 60))
	return fmt.Sprintf("%02d:%02d", m, s)
}

// Center a buffer vertically an horizontally in a view
func centeredString(v *gocui.View, s string) string {
	width, height := v.Size()
	vPadding := strings.Repeat("\n", height/2+1)
	hPadding := strings.Repeat(" ", (width-len(s))/2)
	return fmt.Sprintf("%s%s%s", vPadding, hPadding, s)
}

func formatState(s timer.WorkState) string {
	var stateName string
	switch s {
	case timer.IDLE:
		stateName = "IDLE"
	case timer.LONGBREAK:
		stateName = "LONGBREAK"
	case timer.SHORTBREAK:
		stateName = "SHORTBREAK"
	case timer.WORK:
		stateName = "WORK"
	}
	return stateName
}

type Point struct {
	x int
	y int
}

func centeredView(winWidth, winHeight, width, height, offsetX, offsetY int) (topPoint, bottomPoint Point) {
	midPoint := Point{winWidth / 2, winHeight / 2}
	sPoint := Point{midPoint.x - width/2, midPoint.y - height/2}
	ePoint := Point{midPoint.x + width/2, midPoint.y + height/2}

	if offsetX != 0 {
		sPoint.x += offsetX
		ePoint.x += offsetX
	}
	if offsetY != 0 {
		sPoint.y += offsetY
		ePoint.y += offsetY
	}
	return sPoint, ePoint
}

func footerViewPosition(winWidth, winHeight, width, height, offsetY int) (topPoint, bottomPoint Point) {
	bottomOffset := winHeight/2 - height/2.0 + offsetY
	return centeredView(winWidth, winHeight, width, height, 0, bottomOffset)
}
