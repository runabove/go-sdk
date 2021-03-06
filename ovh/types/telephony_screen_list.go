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

// TelephonyScreenList Screen list
type TelephonyScreenList struct {
	CallNumber string `json:"callNumber,omitempty"`

	ID int64 `json:"id,omitempty"`

	Nature string `json:"nature,omitempty"`

	Status string `json:"status,omitempty"`

	TType string `json:"type,omitempty"`
}
