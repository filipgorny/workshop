package views

import "github.com/rivo/tview"

type Size struct {
	Width, Height int
}

type Position struct {
	X, Y int
}

type TermBox struct {
	position *Position
	size     *Size

	*tview.Box
}

type TermView struct {
	*TermBox
}

type FullScreenView struct {
	*TermBox
}

func NewTermBox() *TermBox {
	box := tview.NewBox()

	tb := TermBox{
		position: &Position{0, 0},
		size:     &Size{64, 64},
		Box:      box,
	}

	return &tb
}

func NewTermView() *TermView {
	tv := TermView{
		TermBox: NewTermBox(),
	}

	return &tv
}

func NewFullScreenView() *FullScreenView {
	fsv := FullScreenView{
		TermBox: NewTermBox(),
	}

	return &fsv
}

func (tv *TermView) Resize(width, height int) {
	tv.size = &Size{width, height}
}

func (tv *TermView) SetPosition(x, y int) {
	tv.position = &Position{x, y}
}

func (tv *TermView) GetPosition() *Position {
	return tv.position
}

func (tv *TermView) GetSize() *Size {
	return tv.size
}
