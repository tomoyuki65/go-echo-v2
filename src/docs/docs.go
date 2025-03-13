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
        "/api/v1/": {
            "get": {
                "description": "テキスト「Hello World !!」を出力する。",
                "tags": [
                    "index"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/healthcheck": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "APIとDBの接続確認をするためのヘルスチェックAPI",
                "tags": [
                    "healthcheck"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-echo-v2_internal_handlers_healthcheck.OKResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/healthcheck.InternalServerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "go-echo-v2_internal_handlers_healthcheck.OKResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Health Check OK !!"
                }
            }
        },
        "healthcheck.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Failed to health check: error message"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "go-echo-v2 API",
	Description:      "Go言語（Golang）のフレームワーク「Echo」によるAPI開発サンプルのバージョン２",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
