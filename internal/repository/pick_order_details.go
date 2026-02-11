package repository

import (
	"delivery/internal/db"
	"delivery/internal/model"
	"errors"
	"time"
)

func PickOrder(details model.DeliveryPick) (*model.Order, error) {

	// 1️⃣ Atomic update (no preload, no fetch)
	result := db.DB.
		Model(&model.Order{}).
		Where(
			"order_number = ? AND status IN ? AND (picked_delivery_boy_id = ? OR picked_delivery_boy_id IS NULL)",
			details.OrderNumber,
			[]string{"READY", "CREATED"},
			0,
		).
		Updates(map[string]interface{}{
			"picked_delivery_boy_id": details.DeliveryBoyID,
			"status":                 "PICKED",
			"updated_at":             time.Now(),
		})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("order already picked or not available")
	}

	// 2️⃣ Fetch updated order (NOW preload works)
	var order model.Order
	err := db.DB.
		Preload("Items").
		Preload("Customer").
		Preload("Store").
		Where("order_number = ?", details.OrderNumber).
		First(&order).
		Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}
