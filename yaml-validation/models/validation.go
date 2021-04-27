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

// Validates the content (syntax wise) checks the constraints
// Parameter: content (string) represents the content of the yaml file,
// which will be validated.
// returns all constraint-errors in []ValidationError and the kubeval error
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

	groupVersionSplit := strings.Split(groupversion, "/")
	if len(groupVersionSplit) == 1 {
		groupKindVersion.Version = groupVersionSplit[0]
	} else {
		groupKindVersion.Group = groupVersionSplit[0]
		groupKindVersion.Version = groupVersionSplit[1]
	}

	constraints := GetConstraintsByGKV(&groupKindVersion)

	for _, currentConstraint := range constraints {
		errorDescription := ""
		value := getValueFromPath(yamlMap, currentConstraint.Path)

		var actual string
		var ok bool
		isArray := false

		if actual, ok = value.(string); !ok {
			if number, ok := value.(int); ok {
				actual = strconv.Itoa(number)
			} else if arr, ok := value.([]interface{}); ok {
				actual = strings.Join(strings.Fields(fmt.Sprint(arr)), ", ")
				isArray = true
			} else if boolValue, ok := value.(bool); ok {
				actual = strconv.FormatBool(boolValue)
			}
		}

		if currentConstraint.Max != nil {
			if isArray {
				for _, currentValue := range value.([]interface{}) {
					if !isBetweenMinMax(currentConstraint, currentValue.(int)) {
						errorDescription = "Given value out of range"
						break
					}
				}
			} else if !isBetweenMinMax(currentConstraint, value.(int)) {
				errorDescription = "Given value out of range"
			}

		} else if currentConstraint.Enum != nil {
			if isArray {
				isValid := true
				actualValues := strings.Split(actual, ", ")
				for _, currentValue := range actualValues[1 : len(actualValues)-2] {
					if !contains(currentConstraint.Enum, currentValue) {
						isValid = false
					}
				}

				if !isValid {
					errorDescription = "Constraint enum does not contain given one or more of the given values"
				}
			} else if !contains(currentConstraint.Enum, actual) {
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

// Checks whether the given string array contains the given searchText
// Parameters:
// 		- enum ([]string): array which we search through
// 		- searchText (string): the string we look for in the array
// Returns boolean: true if the array contains the searchText
func contains(enum []string, searchText string) bool {
	for _, currentValue := range enum {
		if currentValue == searchText {
			return true
		}
	}
	return false
}

// Checks whether the given value is between the min and max values given within the currentConstraint
// Parameters:
// 		- currentConstraint (*Constraint): Contains the min and max values
// 		- value (int): integer which should be between min and max
// Returns: bool: true if value is between min and max, otherwise false
func isBetweenMinMax(currentConstraint *Constraint, value int) bool {
	actualFloat := float64(value)
	return actualFloat <= float64(*currentConstraint.Max) && actualFloat >= float64(*currentConstraint.Min)
}

// Gets the value of the property by the given path from the given k8s specification (map)
// Parameters:
// 		- m (map[interface{}]interface{}): Represents the content of the given yaml file as a map
//		- path (string): Represents the
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
