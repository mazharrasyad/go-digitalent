// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://google.com",
        "contact": {
            "name": "API Support",
            "email": "muhazharrasyad@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comments": {
            "get": {
                "description": "Get details of all comments data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Get all comments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Create comment",
                "parameters": [
                    {
                        "description": "Create comment",
                        "name": "models.Comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comments/{Id}": {
            "patch": {
                "description": "Update comment data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Update comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comments/{commentId}": {
            "get": {
                "description": "Get details of specific comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Get comment by comment id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete comment data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Delete comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment",
                        "name": "commentId",
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
        "/photos": {
            "get": {
                "description": "Get details of all photos data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Get all photos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Create photo",
                "parameters": [
                    {
                        "description": "Create photo",
                        "name": "models.Photo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photos/{Id}": {
            "patch": {
                "description": "Update photo data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Update photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photos/{photoId}": {
            "get": {
                "description": "Get details of specific photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Get photo by photo id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete photo data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Delete photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo",
                        "name": "photoId",
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
        "/register": {
            "post": {
                "description": "User Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "User Login",
                        "name": "models.User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/social-medias": {
            "get": {
                "description": "Get details of all socialMedias data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Get all socialMedias",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new socialMedia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Create socialMedia",
                "parameters": [
                    {
                        "description": "Create socialMedia",
                        "name": "models.SocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/social-medias/{Id}": {
            "patch": {
                "description": "Update socialMedia data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Update socialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the socialMedia",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/social-medias/{socialMediaId}": {
            "get": {
                "description": "Get details of specific socialMedia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Get socialMedia by socialMedia id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the socialMedia",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete socialMedia data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "json"
                ],
                "summary": "Delete socialMedia",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the socialMedia",
                        "name": "socialMediaId",
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
        "models.Comment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo": {
                    "$ref": "#/definitions/models.Photo"
                },
                "photo_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Photo"
                    }
                },
                "socialmedias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMedia"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
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
	Schemes:          []string{},
	Title:            "MyGram API",
	Description:      "Service to manage data",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
