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

// SmsPhonebooksPhonebookContactPost ...
type SmsPhonebooksPhonebookContactPost struct {
	Group string `json:"group,omitempty"`

	HomeMobile string `json:"homeMobile,omitempty"`

	HomePhone string `json:"homePhone,omitempty"`

	Name string `json:"name,omitempty"`

	Surname string `json:"surname,omitempty"`

	WorkMobile string `json:"workMobile,omitempty"`

	WorkPhone string `json:"workPhone,omitempty"`
}
