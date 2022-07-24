package views

type DrawContext struct {
  Draw func (target interface{})
}

func NewDrawContext() *DrawContext {
  return &DrawContext{
    Draw:  func (target interface{}) {
    switch (target.(type)) {
    case tview.Box:
      tview.Draw(target.(tview.Box))
      break;
    case tcell.Screen:
      tcell.Draw(target.(tcell.Screen))
      break;
    default:
      log.Printf("Unsupported target type: %T", target)
    }
  },
}
