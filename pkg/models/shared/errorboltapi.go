// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ErrorBoltAPI - Error object containing custom error information
type ErrorBoltAPI struct {
	// Custom-defined Bolt error code. This can be used to programmatically react to specific errors.
	Code *int64 `json:"code,omitempty"`
	// Human-readable description of the error for developers. Should not be shown to users and is not localized.
	Message *string `json:"message,omitempty"`
}

func (o *ErrorBoltAPI) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ErrorBoltAPI) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
