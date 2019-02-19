/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type EncryptionStatus string

// List of EncryptionStatus
const (
	EncryptionStatusUNECRYPTED EncryptionStatus = "Unecrypted"
	EncryptionStatusUNLOCKED EncryptionStatus = "Unlocked"
	EncryptionStatusLOCKED EncryptionStatus = "Locked"
	EncryptionStatusFOREIGN EncryptionStatus = "Foreign"
	EncryptionStatusUNENCRYPTED EncryptionStatus = "Unencrypted"
)