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

// Show the deconsolidation terms
type XdslDeconsolidationTerms struct {

	// Duration of month the access will be engaged
	Engagement float64 `json:"engagement,omitempty"`

	MonthlyPrice OrderPrice `json:"monthlyPrice,omitempty"`

	Price OrderPrice `json:"price,omitempty"`
}