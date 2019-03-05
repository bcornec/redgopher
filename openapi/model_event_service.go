/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The Event Service resource contains properties for managing event subcriptions and generates the events sent to subscribers.  The resource has links to the actual collection of subscriptions (called Event Destinations).
type EventService struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// This is the number of attempts an event posting is retried before the subscription is terminated.  This retry is at the service level, meaning the HTTP POST to the Event Destination was returned by the HTTP operation as unsuccessful (4xx or 5xx return code) or an HTTP timeout occurred this many times before the Event Destination subscription is terminated.
	DeliveryRetryAttempts int32 `json:"DeliveryRetryAttempts,omitempty"`
	// This represents the number of seconds between retry attempts for sending any given Event.
	DeliveryRetryIntervalSeconds int32 `json:"DeliveryRetryIntervalSeconds,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Indicates the content types of the message that this service can send to the event destination.
	EventFormatTypes []EventFormatType `json:"EventFormatTypes,omitempty"`
	// This is the types of Events that can be subscribed to.
	EventTypesForSubscription []EventType `json:"EventTypesForSubscription,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// A list of the Prefixes of the Message Registries that can be used for the RegistryPrefix property on a subscription.
	RegistryPrefixes []string `json:"RegistryPrefixes,omitempty"`
	// A list of @odata.type values (Schema names) that can be specified in a ResourceType on a subscription.
	ResourceTypes []string `json:"ResourceTypes,omitempty"`
	SSEFilterPropertiesSupported SseFilterPropertiesSupported `json:"SSEFilterPropertiesSupported,omitempty"`
	// Link to a URI for receiving Sever Sent Event representations of the events generated by this service.
	ServerSentEventUri string `json:"ServerSentEventUri,omitempty"`
	// This indicates whether this service is enabled.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`
	Status Status `json:"Status,omitempty"`
	// This indicates if the service supports the SubordinateResource property on Event Subscriptions.
	SubordinateResourcesSupported bool `json:"SubordinateResourcesSupported,omitempty"`
	Subscriptions IdRef `json:"Subscriptions,omitempty"`
}
