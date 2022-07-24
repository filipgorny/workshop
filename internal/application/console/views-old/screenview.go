package views 

type ScreenView struct { 
  name string *TermContainer
  *FullScreenView
}

func NewScreenView(name string) *ScreenView {
	return &ScreenView{
		name:           name,
		TermContainer:  NewTermContainer(),
		FullScreenView: NewFullScreenView(),
	}
}

func (sv *ScreenView) GetName() string {
	return sv.name
}
