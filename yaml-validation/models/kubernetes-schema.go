package models

import (
	"encoding/json"
	"net/http"
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
}

type Property struct {
	Description string        `json:"description,omitempty"`
	Type        string        `json:"type,omitempty"`
	Format      string        `json:"format,omitempty"`
	Reference   string        `json:"$ref,omitempty"`
	Items       *PropertyItem `json:"items,omitempty"`
	Enum        []string      `json:"enum,omitempty"`
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
