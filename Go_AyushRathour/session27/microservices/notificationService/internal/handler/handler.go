package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendEmail(c *gin.Context) {
	c.String(http.StatusOK, "Email sent")
}
func SendText(c *gin.Context) {
	c.String(http.StatusOK, "Text sent")
}

func SendWhatsApp(c *gin.Context) {
	c.String(http.StatusOK, "WhatsApp sent")
}
