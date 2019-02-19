/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The BootOption resource reports information about a single BootOption contained within a system.
type BootOption struct {
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id,omitempty"`
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type,omitempty"`
	Actions Actions2 `json:"Actions,omitempty"`
	Alias BootSource `json:"Alias,omitempty"`
	// A flag that shows if the Boot Option is enabled.
	BootOptionEnabled bool `json:"BootOptionEnabled,omitempty"`
	// The unique boot option string that is referenced in the BootOrder.
	BootOptionReference string `json:"BootOptionReference"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// The user-readable display string of the Boot Option.
	DisplayName string `json:"DisplayName,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem string `json:"Oem,omitempty"`
	// The ID(s) of the resources associated with this Boot Option.
	RelatedItem []IdRef2 `json:"RelatedItem,omitempty"`
	// The number of items in a collection.
	RelatedItemodataCount float32 `json:"RelatedItem@odata.count,omitempty"`
	// The UEFI device path used to access this UEFI Boot Option.
	UefiDevicePath string `json:"UefiDevicePath,omitempty"`
}
