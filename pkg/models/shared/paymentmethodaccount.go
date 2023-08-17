// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PaymentMethodAccountNetwork string

const (
	PaymentMethodAccountNetworkUnknown        PaymentMethodAccountNetwork = "unknown"
	PaymentMethodAccountNetworkVisa           PaymentMethodAccountNetwork = "visa"
	PaymentMethodAccountNetworkMastercard     PaymentMethodAccountNetwork = "mastercard"
	PaymentMethodAccountNetworkAmex           PaymentMethodAccountNetwork = "amex"
	PaymentMethodAccountNetworkDiscover       PaymentMethodAccountNetwork = "discover"
	PaymentMethodAccountNetworkDinersClubUsCa PaymentMethodAccountNetwork = "diners_club_us_ca"
	PaymentMethodAccountNetworkJcb            PaymentMethodAccountNetwork = "jcb"
	PaymentMethodAccountNetworkUnionpay       PaymentMethodAccountNetwork = "unionpay"
	PaymentMethodAccountNetworkAlliancedata   PaymentMethodAccountNetwork = "alliancedata"
	PaymentMethodAccountNetworkCitiplcc       PaymentMethodAccountNetwork = "citiplcc"
)

func (e PaymentMethodAccountNetwork) ToPointer() *PaymentMethodAccountNetwork {
	return &e
}

func (e *PaymentMethodAccountNetwork) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		fallthrough
	case "visa":
		fallthrough
	case "mastercard":
		fallthrough
	case "amex":
		fallthrough
	case "discover":
		fallthrough
	case "diners_club_us_ca":
		fallthrough
	case "jcb":
		fallthrough
	case "unionpay":
		fallthrough
	case "alliancedata":
		fallthrough
	case "citiplcc":
		*e = PaymentMethodAccountNetwork(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodAccountNetwork: %v", v)
	}
}

// PaymentMethodAccountPriority - Used to indicate the card's priority. '1' indicates primary, while '2' indicates a secondary card.
type PaymentMethodAccountPriority int64

const (
	PaymentMethodAccountPriorityOne PaymentMethodAccountPriority = 1
	PaymentMethodAccountPriorityTwo PaymentMethodAccountPriority = 2
)

func (e PaymentMethodAccountPriority) ToPointer() *PaymentMethodAccountPriority {
	return &e
}

func (e *PaymentMethodAccountPriority) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = PaymentMethodAccountPriority(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodAccountPriority: %v", v)
	}
}

// PaymentMethodAccountTokenType - Used to define which payment processor generated the token for this credit card.  For those using Bolt's tokenizer, the value must be `bolt`.
type PaymentMethodAccountTokenType string

const (
	PaymentMethodAccountTokenTypeVantiv   PaymentMethodAccountTokenType = "vantiv"
	PaymentMethodAccountTokenTypeApplepay PaymentMethodAccountTokenType = "applepay"
	PaymentMethodAccountTokenTypeBolt     PaymentMethodAccountTokenType = "bolt"
	PaymentMethodAccountTokenTypeStripe   PaymentMethodAccountTokenType = "stripe"
	PaymentMethodAccountTokenTypePlcc     PaymentMethodAccountTokenType = "plcc"
)

func (e PaymentMethodAccountTokenType) ToPointer() *PaymentMethodAccountTokenType {
	return &e
}

func (e *PaymentMethodAccountTokenType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vantiv":
		fallthrough
	case "applepay":
		fallthrough
	case "bolt":
		fallthrough
	case "stripe":
		fallthrough
	case "plcc":
		*e = PaymentMethodAccountTokenType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodAccountTokenType: %v", v)
	}
}

