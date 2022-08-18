package ui

import (
	"log"

	"github.com/jroimartin/gocui"
)

func (u *UI) render() {
	u.renderTitle()
	u.renderFooter()
	u.renderClock()
}

func (u *UI) layout(g *gocui.Gui) error {
	// Layout footer view
	if err := u.layoutClock(); err != nil {
		log.Panicln(err)
		return err
	}

	// Layout state View
	if err := u.layoutTitle(); err != nil {
		log.Panicln(err)
		return err

	}
	// Layout footer view
	if err := u.layoutFooter(); err != nil {
		log.Panicln(err)
		return err
	}

	return nil
}
