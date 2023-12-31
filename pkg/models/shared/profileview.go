// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProfileView - The shopper's account profile.
type ProfileView struct {
	// An email address.
	Email *string `json:"email,omitempty"`
	// The given name of the person associated with this record.
	FirstName *string `json:"first_name,omitempty"`
	// The surname of the person associated with this record.
	LastName *string `json:"last_name,omitempty"`
	// A key-value pair object that allows users to store arbitrary information associated with an object.  For any individual account object, we allow up to 50 keys. Keys can be up to 40 characters long and values can be up to 500 characters long.  Metadata should not contain any sensitive customer information, like PII (Personally Identifiable Information). For more information about metadata, see our [documentation](https://help.bolt.com/developers/references/embedded-metadata/).
	//
	Metadata *Metadata `json:"metadata,omitempty"`
	// The given and surname of the person associated with this address.
	Name *string `json:"name,omitempty"`
	// A phone number following E164 standards, in its globalized format, i.e. prepended with a plus sign.
	Phone *string `json:"phone,omitempty"`
}

func (o *ProfileView) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ProfileView) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *ProfileView) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *ProfileView) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ProfileView) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ProfileView) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}
