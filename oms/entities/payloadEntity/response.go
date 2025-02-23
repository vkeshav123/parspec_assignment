package payloadEntity

import "time"

type CreateOrderResponse struct {
	OrderId int    `json:"order_id"`
	Status  string `json:"status"`
}

type GetOrderResponse struct {
	Items       []OrderItem `json:"items"`
	TotalAmount float64     `json:"total_amount"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	CreatedBy   int         `json:"created_by"`
}

type GetOrderMetricsResponse struct {
	TotalOrders       int `json:"total_orders"`
	PendingOrders     int `json:"pending_orders"`
	ProcessingOrders  int `json:"processing_orders"`
	CompletedOrders   int `json:"completed_orders"`
	AvgPendingTime    int `json:"average_pending_time"`
	AvgProcessingTime int `json:"average_processing_time"`
	AvgCompletionTime int `json:"average_completion_time"`
}
