// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/notification": {
            "post": {
                "description": "Create an Notification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Create Notification",
                "parameters": [
                    {
                        "description": "Notification request body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNotification"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/notification/device/{device_reference}/{page}": {
            "get": {
                "description": "Retrieve the Notification list for a user's device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Get Notification list by user device reference",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device Reference",
                        "name": "device_reference",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page",
                        "name": "page",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/notification/user/{user_reference}/{page}": {
            "get": {
                "description": "Retrieve the Notification list for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Get Notification list by user reference",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Reference",
                        "name": "user_reference",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page",
                        "name": "device_reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/notification/{reference}": {
            "get": {
                "description": "Retrieve the notification by reference",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Get Notification by reference",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reference",
                        "name": "Reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/helper.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateNotification": {
            "type": "object",
            "required": [
                "channel",
                "contact",
                "message_body",
                "notification_type",
                "notified_by",
                "notify_on",
                "subject",
                "user_reference"
            ],
            "properties": {
                "channel": {
                    "$ref": "#/definitions/shared.Channel"
                },
                "contact": {
                    "type": "string"
                },
                "device_reference": {
                    "type": "string"
                },
                "message_body": {
                    "type": "string"
                },
                "notification_type": {
                    "$ref": "#/definitions/shared.NotificationType"
                },
                "notified_by": {
                    "type": "string",
                    "maxLength": 38,
                    "minLength": 26
                },
                "notify_on": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "user_reference": {
                    "type": "string",
                    "maxLength": 38,
                    "minLength": 26
                }
            }
        },
        "helper.ErrorBody": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Code    string      ` + "`" + `json:\"code\"` + "`" + `"
                }
            }
        },
        "helper.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error_reference": {
                    "type": "string"
                },
                "error_type": {
                    "type": "string"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/helper.ErrorBody"
                    }
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "shared.Channel": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "Sms",
                "Email"
            ]
        },
        "shared.NotificationType": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "Instant",
                "Scheduled"
            ]
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
