/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This defines the action for securely erasing given regions using the NIST SP800-88 Purge: Overwrite.
type OverwriteUnitRequestBody struct {
	// Passphrase for doing the operation.
	Passphrase string `json:"Passphrase"`
	// Memory region ID for which this action to be applied.
	RegionId string `json:"RegionId"`
}