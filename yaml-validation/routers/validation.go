package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yaml-validation/models"
)

// Validates the given content
// Parameter: c (*gin.Context): contains the content that should be validated
// Possible status codes:
// 		- 422, if the content is empty or not valid
// 		- 200, if no validation errors occurred
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
