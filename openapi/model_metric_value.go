/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

// A metric value.
type MetricValue struct {
	MetricDefinition IdRef `json:"MetricDefinition,omitempty"`
	// The metric definitions identifier for this metric.
	MetricId string `json:"MetricId,omitempty"`
	// The URI for the property from which this metric is derived.
	MetricProperty string `json:"MetricProperty,omitempty"`
	// The value identifies this resource.
	MetricValue string `json:"MetricValue,omitempty"`
	// The time when the value of the metric is obtained. A management application may establish a time series of metric data by retrieving the instances of metric value and sorting them according to their Timestamp.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}