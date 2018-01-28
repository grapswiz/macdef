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
	"fmt"

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
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file := filepath.Join(home, ".macdef", "macdef.toml")
		if !util.ExistsFile(file) {
			ioutil.WriteFile(file, []byte(""), os.ModePerm)
		}
		var s setting.Setting
		_, err = toml.DecodeFile(file, &s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		item, err := definition.GetItem(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		split := strings.Split(args[0], ".")
		setting.Put(&s, setting.Item{
			Name:     split[1],
			Category: split[0],
			Value:    args[1],
			Type:     item.Type,
		})
		var buffer bytes.Buffer
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		encoder := toml.NewEncoder(&buffer)
		err = encoder.Encode(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(file, buffer.Bytes(), os.ModePerm)
		// FIXME ファイルが存在しないとき例えば set dock.showhidden true した場合 macdef.toml が空になってしまっている
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
