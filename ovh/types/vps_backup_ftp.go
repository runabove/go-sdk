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

// VpsBackupFtp Backup Ftp assigned to this VPS
type VpsBackupFtp struct {

	// FtpBackupName The backup FTP server name
	FtpBackupName string `json:"ftpBackupName,omitempty"`

	Quota *VpsBackupFtpQuota `json:"quota,omitempty"`

	// ReadOnlyDate If not-null, gives the date since when your account was set in read-only mode because the quota was exceeded
	ReadOnlyDate *time.Time `json:"readOnlyDate,omitempty"`

	// TType The backup FTP type
	TType string `json:"type,omitempty"`

	Usage *VpsBackupFtpUsage `json:"usage,omitempty"`
}
