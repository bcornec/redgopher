/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This action is used to replace an existing certificate.
type ReplaceCertificateRequestBody struct {
	// The string for the certificate.
	CertificateString string `json:"CertificateString"`
	CertificateType CertificateType `json:"CertificateType"`
	CertificateUri IdRef `json:"CertificateUri"`
}
