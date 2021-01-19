package models

import (
	"github.com/instrumenta/kubeval/kubeval"
)

type ValidationError struct {
	Description string `json:"description"`
	Value       string `json:"value"`
	Field       string `json:"field"`
}

func ValidateContent(content string) ([]ValidationError, error) {
	config := kubeval.NewDefaultConfig()

	contentBytes := []byte(content)

	validationResults, err := kubeval.Validate(contentBytes, config)
	if err != nil {
		return nil, err
	}

	var validationError []ValidationError
	for _, result := range validationResults {
		for _, resultError := range result.Errors {
			fieldDetail := resultError.Details()["field"]
			var field string
			if fieldDetail != nil {
				field = fieldDetail.(string)
			} else {
				field = ""
			}

			validationError = append(validationError, ValidationError{
				Description: resultError.Description(),
				Value:       resultError.Value().(string),
				Field:       field,
			})
		}
	}

	return validationError, nil
}
