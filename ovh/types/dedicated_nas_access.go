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

// DedicatedNasAccess Define Acl for partition
type DedicatedNasAccess struct {

	// AccessID the id of the access
	AccessID int64 `json:"accessId,omitempty"`

	// IP the ip in root on storage
	IP string `json:"ip,omitempty"`
}
