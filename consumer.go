package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Consumer struct {
	Id int
}

func (w *Consumer) Exec(ctx context.Context, consumerGroup *ConsumerGroup, p *Partition) error {
	for {
		select {
		case <-ctx.Done(): // Terminate worker if the context is canceled
			return ctx.Err()
		default:
			time.Sleep(1 * time.Second)
			p.lock.Lock()
			offset, ok := consumerGroup.Offset[p.PartitionNumber.ToInt()]
			if !ok {
				// Offset not found
				log.Fatalf("offset for partition %d not found. Please restart", p.PartitionNumber)
				break
			}
			for offset >= len(p.Val) { // Wait if no data is available
				p.ReachedEnd.Wait() // Wait for the producer to signal new data
			}

			// Process the next value
			message := p.Val[offset]
			fmt.Printf("received for consumer group %d consumer :%d with data :%s\n", consumerGroup.Id, w.Id, message)

			// Process it --> if fails --> Put it to dlq
			/*
				Replace it with the logic of what to do after the message is recived
			*/

			(*consumerGroup).Offset[p.PartitionNumber.ToInt()]++ // Move to the next value
			p.lock.Unlock()
		}
	}
}

func CreateConsumers(count int) []*Consumer {
	var Workers []*Consumer
	for i := 0; i < count; i++ {
		val := new(Consumer)
		val.Id = i
		Workers = append(Workers, val)
	}
	return Workers
}
