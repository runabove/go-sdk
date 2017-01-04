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

import (
	"time"
)

// List of documents added on your account
type NichandleDocumentDocument struct {

	// Document creation
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Document expiration
	ExpirationDate time.Time `json:"expirationDate,omitempty"`

	// URL used to get document
	GetUrl string `json:"getUrl,omitempty"`

	// Document id
	Id string `json:"id,omitempty"`

	// Document name
	Name string `json:"name,omitempty"`

	// URL used to put document
	PutUrl string `json:"putUrl,omitempty"`

	// Document size (in bytes)
	Size int64 `json:"size,omitempty"`

	// Document tags
	Tags []ComplexTypeSafeKeyValueString `json:"tags,omitempty"`

	// Document validation
	ValidationDate time.Time `json:"validationDate,omitempty"`
}