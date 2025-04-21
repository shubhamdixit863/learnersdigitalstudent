package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
	"orderService/internal/handler"
	"orderService/internal/services"
)

const TOPIC = "logs"

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", TOPIC, 3)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	kafkaClient := services.KafkaClient{
		KafkaConn: conn,
	}

	err = kafkaClient.PublishMessages("hey there ")
	if err != nil {
		log.Println(err)
		return
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.POST("/createOrder", handler.CreateOrder)
	v1.POST("/deliverOrder", handler.DeliverOrder)

	err = r.Run("localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
