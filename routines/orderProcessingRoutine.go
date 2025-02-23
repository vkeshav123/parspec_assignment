package routines

import (
	"log"
	"oms/oms/globals"
	"oms/oms/repo"
	"oms/providers"
	"time"
)

func StartPendingOrderUpdater(repoInstance repo.Repo, config providers.Order) {
	ticker := time.NewTicker(time.Duration(config.OrderProcessingInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ProcessOrders(repoInstance, config.OrderProcessingBatchSize)
	}

}

// Process orders in FIFO order
func ProcessOrders(repoInstance repo.Repo, batchSize int) {
	log.Println("Checking for new orders...")

	orders, err := repoInstance.GetOrders([]string{globals.OrderStatusPending})
	if err != nil {
		log.Printf("Error fetching orders: %v\n", err)
		return
	}
	if len(orders) == 0 {
		log.Println("No pending orders found.")
		return
	}
	for i, order := range orders {
		if i >= batchSize {
			break
		}
		err := repoInstance.UpdateOrder(order.ID, globals.OrderStatusProcessing)
		if err != nil {
			log.Printf("Error updating order %d: %v\n", order.ID, err)
		}
	}
}
