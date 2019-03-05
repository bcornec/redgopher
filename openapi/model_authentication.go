/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Contains the authentication information for the external service.
type Authentication struct {
	AuthenticationType AuthenticationTypes `json:"AuthenticationType,omitempty"`
	// This property is used with a PATCH or PUT to write a base64 encoded version of the kerberos keytab for the account.  This property is null on a GET.
	KerberosKeytab string `json:"KerberosKeytab,omitempty"`
	Oem Oem `json:"Oem,omitempty"`
	// This property is used with a PATCH or PUT to write the password for the account service.  This property is null on a GET.
	Password string `json:"Password,omitempty"`
	// This property is used with a PATCH or PUT to write the token for the account.  This property is null on a GET.
	Token string `json:"Token,omitempty"`
	// This property contains the user name for the account service.
	Username string `json:"Username,omitempty"`
}
