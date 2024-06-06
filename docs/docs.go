// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Figarillo",
            "email": "axel.leonardi.22@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/categories": {
            "get": {
                "description": "Get a list of all categories with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "List all categories with pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Categories retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Category"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new category with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Create a new category",
                "parameters": [
                    {
                        "description": "Category data",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Category created successfully",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                }
            }
        },
        "/api/categories/{id}": {
            "get": {
                "description": "Retrieve a category using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get a category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing category by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Update a category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category data",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category updated successfully",
                        "schema": {
                            "$ref": "#/definitions/entity.Category"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing category using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Delete a category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category deleted successfully"
                    }
                }
            }
        },
        "/api/clients": {
            "get": {
                "description": "Get a list of all clients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "List all clients",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
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
                                "$ref": "#/definitions/entity.Client"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Create a client",
                "parameters": [
                    {
                        "description": "Client",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Client"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/api/clients/{id}": {
            "get": {
                "description": "Get a client by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Get a client",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Client"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Update a client",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Client",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "clients"
                ],
                "summary": "Delete a client",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Client ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/orders": {
            "get": {
                "description": "Get a list of orders with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "List orders with pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Orders retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new order with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "Order data",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Order"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Order created successfully",
                        "schema": {
                            "$ref": "#/definitions/entity.Order"
                        }
                    }
                }
            }
        },
        "/api/orders/client/{id}": {
            "get": {
                "description": "Retrive orders using its client ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get orders by client ID",
                "responses": {
                    "200": {
                        "description": "Orders retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Order"
                            }
                        }
                    }
                }
            }
        },
        "/api/orders/{id}": {
            "get": {
                "description": "Retrive a order using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get a order by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/entity.Order"
                        }
                    }
                }
            },
            "put": {
                "description": "Set order status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Set status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order data",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order status updated successfully",
                        "schema": {
                            "$ref": "#/definitions/entity.Order"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a order by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete a order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order deleted successfully"
                    }
                }
            }
        },
        "/api/products": {
            "get": {
                "description": "Get a list of products with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "List products with pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
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
                                "$ref": "#/definitions/entity.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create a product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/api/products/{id}": {
            "get": {
                "description": "Get a product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Delete a product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
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
        "entity.Category": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Product"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Client": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 120,
                    "minimum": 0
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Order"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Order": {
            "type": "object",
            "required": [
                "client_id",
                "status"
            ],
            "properties": {
                "client": {
                    "$ref": "#/definitions/entity.Client"
                },
                "client_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Product"
                    }
                },
                "status": {
                    "type": "string"
                },
                "total": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entity.Product": {
            "type": "object",
            "required": [
                "category_id"
            ],
            "properties": {
                "category": {
                    "$ref": "#/definitions/entity.Category"
                },
                "category_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number",
                    "minimum": 0
                },
                "stock": {
                    "type": "integer",
                    "minimum": 0
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "golerplate API",
	Description:      "GOlerplate is boilerpalte for GO",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
