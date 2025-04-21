package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/segmentio/kafka-go"
	"log"
	"microservices/internal/services"

	"sync"
)

func connectKafka() (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", TOPIC, 0)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func rabbitMqConnect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	return conn, nil

}

const TOPIC = "logs"

func main() {
	var wg sync.WaitGroup

	var broker services.Broker

	//for kafka
	//broker = services.NewKafkaClient(conn)
	//conn, err := connectKafka()
	//if err != nil {
	//	return
	//}

	// for rabbit mq
	connect, err := rabbitMqConnect()
	if err != nil {
		log.Fatalf("Error connecting with rabbit mq %s", err)
		return
	}
	broker = services.NewRabbitMqClient(connect)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			err := broker.PublishMessages("This is a published message")
			if err != nil {
				log.Println(err)
				return
			}
		}

	}()

	go func() {
		defer wg.Done()
		for {
			err := broker.ConsumeMessages()
			if err != nil {
				return
			}

		}

	}()

	wg.Wait()
	//err = broker.ConsumeMessages()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
}
