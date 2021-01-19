package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "yaml-validation/models"
)

func getValidationResult(c *gin.Context) {
	data, _ := c.GetRawData()
	yamlContent := string(data)
	results, err := ValidateContent(yamlContent)

	// TODO: handle error
	print(err)

	if len(results) > 0 {
		c.JSON(http.StatusUnprocessableEntity, results)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
