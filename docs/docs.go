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
            "name": "Joshua Ong",
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
        "/addmovie/": {
            "post": {
                "description": "Create new movie based on parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "Movie Data",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully create a new movie with the specified movieid and moviename",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failed to create a new movie because at least one of the parameters is missing",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/deletemovie/{movieid}": {
            "delete": {
                "description": "Delete movie based on movieid",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie ID",
                        "name": "movieid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully delete a movie with the specified movieid",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "movieid was not provided",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Failed to delete a movie with the specified movieid because no movie has that id",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/deletemovies/": {
            "delete": {
                "description": "Delete all movies from database",
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "Succesfully delete all movies",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Failed to delete all movies",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/getmovie/{movieid}/": {
            "get": {
                "description": "Get a movie by its movieid",
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie ID",
                        "name": "movieid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get a movie with the specified movieid",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "movieid was not provided",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Failed to get a movie with the specified movieid because no movie has that id",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/movies/": {
            "get": {
                "description": "Get all movies from the database",
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get all movies",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Failed to get all movies",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Movie"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "type": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.JsonResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Movie"
                    }
                },
                "message": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "main.Movie": {
            "type": "object",
            "properties": {
                "movieid": {
                    "type": "string"
                },
                "moviename": {
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
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Golang Mux Movies API",
	Description:      "This is a movies API server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
