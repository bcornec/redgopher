/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type MemoryClassification string

// List of MemoryClassification
const (
	MemoryClassificationVOLATILE MemoryClassification = "Volatile"
	MemoryClassificationBYTE_ACCESSIBLE_PERSISTENT MemoryClassification = "ByteAccessiblePersistent"
	MemoryClassificationBLOCK MemoryClassification = "Block"
)