package views

import "github.com/gdamore/tcell/v2"

type TermContainer struct {
	children []*TermContainerChild
}

type TermContainerChild struct {
	view TermPrimitive
}

func NewTermContainer() *TermContainer {
	termContainer := TermContainer{
		children: make([]*TermContainerChild, 0),
	}

	return &termContainer
}

func NewTermContainerChild(view TermPrimitive) *TermContainerChild {
	return &TermContainerChild{
		view: view,
	}
}

func (tc *TermContainer) AddContainerChild(child *TermContainerChild) {
	tc.children = append(tc.children, child)
}

func (tc *TermContainer) AddChild(view TermPrimitive) {
	tc.AddContainerChild(NewTermContainerChild(view))
}

func (tc *TermContainer) Draw(screen tcell.Screen) {
	for _, child := range tc.children {
		child.view.Draw(screen)
	}
}

func (tc *TermContainer) GetChildren() []TermPrimitive {
	children := []TermPrimitive{}

	for _, child := range tc.children {
		children = append(children, child.view)
	}

	return children
}
