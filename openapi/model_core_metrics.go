/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The processor core metrics.
type CoreMetrics struct {
	// The C-state residency of this core in the processor.
	CStateResidency []CStateResidency `json:"CStateResidency,omitempty"`
	// The cache metrics of this core in the processor.
	CoreCache []CacheMetrics `json:"CoreCache,omitempty"`
	// The processor core identifier.
	CoreId string `json:"CoreId,omitempty"`
	// The number of stalled cycles due to I/O operations.
	IOStallCount float32 `json:"IOStallCount,omitempty"`
	// The number of instructions per clock cycle of this core.
	InstructionsPerCycle float32 `json:"InstructionsPerCycle,omitempty"`
	// The number of stalled cycles due to memory operations.
	MemoryStallCount float32 `json:"MemoryStallCount,omitempty"`
	// The unhalted cycles count of this core.
	UnhaltedCycles float32 `json:"UnhaltedCycles,omitempty"`
}