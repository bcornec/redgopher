/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The Fabric schema represents a simple fabric consisting of one or more switches, zero or more endpoints, and zero or more zones.
type Fabric struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	Endpoints IdRef `json:"Endpoints,omitempty"`
	FabricType IdRef `json:"FabricType,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	Links Links2 `json:"Links,omitempty"`
	// The value of this property shall contain the maximum number of zones the switch can currently configure.
	MaxZones int32 `json:"MaxZones,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	Status Status `json:"Status,omitempty"`
	Switches IdRef `json:"Switches,omitempty"`
	Zones IdRef `json:"Zones,omitempty"`
}
