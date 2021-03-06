/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This object describes the central processors of the system in general detail.
type ProcessorSummary struct {
	// The number of physical processors in the system.
	Count int32 `json:"Count,omitempty"`
	// The number of logical processors in the system.
	LogicalProcessorCount int32 `json:"LogicalProcessorCount,omitempty"`
	// The processor model for the primary or majority of processors in this system.
	Model string `json:"Model,omitempty"`
	Status Status `json:"Status,omitempty"`
}
