// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Credit struct {
	Status *string `json:"status,omitempty"`
}

func (o *Credit) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
