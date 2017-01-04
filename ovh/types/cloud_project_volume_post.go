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

type CloudProjectVolumePost struct {

	Bootable bool `json:"bootable,omitempty"`

	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`

	Region string `json:"region,omitempty"`

	Size int64 `json:"size,omitempty"`

	SnapshotId string `json:"snapshotId,omitempty"`

	Type_ string `json:"type,omitempty"`
}