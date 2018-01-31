package setting

import "testing"

func TestPut(t *testing.T) {
	i := Item{
		Name:	"b",
	}
	s := Setting{
		Items: []Item{
			{Name:	"a", Value: "a"},
		},
	}
	if len(s.Items) != 1 {
		t.Errorf("前提条件ミス: %v", len(s.Items))
	}
	s.Put(i)
	if len(s.Items) != 2 {
		t.Errorf("nameが異なる場合は挿入されないといけない: %v", len(s.Items))
	}
	i.Name = "a"
	i.Value = "overwrite"
	s.Put(i)
	if s.Items[0].Name != "a" || s.Items[0].Value != "overwrite" {
		t.Errorf("nameが同一の場合はitemの値で上書きされないといけない")
	}
}