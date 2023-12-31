// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestResult - Custom-defined Bolt result object.
type RequestResult struct {
	// Indicates that the request failed. This value is always false.
	Success *bool `json:"success,omitempty"`
}

func (o *RequestResult) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}
