package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	// brokers = "kfk1-stg.myhll.cn:9092,kfk2-stg.myhll.cn:9092,kfk3-stg.myhll.cn:9092" // 填写你的broker地址
	brokers = "kfk1-pre.myhll.cn:9092,kfk2-pre.myhll.cn:9092,kfk3-pre.myhll.cn:9092" // 填写你的broker地址
	group   = "ryan-example-group"                                                   // 填写你的consumer group id
	// topics  = "ios_egp_ysh_asr_result"                                               // 填写你要消费的topic
	topics = "vn_ticket_voice" // 填写你要消费的topic
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, offset=%d, partition=%d",
			string(msg.Value), msg.Timestamp, msg.Topic, msg.Offset, msg.Partition)
		session.MarkMessage(msg, "")
	}

	return nil
}

func main() {
	// Set up a new Sarama configuration.
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2 // 确保Kafka集群使用的版本
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Start with a context to manage cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new consumer group
	consumer := &Consumer{}

	// Initialize the consumer group
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}

	// Capture interrupt signals to release resources gracefully
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)

	// Start the consumer loop
	go func() {
		for {
			if err := client.Consume(ctx, strings.Split(topics, ","), consumer); err != nil {
				log.Fatalf("Error during consuming: %v", err)
			}
			// Check if context was cancelled, signaling the end of the consumer loop
			if ctx.Err() != nil {
				return
			}
		}
	}()

	log.Println("Consumer up and running!...")
	<-sigterm
	log.Println("Interrupt is detected. Exiting...")

	// Ensure that the consumer group is closed gracefully
	if err = client.Close(); err != nil {
		log.Fatalf("Error closing client: %v", err)
	}
}
