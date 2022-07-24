package content

type Text struct {
	Content
	Load func(string) (string, error)
}
