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
type PcIeLinks struct {
	Oem string `json:"Oem,omitempty"`
	// An array of references to the PCIe Devices contained in this slot.
	PCIeDevice []IdRef `json:"PCIeDevice,omitempty"`
	// The number of items in a collection.
	PCIeDeviceodataCount int32 `json:"PCIeDevice@odata.count,omitempty"`
}
