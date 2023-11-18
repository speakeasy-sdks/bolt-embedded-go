// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

// ErrorsBoltAPIResponse - Forbidden
type ErrorsBoltAPIResponse struct {
	Errors []shared.ErrorBoltAPI `json:"errors,omitempty"`
	// Custom-defined Bolt result object.
	Result *shared.RequestResult `json:"result,omitempty"`
}

var _ error = &ErrorsBoltAPIResponse{}

func (e *ErrorsBoltAPIResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
