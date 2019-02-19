/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type BaseModuleType string

// List of BaseModuleType
const (
	BaseModuleTypeRDIMM BaseModuleType = "RDIMM"
	BaseModuleTypeUDIMM BaseModuleType = "UDIMM"
	BaseModuleTypeSO_DIMM BaseModuleType = "SO_DIMM"
	BaseModuleTypeLRDIMM BaseModuleType = "LRDIMM"
	BaseModuleTypeMINI_RDIMM BaseModuleType = "Mini_RDIMM"
	BaseModuleTypeMINI_UDIMM BaseModuleType = "Mini_UDIMM"
	BaseModuleTypeSO_RDIMM_72B BaseModuleType = "SO_RDIMM_72b"
	BaseModuleTypeSO_UDIMM_72B BaseModuleType = "SO_UDIMM_72b"
	BaseModuleTypeSO_DIMM_16B BaseModuleType = "SO_DIMM_16b"
	BaseModuleTypeSO_DIMM_32B BaseModuleType = "SO_DIMM_32b"
	BaseModuleTypeDIE BaseModuleType = "Die"
)