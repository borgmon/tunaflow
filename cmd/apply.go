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
	"fmt"
	"os"
	"path/filepath"

	"github.com/borgmon/tunaflow/generator"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply your Tunaflow config",
	Run: func(cmd *cobra.Command, args []string) {
		configFileName := "app.yaml"
		if configOverride := cmd.Flag("config").Value; configOverride.String() != "" {
			configFileName = configOverride.String()
		}

		configD, err := os.ReadFile(filepath.Join(generator.GetPwd(), configFileName))
		CheckErr(err)
		config := &generator.AppConfig{}
		err = yaml.Unmarshal(configD, config)
		CheckErr(err)

		g := generator.Generator{
			Config: config,
		}

		err = g.GenerateBase()
		CheckErr(err)

		err = g.GenerateSchema()
		CheckErr(err)

		err = g.GenerateTransformer()
		CheckErr(err)

		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringP("config", "c", "", "specify a config file(other than app.config)")
}
