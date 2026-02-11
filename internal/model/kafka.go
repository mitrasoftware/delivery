package model

import "time"

type DeliveryStartedEvent struct {
	Event string `json:"event"`

	OrderID     uint   `json:"order_id"`
	OrderNumber string `json:"order_number"`

	DeliveryBoyID  uint    `json:"delivery_boy_id"`
	DeliveryBoyLat float64 `json:"delivery_boy_lat"`
	DeliveryBoyLng float64 `json:"delivery_boy_lng"`

	StoreID  uint    `json:"store_id"`
	StoreLat float64 `json:"store_lat"`
	StoreLng float64 `json:"store_lng"`

	CustomerID  uint    `json:"customer_id"`
	CustomerLat float64 `json:"customer_lat"`
	CustomerLng float64 `json:"customer_lng"`

	Timestamp time.Time `json:"timestamp"`
}
