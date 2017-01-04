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

// Server attached to an SSL Gateway
type SslGatewayServer struct {

	// IP address of the server attached to your SSL Gateway
	Address string `json:"address,omitempty"`

	// Id of your server
	Id int64 `json:"id,omitempty"`

	// Port of your server attached to your SSL Gateway
	Port int64 `json:"port,omitempty"`

	// Server state
	State string `json:"state,omitempty"`
}