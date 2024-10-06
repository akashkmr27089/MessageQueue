package main

import "hash/fnv"

func stringToInt64(s string) int64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(s))
	return int64(hasher.Sum64() & 0x7FFFFFFFFFFFFFFF)
}

// This is for getting the partition number by partition key
func getPartitionNumberByPartitionKey(val string, partitions []PartitionNumber) PartitionNumber {
	lenOfPartition := len(partitions)
	index := stringToInt64(val) % int64(lenOfPartition)
	return partitions[index]
}
