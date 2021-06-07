package gui

type Box struct {
	x int
	y int
	w int
	h int
}

func NewBox() Box {
	return Box{
		x: 0,
		y: 0,
		w: 0,
		h: 0,
	}
}

func (b *Box) MoveRight() {
	b.x++
}
