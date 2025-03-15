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
        },
        "/api/v1/user": {
            "post": {
                "description": "ユーザー作成API",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "作成するユーザー情報",
                        "name": "CreateUserRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/go-echo-v2_internal_handlers_user.CreateUserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/go-echo-v2_internal_handlers_user.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/user.BadRequestResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/user.UnprocessableEntityResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/:uid": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "有効な対象ユーザー取得API",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "対象データが存在しない場合は空のオブジェクト「{}」を返す。",
                        "schema": {
                            "$ref": "#/definitions/go-echo-v2_internal_handlers_user.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/user.UnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "対象ユーザー更新API",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新するユーザー情報",
                        "name": "UpdateUserRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/go-echo-v2_internal_handlers_user.UpdateUserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/go-echo-v2_internal_handlers_user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/user.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/user.UnauthorizedResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/user.UnprocessableEntityResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "対象ユーザー削除API",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.OKResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/user.UnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "全てのユーザー取得API \u003cbr/\u003e ※削除済みユーザー含む",
                "tags": [
                    "user"
                ],
                "responses": {
                    "200": {
                        "description": "対象データが存在しない場合は空の配列「[]」を返す。",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/go-echo-v2_internal_handlers_user.UserResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/user.UnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/user.InternalServerErrorResponse"
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
        "go-echo-v2_internal_handlers_user.CreateUserRequestBody": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "t.yamada@example.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "太郎"
                },
                "last_name": {
                    "type": "string",
                    "example": "山田"
                }
            }
        },
        "go-echo-v2_internal_handlers_user.CreateUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "t.yamada@example.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "太郎"
                },
                "last_name": {
                    "type": "string",
                    "example": "山田"
                },
                "uid": {
                    "type": "string",
                    "example": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
                }
            }
        },
        "go-echo-v2_internal_handlers_user.UpdateUserRequestBody": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "t.sato@example.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "太郎"
                },
                "last_name": {
                    "type": "string",
                    "example": "佐藤"
                }
            }
        },
        "go-echo-v2_internal_handlers_user.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2025-03-15 18:08:00"
                },
                "deleted_at": {
                    "type": "string",
                    "example": ""
                },
                "email": {
                    "type": "string",
                    "example": "t.yamada@example.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "太郎"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_name": {
                    "type": "string",
                    "example": "山田"
                },
                "uid": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2025-03-15 18:08:00"
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
        },
        "user.BadRequestResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "リクエストボディが不正です。: error message"
                }
            }
        },
        "user.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Internal Server Error: error message"
                }
            }
        },
        "user.OKResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "user.UnauthorizedResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Unauthorized"
                }
            }
        },
        "user.UnprocessableEntityResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "バリデーションエラー: error message"
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
