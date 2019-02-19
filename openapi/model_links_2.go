/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Contains references to other resources that are related to this resource.
type Links2 struct {
	// This property is an array of references to the certificates installed on this service.
	Certificates []IdRef `json:"Certificates,omitempty"`
	// The number of items in a collection.
	CertificatesodataCount int32 `json:"Certificates@odata.count,omitempty"`
	Oem string `json:"Oem,omitempty"`
}
