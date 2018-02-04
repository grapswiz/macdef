// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"github.com/grapswiz/macdef/pkg/setting"
	"github.com/grapswiz/macdef/pkg/util"
	"github.com/grapswiz/macdef/pkg/definition"
	"strings"
	"bytes"
	"fmt"
)

// setCmd represents the set command
// .macdef/macdef.toml があればそこに上書き、無ければ作成
var setCmd = &cobra.Command{
	Use:   "set [KEY] [VALUE]",
	Short: "Write a setting to $HOME/.macdef/macdef.toml",
	Long: `Write a setting to $HOME/.macdef/macdef.toml.
This command just for setting. So you want to reflect,
you should use apply command.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := homedir.Dir()
		file := filepath.Join(home, ".macdef", "macdef.toml")
		if !util.ExistsFile(file) {
			ioutil.WriteFile(file, []byte(""), os.ModePerm)
		}
		s := setting.Setting{
			Items: map[string]setting.Item{},
		}
		_, _ = toml.DecodeFile(file, &s)
		item, err := definition.GetItem(args[0])
		if err != nil {
			fmt.Printf("%v", err)
		}
		split := strings.Split(args[0], ".")
		s.Items[split[1]] = setting.Item{
			Category: split[0],
			Value: args[1],
			Type: item.Type,
		}
		var buffer bytes.Buffer
		encoder := toml.NewEncoder(&buffer)
		_ = encoder.Encode(s)
		ioutil.WriteFile(file, buffer.Bytes(), os.ModePerm)
	},
}

func init() {
	RootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
