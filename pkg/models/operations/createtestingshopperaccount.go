// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"net/http"
)

type CreateTestingShopperAccountSecurity struct {
	XAPIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *CreateTestingShopperAccountSecurity) GetXAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XAPIKey
}

// CreateTestingShopperAccountRequestBodyEmailState - The status of the shopper account identifier (email or phone). If the account does not have this identifier, the status is "missing"; If the identifier has been used to receive an OTP code, the status is "verified"; If the identifier has not been used to receive an OTP code, the status is "unverified".
type CreateTestingShopperAccountRequestBodyEmailState string

const (
	CreateTestingShopperAccountRequestBodyEmailStateMissing    CreateTestingShopperAccountRequestBodyEmailState = "missing"
	CreateTestingShopperAccountRequestBodyEmailStateVerified   CreateTestingShopperAccountRequestBodyEmailState = "verified"
	CreateTestingShopperAccountRequestBodyEmailStateUnverified CreateTestingShopperAccountRequestBodyEmailState = "unverified"
)

func (e CreateTestingShopperAccountRequestBodyEmailState) ToPointer() *CreateTestingShopperAccountRequestBodyEmailState {
	return &e
}

func (e *CreateTestingShopperAccountRequestBodyEmailState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "missing":
		fallthrough
	case "verified":
		fallthrough
	case "unverified":
		*e = CreateTestingShopperAccountRequestBodyEmailState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTestingShopperAccountRequestBodyEmailState: %v", v)
	}
}

type CreateTestingShopperAccountRequestBody struct {
	// Number of days after which the test account is deactivated. Default: 30 days. Maximum: 180 days.
	DeactivateInDays *int64 `json:"deactivate_in_days,omitempty"`
	// Deprecated. Please leave this field absent and let the API automatically generate a random email.
	Email *string `json:"email,omitempty"`
	// The status of the shopper account identifier (email or phone). If the account does not have this identifier, the status is "missing"; If the identifier has been used to receive an OTP code, the status is "verified"; If the identifier has not been used to receive an OTP code, the status is "unverified".
	EmailState *CreateTestingShopperAccountRequestBodyEmailState `json:"email_state,omitempty"`
	// Add a random U.S. address to the created account if set to `true`
	HasAddress *bool `json:"has_address,omitempty"`
	// Set this account as migrated by the merchant in the request
	Migrated *bool `json:"migrated,omitempty"`
	// Deprecated. Please leave this field absent and let the API automatically generate a random phone number.
	Phone *string `json:"phone,omitempty"`
	// The status of the shopper account identifier (email or phone). If the account does not have this identifier, the status is "missing"; If the identifier has been used to receive an OTP code, the status is "verified"; If the identifier has not been used to receive an OTP code, the status is "unverified".
	PhoneState *shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailState `json:"phone_state,omitempty"`
}

func (o *CreateTestingShopperAccountRequestBody) GetDeactivateInDays() *int64 {
	if o == nil {
		return nil
	}
	return o.DeactivateInDays
}

func (o *CreateTestingShopperAccountRequestBody) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *CreateTestingShopperAccountRequestBody) GetEmailState() *CreateTestingShopperAccountRequestBodyEmailState {
	if o == nil {
		return nil
	}
	return o.EmailState
}

func (o *CreateTestingShopperAccountRequestBody) GetHasAddress() *bool {
	if o == nil {
		return nil
	}
	return o.HasAddress
}

func (o *CreateTestingShopperAccountRequestBody) GetMigrated() *bool {
	if o == nil {
		return nil
	}
	return o.Migrated
}

func (o *CreateTestingShopperAccountRequestBody) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *CreateTestingShopperAccountRequestBody) GetPhoneState() *shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailState {
	if o == nil {
		return nil
	}
	return o.PhoneState
}

type CreateTestingShopperAccountRequest struct {
	RequestBody *CreateTestingShopperAccountRequestBody `request:"mediaType=application/json"`
	// The publicly viewable identifier used to identify a merchant division. This key is found in the Developer > API section of the Bolt Merchant Dashboard [RECOMMENDED].
	XPublishableKey *string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *CreateTestingShopperAccountRequest) GetRequestBody() *CreateTestingShopperAccountRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *CreateTestingShopperAccountRequest) GetXPublishableKey() *string {
	if o == nil {
		return nil
	}
	return o.XPublishableKey
}

// CreateTestingShopperAccount200ApplicationJSON - Testing Account Created
type CreateTestingShopperAccount200ApplicationJSON struct {
	// An email address.
	Email *string `json:"email,omitempty"`
	// The status of the shopper account identifier (email or phone). If the account does not have this identifier, the status is "missing"; If the identifier has been used to receive an OTP code, the status is "verified"; If the identifier has not been used to receive an OTP code, the status is "unverified".
	EmailState *shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailState `json:"email_state,omitempty"`
	// The merchant's public id if the account is migrated
	MigratedMerchantOwnerID *string `json:"migrated_merchant_owner_id,omitempty"`
	// OAuth code that is associated with this account and can be used to exchange for an access token
	OauthCode *string `json:"oauth_code,omitempty"`
	// Fixed OTP code that can be used to login to the created account
	OtpCode *string `json:"otp_code,omitempty"`
	// A phone number following E164 standards, in its globalized format, i.e. prepended with a plus sign.
	Phone *string `json:"phone,omitempty"`
	// The status of the shopper account identifier (email or phone). If the account does not have this identifier, the status is "missing"; If the identifier has been used to receive an OTP code, the status is "verified"; If the identifier has not been used to receive an OTP code, the status is "unverified".
	PhoneState *shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailState `json:"phone_state,omitempty"`
	// The created testing account will be deactivated after this date
	WillDeactivateAt *string `json:"will_deactivate_at,omitempty"`
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetEmailState() *shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailState {
	if o == nil {
		return nil
	}
	return o.EmailState
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetMigratedMerchantOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.MigratedMerchantOwnerID
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetOauthCode() *string {
	if o == nil {
		return nil
	}
	return o.OauthCode
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetOtpCode() *string {
	if o == nil {
		return nil
	}
	return o.OtpCode
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetPhoneState() *shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailState {
	if o == nil {
		return nil
	}
	return o.PhoneState
}

func (o *CreateTestingShopperAccount200ApplicationJSON) GetWillDeactivateAt() *string {
	if o == nil {
		return nil
	}
	return o.WillDeactivateAt
}

type CreateTestingShopperAccountResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Testing Account Created
	CreateTestingShopperAccount200ApplicationJSONObject *CreateTestingShopperAccount200ApplicationJSON
}

func (o *CreateTestingShopperAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTestingShopperAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTestingShopperAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTestingShopperAccountResponse) GetCreateTestingShopperAccount200ApplicationJSONObject() *CreateTestingShopperAccount200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateTestingShopperAccount200ApplicationJSONObject
}