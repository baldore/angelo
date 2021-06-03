package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
)

// ❯ ffplay -v warning -nodisp -autoexit -af "atempo=1" -ss 4 -t 5 -loop 0 samples/ton-doux-sourire.mp3
// cmd := exec.Command( "ffplay", "-v", "warning", "-nodisp", "-autoexit", "-af", "atempo=0.8", "-ss", "4", "-t", "5", "-loop", "0", "samples/ton-doux-sourire.mp3",)

// if err := cmd.Run(); err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("DONE!!!")
// fmt.Printf("what??? %v\n", cmd.Process.Pid)

type UI struct {
	screen tcell.Screen
	quitCh chan struct{}
	keysCh chan *tcell.EventKey
}

func newUI() UI {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	return UI{
		screen: screen,
		quitCh: make(chan struct{}),
		keysCh: make(chan *tcell.EventKey),
	}
}

// Initializes the screen.
func (ui UI) initScreen() {
	if err := ui.screen.Init(); err != nil {
		log.Fatal(err)
	}

	ui.screen.SetStyle(tcell.StyleDefault)
	ui.screen.Clear()
}

// Starts the application.
func (ui UI) startApp() {
	for {
		select {
		case <-ui.quitCh:
			return
		case key := <-ui.keysCh:
			handleKeyEvent(key)
		}
	}
}

// Liberate resources when the app finishes.
func (ui UI) cleanUp() {
	ui.screen.Fini()
}

// Listens the events. It's recommended to run this in a goroutine.
func (ui UI) listenEvents() {
	for {
		ev := ui.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
				close(ui.quitCh)

				return
			default:
				ui.keysCh <- ev
			}
		case *tcell.EventResize:
			ui.screen.Sync()
		}
	}
}

func handleKeyEvent(key *tcell.EventKey) {
	switch key.Key() {
	case tcell.KeyLeft:
		fmt.Printf("%v\n", "left")
	case tcell.KeyRight:
		fmt.Printf("%v\n", "right")
	}
}

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	ui := newUI()
	ui.initScreen()
	go ui.listenEvents()
	ui.startApp()
	ui.cleanUp()
}
