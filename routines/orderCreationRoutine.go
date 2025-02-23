package routines

import (
	"fmt"
	"oms/oms/entities/gormEntity"
	"oms/oms/repo"
	"oms/providers"
	"sync"
	"time"
)

func StartOrderCreation(repoInstance repo.Repo, config providers.Order) {
	startTime := time.Now()

	// Create order queue (channel)
	orderCount := config.LoadCreationBatchSize
	workers := 10
	jobs := make(chan gormEntity.Order, orderCount)
	var wg sync.WaitGroup

	// Launch worker pool
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go CreateOrder(i, jobs, &wg, repoInstance)
	}

	// Generate orders and send them to the queue
	for i := 1; i <= orderCount; i++ {
		order := gormEntity.Order{
			ProductID:   i % 10,
			Quantity:    1 + (i % 5),
			OrderAmount: 100,
			Status:      "pending",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		jobs <- order
	}

	close(jobs) // Close the channel after dispatching all orders
	wg.Wait()   // Wait for all workers to finish

	fmt.Println("All orders processed in", time.Since(startTime))
}

func CreateOrder(id int, orders <-chan gormEntity.Order, wg *sync.WaitGroup, repoInstance repo.Repo) {
	defer wg.Done()

	for order := range orders {
		// Retry logic with exponential backoff
		var err error
		for attempt := 1; attempt <= 3; attempt++ {
			order, err = repoInstance.CreateOrder(order)
			if err == nil {
				fmt.Println("Order created successfully with id: ", order.ID)
				break
			}

			fmt.Printf("Error creating order (attempt %d): %v\n", attempt, err)
			time.Sleep(time.Duration(attempt) * time.Second) // Exponential backoff
		}

		if err != nil {
			fmt.Println("Failed to create order after retries:", order)
		}
	}
}
