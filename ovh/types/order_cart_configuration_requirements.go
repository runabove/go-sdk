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

// OrderCartConfigurationRequirements Information about a configuration
type OrderCartConfigurationRequirements struct {

	// Fields Indicates if some particular fields have to be inputed during the creation of `type` resource
	Fields []string `json:"fields,omitempty"`

	// Label Label for your configuration item
	Label string `json:"label,omitempty"`

	// Required Indicates if the configuration item is required
	Required bool `json:"required,omitempty"`

	// TType Type of the configuration item
	TType string `json:"type,omitempty"`
}
