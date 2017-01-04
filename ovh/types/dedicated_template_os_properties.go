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

// A structure describing properties customizables about this dedicated installation template
type DedicatedTemplateOsProperties struct {

	// Template change log details
	ChangeLog string `json:"changeLog,omitempty"`

	// Set up the server using the provided hostname instead of the default hostname
	CustomHostname string `json:"customHostname,omitempty"`

	// Indicate the URL where your postinstall customisation script is located
	PostInstallationScriptLink string `json:"postInstallationScriptLink,omitempty"`

	// indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is \"loh1Xee7eo OK OK OK UGh8Ang1Gu
	PostInstallationScriptReturn string `json:"postInstallationScriptReturn,omitempty"`

	Rating int64 `json:"rating,omitempty"`

	// Name of the ssh key that should be installed. Password login will be disabled
	SshKeyName string `json:"sshKeyName,omitempty"`

	// Use the distribution's native kernel instead of the recommended OVH Kernel
	UseDistributionKernel bool `json:"useDistributionKernel,omitempty"`
}