package main

// getSampleMessage
type Message struct {
	Id           int
	TopicName    TopicName
	PartitionKey string // Generally this is taken from the Data itself. This is done to prevent un_marshaling of data in this case
	Data         string
}

func getSampleTopicsAndItsMapping() []Topic {
	topic := []Topic{
		{
			Name:                 "score",
			PartitionKey:         "match",
			AssociatedPartitions: []PartitionNumber{PartitionNumber(0), PartitionNumber(1)},
		},
		{
			Name:                 "playoff",
			PartitionKey:         "match",
			AssociatedPartitions: []PartitionNumber{PartitionNumber(0), PartitionNumber(1)},
		},
	}
	return topic
}

func getSampleMessageWithTopic() []Message {
	data := []Message{
		Message{
			Id:           1,
			TopicName:    "score",
			PartitionKey: "india_vs_pak",
			Data:         "23",
		},
		Message{
			Id:           2,
			TopicName:    "score",
			PartitionKey: "pak23",
			Data:         "1223",
		},
		Message{
			Id:           3,
			TopicName:    "playoff",
			PartitionKey: "india32",
			Data:         "321123",
		},
		Message{
			Id:           4,
			TopicName:    "score",
			PartitionKey: "sa123",
			Data:         "232",
		},
		Message{
			Id:           5,
			TopicName:    "playoff",
			PartitionKey: "aus432",
			Data:         "23123132",
		},
		Message{
			Id:           6,
			TopicName:    "score",
			PartitionKey: "aus432",
			Data:         "123456789",
		},
	}

	return data
}

func getSampleCGSettings() []CGSettings {
	return []CGSettings{
		{
			NumOfConsumer: 2,
			Id:            1,
		},
		{
			NumOfConsumer: 2,
			Id:            2,
		},
	}
}
