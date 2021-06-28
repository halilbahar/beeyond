package routers

import (
	"github.com/gin-gonic/gin"
	"kubernetes-validation-beeyond/models"
	"net/http"
)

// @Summary Validate content
// @Description Validates the given content
// @Tags Validation
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /api/validate/ [post]
func getValidationResult(c *gin.Context) {
	data, _ := c.GetRawData()
	yamlContent := string(data)
	results, err := models.ValidateContent(yamlContent)

	if err != nil {
		// TODO: find what errors can occur and return them if ok
		results = append(results, models.ValidationError{
			Message: "YAML-Format not valid",
			Value:       yamlContent,
			Key:       "content",
		})
	}

	c.JSON(http.StatusOK, results)
}
