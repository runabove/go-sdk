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

// OrderEmailExchangeServiceAccountPost ...
type OrderEmailExchangeServiceAccountPost struct {
	Licence string `json:"licence,omitempty"`

	Number int64 `json:"number,omitempty"`

	StorageQuota int64 `json:"storageQuota,omitempty"`
}
