/*
Copyright © 2020 Rakesh Das <rakesh@rakeshdas.com>

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

	"github.com/spf13/cobra"
)

// NotificationTitle holds the notification title gotten from the flag
var NotificationTitle string

// sendNotificationCmd represents the sendNotification command
var sendNotificationCmd = &cobra.Command{
	Use:   "sendNotification",
	Short: "Send a notification to a device on your Join account",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sending a notification with the title of \"%s\"\n", args)
	},
}

func init() {
	rootCmd.AddCommand(sendNotificationCmd)
	sendNotificationCmd.Flags().StringVarP(&NotificationTitle, "notification-title", "nt", "join-cli", "Set the title of the notification")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendNotificationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendNotificationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
