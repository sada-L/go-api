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
        "/image/multi": {
            "post": {
                "security": [
                    {
                        "Bearer Authentication": []
                    }
                ],
                "description": "Загружает несколько выбранных изображений",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Загрузка нескольких изображений",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": "Image files",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "'Bearer _YOUR_TOKEN_'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Images successfully uploaded",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/image/multi/zip": {
            "post": {
                "security": [
                    {
                        "Bearer Authentication": []
                    }
                ],
                "description": "Загружает несколько выбранных изображений с последующем сжатием",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Загрузка нескольких изображений с сжатием",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": "Image files",
                        "name": "files",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "'Bearer _YOUR_TOKEN_'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Images successfully uploaded",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/image/single": {
            "post": {
                "security": [
                    {
                        "Bearer Authentication": []
                    }
                ],
                "description": "Загружает выбранное изображение.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Загрузить изображение.",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "'Bearer _YOUR_TOKEN_'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/image/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer Authentication": []
                    }
                ],
                "description": "Возвращает изображение по ID.",
                "produces": [
                    "image/png",
                    "image/jpeg"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Получить изображение по ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "'Bearer _YOUR_TOKEN_'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image received",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer Authentication": []
                    }
                ],
                "description": "Удаляет изображение по ID.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Удалить изображение по ID.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Image ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "'Bearer _YOUR_TOKEN_'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Авторизация по логину и паролю.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Авторизация.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT tokens",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "Bearer Authentication": []
                    }
                ],
                "description": "Просморт профиля пользователя по токену доступа.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Просмотреть профиль.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "'Bearer _YOUR_TOKEN_'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile",
                        "schema": {
                            "$ref": "#/definitions/domain.Profile"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "Обновить токены с помощью токена обновления.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Обновление токенов.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Refresh token",
                        "name": "refreshToken",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Tokens",
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Регистарация по логину и паролю.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Регистрация.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT tokens",
                        "schema": {
                            "$ref": "#/definitions/domain.SignupResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "domain.Profile": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.RefreshTokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "domain.SignupResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "GO API",
	Description:      "Server for an image management API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
