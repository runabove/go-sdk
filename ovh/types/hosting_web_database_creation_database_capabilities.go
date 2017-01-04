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

// Struct which describs quota and available for a specific type of database
type HostingWebDatabaseCreationDatabaseCapabilities struct {

	// Number of database left for creation
	Available int64 `json:"available,omitempty"`

	Quota DedicatedCloudFilerProfileSize `json:"quota,omitempty"`

	Type_ string `json:"type,omitempty"`
}