package config

var schemav1beta1 string = `
{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/Config",
    "definitions": {
        "Config": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "apiVersion": {
                    "type": "string"
                },
                "kind": {
                    "type": "string"
                },
                "metadata": {
                    "$ref": "#/definitions/Metadata"
                },
                "spec": {
                    "$ref": "#/definitions/Spec"
                }
            },
            "required": [
                "apiVersion",
                "kind",
                "metadata",
                "spec"
            ],
            "title": "Config"
        },
        "Metadata": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "name": {
                    "type": "string"
                },
                "environment": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            },
            "required": [
                "description",
                "environment",
                "name"
            ],
            "title": "Metadata"
        },
        "Spec": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "backend": {
                    "$ref": "#/definitions/Backend"
                }
            },
            "required": [
                "backend"
            ],
            "title": "Spec"
        },
        "Backend": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "azure": {
                    "$ref": "#/definitions/Azure"
                }
            },
            "title": "Backend"
        },
        "Azure": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "autoCreateStorage": {
                    "type": "boolean"
                },
                "storageAccountRg": {
                    "type": "string"
                },
                "storageAccountName": {
                    "type": "string"
                },
                "storageAccountContainer": {
                    "type": "string"
                },
                "stateFileName": {
                    "type": "string"
                },
                "subscriptionId": {
                    "type": "string",
                    "format": "uuid"
                },
                "tenantId": {
                    "type": "string",
                    "format": "uuid"
                },
                "credentials": {
                    "$ref": "#/definitions/Credentials"
                }
            },
            "required": [
                "autoCreateStorage",
                "credentials",
                "stateFileName",
                "storageAccountContainer",
                "storageAccountName",
                "storageAccountRg",
                "subscriptionId",
                "tenantId"
            ],
            "title": "Azure"
        },
        "Credentials": {
            "type": "object",
            "additionalProperties": false,
            "properties": {
                "useAzLogin": {
                    "type": "string"
                },
                "fromEnvironment": {
                    "type": "string"
                },
                "clientIdEnvName": {
                    "type": "string"
                },
                "clientSecretEnvName": {
                    "type": "string"
                }
            },
            "required": [
                "useAzLogin"
            ],
            "title": "Credentials"        }
    }
}
`
