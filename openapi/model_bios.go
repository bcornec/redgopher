/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The Bios schema contains properties related to the BIOS Attribute Registry.  The Attribute Registry describes the system-specific BIOS attributes and Actions for changing to BIOS settings.  Changes to the BIOS typically require a system reset before they take effect.  It is likely that a client will find the @Redfish.Settings term in this resource, and if it is found, the client makes requests to change BIOS settings by modifying the resource identified by the @Redfish.Settings term.
type Bios struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// The Resource ID of the Attribute Registry that has the system-specific information about a BIOS resource.
	AttributeRegistry string `json:"AttributeRegistry,omitempty"`
	Attributes Attributes `json:"Attributes,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem string `json:"Oem,omitempty"`
}
