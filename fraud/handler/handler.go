package fraud

import (
	"net/http"

	models "fraud/models"
	service "fraud/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/fraud/check", CheckFraudHandler)
}

func CheckFraudHandler(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := service.CheckFraud(payment)

	c.JSON(http.StatusOK, result)
}
