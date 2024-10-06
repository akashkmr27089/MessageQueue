package main

import (
	"fmt"
	"sync"
)

type PartitionNumber int

func (p PartitionNumber) ToInt() int {
	return int(p)
}

type Partition struct {
	PartitionNumber PartitionNumber
	Val             []any
	lock            sync.Mutex
	ReachedEnd      *sync.Cond
}

func (p *Partition) PartitionConstructor(index PartitionNumber) *Partition {
	p.Val = make([]any, 0)
	p.ReachedEnd = sync.NewCond(&p.lock)
	p.PartitionNumber = index
	return p
}

func (p *Partition) Producer(data chan string) {
	// Listener for each partition should keep listening for keeping the lock intact
	for i := 1; ; i++ {
		p.lock.Lock()
		select {
		case value := <-data:
			p.Val = append(p.Val, fmt.Sprintf("Partition :%d, message %d, value %s", p.PartitionNumber, i, value))
		default:
			if p.lock.TryLock() {
				p.lock.Unlock()
			}
		}
		// Wake up all waiting workers
		p.ReachedEnd.Signal() //: Try Broadcast()
		p.lock.Unlock()
	}
}

func CreatePartitions(count int) []*Partition {
	var partitions []*Partition
	for i := 0; i < count; i += 1 {
		partition := new(Partition)
		partition.PartitionConstructor(PartitionNumber(i))
		partitions = append(partitions, partition)
	}

	return partitions
}

func createPartitions(count int) ([]*Partition, []chan string) {
	partitions := CreatePartitions(count)
	partitionChans := []chan string{}
	for i := 0; i < count; i += 1 {
		val := make(chan string)
		partitionChans = append(partitionChans, val)
		go partitions[i].Producer(val)
	}

	return partitions, partitionChans
}

//func (p *Partition) TestProducer() {
//	// Listener for each partition should keep listening for keeping the lock intact
//	for i := 1; ; i++ {
//		p.lock.Lock()
//		time.Sleep(1 * time.Millisecond)
//
//		// Append the new value
//		if p.PartitionNumber%2 == 0 {
//			i += 1000
//		}
//		p.Val = append(p.Val, fmt.Sprintf("Partition :%d, message %d", p.PartitionNumber, i))
//
//		// Wake up all waiting workers
//		p.ReachedEnd.Signal() //: Try Broadcast()
//		p.lock.Unlock()
//	}
//}
