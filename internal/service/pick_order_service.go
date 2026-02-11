package service

import (
	"delivery/internal/kafka"
	"delivery/internal/model"
	"delivery/internal/repository"
	"time"
)

func PickOrderChangeStatus(orderDetails model.DeliveryPick) (*model.Order, error) {

	// here i need to write business logic to change order status

	order, err := repository.PickOrder(orderDetails)

	if err != nil {
		return nil, err
	}
	kafka.InitProducer()
	kafka.StartConsumer()

	kafka.Publish("boy"+orderDetails.OrderNumber, model.DeliveryStartedEvent{
		Event: "delivery.started",

		OrderID:     order.ID,
		OrderNumber: order.OrderNumber,

		DeliveryBoyID:  order.PickedDeliveryBoyID,
		DeliveryBoyLat: orderDetails.DeliveryBoyLatitude,
		DeliveryBoyLng: orderDetails.DeliveryBoyLongitude,

		StoreID:  order.StoreID,
		StoreLat: order.Store.Lat,
		StoreLng: order.Store.Lng,

		CustomerID:  order.CustomerID,
		CustomerLat: order.Customer.Lat,
		CustomerLng: order.Customer.Lng,
		Timestamp:   time.Now(),
	})

	return order, nil
	// here have to get all order details and tracking kafka
}
