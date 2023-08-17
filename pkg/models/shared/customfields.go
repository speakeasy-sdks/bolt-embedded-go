// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CustomFieldsCheckoutSetup string

const (
	CustomFieldsCheckoutSetupShippingStep           CustomFieldsCheckoutSetup = "shipping_step"
	CustomFieldsCheckoutSetupDeliveryStep           CustomFieldsCheckoutSetup = "delivery_step"
	CustomFieldsCheckoutSetupPaymentStep            CustomFieldsCheckoutSetup = "payment_step"
	CustomFieldsCheckoutSetupAccountRegistrationSso CustomFieldsCheckoutSetup = "account_registration_sso"
)

func (e CustomFieldsCheckoutSetup) ToPointer() *CustomFieldsCheckoutSetup {
	return &e
}

func (e *CustomFieldsCheckoutSetup) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "shipping_step":
		fallthrough
	case "delivery_step":
		fallthrough
	case "payment_step":
		fallthrough
	case "account_registration_sso":
		*e = CustomFieldsCheckoutSetup(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomFieldsCheckoutSetup: %v", v)
	}
}

type CustomFields struct {
	CheckoutSetup *CustomFieldsCheckoutSetup `json:"checkout_setup,omitempty"`
	// Defines whether the field is dynamic.
	Dynamic *bool `json:"dynamic,omitempty"`
	// The external ID for the custom field.
	ExternalID *string `json:"external_id,omitempty"`
	FieldSetup *string `json:"field_setup,omitempty"`
	// The displayed label for the custom field, seen by the shopper.
	Label    *string `json:"label,omitempty"`
	Position *int64  `json:"position,omitempty"`
	// The internal ID for the custom field.
	PublicID *string `json:"public_id,omitempty"`
	// Defines if the field must be completed to check out.
	Required *bool `json:"required,omitempty"`
	// Defines whether the shopper is opted into a newsletter or not.
	SubscribeToNewsletter *bool `json:"subscribe_to_newsletter,omitempty"`
}

func (o *CustomFields) GetCheckoutSetup() *CustomFieldsCheckoutSetup {
	if o == nil {
		return nil
	}
	return o.CheckoutSetup
}

func (o *CustomFields) GetDynamic() *bool {
	if o == nil {
		return nil
	}
	return o.Dynamic
}

func (o *CustomFields) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *CustomFields) GetFieldSetup() *string {
	if o == nil {
		return nil
	}
	return o.FieldSetup
}

func (o *CustomFields) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *CustomFields) GetPosition() *int64 {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *CustomFields) GetPublicID() *string {
	if o == nil {
		return nil
	}
	return o.PublicID
}

func (o *CustomFields) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *CustomFields) GetSubscribeToNewsletter() *bool {
	if o == nil {
		return nil
	}
	return o.SubscribeToNewsletter
}
