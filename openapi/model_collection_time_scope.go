/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CollectionTimeScope : The time scope of the related metric values.
type CollectionTimeScope string

// List of CollectionTimeScope
const (
	CollectionTimeScopePOINT CollectionTimeScope = "Point"
	CollectionTimeScopeINTERVAL CollectionTimeScope = "Interval"
	CollectionTimeScopeSTARTUP_INTERVAL CollectionTimeScope = "StartupInterval"
)