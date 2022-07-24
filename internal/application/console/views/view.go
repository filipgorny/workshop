package views

type VarState string

type View struct {
	State VarState
	Draw  *DrawContext
}
