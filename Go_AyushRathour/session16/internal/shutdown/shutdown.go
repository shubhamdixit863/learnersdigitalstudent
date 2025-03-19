package shutdown

import (
	"fmt"
	"session16/internal/order"
	"session16/internal/status"
)

func GracefulShutdown() {
	fmt.Println("Graceful shutdown initiated...")
	close(order.OrderQueue)
	close(status.Queue)
}
