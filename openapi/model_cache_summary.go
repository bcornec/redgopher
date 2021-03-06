/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This object describes the cache memory of the storage controller in general detail.
type CacheSummary struct {
	// The portion of the cache memory that is persistent, measured in MiB.
	PersistentCacheSizeMiB int32 `json:"PersistentCacheSizeMiB,omitempty"`
	Status Status `json:"Status,omitempty"`
	// The total configured cache memory, measured in MiB.
	TotalCacheSizeMiB int32 `json:"TotalCacheSizeMiB"`
}
