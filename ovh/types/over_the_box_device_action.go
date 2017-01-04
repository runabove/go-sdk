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

// Device action
type OverTheBoxDeviceAction struct {

	// The id of the action
	ActionId string `json:"actionId,omitempty"`

	// The name of the action
	Name string `json:"name,omitempty"`

	// The status of the action
	Status string `json:"status,omitempty"`
}