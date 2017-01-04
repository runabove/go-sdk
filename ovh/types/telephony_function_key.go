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

// Plug & Phone function key
type TelephonyFunctionKey struct {

	// The default function used by the key
	Default_ string `json:"default,omitempty"`

	// The function active on the key
	Function string `json:"function,omitempty"`

	// The number of the function key
	KeyNum int64 `json:"keyNum,omitempty"`

	// The key label
	Label string `json:"label,omitempty"`

	// The function parameter
	Parameter string `json:"parameter,omitempty"`

	// The key type
	Type_ string `json:"type,omitempty"`
}