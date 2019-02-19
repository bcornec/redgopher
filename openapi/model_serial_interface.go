/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This schema defines an asynchronous serial interface resource.
type SerialInterface struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	BitRate BitRate `json:"BitRate,omitempty"`
	ConnectorType ConnectorType `json:"ConnectorType,omitempty"`
	DataBits DataBits `json:"DataBits,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	FlowControl FlowControl2 `json:"FlowControl,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// This indicates whether this interface is enabled.
	InterfaceEnabled bool `json:"InterfaceEnabled,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem string `json:"Oem,omitempty"`
	Parity Parity `json:"Parity,omitempty"`
	PinOut PinOut `json:"PinOut,omitempty"`
	SignalType SignalType `json:"SignalType,omitempty"`
	StopBits StopBits `json:"StopBits,omitempty"`
}
