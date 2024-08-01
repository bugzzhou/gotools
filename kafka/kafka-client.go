package kafka

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

// kafka consumer

func Get() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{"192.168.0.233:9092"}, config)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	defer consumer.Close()

	partitions, err := consumer.Partitions("web_log")
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	for partition := range partitions {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d, err: %v\n", partition, err)
			continue
		}
		defer pc.Close()

		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case msg := <-pc.Messages():
					fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				case <-pc.Errors():
					fmt.Println("Consumer error occurred")
				case <-time.After(10 * time.Second): // Timeout to handle no new messages
					fmt.Println("No new messages after 10 seconds")
				}
			}
		}(pc)
	}

	// Keep the main thread running
	select {}
}
