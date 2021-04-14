package models

import (
	"encoding/json"
	"fmt"
	"github.com/instrumenta/kubeval/kubeval"
	"gopkg.in/yaml.v2"
	"regexp"
	"strconv"
	"strings"
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

			bytes, _ := json.Marshal(resultError.Value())

			validationError = append(validationError, ValidationError{
				Description: resultError.Description(),
				Value:       string(bytes),
				Field:       field,
			})
		}
	}

	var groupKindVersion GroupKindVersion

	yamlMap := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(content), &yamlMap)

	groupKindVersion.Kind = getValueFromPath(yamlMap, "kind").(string)
	groupversion := getValueFromPath(yamlMap, "apiVersion").(string)

	groupKindVersion.Group = strings.Split(groupversion, "/")[0]
	groupKindVersion.Version = strings.Split(groupversion, "/")[1]

	constraints := GetConstraintsByGKV(&groupKindVersion)

	for _, currentConstraint := range constraints {
		errorDescription := ""
		value := getValueFromPath(yamlMap, currentConstraint.Path)

		var actual string
		var ok bool
		if actual, ok = value.(string); !ok {
			actual = strconv.Itoa(value.(int))
		}

		if currentConstraint.Max != nil {
			actualFloat := float64(getValueFromPath(yamlMap, currentConstraint.Path).(int))
			fmt.Print(err)
			if actualFloat > float64(*currentConstraint.Max) || actualFloat < float64(*currentConstraint.Min) {
				errorDescription = "Given value out of range"
			}
		} else if currentConstraint.Enum != nil {
			found := false
			for _, s := range currentConstraint.Enum {
				if s == actual {
					found = true
				}
			}
			if !found {
				errorDescription = "Constraint enum does not contain given value"
			}
		} else {
			// TODO: "^"+*currentConstraint.Regex+"$"
			matched, _ := regexp.MatchString("^"+*currentConstraint.Regex+"$", actual)

			if !matched {
				errorDescription = "Given value does not match regex"
			}

		}

		if errorDescription != "" {
			validationError = append(validationError, ValidationError{
				Description: errorDescription,
				Value:       actual,
				Field:       currentConstraint.Path,
			})
		}
	}

	return validationError, nil
}

func getValueFromPath(m map[interface{}]interface{}, path string) interface{} {
	var obj interface{} = m
	var val interface{} = nil

	parts := strings.Split(path, ".")
	for _, p := range parts {
		if v, ok := obj.(map[interface{}]interface{}); ok {
			obj = v[p]
			val = obj
		} else {
			return nil
		}
	}

	return val
}
