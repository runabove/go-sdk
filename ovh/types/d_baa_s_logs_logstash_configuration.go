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

// DBaasLogsLogstashConfiguration Logstash configuration
type DBaasLogsLogstashConfiguration struct {

	// FilterSection The filter section of logstash.conf
	FilterSection string `json:"filterSection,omitempty"`

	// InputSection The filter section of logstash.conf
	InputSection string `json:"inputSection,omitempty"`

	// PatternSection The list of customs Grok patterns
	PatternSection string `json:"patternSection,omitempty"`
}
