/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Contains settings for parsing a generic LDAP service.
type LdapService struct {
	Oem Oem `json:"Oem,omitempty"`
	SearchSettings LdapSearchSettings `json:"SearchSettings,omitempty"`
}
