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
	"gopkg.in/src-d/go-git.v4"
	"os"
	"os/user"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates definitions",
	Long:  `Updates definitions`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
		dir := usr.HomeDir + "/.dfs"
		if _, err = os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, 0755)
		}
		repo := dir + "/repo"
		if _, err = os.Stat(repo); os.IsNotExist(err) {
			os.Mkdir(repo, 0755)
		}
		var r *git.Repository
		if r, err = git.PlainClone(repo, false, &git.CloneOptions{
			URL:      "https://github.com/grapswiz/dfs",
			Progress: os.Stdout,
		}); err == git.ErrRepositoryAlreadyExists {
			r, err = git.PlainOpen(repo)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		w, err := r.Worktree()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
