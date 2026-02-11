package repository

import (
	"delivery/internal/db"
	"delivery/internal/model"
)

func GetOrders(latLng model.LatLang) (*[]model.Order, error) {

	var order []model.Order

	err := db.DB.
		Preload("Items").
		Preload("Customer").
		Preload("Store").
		Where("status IN ? AND (picked_delivery_boy_id = ? OR picked_delivery_boy_id IS NULL)", []string{"READY", "CREATED"},
			0).
		Order("created_at DESC").
		Find(&order).
		Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
