package definition

import (
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"strings"
	"github.com/grapswiz/macdef/pkg/util"
	"github.com/BurntSushi/toml"
	"os/exec"
	"github.com/pkg/errors"
	"github.com/k0kubun/pp"
)

type Item struct {
	Description string
	Type     string
	Values	[]string
	Commands	[]string
}

type Definition struct {
	Items map[string]Item
}

func GetItem(key string) (item Item, err error) {
	var i Item
	split := strings.Split(key, ".")
	home, err := homedir.Dir()
	if err != nil {
		return i, errors.Cause(err)
	}
	version, err := exec.Command("sw_vers", "-productVersion").CombinedOutput()
	if err != nil {
		return i, errors.Cause(err)
	}
	major := strings.Join(strings.Split(string(version), ".")[:2], ".")
	fileName := filepath.Join(home, ".macdef", "repo", "definitions", major, split[0]+".toml")
	if !util.ExistsFile(fileName) {
		return i, errors.New("path: " + fileName +  " is not found")
	}
	d := Definition{
		map[string]Item{},
	}
	meta, err := toml.DecodeFile(fileName, &d) // TODO デコードに失敗してる？？
	pp.Print(meta)
	value, ok := d.Items[split[1]]
	if ok {
		return value, nil
	}
	return i, errors.New("definition " + key + " is not found")
}
