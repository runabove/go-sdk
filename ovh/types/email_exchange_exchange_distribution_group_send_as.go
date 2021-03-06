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

// EmailExchangeDistributionGroupSendAs Get users authorized to Send On Behalf To mails from this mailbox
type EmailExchangeDistributionGroupSendAs struct {

	// AllowedAccountID Account id to give send on behalf to
	AllowedAccountID int64 `json:"allowedAccountId,omitempty"`

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// TaskPendingID Pending task id
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}
