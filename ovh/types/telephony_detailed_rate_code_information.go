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

// Detailed informations related to this number
type TelephonyDetailedRateCodeInformation struct {

	// Cancelable datetime deadline for the new scheduled rateCode
	CancelLimitDatetime time.Time `json:"cancelLimitDatetime,omitempty"`

	// Effective datetime
	EffectiveDatetime time.Time `json:"effectiveDatetime,omitempty"`

	PricePerCallWithoutTax OrderPrice `json:"pricePerCallWithoutTax,omitempty"`

	PricePerMinuteWithoutTax OrderPrice `json:"pricePerMinuteWithoutTax,omitempty"`

	// Scheduled rate code
	RateCode string `json:"rateCode,omitempty"`

	RepaymentPricePerCallWithoutTax OrderPrice `json:"repaymentPricePerCallWithoutTax,omitempty"`

	RepaymentPricePerMinuteWithoutTax OrderPrice `json:"repaymentPricePerMinuteWithoutTax,omitempty"`

	UpdateRateCodePriceWithoutTax OrderPrice `json:"updateRateCodePriceWithoutTax,omitempty"`
}