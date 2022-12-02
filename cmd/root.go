/*
Copyright © 2022 Omer Faruk Urhan

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
	"cflogger/pkg"
	"errors"
	"os"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

var CustomTime string
var Timeout string
var StackName string
var Region string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cflogger",
	Short: "CloudFormation Logger App",
	Long: `This application display CloudFormation stack events and events details.

	Find more information at: https://github.com/omerurhan/cflogger
	

	`,
	Args: func(cmd *cobra.Command, args []string) error {
		// cobra validation
		//var err error
		//if err = cobra.MaximumNArgs(0)(cmd, args); err != nil {
		//	return err
		//}
		var err error
		var i int
		// Time validation
		if CustomTime != "" {
			err = pkg.GetTime(CustomTime)
			if err != nil {
				return err
			}
		}
		// Timeout validation
		i, err = strconv.Atoi(Timeout)
		if err != nil {
			return err
		}
		pkg.GetTimeout(i)

		// StackName Validation
		if StackName == "" {
			err = errors.New("--stack-name can not be empty string.")
			return err
		}
		pkg.GetData(StackName)

		//// Region Validation
		match, err := regexp.MatchString(("(us|ap|ca|cn|eu|sa|me|af)-(central|(north|south)?(east|west)?)-(1|2|3)"), Region)
		if err != nil {
			return err
		}
		if !match || len(Region) == 0 {
			err = errors.New("--region name not valid.")
			return err
		}
		pkg.GetRegion(Region)

		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		pkg.Start()
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&CustomTime, "since", "", "Get events since yyyy-mm-dd hh:mm")
	rootCmd.PersistentFlags().StringVar(&Timeout, "timeout", "15", "Log trace timeout in minutes. Default 15 minute.")
	rootCmd.PersistentFlags().StringVar(&StackName, "stack-name", "", "AWS Cloudformation stack name or stack id. If you set \"-\", input waiting from stdin. (required)")
	rootCmd.MarkPersistentFlagRequired("stack-name")
	rootCmd.PersistentFlags().StringVar(&Region, "region", "", "AWS region. (required)")
	rootCmd.MarkPersistentFlagRequired("region")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
