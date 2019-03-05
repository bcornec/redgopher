/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Contains references to other resources that are related to this resource.
type ManagerAccount2Links struct {
	Oem Oem `json:"Oem,omitempty"`
	Role IdRef `json:"Role,omitempty"`
}
