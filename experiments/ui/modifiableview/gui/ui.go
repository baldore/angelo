package gui

import (
	"errors"
	"log"

	"github.com/awesome-gocui/gocui"
)

type UI struct {
	box Box
}

func NewUI() UI {
	return UI{
		box: NewBox(),
	}
}

func (ui UI) Start() error {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(ui.Layout)

	// TODO: set bindings somewhere else
	err = g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return nil
	})
	if err != nil {
		log.Panicln(err)
	}

	if err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}

	return nil
}

func (ui UI) StartKeybindings() {

}

func (ui UI) Layout(g *gocui.Gui) error {
	// fmt.Printf("rendering layout: %v\n", ui)
	// maxX, maxY := g.Size()

	if _, err := g.SetView("v1", 0, 0, 1, 1, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
