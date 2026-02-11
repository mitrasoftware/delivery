package orders

import (
	"delivery/internal/model"
	"delivery/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PickOrders(c *gin.Context) {

	var deliveryPick = model.DeliveryPick{}

	if err := c.ShouldBindJSON(&deliveryPick); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.PickOrderChangeStatus(deliveryPick)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}
