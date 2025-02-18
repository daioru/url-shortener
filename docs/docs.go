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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shorten": {
            "post": {
                "description": "Принимает оригинальный URL и возвращает сокращённый",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Создать короткий URL",
                "parameters": [
                    {
                        "description": "URL для сокращения",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ShortenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный ответ с сокращённым URL",
                        "schema": {
                            "$ref": "#/definitions/service.ShortenResponse"
                        }
                    },
                    "400": {
                        "description": "Ошибка валидации",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/service.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "description": "Оставляем пустым, если нет доп. сообщения",
                    "type": "string"
                }
            }
        },
        "service.ShortenRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "service.ShortenResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "short_url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "172.27.227.76:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "URL Shortener API",
	Description:      "API для сокращения URL",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
