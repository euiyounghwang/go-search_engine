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
        "/": {
            "get": {
                "description": "test swagger api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "test swagger api",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/es/search": {
            "post": {
                "description": "search engine api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "search engine api",
                "parameters": [
                    {
                        "description": "Search Info Body",
                        "name": "search",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repository.Search"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Search"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "search engine health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "search engine health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/hello/{name}": {
            "get": {
                "description": "test swagger api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "test swagger api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Users name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Users"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Users": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "Age",
                    "type": "integer",
                    "example": 10
                },
                "id": {
                    "description": "UserId",
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "description": "Name",
                    "type": "string",
                    "example": "John"
                }
            }
        },
        "repository.Search": {
            "type": "object",
            "properties": {
                "ids_filter": {
                    "description": "IdsFilter        []string ` + "`" + `json:\"ids_filter\"\"` + "`" + `",
                    "type": "string",
                    "example": "111,222"
                },
                "include_basic_aggs": {
                    "type": "boolean",
                    "example": true
                },
                "pit": {
                    "type": "string",
                    "example": ""
                },
                "query_string": {
                    "type": "string",
                    "example": "performance"
                },
                "size": {
                    "type": "integer",
                    "example": 10
                },
                "sort_order": {
                    "type": "string",
                    "example": "DESC"
                },
                "source_fields": {
                    "type": "string",
                    "example": "*"
                },
                "start_date": {
                    "type": "string",
                    "example": "2021 01-01 00:00:00"
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
