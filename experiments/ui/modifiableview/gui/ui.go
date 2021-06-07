package gui

import (
	"errors"

	"github.com/awesome-gocui/gocui"
)

type UI struct {
	box *Box
	g   *gocui.Gui
}

func NewUI() UI {
	return UI{
		box: NewBox(),
	}
}

func (ui UI) Start() error {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		return err
	}
	defer g.Close()

	ui.g = g
	g.SetManagerFunc(ui.Layout)
	if ui.StartKeybindings(); err != nil {
		return err
	}
	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		return err
	}

	return nil
}

func (ui *UI) Layout(g *gocui.Gui) error {
	b := ui.box
	if _, err := g.SetView("box", b.x, b.y, b.x+b.w, b.y+b.h, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
	}

	return nil
}
