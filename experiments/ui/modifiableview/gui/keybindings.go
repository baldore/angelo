package gui

import (
	"github.com/awesome-gocui/gocui"
)

type Binding struct {
	// TODO: not sure if there's a better way for this
	Key      interface{}
	Handler  func() error
	Modifier gocui.Modifier
}

func wrapKeybindingHandler(f func() error) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return f()
	}
}

func (ui UI) getInitialBindings() []Binding {
	return []Binding{
		{
			Key: gocui.KeyArrowLeft,
			Handler: func() error {
				ui.box.MoveLeft()
				return nil
			},
		},
		{
			Key: gocui.KeyArrowRight,
			Handler: func() error {
				ui.box.MoveRight()
				return nil
			},
		},
		{
			Key: gocui.KeyArrowUp,
			Handler: func() error {
				ui.box.MoveUp()
				return nil
			},
		},
		{
			Key: gocui.KeyArrowDown,
			Handler: func() error {
				ui.box.MoveDown()
				return nil
			},
		},
		{
			Key:     gocui.KeyCtrlC,
			Handler: quit,
		},
	}
}

func (ui UI) StartKeybindings() error {
	bindings := ui.getInitialBindings()
	for _, binding := range bindings {
		err := ui.g.SetKeybinding("", binding.Key, binding.Modifier, wrapKeybindingHandler(binding.Handler))
		if err != nil {
			return err
		}
	}

	return nil
}

func quit() error {
	return gocui.ErrQuit
}
