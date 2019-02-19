/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type CompositionState string

// List of CompositionState
const (
	CompositionStateCOMPOSING CompositionState = "Composing"
	CompositionStateCOMPOSED_AND_AVAILABLE CompositionState = "ComposedAndAvailable"
	CompositionStateCOMPOSED CompositionState = "Composed"
	CompositionStateUNUSED CompositionState = "Unused"
	CompositionStateFAILED CompositionState = "Failed"
	CompositionStateUNAVAILABLE CompositionState = "Unavailable"
)