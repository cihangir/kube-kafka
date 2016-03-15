package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	kafka "github.com/Shopify/sarama"
)

func main() {

	producer, err := kafka.NewAsyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, errors int
ProducerLoop:
	for {
		select {
		case producer.Input() <- &kafka.ProducerMessage{
			Topic: "my_topic",
			Key:   nil,
			Value: kafka.StringEncoder(fmt.Sprintf("enqueued: %d", enqueued)),
		}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		case <-signals:
			break ProducerLoop
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
