// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"net/http"
)

type AddPaymentMethodSecurity struct {
	OAuth   string `security:"scheme,type=oauth2,name=Authorization"`
	XAPIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *AddPaymentMethodSecurity) GetOAuth() string {
	if o == nil {
		return ""
	}
	return o.OAuth
}

func (o *AddPaymentMethodSecurity) GetXAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XAPIKey
}

type AddPaymentMethodRequestBodyNetwork string

const (
	AddPaymentMethodRequestBodyNetworkUnknown        AddPaymentMethodRequestBodyNetwork = "unknown"
	AddPaymentMethodRequestBodyNetworkVisa           AddPaymentMethodRequestBodyNetwork = "visa"
	AddPaymentMethodRequestBodyNetworkMastercard     AddPaymentMethodRequestBodyNetwork = "mastercard"
	AddPaymentMethodRequestBodyNetworkAmex           AddPaymentMethodRequestBodyNetwork = "amex"
	AddPaymentMethodRequestBodyNetworkDiscover       AddPaymentMethodRequestBodyNetwork = "discover"
	AddPaymentMethodRequestBodyNetworkDinersClubUsCa AddPaymentMethodRequestBodyNetwork = "diners_club_us_ca"
	AddPaymentMethodRequestBodyNetworkJcb            AddPaymentMethodRequestBodyNetwork = "jcb"
	AddPaymentMethodRequestBodyNetworkUnionpay       AddPaymentMethodRequestBodyNetwork = "unionpay"
	AddPaymentMethodRequestBodyNetworkAlliancedata   AddPaymentMethodRequestBodyNetwork = "alliancedata"
	AddPaymentMethodRequestBodyNetworkCitiplcc       AddPaymentMethodRequestBodyNetwork = "citiplcc"
)

func (e AddPaymentMethodRequestBodyNetwork) ToPointer() *AddPaymentMethodRequestBodyNetwork {
	return &e
}

func (e *AddPaymentMethodRequestBodyNetwork) UnmarshalJSON(data []byte) error {
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
		*e = AddPaymentMethodRequestBodyNetwork(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddPaymentMethodRequestBodyNetwork: %v", v)
	}
}

// AddPaymentMethodRequestBodyPriority - Used to indicate the card's priority. '1' indicates primary, while '2' indicates a secondary card.
type AddPaymentMethodRequestBodyPriority int64

const (
	AddPaymentMethodRequestBodyPriorityOne AddPaymentMethodRequestBodyPriority = 1
	AddPaymentMethodRequestBodyPriorityTwo AddPaymentMethodRequestBodyPriority = 2
)

func (e AddPaymentMethodRequestBodyPriority) ToPointer() *AddPaymentMethodRequestBodyPriority {
	return &e
}

func (e *AddPaymentMethodRequestBodyPriority) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = AddPaymentMethodRequestBodyPriority(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddPaymentMethodRequestBodyPriority: %v", v)
	}
}

// AddPaymentMethodRequestBodyTokenType - Used to define which payment processor generated the token for this credit card.  For those using Bolt's tokenizer, the value must be `bolt`.
type AddPaymentMethodRequestBodyTokenType string

const (
	AddPaymentMethodRequestBodyTokenTypeVantiv   AddPaymentMethodRequestBodyTokenType = "vantiv"
	AddPaymentMethodRequestBodyTokenTypeApplepay AddPaymentMethodRequestBodyTokenType = "applepay"
	AddPaymentMethodRequestBodyTokenTypeBolt     AddPaymentMethodRequestBodyTokenType = "bolt"
	AddPaymentMethodRequestBodyTokenTypeStripe   AddPaymentMethodRequestBodyTokenType = "stripe"
	AddPaymentMethodRequestBodyTokenTypePlcc     AddPaymentMethodRequestBodyTokenType = "plcc"
)

func (e AddPaymentMethodRequestBodyTokenType) ToPointer() *AddPaymentMethodRequestBodyTokenType {
	return &e
}

func (e *AddPaymentMethodRequestBodyTokenType) UnmarshalJSON(data []byte) error {
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
		*e = AddPaymentMethodRequestBodyTokenType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddPaymentMethodRequestBodyTokenType: %v", v)
	}
}

