// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/turno": {
            "get": {
                "description": "Get turnos by Dentista ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "Get turnos by Dentista ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "dni",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/turno.Turno"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "Add a turno",
                "parameters": [
                    {
                        "description": "Turno",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            }
        },
        "/turno/{id}": {
            "get": {
                "description": "Get turno by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "Get turno by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "Update a turno",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Turno",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a turno",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "Delete a turno",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "turno deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update a turno with a patch",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "Update a turno with a patch",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Turno",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/turno.Turno"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "turno.Turno": {
            "type": "object",
            "properties": {
                "dentistaId": {
                    "type": "string"
                },
                "descripcion": {
                    "type": "string"
                },
                "fechaHora": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "pacienteId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Certified Tech Developer - Dentistas",
	Description:      "This API Handle Products.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
