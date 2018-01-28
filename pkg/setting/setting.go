package setting

type Item struct {
	Name     string
	Category string
	Type     string
	Value    string
}

type Setting struct {
	Items []Item
}

func Put(s *Setting, item Item) {
	for i, value := range s.Items {
		if value.Name == item.Name {
			s.Items[i] = item
			return
		}
		s.Items = append(s.Items, item)
	}
}