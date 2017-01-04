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

// SshKey
type CloudSshkeySshKey struct {

	// SSH key id
	Id string `json:"id,omitempty"`

	// SSH key name
	Name string `json:"name,omitempty"`

	// SSH public key
	PublicKey string `json:"publicKey,omitempty"`

	// SSH key regions
	Regions []string `json:"regions,omitempty"`
}