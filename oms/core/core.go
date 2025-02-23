package core

import (
	"oms/oms/entities/payloadEntity"
	"oms/oms/globals"
	"oms/oms/repo"
	"time"
)

type Core struct {
	Repo repo.Repo
}

// func (c *Core) CreateOrder(payload payloadEntity.CreateOrderRequest) (payloadEntity.CreateOrderResponse, error) {
// 	var err error
// 	var totalAmount float64
// 	for _, item := range payload.Items {
// 		totalAmount += item.Price
// 	}
// 	orderEntity := gormEntity.Order{
// 		OrderAmount: totalAmount,
// 		Status:      "pending", // since logic for order status is not specified
// 		CreatedAt:   time.Now(),
// 	}
// 	orderEntity, err = c.Repo.CreateOrder(orderEntity)
// 	if err != nil {
// 		return payloadEntity.CreateOrderResponse{}, err
// 	}
// 	return payloadEntity.CreateOrderResponse{OrderId: orderEntity.ID, Status: orderEntity.Status}, nil
// }

// func (c *Core) GetOrder(payload payloadEntity.GetOrderRequest) (payloadEntity.GetOrderResponse, error) {
// 	orderEntity, err := c.Repo.GetOrder(payload.OrderId)
// 	if err != nil {
// 		return payloadEntity.GetOrderResponse{}, err
// 	}
// 	orderResponse := payloadEntity.GetOrderResponse{
// 		TotalAmount: orderEntity.OrderAmount,
// 		Status:      orderEntity.Status,
// 		CreatedAt:   orderEntity.CreatedAt,
// 	}
// 	return orderResponse, nil
// }

// func (c *Core) UpdateOrder(payload payloadEntity.UpdateOrderRequest) error {
// 	// Check if order exists
// 	_, err := c.Repo.GetOrder(payload.OrderId)
// 	if err != nil {
// 		return err
// 	}
// 	// Update order status
// 	err = c.Repo.UpdateOrder(payload.OrderId, payload.Status)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (c *Core) GetOrderMetrics() (payloadEntity.GetOrderMetricsResponse, error) {
	orders, err := c.Repo.GetOrders(globals.OrderStatuses)
	if err != nil {
		return payloadEntity.GetOrderMetricsResponse{}, err
	}
	var totalAmount float64
	var totalTimeTaken time.Duration
	var pendingOrders, completedOrders, processingOrders int
	for _, order := range orders {
		totalAmount += order.OrderAmount
		switch order.Status {
		case globals.OrderStatusPending:
			pendingOrders++
		case globals.OrderStatusCompleted:
			timeTaken := order.CompletedAt.Sub(order.CreatedAt)
			totalTimeTaken += timeTaken
			completedOrders++
		case globals.OrderStatusProcessing:
			processingOrders++
		}
	}
	averageProcessingTime := int(totalTimeTaken.Seconds()) / len(orders)
	totalOrders := pendingOrders + completedOrders + processingOrders
	return payloadEntity.GetOrderMetricsResponse{
		TotalOrders:       totalOrders,
		PendingOrders:     pendingOrders,
		ProcessingOrders:  processingOrders,
		CompletedOrders:   completedOrders,
		AvgProcessingTime: averageProcessingTime,
	}, nil
}
