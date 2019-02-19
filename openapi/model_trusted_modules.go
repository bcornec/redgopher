/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This object describes the inventory of a Trusted Modules installed in the system.
type TrustedModules struct {
	// The firmware version of this Trusted Module.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`
	// The 2nd firmware version of this Trusted Module, if applicable.
	FirmwareVersion2 string `json:"FirmwareVersion2,omitempty"`
	InterfaceType InterfaceType `json:"InterfaceType,omitempty"`
	InterfaceTypeSelection InterfaceTypeSelection `json:"InterfaceTypeSelection,omitempty"`
	Oem string `json:"Oem,omitempty"`
	Status Status `json:"Status,omitempty"`
}
