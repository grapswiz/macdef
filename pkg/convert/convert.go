package convert

import (
	"github.com/grapswiz/macdef/pkg/setting"
	"github.com/grapswiz/macdef/pkg/definition"
)

// TODO
func SettingToDefinition(setting setting.Setting) (def definition.Definition) {
	return definition.Definition{
		Items: []definition.Item{
			{
				Name: "",
				Category: "",
				Type: "",
				Value: "",
			},
		},
	}
}

// TODO
//func DefinitionsToBashScript(defs []definition.Definition) (script string) {
//}