// PaymentMethodAccount - The `credit_card` object is used to to pay for guest checkout transactions or save payment method details to an account. Once saved, you can reference the credit card with the associated `credit_card_id` for future transactions.
type PaymentMethodAccount struct {
	// The Address object is used for billing, shipping, and physical store address use cases.
	BillingAddress Address `json:"billing_address"`
	// The unique Bolt ID associated with a saved shopper address. This can be obtained by accessing a shopper's account details. If you use this field, you do not need to use `billing_address`.
	//
	BillingAddressID *string `json:"billing_address_id,omitempty"`
	// The Bank Identification Number for the credit card. This is typically the first 4-6 digits of the credit card number.
	Bin        *string `json:"bin,omitempty"`
	Cryptogram *string `json:"cryptogram,omitempty"`
	// Set this to true to make this the default payment method. There can be only one payment method with default set to true.
	Default *bool   `json:"default,omitempty"`
	Eci     *string `json:"eci,omitempty"`
	// The expiration date of the credit card.
	Expiration string `json:"expiration"`
	// The last 4 digits of the credit card number.
	Last4 *string `json:"last4,omitempty"`
	// A key-value pair object that allows users to store arbitrary information associated with an object.  For any individual account object, we allow up to 50 keys. Keys can be up to 40 characters long and values can be up to 500 characters long.  Metadata should not contain any sensitive customer information, like PII (Personally Identifiable Information). For more information about metadata, see our [documentation](https://help.bolt.com/developers/references/embedded-metadata/).
	//
	Metadata *Metadata                    `json:"metadata,omitempty"`
	Network  *PaymentMethodAccountNetwork `json:"network,omitempty"`
	// Used to provide ApplePay DPAN or private label credit card PAN when applicable. Required when charging a private label credit card.
	Number *string `json:"number,omitempty"`
	// Used for the postal or zip code associated with the credit card.
	PostalCode *string `json:"postal_code,omitempty"`
	// Used to indicate the card's priority. '1' indicates primary, while '2' indicates a secondary card.
	//
	Priority *PaymentMethodAccountPriority `json:"priority,omitempty"`
	// Determines whether or not the credit card will be saved to the shopper's account. Defaults to `true`.
	//
	Save *bool `json:"save,omitempty"`
	// The Bolt token associated to the credit card.
	Token string `json:"token"`
	// Used to define which payment processor generated the token for this credit card.  For those using Bolt's tokenizer, the value must be `bolt`.
	//
	TokenType *PaymentMethodAccountTokenType `json:"token_type,omitempty"`
}

func (o *PaymentMethodAccount) GetBillingAddress() Address {
	if o == nil {
		return Address{}
	}
	return o.BillingAddress
}

func (o *PaymentMethodAccount) GetBillingAddressID() *string {
	if o == nil {
		return nil
	}
	return o.BillingAddressID
}

func (o *PaymentMethodAccount) GetBin() *string {
	if o == nil {
		return nil
	}
	return o.Bin
}

func (o *PaymentMethodAccount) GetCryptogram() *string {
	if o == nil {
		return nil
	}
	return o.Cryptogram
}

func (o *PaymentMethodAccount) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *PaymentMethodAccount) GetEci() *string {
	if o == nil {
		return nil
	}
	return o.Eci
}

func (o *PaymentMethodAccount) GetExpiration() string {
	if o == nil {
		return ""
	}
	return o.Expiration
}

func (o *PaymentMethodAccount) GetLast4() *string {
	if o == nil {
		return nil
	}
	return o.Last4
}

func (o *PaymentMethodAccount) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *PaymentMethodAccount) GetNetwork() *PaymentMethodAccountNetwork {
	if o == nil {
		return nil
	}
	return o.Network
}

func (o *PaymentMethodAccount) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *PaymentMethodAccount) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PaymentMethodAccount) GetPriority() *PaymentMethodAccountPriority {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *PaymentMethodAccount) GetSave() *bool {
	if o == nil {
		return nil
	}
	return o.Save
}

func (o *PaymentMethodAccount) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *PaymentMethodAccount) GetTokenType() *PaymentMethodAccountTokenType {
	if o == nil {
		return nil
	}
	return o.TokenType
}
