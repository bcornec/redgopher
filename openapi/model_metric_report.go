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

// The metric definitions used to create a metric report.
type MetricReport struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	MetricReportDefinition IdRef `json:"MetricReportDefinition,omitempty"`
	// An array of metric values for the metered items of this Metric.
	MetricValues []MetricValue `json:"MetricValues,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// The current sequence identifier for this metric report.
	ReportSequence string `json:"ReportSequence"`
	// The time associated with the metric report in its entirety.  The time of the metric report may be relevant when the time of individual metrics are minimally different.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}
