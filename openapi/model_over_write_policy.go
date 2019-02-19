/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type OverWritePolicy string

// List of OverWritePolicy
const (
	OverWritePolicyUNKNOWN OverWritePolicy = "Unknown"
	OverWritePolicyWRAPS_WHEN_FULL OverWritePolicy = "WrapsWhenFull"
	OverWritePolicyNEVER_OVER_WRITES OverWritePolicy = "NeverOverWrites"
)