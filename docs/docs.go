// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"


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
}

func init() {
    swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

const docTemplate = `{
  "swagger": "2.0",
  "info": {
    "version": "1.0",
    "title": "goapp",
    "contact": {}
  },
  "host": "localhost:1323",
  "basePath": "/api",
  "securityDefinitions": {},
  "schemes": [
    "http"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/user/create": {
      "post": {
        "summary": "register",
        "tags": [
          "auth"
        ],
        "operationId": "register",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "email",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "password",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/login": {
      "post": {
        "summary": "login",
        "tags": [
          "auth"
        ],
        "operationId": "login",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "parameters": [
          {
            "name": "email",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          },
          {
            "name": "password",
            "in": "formData",
            "required": true,
            "type": "string",
            "description": ""
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    },
    "/admin/page": {
      "get": {
        "summary": "rute operator",
        "tags": [
          "operator"
        ],
        "operationId": "ruteoperator",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [

          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          },
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        },
        "security": []
      }
    },
    "/user/all": {
      "get": {
        "summary": "Get users",
        "tags": [
          "Misc"
        ],
        "operationId": "Getusers",
        "deprecated": false,
        "produces": [
          "application/json"
        ],
        "parameters": [
          {
            "name": "Authorization",
            "in": "header",
            "required": false,
            "default": "Bearer {token}",
            "type": "string"
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "headers": {}
          }
        }
      }
    }
  },
  "security": [],
  "tags": [
    {
      "name": "auth"
    },
    {
      "name": "operator"
    },
    {
      "name": "Misc",
      "description": ""
    }
  ]
}`