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

// HourlyVolume
type CloudBillingViewHourlyVolume struct {

	// Detail about volume consumption
	Details []CloudBillingViewHourlyVolumeDetail `json:"details,omitempty"`

	Quantity CloudBillingViewQuantity `json:"quantity,omitempty"`

	// Region
	Region string `json:"region,omitempty"`

	// Total price
	TotalPrice float64 `json:"totalPrice,omitempty"`

	// Volume type
	Type_ string `json:"type,omitempty"`
}