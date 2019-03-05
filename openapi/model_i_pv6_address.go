/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This type describes an IPv6 Address.
type IPv6Address struct {
	// This is the IPv6 Address.
	Address string `json:"Address,omitempty"`
	AddressOrigin IPv6AddressOrigin `json:"AddressOrigin,omitempty"`
	AddressState AddressState `json:"AddressState,omitempty"`
	Oem Oem `json:"Oem,omitempty"`
	PrefixLength float32 `json:"PrefixLength,omitempty"`
}
