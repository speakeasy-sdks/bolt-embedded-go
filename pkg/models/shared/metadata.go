// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Metadata - A key-value pair object that allows users to store arbitrary information associated with an object.  For any individual account object, we allow up to 50 keys. Keys can be up to 40 characters long and values can be up to 500 characters long.  Metadata should not contain any sensitive customer information, like PII (Personally Identifiable Information). For more information about metadata, see our [documentation](https://help.bolt.com/developers/references/embedded-metadata/).
type Metadata struct {
	AdditionalProperties *string `json:"additionalProperties,omitempty"`
}

func (o *Metadata) GetAdditionalProperties() *string {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
