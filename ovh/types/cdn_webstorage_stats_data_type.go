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

// A structure describing type of a stats hash
type CdnWebstorageStatsDataType struct {

	Date time.Time `json:"date,omitempty"`

	Value int64 `json:"value,omitempty"`
}