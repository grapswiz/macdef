package setting

type Item struct {
	Category string
	Type     string
	Value    string
}

type Setting struct {
	Items map[string]Item
}
