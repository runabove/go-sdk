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

// CloudInstanceNetworkParams NetworkParams
type CloudInstanceNetworkParams struct {

	// IP Static ip (Can only be defined for private networks)
	IP string `json:"ip,omitempty"`

	// NetworkID Private or public network Id
	NetworkID string `json:"networkId,omitempty"`
}
