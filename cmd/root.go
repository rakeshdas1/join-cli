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
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// APIKey holds the join api key to authenticate with the join api
	APIKey string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "join-cli",
	Short: "A CLI application to interact with your devices registered on Join",
	Long: `join-cli is a small utility to interact with your devices such as your other computers,
	phones, tablets, etc directly from the command line. You can use this to send a notification to your phone whenever a long 
	running task finally finishes running, such as a big download or a complex build.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.join-cli.json)")
	rootCmd.PersistentFlags().StringVarP(&APIKey, "api-key", "k", "", "Join api key to authenticate with the join api (required)")
	// look for api key in the config file
	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".join-cli" (without extension).
		viper.AddConfigPath(home + "/.config/join-cli")
		viper.SetConfigName("join-cli")
		viper.SetConfigType("json")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Did not find a config file, writing a new one...")
		// viper.WriteConfigAs("~/.config/join-cli/join-cli.json")
		viper.SafeWriteConfig()
	}
	fmt.Printf("Got api key of %s\n", viper.GetString("api-key"))
}
