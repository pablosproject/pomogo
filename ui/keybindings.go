package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/pablosproject/pomogo/timer"
)

func (u *UI) registerKeybindings() error {
	if err := u.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, u.quit); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding("", 'q', gocui.ModNone, u.quit); err != nil {
		return err
	}
	if err := u.gui.SetKeybinding("", 's', gocui.ModNone, u.keyS); err != nil {
		return err
	}
	return nil
}

func (u *UI) keyS(gui *gocui.Gui, v *gocui.View) error {
	state := u.timer.State()
	if state == timer.IDLE {
		u.timer.Start()
	}
	if state == timer.LONGBREAK || state == timer.SHORTBREAK || state == timer.WORK {
		u.timer.Stop()
	}

	return nil
}

func (ui *UI) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
