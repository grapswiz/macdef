package convert

import (
	"github.com/grapswiz/macdef/pkg/setting"
	"github.com/grapswiz/macdef/pkg/definition"
	"strings"
	"github.com/k0kubun/pp"
	"text/template"
	"bytes"
)

func SettingToBashScript(setting setting.Setting) string {
	var commands []string
	for _, sItem := range setting.Items {
		def, err := definition.GetByCategory(sItem.Category)
		if err != nil {
			continue
		}
		for _, dItem := range def.Items {
			if sItem.Name == dItem.Name {
				var kore []string
				for _, command := range dItem.Commands {
					tpl, err := template.New("mytemplate").Parse(command)
					if err != nil {
						pp.Errorf("%v", err)
						continue
					}
					buf := new(bytes.Buffer)
					tpl.Execute(buf, map[string]interface{}{
						"Value": sItem.Value,
					})
					kore = append(kore, buf.String())
				}
				commands = append(commands, strings.Join(kore, "\n"))
			}
		}
	}
	commands = append(commands, "killall Dock")
	return strings.Join(commands, "\n")
}
