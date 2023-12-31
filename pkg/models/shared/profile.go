// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProfileMetadata - A key-value pair object that allows users to store arbitrary information associated with an object.  For any individual account object, we allow up to 50 keys. Keys can be up to 40 characters long and values can be up to 500 characters long.  Metadata should not contain any sensitive customer information, like PII (Personally Identifiable Information). For more information about metadata, see our [documentation](https://help.bolt.com/developers/references/embedded-metadata/).
type ProfileMetadata struct {
	AdditionalProperties *string `json:"additionalProperties,omitempty"`
}

func (o *ProfileMetadata) GetAdditionalProperties() *string {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// Profile - The first name, last name, email address, and phone number of a shopper.
type Profile struct {
	// An email address.
	Email string `json:"email"`
	// The given name of the person associated with this record.
	FirstName string `json:"first_name"`
	// The surname of the person associated with this record.
	LastName string `json:"last_name"`
	// A key-value pair object that allows users to store arbitrary information associated with an object.  For any individual account object, we allow up to 50 keys. Keys can be up to 40 characters long and values can be up to 500 characters long.  Metadata should not contain any sensitive customer information, like PII (Personally Identifiable Information). For more information about metadata, see our [documentation](https://help.bolt.com/developers/references/embedded-metadata/).
	//
	Metadata *ProfileMetadata `json:"metadata,omitempty"`
	// A phone number following E164 standards, in its globalized format, i.e. prepended with a plus sign.
	Phone string `json:"phone"`
}

func (o *Profile) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *Profile) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *Profile) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *Profile) GetMetadata() *ProfileMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Profile) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}
