/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CollectionFunction : An operation to perform over the sample.
type CollectionFunction string

// List of CollectionFunction
const (
	CollectionFunctionAVERAGE CollectionFunction = "Average"
	CollectionFunctionMAXIMUM CollectionFunction = "Maximum"
	CollectionFunctionMINIMUM CollectionFunction = "Minimum"
	CollectionFunctionSUMMATION CollectionFunction = "Summation"
)