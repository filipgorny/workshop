package views

type ContainerView struct {
	*TermContainer
	*TermView
}

func NewContainerView() *ContainerView {
	sv := ContainerView{
		TermView:      NewTermView(),
		TermContainer: NewTermContainer(),
	}

	return &sv
}
