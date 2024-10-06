package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	// Take the Message based on Topic, Message Data, Consumer

	/*
		Kafka Producer
		Initialize
		Map topic, partition key, and which all partition where you can send the data
		Also write the logic passing data to partion

		- Processing the incoming message : Server to listen the message
		- Based on Partition Key, and Rule Engine --> Pass the value to a specific Partition
	*/

	/*
		Consumer and Worker
		-- Consumer the Data from the partition
		-- Find Ways to make the data persistent as well
		-- Consume and do a certain function --> Make a function class later on
	*/

	PARTITION_COUNT := 2

	topics := getSampleTopicsAndItsMapping()
	topicPartitionMap := make(map[TopicName][]PartitionNumber)
	for _, val := range topics {
		topicPartitionMap[val.Name] = val.AssociatedPartitions
	}
	messages := getSampleMessageWithTopic() // Get Sample Message Data
	partitions, partitionChannels := createPartitions(PARTITION_COUNT)

	//// Getting Messages, passing it to concerned partition
	for _, message := range messages {
		partitionNumber := getPartitionNumberByPartitionKey(message.PartitionKey, topicPartitionMap[message.TopicName])
		partitionChannels[partitionNumber] <- message.Data
	}

	// Start Process for Consumer
	func(partitions []*Partition) {
		cgSettings := getSampleCGSettings()
		consumerGroups := CreateConsumerGroups(cgSettings, partitions)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		g, newCtx := errgroup.WithContext(ctx)
		for _, partition := range partitions {
			for _, cgGroup := range consumerGroups {
				for _, val := range cgGroup.Consumers {
					consumer := *val
					g.Go(
						func() error {
							return consumer.Exec(newCtx, cgGroup, partition)
						},
					)
				}
			}
		}

		if err := g.Wait(); err != nil {
			log.Fatalln("error while processing:", err)
		}

		fmt.Println("All workers have finished processing.")
	}(partitions)
}
