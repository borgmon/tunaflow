/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/borgmon/tunaflow/cmd"
	"github.com/borgmon/tunaflow/generator"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func main1() {
	cmd.Execute()
}
func check(e error) {
	if e != nil {
		fmt.Printf("%+v", e)
	}
}

func main() {
	basePath := "example"
	err := clean(basePath)
	check(err)

	configD, err := os.ReadFile(basePath + "/app.yaml")
	check(err)
	config := &generator.AppConfig{}
	err = yaml.Unmarshal(configD, config)
	check(err)

	g := generator.Generator{
		Config:   config,
		BasePath: basePath,
	}
	err = g.Prepare()
	check(err)

	err = g.GenTemplate()
	check(err)

	log.Println("Done")
}

func clean(basePath string) error {
	if err := os.RemoveAll(basePath + "/build"); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
