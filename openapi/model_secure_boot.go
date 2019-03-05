/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This resource contains UEFI Secure Boot information. It represents properties for managing the UEFI Secure Boot functionality of a system.
type SecureBoot struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	SecureBootCurrentBoot SecureBootCurrentBootType `json:"SecureBootCurrentBoot,omitempty"`
	// Enable or disable UEFI Secure Boot (takes effect on next boot).
	SecureBootEnable bool `json:"SecureBootEnable,omitempty"`
	SecureBootMode SecureBootModeType `json:"SecureBootMode,omitempty"`
}
