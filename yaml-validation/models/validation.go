package models

import (
	"github.com/instrumenta/kubeval/kubeval"
	"github.com/smallfish/simpleyaml"
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

			validationError = append(validationError, ValidationError{
				Description: resultError.Description(),
				Value:       resultError.Value().(string),
				Field:       field,
			})
		}
	}

	y, err := simpleyaml.NewYaml([]byte(content))
	if err != nil {
		panic(err)
	}

	var gkv GroupKindVersion

	gkv.Kind, _ = y.Get("kind").String()
	groupversion, _ := y.Get("apiVersion").String()
	gkv.Group = strings.Split(groupversion, "/")[0]
	gkv.Version = strings.Split(groupversion, "/")[1]

	constraints := GetConstraintsByGKV(&gkv)

	for _, cur := range constraints {
		pathSegments := strings.Split(cur.Path, ".")
		p := y
		for _, seg := range pathSegments {
			p = p.Get(seg)
		}
		actual, err := p.String()
		if err != nil {
			actualInt, _ := p.Int()
			actual = strconv.Itoa(actualInt)
		}

		errorDescription := ""

		if cur.Max != nil {
			actualFloat, _ := strconv.ParseFloat(actual, 32)
			if actualFloat > float64(*cur.Max) || actualFloat < float64(*cur.Min) {
				errorDescription = "Given value out of range"
			}
		} else if cur.Enum != nil {
			found := false
			for _, s := range cur.Enum {
				if s == actual {
					found = true
				}
			}
			if !found {
				errorDescription = "Constraint enum does not contain given value"
			}
		} else {
			matched, _ := regexp.MatchString("^"+*cur.Regex+"$", actual)

			if !matched {
				errorDescription = "Given value does not match regex"
			}

		}

		if errorDescription != "" {
			validationError = append(validationError, ValidationError{
				Description: errorDescription,
				Value:       actual,
				Field:       cur.Path,
			})
		}
	}

	return validationError, nil
}
