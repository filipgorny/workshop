package theme

import (
	"github.com/filipgorny/workshop/internal/application/console/style"
)

type ThemeValueCreator func(typeName string, value interface{}) ThemeValue

type ThemeValueFactory func(values ...interface{}) ThemeValue

type Theme struct {
	Name            string
	BackgroundColor style.Color
	ForegroundColor style.Color
	HighlightColor  style.Color
	BorderColor     style.Color
	BorderSize      int
}
