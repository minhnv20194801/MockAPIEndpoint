package routers

import (
	"rezza-mock/controllers/visacontroller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	apiv1 := router.Group("/api/v1")

	apiv1Visa := apiv1.Group("/visa-tubudd")
	{
		apiv1Visa.POST("/finish-payment", visacontroller.FinishPayment)
	}
	return router
}
