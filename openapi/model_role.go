/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This resource defines a user role to be used in conjunction with a Manager Account.
type Role struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// The redfish privileges that this role includes.
	AssignedPrivileges []PrivilegeType `json:"AssignedPrivileges,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// This property is used to indicate if the Role is one of the Redfish Predefined Roles vs a Custom role.
	IsPredefined bool `json:"IsPredefined,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem string `json:"Oem,omitempty"`
	// The OEM privileges that this role includes.
	OemPrivileges []string `json:"OemPrivileges,omitempty"`
	// This property contains the name of the Role.
	RoleId string `json:"RoleId,omitempty"`
}
