/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This type describes the location of a resource.
type Location struct {
	// This indicates the location of the resource.
	Info string `json:"Info,omitempty"`
	// This represents the format of the Info property.
	InfoFormat string `json:"InfoFormat,omitempty"`
	Oem Oem `json:"Oem,omitempty"`
}
