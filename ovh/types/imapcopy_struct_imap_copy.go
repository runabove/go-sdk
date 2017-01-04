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

// Structure of imapCopy
type ImapcopyStructImapCopy struct {

	// If true, IMAP Serveur of mailbox
	SSL bool `json:"SSL,omitempty"`

	// Account of mailbox
	Account string `json:"account,omitempty"`

	// Password of mailbox
	Password string `json:"password,omitempty"`

	// IMAP Serveur of mailbox
	ServerIMAP string `json:"serverIMAP,omitempty"`
}