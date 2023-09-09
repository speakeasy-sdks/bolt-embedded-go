// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GiftOptions - Contains the gift option settings for wrapping and custom messages.
type GiftOptions struct {
	// Includes the gift message written by the shopper.
	Message *string `json:"message,omitempty"`
	// Defines whether gift wrapping was requested.
	Wrap *bool `json:"wrap,omitempty"`
}

func (o *GiftOptions) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GiftOptions) GetWrap() *bool {
	if o == nil {
		return nil
	}
	return o.Wrap
}