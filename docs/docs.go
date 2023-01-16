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
            "name": "qinguoyi",
            "email": "qinguoyiwork@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/storage/v0/ping": {
            "get": {
                "description": "LangGo的测试接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试"
                ],
                "summary": "测试接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/langgo_app_pkg_response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "langgo_app_pkg_response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "自定义错误码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据"
                },
                "message": {
                    "description": "信息",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8890",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "业务框架LangGo接口",
	Description:      "LangGo相关接口",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}