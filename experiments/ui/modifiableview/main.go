package main

import (
	"log"

	"github.com/baldore/angelo/modifiableview/gui"
)

func main() {
	ui := gui.NewUI()
	if err := ui.Start(); err != nil {
		log.Panicln(err)
	}
}
