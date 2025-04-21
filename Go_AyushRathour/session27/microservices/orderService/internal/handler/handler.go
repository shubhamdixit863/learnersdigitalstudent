package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendNotification(url string, content string) {
	jsonValue, _ := json.Marshal(content)
	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
}

func CreateOrder(c *gin.Context) {
	sendNotification("http://localhost:5000/sendEmail", " Order created Email sent")
	sendNotification("http://localhost:5000/sendText", "Order created text sent")
	sendNotification("http://localhost:5000/sendWhatsapp", "Order created WhatsApp msg sent")
	c.String(http.StatusOK, "Ordered Successfully & notifications sent")
}

func DeliverOrder(c *gin.Context) {
	sendNotification("http://localhost:5000/sendEmail", " Order  delivered Email sent")
	sendNotification("http://localhost:5000/sendText", "Order delivered text sent")
	sendNotification("http://localhost:5000/sendWhatsapp", "Order delivered WhatsApp msg sent")
	c.String(http.StatusOK, "Order delivered Successfully & notifications sent")
}
