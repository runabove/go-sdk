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

// TelephonyRegistrationInformations Informations about a SIP registration (i.e. IP, port, User-Agent...)
type TelephonyRegistrationInformations struct {

	// Datetime SIP registration's date
	Datetime *time.Time `json:"datetime,omitempty"`

	// Domain SIP registration's domain
	Domain string `json:"domain,omitempty"`

	// IP SIP registration's IP
	IP string `json:"ip,omitempty"`

	// Port SIP registration's port
	Port int64 `json:"port,omitempty"`

	// UserAgent SIP registration's User-Agent
	UserAgent string `json:"userAgent,omitempty"`
}
