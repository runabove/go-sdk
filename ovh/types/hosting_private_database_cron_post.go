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

type HostingPrivateDatabaseCronPost struct {

	Command string `json:"command,omitempty"`

	DatabaseName string `json:"databaseName,omitempty"`

	Description string `json:"description,omitempty"`

	Email string `json:"email,omitempty"`

	Frequency string `json:"frequency,omitempty"`

	Status string `json:"status,omitempty"`
}