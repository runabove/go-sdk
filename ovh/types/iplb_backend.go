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

// Backend
type IplbBackend struct {

	// Load balancing algorithm. 'roundrobin' if null
	Balance string `json:"balance,omitempty"`

	// Id of your backend
	Id int64 `json:"id,omitempty"`

	// Backend name
	Name string `json:"name,omitempty"`

	// Port attached to your backend. Inherited from frontend if null
	Port int64 `json:"port,omitempty"`

	// Probe type. 'tcp' if null
	Probe string `json:"probe,omitempty"`

	// Stickiness type. No stickyness if null
	Stickiness string `json:"stickiness,omitempty"`

	// Type of your backend
	Type_ string `json:"type,omitempty"`

	// Zone of you backend.
	Zone string `json:"zone,omitempty"`
}