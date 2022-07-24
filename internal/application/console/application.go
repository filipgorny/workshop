package console

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/filipgorny/workshop/internal/application/console/views"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ConsoleApplication struct {
	application     *tview.Application
	applicationView *views.ApplicationView
	options         *ConsoleApplicationOptions
}

func NewConsoleApplication() *ConsoleApplication {
	tviewApplication := tview.NewApplication()

	applicationView := views.NewApplicationView()

	app := ConsoleApplication{
		application:     tviewApplication,
		applicationView: applicationView,
		options:         NewConsoleApplicationOptions(),
	}

	app.application.SetRoot(app.applicationView, true)

	return &app
}

func (a *ConsoleApplication) Run() {
	a.application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			a.Stop()
		}

		return event
	}).SetFocus(a.applicationView).Run()
}

func (a *ConsoleApplication) Stop() {
	a.application.Stop()
}

type ConsoleApplicationVersion struct {
	major int64
	minor int64
	patch int64
}

func (cav *ConsoleApplicationVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", cav.major, cav.minor, cav.patch)
}

func NewConsoleApplicationVersion(version string) *ConsoleApplicationVersion {
	segments := strings.Split(version, ".")

	segment := func(no int) int64 {
		segmentNo := 2 - no

		if segmentNo < 0 || segmentNo > len(segments) {
			return 0
		}

		number, _ := strconv.ParseInt(segments[segmentNo], 0, 64)

		return number
	}

	return &ConsoleApplicationVersion{
		major: segment(2),
		minor: segment(1),
		patch: segment(0),
	}
}

type ConsoleApplicationOptions struct {
	name    string
	version ConsoleApplicationVersion
}

func NewConsoleApplicationOptions() *ConsoleApplicationOptions {
	return &ConsoleApplicationOptions{
		name:    "Console Application",
		version: *NewConsoleApplicationVersion("0.0.0"),
	}
}
