package visacontroller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FinishPayment(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	mapbody := make(map[string]interface{})
	json.Unmarshal(body, &mapbody)

	workflow_id, found := mapbody["workflow_id"]
	if !found {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "workflow_id missing"})
		return
	}
	workflow_id_string, ok := workflow_id.(string)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "workflow_id have to be string"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": workflow_id_string, "message": "Success"})
}
