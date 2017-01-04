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

// ContainerDetail
type CloudStorageContainerDetail struct {

	// Whether this is an archive container or not
	Archive bool `json:"archive,omitempty"`

	// Origins allowed to make Cross Origin Requests
	Cors []string `json:"cors,omitempty"`

	// Container name
	Name string `json:"name,omitempty"`

	// Objects stored in container
	Objects []CloudStorageContainerObject `json:"objects,omitempty"`

	// Public container
	Public bool `json:"public,omitempty"`

	// Container region
	Region string `json:"region,omitempty"`

	// Container static URL
	StaticUrl string `json:"staticUrl,omitempty"`

	// Total bytes stored
	StoredBytes int64 `json:"storedBytes,omitempty"`

	// Total objects stored
	StoredObjects int64 `json:"storedObjects,omitempty"`
}