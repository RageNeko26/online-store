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
        "/api/v1/carts": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Listing all of products that customer have been add.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shopping Cart"
                ],
                "summary": "Find existing cart.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CartResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Add Cart using JSON",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shopping Cart"
                ],
                "summary": "Add product item into cart",
                "parameters": [
                    {
                        "description": "add cart",
                        "name": "cart",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateCartRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.CartResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/carts/{cart_id}": {
            "delete": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "Delete selected cart item by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shopping Cart"
                ],
                "summary": "Delete existing cart.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cart ID",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CartResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CartResponse": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.FindCartRow"
                    }
                },
                "total_quantity": {
                    "type": "integer"
                }
            }
        },
        "controller.CreateCartRequest": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                }
            }
        },
        "controller.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "db.FindCartRow": {
            "type": "object",
            "properties": {
                "cart_id": {
                    "type": "string"
                },
                "customer_name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
