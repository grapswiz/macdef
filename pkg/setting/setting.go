package setting

import "github.com/BurntSushi/toml"

type Item struct {
	Name     string
	Category string
	Type     string
	Value    string
}

type Setting struct {
	Items []Item
}

func Get(fileName string) (Setting, error) {
	var s Setting
	_, err := toml.DecodeFile(fileName, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (s *Setting) Put(item Item) {
	for i, value := range s.Items {
		if value.Name == item.Name {
			s.Items[i] = item
			return
		}
	}
	s.Items = append(s.Items, item)
}