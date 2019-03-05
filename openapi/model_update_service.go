/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This is the schema definition for the Update Service. It represents the properties for the service itself and has links to collections of firmware and software inventory.
type UpdateService struct {
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
	FirmwareInventory IdRef `json:"FirmwareInventory,omitempty"`
	// The URI used to perform an HTTP or HTTPS push update to the Update Service.
	HttpPushUri string `json:"HttpPushUri,omitempty"`
	HttpPushUriOptions HttpPushUriOptions `json:"HttpPushUriOptions,omitempty"`
	// This represents if the properties of HttpPushUriOptions are reserved by any client.
	HttpPushUriOptionsBusy bool `json:"HttpPushUriOptionsBusy,omitempty"`
	// The array of URIs indicating the target for applying the update image.
	HttpPushUriTargets []string `json:"HttpPushUriTargets,omitempty"`
	// This represents if the HttpPushUriTargets property is reserved by any client.
	HttpPushUriTargetsBusy bool `json:"HttpPushUriTargetsBusy,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// This indicates whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`
	SoftwareInventory IdRef `json:"SoftwareInventory,omitempty"`
	Status Status `json:"Status,omitempty"`
}
