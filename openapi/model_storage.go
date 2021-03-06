/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This schema defines a storage subsystem and its respective properties.  A storage subsystem represents a set of storage controllers (physical or virtual) and the resources such as volumes that can be accessed from that subsystem.
type Storage struct {
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
	// The set of drives attached to the storage controllers represented by this resource.
	Drives []IdRef `json:"Drives,omitempty"`
	// The number of items in a collection.
	DrivesodataCount int32 `json:"Drives@odata.count,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	Links Links2 `json:"Links,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// Redundancy information for the storage subsystem.
	Redundancy []IdRef `json:"Redundancy,omitempty"`
	// The number of items in a collection.
	RedundancyodataCount int32 `json:"Redundancy@odata.count,omitempty"`
	Status Status `json:"Status,omitempty"`
	// The set of storage controllers represented by this resource.
	StorageControllers []StorageController `json:"StorageControllers,omitempty"`
	// The number of items in a collection.
	StorageControllersodataCount int32 `json:"StorageControllers@odata.count,omitempty"`
	Volumes IdRef `json:"Volumes,omitempty"`
}
