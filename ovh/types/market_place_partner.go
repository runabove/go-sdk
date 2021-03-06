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

// MarketPlacePartner partner
type MarketPlacePartner struct {

	// Category Category
	Category string `json:"category,omitempty"`

	// City City
	City string `json:"city,omitempty"`

	// CompanyNationalIDentificationNumber Company national identification number
	CompanyNationalIDentificationNumber string `json:"companyNationalIdentificationNumber,omitempty"`

	// Contact str
	Contact string `json:"contact,omitempty"`

	// Country Country
	Country string `json:"country,omitempty"`

	// Description Complete description
	Description string `json:"description,omitempty"`

	// Language Language
	Language string `json:"language,omitempty"`

	// LegalForm Legal form
	LegalForm string `json:"legalForm,omitempty"`

	// OrganisationDisplayName Organisation display name
	OrganisationDisplayName string `json:"organisationDisplayName,omitempty"`

	// OrganisationName Organisation name
	OrganisationName string `json:"organisationName,omitempty"`

	// OtherDetails Complementary information
	OtherDetails string `json:"otherDetails,omitempty"`

	// Province Province name
	Province string `json:"province,omitempty"`

	// Street Street address
	Street string `json:"street,omitempty"`

	// URL Website address
	URL string `json:"url,omitempty"`

	// Vat VAT number
	Vat string `json:"vat,omitempty"`

	// Zip ZipCode
	Zip string `json:"zip,omitempty"`
}
