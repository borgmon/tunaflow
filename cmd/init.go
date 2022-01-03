/*
Copyright © 2022 borgmon

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/borgmon/tunaflow/assets"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initiate a new Tunaflow project",
	Long:  `Initiate a new Tunaflow project. This will give you a new YAML config to start with`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t := &assets.Assets{
			FileName: "app.yaml",
			OutPath:  "./",
			Data: &map[string]string{
				"AppName": args[0],
			},
		}
		return t.Generate()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
