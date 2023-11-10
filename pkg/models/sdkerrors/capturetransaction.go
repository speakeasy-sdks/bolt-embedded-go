// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"net/http"
)

type Errors struct {
	Code    *float64 `json:"code,omitempty"`
	Field   *string  `json:"field,omitempty"`
	Message *string  `json:"message,omitempty"`
}

func (o *Errors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Errors) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

func (o *Errors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

type Result struct {
}

// CaptureTransactionResponseBody - Unprocessable Entity
type CaptureTransactionResponseBody struct {
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response `json:"-"`
	Errors      []Errors       `json:"errors,omitempty"`
	Result      *Result        `json:"result,omitempty"`
}

var _ error = &CaptureTransactionResponseBody{}

func (e *CaptureTransactionResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}