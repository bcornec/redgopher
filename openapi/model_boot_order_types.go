/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BootOrderTypes : The enumerations of BootOrderTypes specify the choice of boot order property to use when controller the persistent boot order for this computer system.
type BootOrderTypes string

// List of BootOrderTypes
const (
	BootOrderTypesBOOT_ORDER BootOrderTypes = "BootOrder"
	BootOrderTypesALIAS_BOOT_ORDER BootOrderTypes = "AliasBootOrder"
)