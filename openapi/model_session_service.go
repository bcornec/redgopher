/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This is the schema definition for the Session Service.  It represents the properties for the service itself and has links to the actual list of sessions.
type SessionService struct {
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
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem string `json:"Oem,omitempty"`
	// This indicates whether this service is enabled.  If set to false, the Session Service is disabled, and new sessions cannot be created, old sessions cannot be deleted, and established sessions may continue operating.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`
	// This is the number of seconds of inactivity that a session may have before the session service closes the session due to inactivity.
	SessionTimeout int32 `json:"SessionTimeout,omitempty"`
	Sessions IdRef `json:"Sessions,omitempty"`
	Status Status `json:"Status,omitempty"`
}
