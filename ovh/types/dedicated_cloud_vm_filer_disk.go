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

// DedicatedCloudVMFilerDisk A structure describing filer disks of a the virtual machine
type DedicatedCloudVMFilerDisk struct {
	Capacity float64 `json:"capacity,omitempty"`

	ID string `json:"id,omitempty"`
}
