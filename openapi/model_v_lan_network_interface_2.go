/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This resource contains information for a Virtual LAN (VLAN) network instance available on a manager, system or other device.
type VLanNetworkInterface2 struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// This indicates if this VLAN is enabled.
	VLANEnable bool `json:"VLANEnable,omitempty"`
	VLANId float32 `json:"VLANId,omitempty"`
}
