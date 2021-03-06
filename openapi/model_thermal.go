/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This is the schema definition for the Thermal properties.  It represents the properties for Temperature and Cooling.
type Thermal struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions ThermalActions `json:"Actions,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// This is the definition for fans.
	Fans []Fan `json:"Fans,omitempty"`
	// The number of items in a collection.
	FansodataCount int32 `json:"Fans@odata.count,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups.
	Redundancy []IdRef `json:"Redundancy,omitempty"`
	// The number of items in a collection.
	RedundancyodataCount int32 `json:"Redundancy@odata.count,omitempty"`
	Status Status `json:"Status,omitempty"`
	// This is the definition for temperature sensors.
	Temperatures []Temperature `json:"Temperatures,omitempty"`
	// The number of items in a collection.
	TemperaturesodataCount int32 `json:"Temperatures@odata.count,omitempty"`
}
