// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the local appsody environment for your project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		Info.log("Stopping development environment")
		projectDir := getProjectDir()
		imageName := fmt.Sprintf("%s-dev", filepath.Base(projectDir))

		dockerStop(imageName)
		//dockerRemove(imageName) is not needed due to --rm flag
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

}