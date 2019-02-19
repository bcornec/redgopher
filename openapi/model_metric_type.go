/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MetricType : Specifies the type of metric provided.  The property provides information to the client on how the metric can be handled.
type MetricType string

// List of MetricType
const (
	MetricTypeNUMERIC MetricType = "Numeric"
	MetricTypeDISCRETE MetricType = "Discrete"
	MetricTypeGAUGE MetricType = "Gauge"
	MetricTypeCOUNTER MetricType = "Counter"
	MetricTypeCOUNTDOWN MetricType = "Countdown"
)