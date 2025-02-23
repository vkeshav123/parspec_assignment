package routines

import (
	"log"
	"oms/oms/globals"
	"oms/oms/repo"
	"oms/providers"
	"time"
)

func StartProcessingOrderUpdater(repoInstance repo.Repo, config providers.Order) {
	ticker := time.NewTicker(time.Duration(config.OrderCompletionInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		CompleteOrders(repoInstance, config.OrderCompletionBatchSize)
	}

}

// Process orders in FIFO order
func CompleteOrders(repoInstance repo.Repo, batchSize int) {
	log.Println("Checking for new orders...")

	orders, err := repoInstance.GetOrders([]string{globals.OrderStatusProcessing})
	if err != nil {
		log.Printf("Error fetching orders: %v\n", err)
		return
	}
	if len(orders) == 0 {
		log.Println("No processing orders found.")
		return
	}
	for i, order := range orders {
		if i >= batchSize {
			break
		}
		err := repoInstance.UpdateOrder(order.ID, globals.OrderStatusCompleted)
		if err != nil {
			log.Printf("Error updating order %d: %v\n", order.ID, err)
		}
	}
}
