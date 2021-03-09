package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yaml-validation/models"
)

func getValidationResult(c *gin.Context) {
	data, _ := c.GetRawData()
	yamlContent := string(data)
	results, err := models.ValidateContent(yamlContent)

	if err != nil {
		// TODO: find what errors can occur and return them if ok
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if len(results) > 0 {
		c.JSON(http.StatusUnprocessableEntity, results)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
