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

// Available clusterIp by country
type HostingWebCountriesIp struct {

	// The whois country of the ip
	Country string `json:"country,omitempty"`

	// The cluster ip
	Ip string `json:"ip,omitempty"`

	// The cluster ipv6
	Ipv6 string `json:"ipv6,omitempty"`
}