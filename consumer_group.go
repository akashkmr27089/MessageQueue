package main

type CGSettings struct {
	NumOfConsumer int
	Id            int
}

type ConsumerGroup struct {
	Offset    map[int]int
	Id        int
	Consumers []*Consumer
}

func CreateConsumerGroup(id int, Partitions []*Partition, numberOfConsumers int) *ConsumerGroup {
	val := new(ConsumerGroup)
	val.Offset = make(map[int]int)
	val.Id = id

	for _, partition := range Partitions {
		val.Offset[(*partition).PartitionNumber.ToInt()] = 0 // TODO: Adjust with the last read offset
	}

	// This will create Consumers for the consumer group
	val.Consumers = CreateConsumers(numberOfConsumers)
	return val
}

func CreateConsumerGroups(settings []CGSettings, Partitions []*Partition) []*ConsumerGroup {
	var cgGroups []*ConsumerGroup
	for _, cgSetting := range settings {
		val := CreateConsumerGroup(cgSetting.Id, Partitions, cgSetting.NumOfConsumer)
		cgGroups = append(cgGroups, val)
	}

	return cgGroups
}
