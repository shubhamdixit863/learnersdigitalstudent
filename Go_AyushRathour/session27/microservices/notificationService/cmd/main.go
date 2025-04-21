package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
	"notificationService/internal/handler"
	"notificationService/internal/services"
)

const TOPIC = "logs"

func main() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", TOPIC, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	kafkaClient := services.KafkaClient{
		KafkaConn: conn,
	}

	err = kafkaClient.ConsumeMessages()
	if err != nil {
		log.Println(err)
		return
	}
	r := gin.Default()

	r.POST("/sendEmail", handler.SendEmail)
	r.POST("/sendText", handler.SendText)
	r.POST("/sendWhatsapp", handler.SendWhatsApp)

	err := r.Run("localhost:5000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
