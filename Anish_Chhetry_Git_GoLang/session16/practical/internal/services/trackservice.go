package services

import (
	"fmt"
	"practical/internal/model"
)

func TrackOrder(chStatus chan model.Order) {
	for order := range chStatus {
		fmt.Println("Tracking order: ID:", order.ID, "Status:", order.Status)
	}
}
