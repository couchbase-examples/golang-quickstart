// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/health": {
            "get": {
                "description": "Checks if service is running",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health Check Controller"
                ],
                "summary": "Checks for service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/profile": {
            "post": {
                "description": "Creates the Document with key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile Controller"
                ],
                "summary": "Create Document",
                "parameters": [
                    {
                        "description": "Creates a document",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ProfileResponse"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Page Not found"
                    },
                    "500": {
                        "description": "Error while getting examples"
                    }
                }
            }
        },
        "/api/v1/profile/{id}": {
            "get": {
                "description": "Gets the Document with key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile Controller"
                ],
                "summary": "Get Document",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search document by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ProfileResponse"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Page Not found"
                    },
                    "500": {
                        "description": "Error while getting examples"
                    }
                }
            },
            "put": {
                "description": "Updates the Document with key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile Controller"
                ],
                "summary": "Update Document",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Update document by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Creates a document",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ProfileResponse"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Page Not found"
                    },
                    "500": {
                        "description": "Error while getting examples"
                    }
                }
            },
            "delete": {
                "description": "Deletes the Document with key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile Controller"
                ],
                "summary": "Deletes Document",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Deletes a document with key specified",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.ProfileResponse"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Page Not found"
                    },
                    "500": {
                        "description": "Error while getting examples"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.RequestBody": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                },
                "FirstName": {
                    "type": "string"
                },
                "LastName": {
                    "type": "string"
                },
                "Pid": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "responses.ProfileResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Profile API",
	Description:      "Couchbase Golang Quickstart using Gin Gonic",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
