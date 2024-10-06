package main

type TopicName string

type Topic struct {
	Name                 TopicName
	PartitionKey         string
	AssociatedPartitions []PartitionNumber
}
