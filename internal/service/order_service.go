package service

import (
	"delivery/internal/model"
	"delivery/internal/repository"
)

func GetOrders(latLang model.LatLang) (*[]model.Order, error) {

	order, err := repository.GetOrders(latLang)

	if err != nil {

		return nil, err

	}

	return order, nil
}
