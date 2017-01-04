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

type OrderLicenseWorklightNewPost struct {

	Ip string `json:"ip,omitempty"`

	LessThan1000Users bool `json:"lessThan1000Users,omitempty"`

	Version string `json:"version,omitempty"`
}