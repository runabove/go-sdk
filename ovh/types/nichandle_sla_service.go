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

// NichandleSLAService Describe all services impacted by SLA
type NichandleSLAService struct {

	// Description Service description
	Description string `json:"description,omitempty"`

	// ServiceName Impacted service name
	ServiceName string `json:"serviceName,omitempty"`

	// SLAApplication SLA plan application
	SLAApplication string `json:"slaApplication,omitempty"`

	// SLAPlan SLA plan description
	SLAPlan string `json:"slaPlan,omitempty"`
}
