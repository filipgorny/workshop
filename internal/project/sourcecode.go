package project

import "strings"

type SourceCode struct {
	raw string
}

func NewSourceCode(raw string) *SourceCode {
	return &SourceCode{
		raw: raw,
	}
}

type SourceCodeLine struct {
	content string
	number  int
}

type SourceCodeLines struct {
	lines []SourceCodeLine
}

func (sc *SourceCode) GetLines() SourceCodeLines {
	lines := SourceCodeLines{
		lines: make([]SourceCodeLine, 0),
	}

	for number, line := range strings.Split(sc.raw, "\n") {
		lines.lines = append(lines.lines, SourceCodeLine{content: line, number: number})
	}

	return lines
}

func (sc *SourceCode) GetRuneAt(x, y int) rune {
	lines := sc.GetLines()

	if y >= len(lines.lines) {
		return ' '
	}

	if x >= len(lines.lines[y].content) {
		return ' '
	}

	return rune(sc.GetLines().lines[y].content[x])
}
