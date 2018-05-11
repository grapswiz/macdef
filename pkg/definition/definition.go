package definition

import (
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"fmt"
	"os"
	"strings"
	"github.com/grapswiz/macdef/pkg/util"
	"github.com/BurntSushi/toml"
	"os/exec"
)

type Item struct {
	Name     string
	Description	string
	Type	string
	Values	[]string
	Commands	[]string
}

type Definition struct {
	Items []Item
}

func GetByCategory(category string) (Definition, error) {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var d Definition
	version, err := exec.Command("sw_vers", "-productVersion").CombinedOutput()
	if err != nil {
		return d, err
	}
	major := strings.Join(strings.Split(string(version), ".")[:2], ".")
	fileName := filepath.Join(home, ".macdef", "repo", "definitions", major, category + ".toml")
	if !util.ExistsFile(fileName) {
		return d, fmt.Errorf("path: %s is not found", fileName)
	}
	_, err = toml.DecodeFile(fileName, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func GetItem(key string) (item Item, err error) {
	split := strings.Split(key, ".")
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var i Item
	version, err := exec.Command("sw_vers", "-productVersion").CombinedOutput()
	if err != nil {
		return i, err
	}
	major := strings.Join(strings.Split(string(version), ".")[:2], ".")
	fileName := filepath.Join(home, ".macdef", "repo", "definitions", major, split[0]+".toml")
	var d Definition
	if !util.ExistsFile(fileName) {
		return i, fmt.Errorf("path: %s is not found", fileName)
	}
	_, err = toml.DecodeFile(fileName, &d)
	for _, value := range d.Items {
		if value.Name == split[1] {
			return value, nil
		}
	}
	return i, fmt.Errorf("definition %s is not found", key)
}
