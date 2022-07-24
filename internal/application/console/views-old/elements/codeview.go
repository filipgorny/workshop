package elements

import (
	"github.com/filipgorny/workshop/internal/application/console/views"
	"github.com/filipgorny/workshop/internal/project"
	"github.com/gdamore/tcell/v2"
)

type StyleType string

const (
	STYLE_DEFAULT StyleType = "default"
	STYLE_KEYWORD StyleType = "keyword"
)

type CodeView struct {
	sourceCode *project.SourceCode
	style      tcell.Style

	*views.TermView
}

func NewCodeView() *CodeView {
	cv := CodeView{
		sourceCode: nil,
		style:      tcell.StyleDefault,
		TermView:   views.NewTermView(),
	}
	return &cv
}

func (cv *CodeView) SetSourceCode(sourceCode *project.SourceCode) {
	cv.sourceCode = sourceCode
}

func (cv *CodeView) GetRuneAt(x, y int) rune {
	return cv.sourceCode.GetRuneAt(x, y)
}

func (cv *CodeView) GetStyle(name string) tcell.Style {
	return cv.style
}

func (cv *CodeView) Draw(screen tcell.Screen) {
	x, y, width, height := cv.GetRect()

	for ix := x; ix < width; ix++ {
		for iy := y; iy < height; iy++ {
			screen.SetCell(ix, iy, cv.GetStyle("default"), cv.GetRuneAt(ix, iy))
		}
	}
}
