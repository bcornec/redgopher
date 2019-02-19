/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The properties of the FPGA device.
type Fpga struct {
	// An array of the FPGA external interfaces.
	ExternalInterfaces []FpgaInterface `json:"ExternalInterfaces,omitempty"`
	// The FPGA firmware identifier.
	FirmwareId string `json:"FirmwareId,omitempty"`
	// The FPGA firmware manufacturer.
	FirmwareManufacturer string `json:"FirmwareManufacturer,omitempty"`
	// The FPGA firmware version.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`
	FpgaType FpgaType `json:"FpgaType,omitempty"`
	HostInterface FpgaInterface `json:"HostInterface,omitempty"`
	// The FPGA model.
	Model string `json:"Model,omitempty"`
	Oem string `json:"Oem,omitempty"`
	// The number of the PCIe Virtual Functions.
	PCIeVirtualFunctions int32 `json:"PCIeVirtualFunctions,omitempty"`
	// This flag indicates if the FPGA firmware can be reprogrammed from the host using system software.
	ProgrammableFromHost bool `json:"ProgrammableFromHost,omitempty"`
	// An array of the FPGA reconfiguration slots.  A reconfiguration slot is used by an FPGA to contain an acceleration function that can change as the FPGA is being provisioned.
	ReconfigurationSlots []FpgaReconfigurationSlot `json:"ReconfigurationSlots,omitempty"`
}
