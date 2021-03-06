/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// An Event Destination desribes the target of an event subscription, including the types of events subscribed and context to provide to the target in the Event payload.
type EventDestination struct {
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id,omitempty"`
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type,omitempty"`
	// A client-supplied string that is stored with the event destination subscription.
	Context string `json:"Context,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// The URI of the destination Event Service.
	Destination string `json:"Destination,omitempty"`
	// This property shall contain the types of events that shall be sent to the desination.
	EventTypes []EventType2 `json:"EventTypes,omitempty"`
	// This is for setting HTTP headers, such as authorization information.  This object will be null on a GET.
	HttpHeaders []HttpHeaderProperty `json:"HttpHeaders,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	Protocol EventDestinationProtocol `json:"Protocol,omitempty"`
}
