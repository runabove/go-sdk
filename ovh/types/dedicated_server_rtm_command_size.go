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

// DedicatedServerRtmCommandSize A structure describing informations about RTM
type DedicatedServerRtmCommandSize struct {

	// Command Complete command line used for starting this process
	Command string `json:"command,omitempty"`

	Memory *DedicatedServerRtmCommandSizeMemory `json:"memory,omitempty"`
}
