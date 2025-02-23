package repo

import (
	"errors"
	"oms/oms/entities/gormEntity"
	"oms/oms/globals"
	"time"
)

func (r *Repo) CreateOrder(entity gormEntity.Order) (gormEntity.Order, error) {
	err := r.DB.Table("orders").Create(&entity).Error
	if err != nil {
		return gormEntity.Order{}, err
	}
	return entity, nil
}

// func (r *Repo) GetOrder(orderId int) (gormEntity.Order, error) {
// 	var entity gormEntity.Order
// 	err := r.DB.Table("orders").Where("id = ?", orderId).First(&entity).Error
// 	if err != nil {
// 		return entity, err
// 	}
// 	return entity, nil
// }

func (r *Repo) GetOrders(status []string) ([]gormEntity.Order, error) {
	var entity []gormEntity.Order
	err := r.DB.Table("orders").Where("status IN (?)", status).Find(&entity).Order("id asc").Error
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *Repo) UpdateOrder(orderId int, status string) error {
	gormStatement := r.DB.Table("orders").Where("id = ?", orderId)
	if status == globals.OrderStatusCompleted {
		gormStatement.UpdateColumns(map[string]interface{}{
			"status":       status,
			"completed_at": time.Now(),
		})
	} else if status == globals.OrderStatusProcessing {
		gormStatement.UpdateColumns(map[string]interface{}{
			"status":       status,
			"processed_at": time.Now(),
		})
	} else {
		return errors.New("status cannot be updated")
	}
	err := gormStatement.Error
	if err != nil {
		return err
	}
	return nil
}
