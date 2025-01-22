// Package swagger Code generated by swaggo/swag at 2025-01-22 12:19:57.636596 +0800 +08 m=+1.282696543. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Quddus",
            "url": "http://www.swagger.io/support",
            "email": "ceo@nrmnqdds.com"
        },
        "license": {
            "name": "Bantown Public License",
            "url": "https://github.com/nrmnqdds/gomaluum-api/blob/main/LICENSE.md"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/profile": {
            "get": {
                "description": "Get i-Ma'luum profile",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "scraper"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/result": {
            "get": {
                "description": "Get result from i-Ma'luum",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "scraper"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/api/schedule": {
            "get": {
                "description": "Get schedule from i-Ma'luum",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "scraper"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Logs in the user. Save the token and use it in the Authorization header for future requests.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Login properties",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_proto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "description": "Logs out the user. Clears the token from IIUM's CAS. PASETO token is still valid.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ResponseDTO"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Check the health of the application.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "misc"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "auth_proto.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dtos.ResponseDTO": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Gomaluum API Server",
	Description:      "This is the API server for Gomaluum, an API that serves i-Ma'luum data for ease of developer.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
