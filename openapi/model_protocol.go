/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Settings for a network protocol associated with a manager.
type Protocol struct {
	// Indicates the protocol port.
	Port int32 `json:"Port,omitempty"`
	// Indicates if the protocol is enabled or disabled.
	ProtocolEnabled bool `json:"ProtocolEnabled,omitempty"`
}
