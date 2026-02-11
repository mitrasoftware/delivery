package orders

import (
	"delivery/internal/model"
	"delivery/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {

	var payload = model.LatLang{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	res, err := service.GetOrders(payload)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

	}

	c.JSON(http.StatusOK, gin.H{"orders": res})
}
