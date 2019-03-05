/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Volume contains properties used to describe a volume, virtual disk, LUN, or other logical storage entity for any system.
type Volume2 struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// The size of the smallest addressable unit (Block) of this volume in bytes.
	BlockSizeBytes int32 `json:"BlockSizeBytes,omitempty"`
	// The size in bytes of this Volume.
	CapacityBytes int32 `json:"CapacityBytes,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Is this Volume encrypted.
	Encrypted bool `json:"Encrypted,omitempty"`
	// The types of encryption used by this Volume.
	EncryptionTypes []EncryptionTypes `json:"EncryptionTypes,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The Durable names for the volume.
	Identifiers []Identifier2 `json:"Identifiers,omitempty"`
	Links Links2 `json:"Links,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// The operations currently running on the Volume.
	Operations []Operation `json:"Operations,omitempty"`
	// The size in bytes of this Volume's optimum IO size.
	OptimumIOSizeBytes int32 `json:"OptimumIOSizeBytes,omitempty"`
	Status Status `json:"Status,omitempty"`
	VolumeType VolumeType `json:"VolumeType,omitempty"`
}
