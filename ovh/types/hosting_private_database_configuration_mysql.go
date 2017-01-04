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

// Mysql configuration
type HostingPrivateDatabaseConfigurationMysql struct {

	// If set to True, all changes to a table take effect immediately. If set to 0, you must use COMMIT to accept a transaction or ROLLBACK to cancel it
	AutoCommit bool `json:"autoCommit,omitempty"`

	// The size of the buffer pool, the memory area where InnoDB caches table and index data. ( size in MB )
	InnodbBufferPoolSize int64 `json:"innodbBufferPoolSize,omitempty"`

	// Last update date
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// The maximum size of one packet or any generated/intermediate string. ( size in MB )
	MaxAllowedPacket int64 `json:"maxAllowedPacket,omitempty"`

	// The maximum permitted number of  simultaneous client connections
	MaxConnections int64 `json:"maxConnections,omitempty"`

	// Where was stored temporary files
	Tmpdir string `json:"tmpdir,omitempty"`
}