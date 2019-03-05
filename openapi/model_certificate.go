/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// The Certificate resource describes a certificate used to prove the identify of a component, account, or service.
type Certificate struct {
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// The string for the certificate.
	CertificateString string `json:"CertificateString,omitempty"`
	CertificateType CertificateType `json:"CertificateType,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	Issuer Identifier `json:"Issuer,omitempty"`
	// The usage of the key contained in the certificate.
	KeyUsage []KeyUsage `json:"KeyUsage,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	Subject Identifier `json:"Subject,omitempty"`
	// The date when the certificate is no longer valid.
	ValidNotAfter time.Time `json:"ValidNotAfter,omitempty"`
	// The date when the certificate becomes valid.
	ValidNotBefore time.Time `json:"ValidNotBefore,omitempty"`
}
