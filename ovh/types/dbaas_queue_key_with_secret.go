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

// KeyWithSecret
type DbaasQueueKeyWithSecret struct {

	// Human ID of the key's application
	HumanAppId string `json:"humanAppId,omitempty"`

	// Key ID
	Id string `json:"id,omitempty"`

	// Key name
	Name string `json:"name,omitempty"`

	// Key secret
	Secret string `json:"secret,omitempty"`
}