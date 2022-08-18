package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func (u *UI) layoutTitle() error {
	width, height := u.gui.Size()

	stateWidth := 40
	stateHeight := 5
	paddingTop := 10
	stateTop, stateBottom := centeredView(width, height, stateWidth, stateHeight, 0, -paddingTop)
	if _, err := u.gui.SetView("state", stateTop.x, stateTop.y, stateBottom.x, stateBottom.y); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

func (u *UI) renderTitle() {
	v, _ := u.gui.View("state")

	u.gui.Update(func(gui *gocui.Gui) error {
		v.Clear()
		fmt.Fprint(v, centeredString(v, formatState(u.timer.State())))
		return nil
	})
}
