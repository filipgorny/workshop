package style

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Color struct {
	value tcell.Color
}

func NewColor(color tcell.Color) Color {
	return Color{
		value: color,
	}
}

type ColorCreatingFunction func(...interface{}) Color

var CrateColor map[string]ColorCreatingFunction = map[string]ColorCreatingFunction{
	"RGB": func(values ...interface{}) Color {
		if len(values) != 3 {
			log.Panic("RGB color must have 3 values")
			return NewColor(tcell.ColorDefault)
		}

		r := values[0].(int32)
		g := values[1].(int32)
		b := values[2].(int32)

		return Color{
			value: tcell.NewRGBColor(r, g, b),
		}
	},
	"HEX": func(values ...interface{}) Color {
		if len(values) != 1 {
			log.Panic("HEX color must have 1 value")
			return NewColor(tcell.ColorDefault)
		}

		hex := values[0].(int32)

		return Color{
			value: tcell.NewHexColor(hex),
		}
	},
}

func (c *Color) GetValue() tcell.Color {
	return c.value
}

var ColorWhite = NewColor(tcell.ColorWhite)
var ColorBlack = NewColor(tcell.ColorBlack)
var ColorRed = NewColor(tcell.ColorRed)
var ColorGreen = NewColor(tcell.ColorGreen)
var ColorYellow = NewColor(tcell.ColorYellow)
var ColorBlue = NewColor(tcell.ColorBlue)
var ColorDarkGray = NewColor(tcell.ColorDarkGray)
var ColorLightGray = NewColor(tcell.ColorLightGray)
var ColorDarkRed = NewColor(tcell.ColorDarkRed)
var ColorDarkGreen = NewColor(tcell.ColorDarkGreen)
var ColorDarkBlue = NewColor(tcell.ColorDarkBlue)
var ColorDarkMagenta = NewColor(tcell.ColorDarkMagenta)
var ColorDarkCyan = NewColor(tcell.ColorDarkCyan)
