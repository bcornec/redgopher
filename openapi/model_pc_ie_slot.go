/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This is the definition for a PCI slot information object.
type PcIeSlot struct {
	// This is the number of PCIe lanes supported by this slot.
	Lanes int32 `json:"Lanes,omitempty"`
	Links PcIeLinks `json:"Links,omitempty"`
	Location Location `json:"Location,omitempty"`
	Oem Oem `json:"Oem,omitempty"`
	PCIeType PcIeTypes `json:"PCIeType,omitempty"`
	SlotType SlotTypes `json:"SlotType,omitempty"`
	Status Status `json:"Status,omitempty"`
}
