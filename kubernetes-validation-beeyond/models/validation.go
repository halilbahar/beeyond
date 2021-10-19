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

type NoContentError struct{}

func (e *NoContentError) Error() string {
	return "No Content"
}

type ValidationError struct {
	Message string `json:"message"`
	Value   string `json:"value"`
	Key     string `json:"key"`
}

// Validates the content (syntax wise) checks the constraints
// Parameter: content (string) represents the content of the yaml file,
// which will be validated.
// returns all constraint-errors in []ValidationError and the kubeval error
func ValidateContent(content string) ([]ValidationError, error) {
	if len(content) == 0 {
		return nil, &NoContentError{}
	}
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
				Message: resultError.Description(),
				Value:   string(bytes),
				Key:     field,
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

		if currentConstraint.Disabled && currentConstraint.Path == "" {
			errorDescription = fmt.Sprintf("This root object is disabled")
			actual = fmt.Sprintf("%s", currentConstraint.GroupKindVersion)
		} else if currentConstraint.Disabled && value != nil {
			errorDescription = fmt.Sprintf("Found disabled field (%s)", currentConstraint.Path)
		} else {
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
							errorDescription = fmt.Sprintf("Given value out of range (%.0f-%.0f)", *currentConstraint.Min, *currentConstraint.Max)
							break
						}
					}
				} else if !isBetweenMinMax(currentConstraint, value.(int)) {
					errorDescription = fmt.Sprintf("Given value out of range (%.0f-%.0f)", *currentConstraint.Min, *currentConstraint.Max)
				}

			} else if currentConstraint.Enum != nil {
				if isArray {
					isValid := true
					for _, currentValue := range strings.Split(actual[1:len(actual)-1], ", ") {
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
			} else if currentConstraint.Regex != nil {
				if isArray {
					isValid := true
					for _, currentValue := range strings.Split(actual[1:len(actual)-1], ", ") {
						if !matchesRegex(*currentConstraint.Regex, currentValue) {
							isValid = false
						}
					}

					if !isValid {
						errorDescription = "One or more of the given value does not match the regex"
					}
				} else if !matchesRegex(*currentConstraint.Regex, actual) {
					errorDescription = "Given value does not match regex"
				}
			}
		}

		if errorDescription != "" {
			validationError = append(validationError, ValidationError{
				Message: errorDescription,
				Value:   actual,
				Key:     currentConstraint.Path,
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

// Checks whether the given text matches the given regex
// Parameters:
// 		- regex (string): represents the regex
// 		- text (string): the text that should match the regex
// Returns bool: true if the text matches the regex
func matchesRegex(regex string, text string) bool {
	// TODO: "^"+*currentConstraint.Regex+"$"
	matched, _ := regexp.MatchString("^"+regex+"$", text)
	return matched
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
