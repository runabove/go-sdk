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

import (
	"time"
)

// Backup Ftp ACL for this server and Backup Ftp
type DedicatedServerBackupFtpAcl struct {

	// Wether to allow the CIFS (SMB) protocol for this ACL
	Cifs bool `json:"cifs,omitempty"`

	// Wether to allow the FTP protocol for this ACL
	Ftp bool `json:"ftp,omitempty"`

	// The IP Block specific to this ACL
	IpBlock string `json:"ipBlock,omitempty"`

	// Whether the rule has been applied on the Backup Ftp
	IsApplied bool `json:"isApplied,omitempty"`

	// Date of the last object modification
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Wether to allow the NFS protocol for this ACL
	Nfs bool `json:"nfs,omitempty"`
}