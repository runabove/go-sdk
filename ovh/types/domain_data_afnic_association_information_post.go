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

type DomainDataAfnicAssociationInformationPost struct {

	ContactId int64 `json:"contactId,omitempty"`

	DeclarationDate time.Time `json:"declarationDate,omitempty"`

	PublicationDate time.Time `json:"publicationDate,omitempty"`

	PublicationNumber string `json:"publicationNumber,omitempty"`

	PublicationPageNumber string `json:"publicationPageNumber,omitempty"`
}