package gormEntity

import "time"

type Order struct {
	ID          int       `gorm:"primary_key"`
	ProductID   int       `gorm:"product_id"`
	Quantity    int       `gorm:"quantity"`
	OrderAmount float64   `gorm:"order_amount"`
	Status      string    `gorm:"status"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
	CompletedAt time.Time `gorm:"completed_at"`
}

// type OrderItem struct {
// 	ProductID int     `json:"product_id"`
// 	Qty       int     `json:"qty"`
// 	Price     float64 `json:"price"`
// }
