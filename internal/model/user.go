package model

import "gorm.io/gorm"

// User model definition

type Users struct {
	gorm.Model

	Name     string
	Mobile   string `gorm:"uniqueIndex"`
	Password string
	IsActive bool
	Latitude string
	Logitude string
}

type DeliveryPick struct {
	OrderNumber          string  `json:"order_number"`
	DeliveryBoyID        uint    `json:"delivery_boy_id"`
	Status               string  `json:"status"`
	DeliveryBoyLatitude  float64 `json:"del_boy_lat"`
	DeliveryBoyLongitude float64 `json:"del_boy_lng"`

	StoreLatitude  float64 `json:"store_lat"`
	StoreLongitude float64 `json:"store_lng"`

	CustomerLatitude  float64 `json:"cust_lat"`
	CustomerLongitude float64 `json:"cust_lng"`
}
