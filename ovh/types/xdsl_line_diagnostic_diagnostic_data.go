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

// Diagnostic data and informations
type XdslLineDiagnosticDiagnosticData struct {

	// List of actions already done by customer
	ActionsDone []string `json:"actionsDone,omitempty"`

	// List of actions that must be done by customer
	ActionsToDo []XdslLineDiagnosticCustomerActionToDo `json:"actionsToDo,omitempty"`

	Answers XdslLineDiagnosticAnswers `json:"answers,omitempty"`

	// Diagnostic comment. Can be update during any diagnostic step
	Comment string `json:"comment,omitempty"`

	// Diagnostic creation date
	CreationDate time.Time `json:"creationDate,omitempty"`

	// End of diagnostic date. Will be null until problem totally identified
	DiagnosticDoneDate time.Time `json:"diagnosticDoneDate,omitempty"`

	// Error message
	Error_ string `json:"error,omitempty"`

	// Last diagnostic update date
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	LineDetails XdslLineDiagnosticLineDetails `json:"lineDetails,omitempty"`

	// Current or last robot action
	RobotAction string `json:"robotAction,omitempty"`

	SeltTest XdslLineDiagnosticSeltResult `json:"seltTest,omitempty"`

	// Diagnostic timeout in minutes. Any action restart timeout
	Timeout int64 `json:"timeout,omitempty"`

	// List of questions that must be answered by customer
	ToAnswer []XdslLineDiagnosticQuestion `json:"toAnswer,omitempty"`
}