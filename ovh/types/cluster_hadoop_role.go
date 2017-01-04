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

// Role (ie set of Hadoop services) of the Node
type ClusterHadoopRole struct {

	// ID of the Role
	Id int64 `json:"id,omitempty"`

	// Role name
	Type_ string `json:"type,omitempty"`
}