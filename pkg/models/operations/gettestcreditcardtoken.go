// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// GetTestCreditCardTokenResponseBody - Successfully Fetched Credit Card Token
type GetTestCreditCardTokenResponseBody struct {
	// The credit card bin.
	Bin *string `json:"bin,omitempty"`
	// The date at which the token expires. A token must be used within 15 minutes of creation.
	Expiry *int64 `json:"expiry,omitempty"`
	// The last 4 digits of the card number.
	Last4 *string `json:"last4,omitempty"`
	// The credit card network.
	Network *string `json:"network,omitempty"`
	// The newly generated credit card token.
	Token *string `json:"token,omitempty"`
}

func (o *GetTestCreditCardTokenResponseBody) GetBin() *string {
	if o == nil {
		return nil
	}
	return o.Bin
}

func (o *GetTestCreditCardTokenResponseBody) GetExpiry() *int64 {
	if o == nil {
		return nil
	}
	return o.Expiry
}

func (o *GetTestCreditCardTokenResponseBody) GetLast4() *string {
	if o == nil {
		return nil
	}
	return o.Last4
}

func (o *GetTestCreditCardTokenResponseBody) GetNetwork() *string {
	if o == nil {
		return nil
	}
	return o.Network
}

func (o *GetTestCreditCardTokenResponseBody) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

type GetTestCreditCardTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully Fetched Credit Card Token
	Object *GetTestCreditCardTokenResponseBody
}

func (o *GetTestCreditCardTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTestCreditCardTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTestCreditCardTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTestCreditCardTokenResponse) GetObject() *GetTestCreditCardTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
