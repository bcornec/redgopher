/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type ConnectedVia string

// List of ConnectedVia
const (
	ConnectedViaNOT_CONNECTED ConnectedVia = "NotConnected"
	ConnectedViaURI ConnectedVia = "URI"
	ConnectedViaAPPLET ConnectedVia = "Applet"
	ConnectedViaOEM ConnectedVia = "Oem"
)