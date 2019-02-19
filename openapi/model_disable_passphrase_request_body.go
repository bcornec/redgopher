/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Disable passphrase for given regions.
type DisablePassphraseRequestBody struct {
	// Passphrase for doing the operation.
	Passphrase string `json:"Passphrase"`
	// Memory region ID for which this action to be applied.
	RegionId string `json:"RegionId"`
}
