/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This resource is used to obtain or modify the network services managed by a given manager.
type ManagerNetworkProtocol struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	DHCP Protocol `json:"DHCP,omitempty"`
	DHCPv6 Protocol `json:"DHCPv6,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// This is the fully qualified domain name for the manager obtained by DNS including the host name and top-level domain name.
	FQDN string `json:"FQDN,omitempty"`
	HTTP Protocol `json:"HTTP,omitempty"`
	HTTPS HttpsProtocol `json:"HTTPS,omitempty"`
	// The DNS Host Name of this manager, without any domain information.
	HostName string `json:"HostName,omitempty"`
	IPMI Protocol `json:"IPMI,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	KVMIP Protocol `json:"KVMIP,omitempty"`
	NTP NtpProtocol `json:"NTP,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem string `json:"Oem,omitempty"`
	RDP Protocol `json:"RDP,omitempty"`
	RFB Protocol `json:"RFB,omitempty"`
	SNMP Protocol `json:"SNMP,omitempty"`
	SSDP SsdProtocol `json:"SSDP,omitempty"`
	SSH Protocol `json:"SSH,omitempty"`
	Status Status `json:"Status,omitempty"`
	Telnet Protocol `json:"Telnet,omitempty"`
	VirtualMedia Protocol `json:"VirtualMedia,omitempty"`
}
