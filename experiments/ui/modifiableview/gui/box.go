package gui

type Box struct {
	x int
	y int
	w int
	h int
}

func NewBox() *Box {
	return &Box{
		x: 0,
		y: 0,
		w: 1,
		h: 1,
	}
}

func (b *Box) MoveLeft() {
	if b.x > 0 {
		b.x--
	}
}

func (b *Box) MoveRight() {
	b.x++
}

func (b *Box) MoveUp() {
	if b.y > 0 {
		b.y--
	}
}

func (b *Box) MoveDown() {
	b.y++
}
