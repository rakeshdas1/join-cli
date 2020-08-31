/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"join-cli/models"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listDevicesCmd represents the listDevices command
var listDevicesCmd = &cobra.Command{
	Use:   "listDevices",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listDevices called")
		var joinClient models.JoinAPIClient
		joinClient.BaseURL = "https://joinjoaomgcd.appspot.com/"
		joinClient.APIKey = viper.GetString("api-key")
		joinClient.NewHTTPClient()
		resp, _ := joinClient.GetAllDevices()
		fmt.Printf("%+v\n", resp)
		numOfDevices := len(resp.Records)
		if numOfDevices <= 0 {
			fmt.Printf("Didn't find any devices on this account!\n")
		} else if numOfDevices == 1 {
			fmt.Printf("Found 1 device on this account: \n")
		} else {
			fmt.Printf("Found %d devices on this account: \n", numOfDevices)
		}

		for _, device := range resp.Records {
			fmt.Printf("\t - %s\n", device.DeviceName)
		}
	},
}

func init() {
	rootCmd.AddCommand(listDevicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDevicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDevicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
