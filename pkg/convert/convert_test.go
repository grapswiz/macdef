package convert

import (
	"github.com/grapswiz/macdef/pkg/setting"
	"testing"
)

func TestSettingToDefinition(t *testing.T) {
	s := setting.Setting{
		Items: []setting.Item{
			{
				Name: "showhidden",
				Category: "dock",
				Type: "bool",
				Value: "true",
			},
			{
				Name: "orientation",
				Category: "dock",
				Type: "string",
				Value: "bottom",
			},
		},
	}
	def := SettingToBashScript(s)
	if len(def.Items) != 2 {
		t.Errorf("parse error size: %d, expect: 2", len(def.Items))
		return
	}
	if def.Items[0].Name != s.Items[0].Name {
		t.Errorf("expect: %s, actual: %s", def.Items[0].Name, s.Items[0].Name)
	}
	if def.Items[1].Name != s.Items[1].Name {
		t.Errorf("expect: %s, actual: %s", def.Items[1].Name, s.Items[1].Name)
	}
}