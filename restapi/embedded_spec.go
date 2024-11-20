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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The Plutonium Service API provides endpoints to support the operations of the NFT Marketplace.\nThis document outlines the API's structure, response formats, and capabilities for integration.\n",
    "title": "Plutonium Service API",
    "version": "1.0.0"
  },
  "host": "plutonium",
  "basePath": "/api/v1",
  "paths": {
    "/metrics": {
      "get": {
        "description": "This endpoint provides Prometheus-compatible metrics for monitoring the application. \nIt is typically used by Prometheus or similar monitoring tools to scrape metrics data.\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Monitoring"
        ],
        "summary": "Retrieve Prometheus Metrics",
        "responses": {
          "200": {
            "description": "Metrics retrieved successfully.",
            "schema": {
              "$ref": "#/definitions/PrometheusResponse"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "description": "The ping endpoint provides a simple \"pong\" response to confirm server availability.\nUse this endpoint to:\n- Verify connectivity between the client and server.\n- Measure network latency for diagnostics.\n- Perform quick and reliable health checks over HTTP.\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Public"
        ],
        "summary": "Ping endpoint for server health and latency testing.",
        "responses": {
          "200": {
            "description": "Successful response indicating server availability.",
            "schema": {
              "$ref": "#/definitions/PingResponse"
            }
          },
          "500": {
            "description": "Internal Server Error. This typically indicates a server-side issue or\nan unexpected runtime error preventing proper functionality.\n",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "description": "HTTP status code representing the type of error encountered.",
          "type": "integer",
          "example": 500
        },
        "details": {
          "description": "Additional context or information about the error, if available.",
          "type": "string",
          "example": "Unexpected error while processing the request."
        },
        "message": {
          "description": "Brief explanation of the error encountered.",
          "type": "string",
          "example": "Internal Server Error"
        }
      }
    },
    "PingResponse": {
      "type": "object",
      "properties": {
        "message": {
          "description": "Response message confirming successful server connectivity.",
          "type": "string",
          "enum": [
            "pong"
          ],
          "example": "pong"
        },
        "timestamp": {
          "description": "The timestamp of the server response, useful for tracking latency.",
          "type": "string",
          "format": "date-time",
          "example": "2024-11-19T12:34:56Z"
        }
      }
    },
    "PrometheusResponse": {
      "type": "object",
      "additionalProperties": {
        "type": "string",
        "format": "number"
      }
    }
  },
  "tags": [
    {
      "description": "Endpoints accessible to all clients for general API functionality.",
      "name": "Public"
    },
    {
      "description": "Endpoints accessible to all clients for monitoring API functionality.",
      "name": "Monitoring"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The Plutonium Service API provides endpoints to support the operations of the NFT Marketplace.\nThis document outlines the API's structure, response formats, and capabilities for integration.\n",
    "title": "Plutonium Service API",
    "version": "1.0.0"
  },
  "host": "plutonium",
  "basePath": "/api/v1",
  "paths": {
    "/metrics": {
      "get": {
        "description": "This endpoint provides Prometheus-compatible metrics for monitoring the application. \nIt is typically used by Prometheus or similar monitoring tools to scrape metrics data.\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Monitoring"
        ],
        "summary": "Retrieve Prometheus Metrics",
        "responses": {
          "200": {
            "description": "Metrics retrieved successfully.",
            "schema": {
              "$ref": "#/definitions/PrometheusResponse"
            }
          }
        }
      }
    },
    "/ping": {
      "get": {
        "description": "The ping endpoint provides a simple \"pong\" response to confirm server availability.\nUse this endpoint to:\n- Verify connectivity between the client and server.\n- Measure network latency for diagnostics.\n- Perform quick and reliable health checks over HTTP.\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Public"
        ],
        "summary": "Ping endpoint for server health and latency testing.",
        "responses": {
          "200": {
            "description": "Successful response indicating server availability.",
            "schema": {
              "$ref": "#/definitions/PingResponse"
            }
          },
          "500": {
            "description": "Internal Server Error. This typically indicates a server-side issue or\nan unexpected runtime error preventing proper functionality.\n",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "description": "HTTP status code representing the type of error encountered.",
          "type": "integer",
          "example": 500
        },
        "details": {
          "description": "Additional context or information about the error, if available.",
          "type": "string",
          "example": "Unexpected error while processing the request."
        },
        "message": {
          "description": "Brief explanation of the error encountered.",
          "type": "string",
          "example": "Internal Server Error"
        }
      }
    },
    "PingResponse": {
      "type": "object",
      "properties": {
        "message": {
          "description": "Response message confirming successful server connectivity.",
          "type": "string",
          "enum": [
            "pong"
          ],
          "example": "pong"
        },
        "timestamp": {
          "description": "The timestamp of the server response, useful for tracking latency.",
          "type": "string",
          "format": "date-time",
          "example": "2024-11-19T12:34:56Z"
        }
      }
    },
    "PrometheusResponse": {
      "type": "object",
      "additionalProperties": {
        "type": "string",
        "format": "number"
      }
    }
  },
  "tags": [
    {
      "description": "Endpoints accessible to all clients for general API functionality.",
      "name": "Public"
    },
    {
      "description": "Endpoints accessible to all clients for monitoring API functionality.",
      "name": "Monitoring"
    }
  ]
}`))
}
