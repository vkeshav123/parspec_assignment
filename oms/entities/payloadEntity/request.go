package payloadEntity

type CreateOrderRequest struct {
	UserId int         `json:"user_id"`
	Items  []OrderItem `json:"items"`
}

type GetOrderRequest struct {
	OrderId int `json:"order_id"`
}

type OrderItem struct {
	ProductId uint    `json:"product_id"`
	Qty       uint    `json:"qty"`
	Price     float64 `json:"price"`
}

type UpdateOrderRequest struct {
	OrderId int    `json:"order_id"`
	Status  string `json:"status"`
}
