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
package cmd

import (
	"fmt"
	"multiTopic/config"
	"multiTopic/consumer"

	"github.com/spf13/cobra"
)

// oneConsumerCopyCmd represents the oneConsumerCopy command
var oneConsumerCopyCmd = &cobra.Command{
	Use:   "oneConsumerCopy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oneConsumerCopy called")
		go consumer.NewConsumerCopy(config.TopicTest, config.Earliest).Listen()
		go consumer.NewConsumerCopy(config.TopicTest1, config.Latest).Listen()
		go consumer.NewConsumerCopy(config.TopicTest2, config.Earliest).Listen()
		go consumer.NewConsumerCopy(config.TopicTest3, config.Latest).Listen()
		select {}
	},
}

func init() {
	rootCmd.AddCommand(oneConsumerCopyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oneConsumerCopyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oneConsumerCopyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
./bin/kafka-topics.sh --broker-list 127.0.0.1:909 create test2
./bin/kafka-topics.sh --broker-list 127.0.0.1:909 create test3
*/
