package models

import (
	"encoding/json"
	"net/http"
	"strings"
	"yaml-validation/pkg/setting"
)

type SchemaCollection struct {
	Schemas map[string]*Schema `json:"definitions"`
}

type Schema struct {
	Description      string               `json:"description"`
	Required         []string             `json:"required"`
	Type             string               `json:"type"`
	Properties       map[string]*Property `json:"properties"`
	GroupKindVersion []GroupKindVersion   `json:"x-kubernetes-group-version-kind,omitempty"`
	Constraint       *Constraint          `json:"x-constraint,omitempty"`
}

type Property struct {
	Description        string        `json:"description,omitempty"`
	Type               string        `json:"type,omitempty"`
	Format             string        `json:"format,omitempty"`
	Reference          string        `json:"$ref,omitempty"`
	Items              *PropertyItem `json:"items,omitempty"`
	Enum               []string      `json:"enum,omitempty"`
	Constraint         *Constraint   `json:"x-constraint,omitempty"`
	IsKubernetesObject bool          `json:"x-is-kubernetes-object"`
}

type PropertyItem struct {
	Type      string `json:"type,omitempty"`
	Reference string `json:"$ref,omitempty"`
}

type GroupKindVersion struct {
	Group   string `json:"group"`
	Kind    string `json:"kind"`
	Version string `json:"version"`
}

type PathNotFoundError struct{}

func (p PathNotFoundError) Error() string {
	return "Path not found"
}

func (groupKindVersion GroupKindVersion) ToLower() GroupKindVersion {
	var groupKindVersionLower GroupKindVersion
	groupKindVersionLower.Group = strings.ToLower(groupKindVersion.Group)
	groupKindVersionLower.Kind = strings.ToLower(groupKindVersion.Kind)
	groupKindVersionLower.Version = strings.ToLower(groupKindVersion.Version)

	return groupKindVersionLower
}

func GetSchemaCollection() (*SchemaCollection, error) {
	baseUrl := setting.KubernetesJsonschemaSetting.Url
	kubernetesVersion := setting.KubernetesJsonschemaSetting.KubernetesVersion
	versionType := kubernetesVersion + "-standalone-strict"
	url := baseUrl + "/" + versionType + "/_definitions.json"

	response, _ := http.Get(url)

	collection := &SchemaCollection{}
	err := json.NewDecoder(response.Body).Decode(collection)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func GetSchemaBySegments(segments []string) (*Schema, error) {
	collection, err := GetSchemaCollection()
	if err != nil {
		return nil, err
	}

	var currentSchema *Schema
	for i, segment := range segments {
		// On the first element search for the GroupKindVersion
		if i == 0 {
		schemaLoop:
			for _, schema := range collection.Schemas {
				for _, groupKindVersion := range schema.GroupKindVersion {
					var group string
					if groupKindVersion.Group == "" {
						group = ""
					} else {
						group = "-" + groupKindVersion.Group
					}

					groupKindVersionString := groupKindVersion.Kind + group + "-" + groupKindVersion.Version
					if segment == groupKindVersionString {
						currentSchema = schema
						break schemaLoop
					}
				}
			}

			if currentSchema == nil {
				return nil, PathNotFoundError{}
			}
		} else {
			property := currentSchema.Properties[segment]

			if property == nil {
				return nil, PathNotFoundError{}
			}

			var referencePath string
			if property.Reference != "" {
				referencePath = property.Reference
			} else if property.Items != nil {
				referencePath = property.Items.Reference
			}

			// If the specified path of the user does not exist, return
			// This means the user requested something other than object (string, int, ...)
			if referencePath == "" {
				return nil, PathNotFoundError{}
			}

			// We want the last part of the reference
			// Example: #/definitions/io.k8s.api.apps.v1.DeploymentSpec
			split := strings.Split(referencePath, "/")
			definitionName := split[len(split)-1]

			currentSchema = collection.Schemas[definitionName]
		}
	}

	groupKindVersion, constraintPath := GetGroupKindVersionAndPathFromSegments(segments)

	// Attach constraint to the properties if the exist
	for propertyName, property := range currentSchema.Properties {
		var referencePath string
		if property.Reference != "" {
			referencePath = property.Reference
		} else if property.Type == "array" {
			referencePath = property.Items.Reference
		}

		if referencePath != "" {
			// turn: #/definitions/xxx
			// into this: xxx
			split := strings.Split(referencePath, "/")
			definitionName := split[len(split)-1]

			// If the reference is of type object and has properties we declare it as kubernetes object
			// Add new checks if type object and properties are not enough to determine a kubernetes object
			if collection.Schemas[definitionName].Type == "object" && collection.Schemas[definitionName].Properties != nil {
				property.IsKubernetesObject = true
			}
		}

		if constraintPath == "" {
			property.Constraint = GetConstraint(propertyName, groupKindVersion)
		} else {
			property.Constraint = GetConstraint(constraintPath+"."+propertyName, groupKindVersion)
		}
	}

	return currentSchema, nil
}

// TODO: Based on the groupKindVersion in the segment return all available groupKindVersions instead of the only one in the segment
func GetGroupKindVersionAndPathFromSegments(segments []string) (GroupKindVersion, string) {
	var groupKindVersion GroupKindVersion
	parts := strings.Split(segments[0], "-")
	groupKindVersion.Kind = parts[0]
	if len(parts) == 3 {
		groupKindVersion.Group = parts[1]
		groupKindVersion.Version = parts[2]
	} else {
		groupKindVersion.Version = parts[1]
	}

	constraintPath := strings.Join(segments[1:], ".")
	return groupKindVersion, constraintPath
}

func IsValidConstraintPath(segments []string) bool {
	var lastSegment *string
	if len(segments) != 1 {
		lastSegment = &segments[len(segments)-1]
		segments = segments[0 : len(segments)-1]
	}

	currentSchema, err := GetSchemaBySegments(segments)
	return err == nil && (lastSegment == nil || currentSchema.Properties[*lastSegment] != nil)
}
