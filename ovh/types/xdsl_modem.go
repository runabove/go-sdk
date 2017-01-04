/* 
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

import (
	"time"
)

// Modem
type XdslModem struct {

	BrandName string `json:"brandName,omitempty"`

	Capabilities XdslModemCapabilities `json:"capabilities,omitempty"`

	// IP Address of the DMZ (To modify or delete the DMZ IP on the modem, a re-configuration of your modem has to be made, your configuration will be maintained, but you will lose your connection a few minutes)
	DmzIP string `json:"dmzIP,omitempty"`

	// Level of the FireWall on the modem
	EasyFirewallLevel string `json:"easyFirewallLevel,omitempty"`

	// Whether or not the modem supports IPv6
	Ipv6Support bool `json:"ipv6Support,omitempty"`

	// Whether or not the modem is in bridge mode. To pass from bridge mode to routed mode, a reset is necessary. If the modem is managedByOvh, the bridge state will be kept after a reset
	IsBridged bool `json:"isBridged,omitempty"`

	// Last time the modem made a CWMP request to the Auto Configuration Server
	LastCwmpRequestDate time.Time `json:"lastCwmpRequestDate,omitempty"`

	MacAddress string `json:"macAddress,omitempty"`

	// Whether or not the user can configure his modem via OVH Interface (will lock telnet and local HTTP configuration page)
	ManagedByOvh bool `json:"managedByOvh,omitempty"`

	Model string `json:"model,omitempty"`

	// Size of the Maximum Transmission Unit on the modem's interfaces
	MtuSize int64 `json:"mtuSize,omitempty"`
}