/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The ForceFailover action forces a failover of this manager to the manager used in the parameter.
type ForceFailoverRequestBody struct {
	NewManager IdRef `json:"NewManager"`
}
