/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The state of the physical security sensor.
type PhysicalSecurity struct {
	IntrusionSensor IntrusionSensor `json:"IntrusionSensor,omitempty"`
	// A numerical identifier to represent the physical security sensor.
	IntrusionSensorNumber int32  `json:"IntrusionSensorNumber,omitempty"`
	IntrusionSensorReArm  string `json:"IntrusionSensorReArm,omitempty"`
}
