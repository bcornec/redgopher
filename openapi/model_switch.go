/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Switch contains properties describing a simple fabric switch.
type Switch struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// The user assigned asset tag for this switch.
	AssetTag string `json:"AssetTag,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// The Domain ID for this switch.
	DomainID int32 `json:"DomainID,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	IndicatorLED IndicatorLed2 `json:"IndicatorLED,omitempty"`
	// This indicates whether the switch is in a managed or unmanaged state.
	IsManaged bool `json:"IsManaged,omitempty"`
	Links Links2 `json:"Links,omitempty"`
	Location Location `json:"Location,omitempty"`
	LogServices IdRef `json:"LogServices,omitempty"`
	// This is the manufacturer of this switch.
	Manufacturer string `json:"Manufacturer,omitempty"`
	// The product model number of this switch.
	Model string `json:"Model,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// The part number for this switch.
	PartNumber string `json:"PartNumber,omitempty"`
	Ports IdRef `json:"Ports,omitempty"`
	PowerState PowerState2 `json:"PowerState,omitempty"`
	// Redundancy information for the switches.
	Redundancy []IdRef `json:"Redundancy,omitempty"`
	// The number of items in a collection.
	RedundancyodataCount int32 `json:"Redundancy@odata.count,omitempty"`
	// This is the SKU for this switch.
	SKU string `json:"SKU,omitempty"`
	// The serial number for this switch.
	SerialNumber string `json:"SerialNumber,omitempty"`
	Status Status `json:"Status,omitempty"`
	SwitchType IdRef `json:"SwitchType,omitempty"`
	// The total number of lanes, phys, or other physical transport links that this switch contains.
	TotalSwitchWidth int32 `json:"TotalSwitchWidth,omitempty"`
}
