/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// N_Port ID Virtualization (NPIV) capabilties for a controller.
type Npiv struct {
	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed simultaneously from all ports on this controller.
	MaxDeviceLogins int32 `json:"MaxDeviceLogins,omitempty"`
	// The maximum number of N_Port ID Virtualization (NPIV) logins allowed per physical port on this controller.
	MaxPortLogins int32 `json:"MaxPortLogins,omitempty"`
}
