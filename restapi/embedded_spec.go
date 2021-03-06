// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Go Two API",
    "title": "Go Two",
    "contact": {
      "name": "Go Two API Support"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "operationId": "Default",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "string"
                },
                "success": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Go Two API",
    "title": "Go Two",
    "contact": {
      "name": "Go Two API Support"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "operationId": "Default",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "type": "string"
                },
                "success": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    }
  }
}`))
}
