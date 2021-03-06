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

// DedicatedInstallationTemplateTemplatePartitions  Partitions defined in this partitioning scheme
type DedicatedInstallationTemplateTemplatePartitions struct {

	// Filesystem Partition filesytem
	Filesystem string `json:"filesystem,omitempty"`

	// Mountpoint partition mount point
	Mountpoint string `json:"mountpoint,omitempty"`

	// Order specifies the creation order of the partition on the disk
	Order int64 `json:"order,omitempty"`

	// Raid raid partition type
	Raid string `json:"raid,omitempty"`

	Size *DedicatedInstallationTemplateTemplatePartitionsSize `json:"size,omitempty"`

	TType string `json:"type,omitempty"`

	// VolumeName The volume name needed for proxmox distribution
	VolumeName string `json:"volumeName,omitempty"`
}
