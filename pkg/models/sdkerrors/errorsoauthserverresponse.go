// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type ErrorsOauthServerResponse struct {
	Error_           *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}

var _ error = &ErrorsOauthServerResponse{}

func (e *ErrorsOauthServerResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
