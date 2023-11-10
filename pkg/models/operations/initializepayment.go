// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"net/http"
)

type InitializePaymentSecurity struct {
	OAuth   *string `security:"scheme,type=oauth2,name=Authorization"`
	XAPIKey *string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *InitializePaymentSecurity) GetOAuth() *string {
	if o == nil {
		return nil
	}
	return o.OAuth
}

func (o *InitializePaymentSecurity) GetXAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.XAPIKey
}

// InitializePaymentShopperIdentity - Identification information for the Shopper. This is only required when creating a new Bolt account.
type InitializePaymentShopperIdentity struct {
	// determines whether to create a bolt account for this shopper
	CreateBoltAccount *bool `json:"create_bolt_account,omitempty"`
	// Email address of the shopper
	Email string `json:"email"`
	// First name of the shopper
	FirstName string `json:"first_name"`
	// Last name of the shopper
	LastName string `json:"last_name"`
	// Phone number of the shopper
	Phone string `json:"phone"`
}

func (o *InitializePaymentShopperIdentity) GetCreateBoltAccount() *bool {
	if o == nil {
		return nil
	}
	return o.CreateBoltAccount
}

func (o *InitializePaymentShopperIdentity) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *InitializePaymentShopperIdentity) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *InitializePaymentShopperIdentity) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *InitializePaymentShopperIdentity) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

type InitializePaymentRequestBody struct {
	// The details of the cart being purchased with this payment.
	Cart shared.CartCreate `json:"cart"`
	// Identification information for the Shopper. This is only required when creating a new Bolt account.
	ShopperIdentity *InitializePaymentShopperIdentity `json:"shopper_identity,omitempty"`
}

func (o *InitializePaymentRequestBody) GetCart() shared.CartCreate {
	if o == nil {
		return shared.CartCreate{}
	}
	return o.Cart
}

func (o *InitializePaymentRequestBody) GetShopperIdentity() *InitializePaymentShopperIdentity {
	if o == nil {
		return nil
	}
	return o.ShopperIdentity
}

type InitializePaymentRequest struct {
	// A key created by merchants that ensures `POST` and `PATCH` requests are only performed once. [Read more about Idempotent Requests here](/developers/references/idempotency/).
	IdempotencyKey *string                       `header:"style=simple,explode=false,name=Idempotency-Key"`
	RequestBody    *InitializePaymentRequestBody `request:"mediaType=application/json"`
	// The publicly viewable identifier used to identify a merchant division. This key is found in the Developer > API section of the Bolt Merchant Dashboard [RECOMMENDED].
	XPublishableKey *string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *InitializePaymentRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *InitializePaymentRequest) GetRequestBody() *InitializePaymentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *InitializePaymentRequest) GetXPublishableKey() *string {
	if o == nil {
		return nil
	}
	return o.XPublishableKey
}

// InitializePaymentStatus - The current payment status.
type InitializePaymentStatus string

const (
	InitializePaymentStatusAwaitingUserConfirmation InitializePaymentStatus = "awaiting_user_confirmation"
	InitializePaymentStatusPaymentReady             InitializePaymentStatus = "payment_ready"
	InitializePaymentStatusSuccess                  InitializePaymentStatus = "success"
)

func (e InitializePaymentStatus) ToPointer() *InitializePaymentStatus {
	return &e
}

func (e *InitializePaymentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "awaiting_user_confirmation":
		fallthrough
	case "payment_ready":
		fallthrough
	case "success":
		*e = InitializePaymentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InitializePaymentStatus: %v", v)
	}
}

// InitializePaymentResponseBody - Payment token retrieved.
type InitializePaymentResponseBody struct {
	// The ID for a Payment Attempt
	ID *string `json:"id,omitempty"`
	// The current payment status.
	Status *InitializePaymentStatus `json:"status,omitempty"`
}

func (o *InitializePaymentResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InitializePaymentResponseBody) GetStatus() *InitializePaymentStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type InitializePaymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Payment token retrieved.
	Object *InitializePaymentResponseBody
}

func (o *InitializePaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InitializePaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InitializePaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *InitializePaymentResponse) GetObject() *InitializePaymentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
