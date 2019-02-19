/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This object represents a single IPv6 static address to be assigned on a network interface.
type IPv6StaticAddress struct {
	// A valid IPv6 address.
	Address string `json:"Address"`
	Oem string `json:"Oem,omitempty"`
	PrefixLength float32 `json:"PrefixLength"`
}
