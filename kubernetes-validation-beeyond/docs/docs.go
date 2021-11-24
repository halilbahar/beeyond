// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/constraints/": {
            "get": {
                "description": "Finds all root schemes and their constraints",
                "tags": [
                    "Constraint"
                ],
                "summary": "Find root constraints",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/constraints/{path}": {
            "get": {
                "description": "Finds the schema and its constraints according to the given path",
                "tags": [
                    "Constraint"
                ],
                "summary": "Find constraints by path",
                "parameters": [
                    {
                        "type": "string",
                        "description": "path",
                        "name": "\"path\"",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "creates a new constraint and adds it to the database. If the constraint already exists it gets replaced.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Constraint"
                ],
                "summary": "Creates a new constraint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "path",
                        "name": "\"path\"",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes the constraint with the given path",
                "tags": [
                    "Constraint"
                ],
                "summary": "Delete constraint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "path",
                        "name": "\"path\"",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "no content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Toggles the \"disabled\" from the constraint with the given path. If the given constraint does not exist, it will be created",
                "tags": [
                    "Constraint"
                ],
                "summary": "Toggle disabled on constraint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "path",
                        "name": "\"path\"",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/validate/": {
            "post": {
                "description": "Validates the given content",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Validation"
                ],
                "summary": "Validate content",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "unprocessable entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger Kubernetes Validation Beeyond API",
	Description: "This is an API for the validation of kubernetes specifications (yaml) with constraints.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}