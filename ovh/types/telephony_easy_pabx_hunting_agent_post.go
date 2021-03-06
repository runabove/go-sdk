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

// TelephonyEasyPabxHuntingAgentPost ...
type TelephonyEasyPabxHuntingAgentPost struct {
	AgentNumber string `json:"agentNumber,omitempty"`

	Logged bool `json:"logged,omitempty"`

	NoReplyTimer int64 `json:"noReplyTimer,omitempty"`

	Position int64 `json:"position,omitempty"`
}
