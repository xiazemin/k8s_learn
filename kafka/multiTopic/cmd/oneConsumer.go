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

// oneConsumerCmd represents the oneConsumer command
var oneConsumerCmd = &cobra.Command{
	Use:   "oneConsumer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("oneConsumer called")
		go consumer.NewConsumer(config.TopicTest).Listen()
		go consumer.NewConsumer(config.TopicTest1).Listen() //永远消费不到消息
		//* Call `.Subscribe()` or (`.SubscribeTopics()` to subscribe to multiple topics) to join the group with the specified subscription set. Subscriptions are atomic, calling `.Subscribe*()` again will leave the group and rejoin with the new set of topics.
		//https://pkg.go.dev/github.com/confluentinc/confluent-kafka-go/kafka#section-readme

		select {}
	},
}

func init() {
	rootCmd.AddCommand(oneConsumerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oneConsumerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oneConsumerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
cd software/apache-zookeeper-3.6.2-bin
./bin/zkServer.sh

 cd software/kafka_2.12-2.7.0
 ./bin/kafka-server-start.sh config/server.properties

 ./bin/kafka-console-producer.sh --topic test --broker-list 127.0.0.1:9092
 ./bin/kafka-console-producer.sh --topic test1 --broker-list 127.0.0.1:9092

*/
