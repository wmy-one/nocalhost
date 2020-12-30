/*
Copyright 2020 The Nocalhost Authors.
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

package cmds

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"nocalhost/pkg/nhctl/log"
)

func init() {
	devBuildCmd.Flags().StringVarP(&deployment, "deployment", "d", "", "k8s deployment which your developing service exists")
	debugCmd.AddCommand(devBuildCmd)
}

var devBuildCmd = &cobra.Command{
	Use:   "build [NAME]",
	Short: "Build your code in dev container",
	Long:  `Build your code in dev container`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.Errorf("%q requires at least 1 argument\n", cmd.CommandPath())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//var err error
		applicationName := args[0]
		InitAppAndCheckIfSvcExist(applicationName, deployment)
		if !nocalhostApp.CheckIfSvcIsDeveloping(deployment) {
			log.Fatalf("%s is not in DevMode", deployment)
		}
		profile := nocalhostApp.GetSvcProfile(deployment)
		if profile == nil || len(profile.BuildCommand) == 0 {
			log.Fatal("build command not defined")
		}

		err := nocalhostApp.Exec(deployment, profile.BuildCommand)
		if err != nil {
			log.Fatalf("fail to exec : %s", err.Error())
		}
	},
}
