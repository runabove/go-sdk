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

// Async task
type XdslAsyncTaskPackXdslMigrationMigrationOfferResponse struct {

	// Error
	Error_ string `json:"error,omitempty"`

	Result PackXdslMigrationMigrationOfferResponse `json:"result,omitempty"`

	// Status of the call
	Status string `json:"status,omitempty"`
}