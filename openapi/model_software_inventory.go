/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// This schema defines an inventory of software components.
type SoftwareInventory struct {
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
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// A string representing the lowest supported version of this software.
	LowestSupportedVersion string `json:"LowestSupportedVersion,omitempty"`
	// A string representing the manufacturer/producer of this software.
	Manufacturer string `json:"Manufacturer,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// The ID(s) of the resources associated with this software inventory item.
	RelatedItem []IdRef `json:"RelatedItem,omitempty"`
	// The number of items in a collection.
	RelatedItemodataCount int32 `json:"RelatedItem@odata.count,omitempty"`
	// Release date of this software.
	ReleaseDate time.Time `json:"ReleaseDate,omitempty"`
	// A string representing the implementation-specific ID for identifying this software.
	SoftwareId string `json:"SoftwareId,omitempty"`
	Status Status `json:"Status,omitempty"`
	// A list of strings representing the UEFI Device Path(s) of the component(s) associated with this software inventory item.
	UefiDevicePaths []string `json:"UefiDevicePaths,omitempty"`
	// Indicates whether this software can be updated by the update service.
	Updateable bool `json:"Updateable,omitempty"`
	// A string representing the version of this software.
	Version string `json:"Version,omitempty"`
}
