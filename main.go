package main

import (
	"errors"
	"log"

	"github.com/awesome-gocui/gocui"
)

// ❯ ffplay -v warning -nodisp -autoexit -af "atempo=1" -ss 4 -t 5 -loop 0 samples/ton-doux-sourire.mp3
// cmd := exec.Command( "ffplay", "-v", "warning", "-nodisp", "-autoexit", "-af", "atempo=0.8", "-ss", "4", "-t", "5", "-loop", "0", "samples/ton-doux-sourire.mp3",)

// if err := cmd.Run(); err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("DONE!!!")
// fmt.Printf("what??? %v\n", cmd.Process.Pid)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 20, 10, 50, 50, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}

		v.Title = "Foo"
		v.Editable = true
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
