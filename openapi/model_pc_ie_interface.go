/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This type describes a PCIe Interface.
type PcIeInterface struct {
	// This is the number of PCIe lanes in use by this device.
	LanesInUse int32 `json:"LanesInUse,omitempty"`
	// This is the number of PCIe lanes supported by this device.
	MaxLanes int32 `json:"MaxLanes,omitempty"`
	MaxPCIeType PcIeTypes `json:"MaxPCIeType,omitempty"`
	Oem string `json:"Oem,omitempty"`
	PCIeType PcIeTypes `json:"PCIeType,omitempty"`
}
