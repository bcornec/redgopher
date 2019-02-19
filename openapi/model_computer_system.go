/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This schema defines a computer system and its respective properties.  A computer system represents a machine (physical or virtual) and the local resources such as memory, cpu and other devices that can be accessed from that machine.
type ComputerSystem struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// The user definable tag that can be used to track this computer system for inventory or other client purposes.
	AssetTag string `json:"AssetTag,omitempty"`
	Bios IdRef `json:"Bios,omitempty"`
	// The version of the system BIOS or primary system firmware.
	BiosVersion string `json:"BiosVersion,omitempty"`
	Boot Boot `json:"Boot,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	EthernetInterfaces IdRef `json:"EthernetInterfaces,omitempty"`
	// The DNS Host Name, without any domain information.
	HostName string `json:"HostName,omitempty"`
	HostWatchdogTimer WatchdogTimer `json:"HostWatchdogTimer,omitempty"`
	HostedServices HostedServices `json:"HostedServices,omitempty"`
	// The hosing roles that this computer system supports.
	HostingRoles []HostingRole `json:"HostingRoles,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	IndicatorLED IndicatorLed2 `json:"IndicatorLED,omitempty"`
	Links Links2 `json:"Links,omitempty"`
	LogServices IdRef `json:"LogServices,omitempty"`
	// The manufacturer or OEM of this system.
	Manufacturer string `json:"Manufacturer,omitempty"`
	Memory IdRef `json:"Memory,omitempty"`
	MemoryDomains IdRef `json:"MemoryDomains,omitempty"`
	MemorySummary MemorySummary `json:"MemorySummary,omitempty"`
	// The product name for this system, without the manufacturer name.
	Model string `json:"Model,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	NetworkInterfaces IdRef `json:"NetworkInterfaces,omitempty"`
	Oem string `json:"Oem,omitempty"`
	// A reference to a collection of PCIe Devices used by this computer system.
	PCIeDevices []IdRef `json:"PCIeDevices,omitempty"`
	// The number of items in a collection.
	PCIeDevicesodataCount int32 `json:"PCIeDevices@odata.count,omitempty"`
	// A reference to a collection of PCIe Functions used by this computer system.
	PCIeFunctions []IdRef `json:"PCIeFunctions,omitempty"`
	// The number of items in a collection.
	PCIeFunctionsodataCount int32 `json:"PCIeFunctions@odata.count,omitempty"`
	// The part number for this system.
	PartNumber string `json:"PartNumber,omitempty"`
	PowerRestorePolicy PowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`
	PowerState PowerState2 `json:"PowerState,omitempty"`
	ProcessorSummary ProcessorSummary `json:"ProcessorSummary,omitempty"`
	Processors IdRef `json:"Processors,omitempty"`
	// A reference to a collection of Redundancy entities that each name a set of computer systems that provide redundancy for this ComputerSystem.
	Redundancy []IdRef `json:"Redundancy,omitempty"`
	// The number of items in a collection.
	RedundancyodataCount int32 `json:"Redundancy@odata.count,omitempty"`
	// The manufacturer SKU for this system.
	SKU string `json:"SKU,omitempty"`
	SecureBoot IdRef `json:"SecureBoot,omitempty"`
	// The serial number for this system.
	SerialNumber string `json:"SerialNumber,omitempty"`
	SimpleStorage IdRef `json:"SimpleStorage,omitempty"`
	Status Status `json:"Status,omitempty"`
	Storage IdRef `json:"Storage,omitempty"`
	// The sub-model for this system.
	SubModel string `json:"SubModel,omitempty"`
	SystemType SystemType `json:"SystemType,omitempty"`
	// This object describes the array of Trusted Modules in the system.
	TrustedModules []TrustedModules `json:"TrustedModules,omitempty"`
	UUID string `json:"UUID,omitempty"`
}