// AddPaymentMethodRequestBody - The `credit_card` object is used to to pay for guest checkout transactions or save payment method details to an account. Once saved, you can reference the credit card with the associated `credit_card_id` for future transactions.
type AddPaymentMethodRequestBody struct {
	// The Address object is used for billing, shipping, and physical store address use cases.
	BillingAddress shared.Address `json:"billing_address"`
	// The unique Bolt ID associated with a saved shopper address. This can be obtained by accessing a shopper's account details. If you use this field, you do not need to use `billing_address`.
	//
	BillingAddressID *string `json:"billing_address_id,omitempty"`
	// The Bank Identification Number for the credit card. This is typically the first 4-6 digits of the credit card number.
	Bin        *string `json:"bin,omitempty"`
	Cryptogram *string `json:"cryptogram,omitempty"`
	// This can be left empty. A 3-digit ISO code for currency that will be used in the credit card authorization.
	Currency *string `json:"currency,omitempty"`
	Eci      *string `json:"eci,omitempty"`
	// The expiration date of the credit card.
	Expiration string `json:"expiration"`
	// The last 4 digits of the credit card number.
	Last4 *string `json:"last4,omitempty"`
	// A key-value pair object that allows users to store arbitrary information associated with an object.  For any individual account object, we allow up to 50 keys. Keys can be up to 40 characters long and values can be up to 500 characters long.  Metadata should not contain any sensitive customer information, like PII (Personally Identifiable Information). For more information about metadata, see our [documentation](https://help.bolt.com/developers/references/embedded-metadata/).
	//
	Metadata *shared.Metadata                    `json:"metadata,omitempty"`
	Network  *AddPaymentMethodRequestBodyNetwork `json:"network,omitempty"`
	// Used to provide ApplePay DPAN or private label credit card PAN when applicable. Required when charging a private label credit card.
	Number *string `json:"number,omitempty"`
	// Used for the postal or zip code associated with the credit card.
	PostalCode *string `json:"postal_code,omitempty"`
	// Used to indicate the card's priority. '1' indicates primary, while '2' indicates a secondary card.
	//
	Priority *AddPaymentMethodRequestBodyPriority `json:"priority,omitempty"`
	// Determines whether or not the credit card will be saved to the shopper's account. Defaults to `true`.
	//
	Save *bool `json:"save,omitempty"`
	// The Bolt token associated to the credit card.
	Token string `json:"token"`
	// Used to define which payment processor generated the token for this credit card.  For those using Bolt's tokenizer, the value must be `bolt`.
	//
	TokenType *AddPaymentMethodRequestBodyTokenType `json:"token_type,omitempty"`
}

func (o *AddPaymentMethodRequestBody) GetBillingAddress() shared.Address {
	if o == nil {
		return shared.Address{}
	}
	return o.BillingAddress
}

func (o *AddPaymentMethodRequestBody) GetBillingAddressID() *string {
	if o == nil {
		return nil
	}
	return o.BillingAddressID
}

func (o *AddPaymentMethodRequestBody) GetBin() *string {
	if o == nil {
		return nil
	}
	return o.Bin
}

func (o *AddPaymentMethodRequestBody) GetCryptogram() *string {
	if o == nil {
		return nil
	}
	return o.Cryptogram
}

func (o *AddPaymentMethodRequestBody) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AddPaymentMethodRequestBody) GetEci() *string {
	if o == nil {
		return nil
	}
	return o.Eci
}

func (o *AddPaymentMethodRequestBody) GetExpiration() string {
	if o == nil {
		return ""
	}
	return o.Expiration
}

func (o *AddPaymentMethodRequestBody) GetLast4() *string {
	if o == nil {
		return nil
	}
	return o.Last4
}

func (o *AddPaymentMethodRequestBody) GetMetadata() *shared.Metadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *AddPaymentMethodRequestBody) GetNetwork() *AddPaymentMethodRequestBodyNetwork {
	if o == nil {
		return nil
	}
	return o.Network
}

func (o *AddPaymentMethodRequestBody) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *AddPaymentMethodRequestBody) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *AddPaymentMethodRequestBody) GetPriority() *AddPaymentMethodRequestBodyPriority {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *AddPaymentMethodRequestBody) GetSave() *bool {
	if o == nil {
		return nil
	}
	return o.Save
}

func (o *AddPaymentMethodRequestBody) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *AddPaymentMethodRequestBody) GetTokenType() *AddPaymentMethodRequestBodyTokenType {
	if o == nil {
		return nil
	}
	return o.TokenType
}

type AddPaymentMethodRequest struct {
	// A key created by merchants that ensures `POST` and `PATCH` requests are only performed once. [Read more about Idempotent Requests here](/developers/references/idempotency/).
	IdempotencyKey *string                      `header:"style=simple,explode=false,name=Idempotency-Key"`
	RequestBody    *AddPaymentMethodRequestBody `request:"mediaType=application/json"`
	// The publicly viewable identifier used to identify a merchant division. This key is found in the Developer > API section of the Bolt Merchant Dashboard [RECOMMENDED].
	XPublishableKey *string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *AddPaymentMethodRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *AddPaymentMethodRequest) GetRequestBody() *AddPaymentMethodRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *AddPaymentMethodRequest) GetXPublishableKey() *string {
	if o == nil {
		return nil
	}
	return o.XPublishableKey
}

type AddPaymentMethodResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Payment Method Added
	SavedCreditCardView *shared.SavedCreditCardView
}

func (o *AddPaymentMethodResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddPaymentMethodResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddPaymentMethodResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddPaymentMethodResponse) GetSavedCreditCardView() *shared.SavedCreditCardView {
	if o == nil {
		return nil
	}
	return o.SavedCreditCardView
}
