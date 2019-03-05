/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The metadata information about a metric.
type MetricDefinition struct {
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// Estimated percent error of measured vs. actual values.
	Accuracy float32 `json:"Accuracy,omitempty"`
	Actions Actions2 `json:"Actions,omitempty"`
	Calculable Calculable `json:"Calculable,omitempty"`
	CalculationAlgorithm CalculationAlgorithmEnum `json:"CalculationAlgorithm,omitempty"`
	// Specifies the metric properties which are part of the synthesis calculation.  This property is present when the MetricType property has the value 'Synthesized'.
	CalculationParameters []CalculationParamsType `json:"CalculationParameters,omitempty"`
	// The time interval over which the metric calculation is performed.
	CalculationTimeInterval string `json:"CalculationTimeInterval,omitempty"`
	// Specifies the calibration offset added to the metric reading.
	Calibration float32 `json:"Calibration,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// This array property specifies possible values of a discrete metric.
	DiscreteValues []string `json:"DiscreteValues,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	Implementation ImplementationType `json:"Implementation,omitempty"`
	// Indicates whether the metric values are linear (vs non-linear).
	IsLinear bool `json:"IsLinear,omitempty"`
	// Maximum value for metric reading.
	MaxReadingRange float32 `json:"MaxReadingRange,omitempty"`
	MetricDataType MetricDataType `json:"MetricDataType,omitempty"`
	// A collection of URI for the properties on which this metric definition is defined.
	MetricProperties []string `json:"MetricProperties,omitempty"`
	MetricType MetricType `json:"MetricType,omitempty"`
	// Minimum value for metric reading.
	MinReadingRange float32 `json:"MinReadingRange,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	PhysicalContext IdRef `json:"PhysicalContext,omitempty"`
	// Number of significant digits in the metric reading.
	Precision int32 `json:"Precision,omitempty"`
	// The time interval between when a metric is updated.
	SensingInterval string `json:"SensingInterval,omitempty"`
	// Accuracy of the timestamp.
	TimestampAccuracy string `json:"TimestampAccuracy,omitempty"`
	// The units of measure for this metric.
	Units string `json:"Units,omitempty"`
	// Wildcards used to replace values in AppliesTo and Calculates metric property arrays.
	Wildcards []Wildcard `json:"Wildcards,omitempty"`
}
