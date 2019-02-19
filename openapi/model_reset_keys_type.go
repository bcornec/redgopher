/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type ResetKeysType string

// List of ResetKeysType
const (
	ResetKeysTypeRESET_ALL_KEYS_TO_DEFAULT ResetKeysType = "ResetAllKeysToDefault"
	ResetKeysTypeDELETE_ALL_KEYS ResetKeysType = "DeleteAllKeys"
	ResetKeysTypeDELETE_PK ResetKeysType = "DeletePK"
)