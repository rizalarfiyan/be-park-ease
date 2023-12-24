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
            "name": "Rizal Arfiyan",
            "url": "https://rizalrfiyan.com",
            "email": "rizal.arfiyan.23@gmail.com"
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
        "/": {
            "get": {
                "description": "Base Home",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "home"
                ],
                "summary": "Get Base Home based on parameter",
                "operationId": "get-base-home",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Auth Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Post Auth Login based on parameter",
                "operationId": "post-auth-login",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Auth Me",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get Auth Me based on parameter",
                "operationId": "get-auth-me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/middleware.AuthUserData"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/history": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "All History",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "history"
                ],
                "summary": "Get All History based on parameter",
                "operationId": "get-all-history",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "id",
                            "location_code",
                            "vehicle_type_code",
                            "vehicle_number",
                            "date",
                            "type"
                        ],
                        "type": "string",
                        "description": "Order by",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "Order",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "entry",
                            "exit",
                            "fine"
                        ],
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Vehicle Type",
                        "name": "vehicle_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Location",
                        "name": "location",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.BaseResponsePagination-response_EntryHistory"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "All User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get All User based on parameter",
                "operationId": "get-all-user",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "name",
                            "username",
                            "role",
                            "status"
                        ],
                        "type": "string",
                        "description": "Order by",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "Order",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "admin",
                            "karyawan"
                        ],
                        "type": "string",
                        "description": "Role",
                        "name": "role",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "active",
                            "banned"
                        ],
                        "type": "string",
                        "description": "Status",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.BaseResponsePagination-response_User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Create User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Post Create User based on parameter",
                "operationId": "post-create-user",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "User By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get User By ID based on parameter",
                "operationId": "get-user-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Update User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Post Update User based on parameter",
                "operationId": "put-update-user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "Data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "middleware.AuthUserData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/sql.UserRole"
                },
                "status": {
                    "$ref": "#/definitions/sql.UserStatus"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.AuthLoginRequest": {
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
        "request.CreateUserRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Paijo Royo Royo"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "role": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/sql.UserRole"
                        }
                    ],
                    "example": "karyawan"
                },
                "status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/sql.UserStatus"
                        }
                    ],
                    "example": "active"
                },
                "username": {
                    "type": "string",
                    "example": "paijo"
                }
            }
        },
        "request.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Paijo Royo Royo"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "role": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/sql.UserRole"
                        }
                    ],
                    "example": "karyawan"
                },
                "status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/sql.UserStatus"
                        }
                    ],
                    "example": "active"
                },
                "username": {
                    "type": "string",
                    "example": "paijo"
                }
            }
        },
        "response.BaseMetadataPagination": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.BaseResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 999
                },
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Message!"
                }
            }
        },
        "response.BaseResponsePagination-response_EntryHistory": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.EntryHistory"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/response.BaseMetadataPagination"
                }
            }
        },
        "response.BaseResponsePagination-response_User": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.User"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/response.BaseMetadataPagination"
                }
            }
        },
        "response.EntryHistory": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location_code": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "vehicle_number": {
                    "type": "string"
                },
                "vehicle_type_code": {
                    "type": "string"
                }
            }
        },
        "response.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/sql.UserRole"
                },
                "status": {
                    "$ref": "#/definitions/sql.UserStatus"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "sql.UserRole": {
            "type": "string",
            "enum": [
                "admin",
                "karyawan"
            ],
            "x-enum-varnames": [
                "UserRoleAdmin",
                "UserRoleKaryawan"
            ]
        },
        "sql.UserStatus": {
            "type": "string",
            "enum": [
                "active",
                "banned"
            ],
            "x-enum-varnames": [
                "UserStatusActive",
                "UserStatusBanned"
            ]
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "BE Park Ease API",
	Description:      "This is a API documentation of BE Park Ease",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
