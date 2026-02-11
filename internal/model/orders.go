package model

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	Name   string `gorm:"type:text;not null"`
	Mobile string `gorm:"type:varchar(15)"`
	Email  string `gorm:"type:text"`

	Address string `gorm:"type:text;not null"`

	Lat float64 `gorm:"type:numeric(10,7);not null"`
	Lng float64 `gorm:"type:numeric(10,7);not null"`

	IsActive bool `gorm:"default:true"`
	IsOpen   bool `gorm:"default:true"`

	OpenTime  *time.Time `gorm:"type:time"`
	CloseTime *time.Time `gorm:"type:time"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Order struct {
	gorm.Model
	ID                  uint   `gorm:"primaryKey"`
	OrderNumber         string `gorm:"uniqueIndex;size:30"`
	CustomerID          uint
	PickedDeliveryBoyID uint
	Customer            Customer `gorm:"foreignKey:CustomerID;references:ID"`

	StoreID uint
	Store   Store `gorm:"foreignKey:StoreID;references:ID"`

	TotalAmount   float64 `gorm:"type:numeric(10,2)"`
	PaymentMode   string  `gorm:"size:20"`
	PaymentStatus string  `gorm:"size:20"`

	Status string `gorm:"size:30"`

	PlacedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time

	Items []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint `gorm:"index"`

	ProductID   uint
	ProductName string
	Quantity    int
	Price       float64 `gorm:"type:numeric(10,2)"`
}

type Customer struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:text"`

	Mobile   string `gorm:"type:varchar(15)"`
	Password string `gorm:"type:text"`

	IsActive bool `gorm:"default:true"`

	Address string `gorm:"type:text"`

	Lat float64 `gorm:"type:numeric(10,7)"`
	Lng float64 `gorm:"type:numeric(10,7)"`

	DeletedAt time.Time `gorm:"deleted_at"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type LatLang struct {
	Lat string `json:"latitude"`
	Lng string `json:"longitude"`
}
