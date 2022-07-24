package views

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ApplicationView struct {
	screens *ApplicationScreens

	*FullScreenView
}

type ApplicationScreens struct {
	screens []*ScreenView
	active  string
}

func (as *ApplicationScreens) AddScreen(screen *ScreenView) {
	as.screens = append(as.screens, screen)
}

func (av *ApplicationView) GetActiveScreen() *ScreenView {
	for _, screen := range av.screens.screens {
		if screen.GetName() == av.screens.active {
			return screen
		}
	}

	return nil
}

func (av *ApplicationView) Draw(screen tcell.Screen) {
	activeScreen := av.GetActiveScreen()

	screen.Fill('R', tcell.StyleDefault)
	activeScreen.SetBackgroundColor(tcell.ColorBlue)

	activeScreen.Draw(screen)
}

func NewApplicationView() *ApplicationView {
	defaultScreen := NewScreenView("default")
	defaultScreen.AddChild(tview.NewTextView().SetText("default"))

	application := ApplicationView{
		screens: &ApplicationScreens{
			screens: []*ScreenView{},
			active:  "default",
		},
		FullScreenView: NewFullScreenView(),
	}

	application.screens.screens = make([]*ScreenView, 0)

	application.screens.AddScreen(defaultScreen)

	return &application
}

func (av *ApplicationView) Run() {
	log.Println("ApplicationView.Run()")
}
