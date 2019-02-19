/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// A threshold definition for a sensor.
type Threshold struct {
	Activation ThresholdActivation `json:"Activation,omitempty"`
	// The time interval over which the sensor reading must have passed through this Threshold value before the threshold is considered to be violated.
	DwellTime string `json:"DwellTime,omitempty"`
	// The threshold value.
	Reading float32 `json:"Reading,omitempty"`
}