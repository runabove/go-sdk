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

// OVH easy calls queues
type TelephonyEasyHunting struct {

	// Reject (hangup) anonymous calls
	AnonymousRejection bool `json:"anonymousRejection,omitempty"`

	Description string `json:"description,omitempty"`

	FeatureType string `json:"featureType,omitempty"`

	// Max wait time when caller is in queue (in seconds)
	MaxWaitTime int64 `json:"maxWaitTime,omitempty"`

	// Max number of callers in queue
	QueueSize int64 `json:"queueSize,omitempty"`

	ServiceName string `json:"serviceName,omitempty"`

	ServiceType string `json:"serviceType,omitempty"`

	// The presented number when bridging calls
	ShowCallerNumber string `json:"showCallerNumber,omitempty"`

	// The calls dispatching strategy
	Strategy string `json:"strategy,omitempty"`

	// Tone played just before call is hang up
	ToneOnClosing int64 `json:"toneOnClosing,omitempty"`

	// Tone played when caller is put on hold
	ToneOnHold int64 `json:"toneOnHold,omitempty"`

	// Tone played when call is picked up
	ToneOnOpening int64 `json:"toneOnOpening,omitempty"`

	// The voicemail used by the EasyPABX
	Voicemail string `json:"voicemail,omitempty"`
}