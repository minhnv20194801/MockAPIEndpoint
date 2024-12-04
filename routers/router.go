package routers

import (
	"rezza-mock/controllers/visacontroller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	apiv1 := router.Group("/api/v1")

	apiv1VisaFlow := apiv1.Group("/visa-tubudd/flow")
	{
		apiv1VisaFlow.POST(":workflow_id/finish-payment", visacontroller.FinishPayment)
	}
	return router
}
