/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// A storage device such as a disk drive or optical media device.
type Device struct {
	// The size of the storage device.
	CapacityBytes int32 `json:"CapacityBytes,omitempty"`
	// The name of the manufacturer of this device.
	Manufacturer string `json:"Manufacturer,omitempty"`
	// The product model number of this device.
	Model string `json:"Model,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	Status Status `json:"Status,omitempty"`
}
