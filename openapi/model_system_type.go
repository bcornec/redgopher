/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type SystemType string

// List of SystemType
const (
	SystemTypePHYSICAL SystemType = "Physical"
	SystemTypeVIRTUAL SystemType = "Virtual"
	SystemTypeOS SystemType = "OS"
	SystemTypePHYSICALLY_PARTITIONED SystemType = "PhysicallyPartitioned"
	SystemTypeVIRTUALLY_PARTITIONED SystemType = "VirtuallyPartitioned"
	SystemTypeCOMPOSED SystemType = "Composed"
)