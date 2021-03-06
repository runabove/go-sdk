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

// DockerStackInputCustomSSL A custom SSL certificate associated to a Docker PaaS environment
type DockerStackInputCustomSSL struct {

	// Certificate The custom SSL public certificate
	Certificate string `json:"certificate,omitempty"`

	// Key The custom SSL certificate private key
	Key string `json:"key,omitempty"`
}
