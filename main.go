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

// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// 	"unicode"

// 	"github.com/faiface/beep"
// 	"github.com/faiface/beep/effects"
// 	"github.com/faiface/beep/mp3"
// 	"github.com/faiface/beep/speaker"
// 	"github.com/gdamore/tcell"
// )

// func drawTextLine(screen tcell.Screen, x, y int, s string, style tcell.Style) {
// 	for _, r := range s {
// 		screen.SetContent(x, y, r, nil, style)
// 		x++
// 	}
// }

// type audioPanel struct {
// 	sampleRate beep.SampleRate
// 	streamer   beep.StreamSeeker
// 	ctrl       *beep.Ctrl
// 	resampler  *beep.Resampler
// 	volume     *effects.Volume
// }

// func newAudioPanel(sampleRate beep.SampleRate, streamer beep.StreamSeeker) *audioPanel {
// 	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
// 	resampler := beep.ResampleRatio(4, 1, ctrl)
// 	volume := &effects.Volume{Streamer: resampler, Base: 2}
// 	return &audioPanel{sampleRate, streamer, ctrl, resampler, volume}
// }

// func (ap *audioPanel) play() {
// 	speaker.Play(ap.volume)
// }

// func (ap *audioPanel) draw(screen tcell.Screen) {
// 	mainStyle := tcell.StyleDefault.
// 		Background(tcell.NewHexColor(0x473437)).
// 		Foreground(tcell.NewHexColor(0xD7D8A2))
// 	statusStyle := mainStyle.
// 		Foreground(tcell.NewHexColor(0xDDC074)).
// 		Bold(true)

// 	screen.Fill(' ', mainStyle)

// 	drawTextLine(screen, 0, 0, "Welcome to the Speedy Player!", mainStyle)
// 	drawTextLine(screen, 0, 1, "Press [ESC] to quit.", mainStyle)
// 	drawTextLine(screen, 0, 2, "Press [SPACE] to pause/resume.", mainStyle)
// 	drawTextLine(screen, 0, 3, "Use keys in (?/?) to turn the buttons.", mainStyle)

// 	speaker.Lock()
// 	position := ap.sampleRate.D(ap.streamer.Position())
// 	length := ap.sampleRate.D(ap.streamer.Len())
// 	volume := ap.volume.Volume
// 	speed := ap.resampler.Ratio()
// 	speaker.Unlock()

// 	positionStatus := fmt.Sprintf("%v / %v", position.Round(time.Second), length.Round(time.Second))
// 	volumeStatus := fmt.Sprintf("%.1f", volume)
// 	speedStatus := fmt.Sprintf("%.3fx", speed)

// 	drawTextLine(screen, 0, 5, "Position (Q/W):", mainStyle)
// 	drawTextLine(screen, 16, 5, positionStatus, statusStyle)

// 	drawTextLine(screen, 0, 6, "Volume   (A/S):", mainStyle)
// 	drawTextLine(screen, 16, 6, volumeStatus, statusStyle)

// 	drawTextLine(screen, 0, 7, "Speed    (Z/X):", mainStyle)
// 	drawTextLine(screen, 16, 7, speedStatus, statusStyle)
// }

// func (ap *audioPanel) handle(event tcell.Event) (changed, quit bool) {
// 	switch event := event.(type) {
// 	case *tcell.EventKey:
// 		if event.Key() == tcell.KeyESC {
// 			return false, true
// 		}

// 		if event.Key() != tcell.KeyRune {
// 			return false, false
// 		}

// 		switch unicode.ToLower(event.Rune()) {
// 		case ' ':
// 			speaker.Lock()
// 			ap.ctrl.Paused = !ap.ctrl.Paused
// 			speaker.Unlock()
// 			return false, false

// 		case 'q', 'w':
// 			speaker.Lock()
// 			newPos := ap.streamer.Position()
// 			if event.Rune() == 'q' {
// 				newPos -= ap.sampleRate.N(time.Second)
// 			}
// 			if event.Rune() == 'w' {
// 				newPos += ap.sampleRate.N(time.Second)
// 			}
// 			if newPos < 0 {
// 				newPos = 0
// 			}
// 			if newPos >= ap.streamer.Len() {
// 				newPos = ap.streamer.Len() - 1
// 			}
// 			if err := ap.streamer.Seek(newPos); err != nil {
// 				report(err)
// 			}
// 			speaker.Unlock()
// 			return true, false

// 		case 'a':
// 			speaker.Lock()
// 			ap.volume.Volume -= 0.1
// 			speaker.Unlock()
// 			return true, false

// 		case 's':
// 			speaker.Lock()
// 			ap.volume.Volume += 0.1
// 			speaker.Unlock()
// 			return true, false

// 		case 'z':
// 			speaker.Lock()
// 			ap.resampler.SetRatio(ap.resampler.Ratio() * 15 / 16)
// 			speaker.Unlock()
// 			return true, false

// 		case 'x':
// 			speaker.Lock()
// 			ap.resampler.SetRatio(ap.resampler.Ratio() * 16 / 15)
// 			speaker.Unlock()
// 			return true, false
// 		}
// 	}
// 	return false, false
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Fprintf(os.Stderr, "Usage: %s song.mp3\n", os.Args[0])
// 		os.Exit(1)
// 	}
// 	f, err := os.Open(os.Args[1])
// 	if err != nil {
// 		report(err)
// 	}
// 	streamer, format, err := mp3.Decode(f)
// 	if err != nil {
// 		report(err)
// 	}
// 	defer streamer.Close()

// 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))

// 	screen, err := tcell.NewScreen()
// 	if err != nil {
// 		report(err)
// 	}
// 	err = screen.Init()
// 	if err != nil {
// 		report(err)
// 	}
// 	defer screen.Fini()

// 	ap := newAudioPanel(format.SampleRate, streamer)

// 	screen.Clear()
// 	ap.draw(screen)
// 	screen.Show()

// 	ap.play()

// 	seconds := time.Tick(time.Second)
// 	events := make(chan tcell.Event)
// 	go func() {
// 		for {
// 			events <- screen.PollEvent()
// 		}
// 	}()

// loop:
// 	for {
// 		select {
// 		case event := <-events:
// 			changed, quit := ap.handle(event)
// 			if quit {
// 				break loop
// 			}
// 			if changed {
// 				screen.Clear()
// 				ap.draw(screen)
// 				screen.Show()
// 			}
// 		case <-seconds:
// 			screen.Clear()
// 			ap.draw(screen)
// 			screen.Show()
// 		}
// 	}
// }

// func report(err error) {
// 	fmt.Fprintln(os.Stderr, err)
// 	os.Exit(1)
// }
