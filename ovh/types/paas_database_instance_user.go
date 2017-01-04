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

// Users
type PaasDatabaseInstanceUser struct {

	// Creation date of the user
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Databases granted for this user
	Databases []PaasDatabaseInstanceUserDatabase `json:"databases,omitempty"`

	// The last update date of the user
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// User status
	Status string `json:"status,omitempty"`

	// The id of the task working on this object
	TaskId string `json:"taskId,omitempty"`

	// User name used to connect to your databases
	UserName string `json:"userName,omitempty"`
}