/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The AccountService schema contains properties for managing user accounts. The properties are common to all user accounts, such as password requirements, and control features such as account lockout. The schema also contains links to the collections of Manager Accounts and Roles.
type AccountService struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// The interval of time in seconds between the last failed login attempt and reset of the lockout threshold counter. This value must be less than or equal to AccountLockoutDuration. Reset sets the counter to zero.
	AccountLockoutCounterResetAfter int32 `json:"AccountLockoutCounterResetAfter,omitempty"`
	// The time in seconds an account is locked out. The value must be greater than or equal to the value of the AccountLockoutCounterResetAfter property. If set to 0, no lockout occurs.
	AccountLockoutDuration int32 `json:"AccountLockoutDuration,omitempty"`
	// The number of failed login attempts allowed before a user account is locked for a specified duration. A value of 0 means it is never locked.
	AccountLockoutThreshold int32 `json:"AccountLockoutThreshold,omitempty"`
	Accounts IdRef `json:"Accounts,omitempty"`
	Actions Actions `json:"Actions,omitempty"`
	ActiveDirectory ExternalAccountProvider `json:"ActiveDirectory,omitempty"`
	AdditionalExternalAccountProviders IdRef `json:"AdditionalExternalAccountProviders,omitempty"`
	// The number of authorization failures allowed before the failure attempt is logged to the manager log.
	AuthFailureLoggingThreshold int32 `json:"AuthFailureLoggingThreshold,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	LDAP ExternalAccountProvider `json:"LDAP,omitempty"`
	LocalAccountAuth LocalAccountAuth `json:"LocalAccountAuth,omitempty"`
	// The maximum password length for this service.
	MaxPasswordLength int32 `json:"MaxPasswordLength,omitempty"`
	// The minimum password length for this service.
	MinPasswordLength int32 `json:"MinPasswordLength,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	PrivilegeMap IdRef `json:"PrivilegeMap,omitempty"`
	Roles IdRef `json:"Roles,omitempty"`
	// Indicates whether this service is enabled.  If set to false, the AccountService is disabled.  This means no users can be created, deleted or modified.  Any service attempting to access the AccountService resource (for example, the Session Service) will fail.  New sessions cannot be started when the service is disabled. However, established sessions may still continue operating. This does not affect Basic AUTH connections.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`
	Status Status `json:"Status,omitempty"`
}
