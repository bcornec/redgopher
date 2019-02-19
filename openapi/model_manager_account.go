/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The user accounts, owned by a Manager, are defined in this resource.  Changes to a Manager Account may affect the current Redfish service connection if this manager is responsible for the Redfish service.
type ManagerAccount struct {
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id,omitempty"`
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// This property is used by a User Administrator to disable an account w/o having to delet the user information.  When set to true, the user can login.  When set to false, the account is administratively disabled and the user cannot login.
	Enabled bool `json:"Enabled,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	Links ManagerAccount2Links `json:"Links,omitempty"`
	// This property indicates that the account has been auto-locked by the account service because the lockout threshold has been exceeded.  When set to true, the account is locked. A user admin can write the property to false to manually unlock, or the account service will unlock it once the lockout duration period has passed.
	Locked bool `json:"Locked,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem string `json:"Oem,omitempty"`
	// This property is used with a PATCH or PUT to write the password for the account.  This property is null on a GET.
	Password string `json:"Password,omitempty"`
	// This property contains the Role for this account.
	RoleId string `json:"RoleId,omitempty"`
	// This property contains the user name for the account.
	UserName string `json:"UserName,omitempty"`
}
