/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Specifies a trigger, which apply to metrics.
type Triggers struct {
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	DiscreteTriggerCondition DiscreteTriggerConditionEnum `json:"DiscreteTriggerCondition,omitempty"`
	// List of discrete triggers.
	DiscreteTriggers []DiscreteTrigger `json:"DiscreteTriggers,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// A collection of URI for the properties on which this metric definition is defined.
	MetricProperties []string `json:"MetricProperties,omitempty"`
	MetricType MetricTypeEnum `json:"MetricType,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	NumericThresholds Thresholds2 `json:"NumericThresholds,omitempty"`
	Oem string `json:"Oem,omitempty"`
	Status Status `json:"Status,omitempty"`
	// This property specifies the actions to perform when the trigger occurs.
	TriggerActions []TriggerActionEnum `json:"TriggerActions,omitempty"`
	// Wildcards used to replace values in MetricProperties array property.
	Wildcards []Wildcard `json:"Wildcards,omitempty"`
}
