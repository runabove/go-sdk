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

// Available softwares on a Template
type VpsSoftware struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`

	Type_ string `json:"type,omitempty"`
}