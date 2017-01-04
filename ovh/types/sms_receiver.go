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

// Sms receivers preloaded
type SmsReceiver struct {

	// Download file from URL before sending to contacts
	AutoUpdate bool `json:"autoUpdate,omitempty"`

	// Creation date of the document
	Datetime time.Time `json:"datetime,omitempty"`

	// Description name of the document
	Description string `json:"description,omitempty"`

	// Number of receiver records in the document
	Records int64 `json:"records,omitempty"`

	// Slot number id
	SlotId int64 `json:"slotId,omitempty"`
}