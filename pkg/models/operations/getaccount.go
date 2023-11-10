// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"net/http"
)

type GetAccountSecurity struct {
	OAuth   string `security:"scheme,type=oauth2,name=Authorization"`
	XAPIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *GetAccountSecurity) GetOAuth() string {
	if o == nil {
		return ""
	}
	return o.OAuth
}

func (o *GetAccountSecurity) GetXAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XAPIKey
}

type GetAccountRequest struct {
	// The publicly viewable identifier used to identify a merchant division. This key is found in the Developer > API section of the Bolt Merchant Dashboard [RECOMMENDED].
	XPublishableKey *string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *GetAccountRequest) GetXPublishableKey() *string {
	if o == nil {
		return nil
	}
	return o.XPublishableKey
}

type GetAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Account Details Fetched
	AccountDetails *shared.AccountDetails
}

func (o *GetAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAccountResponse) GetAccountDetails() *shared.AccountDetails {
	if o == nil {
		return nil
	}
	return o.AccountDetails
}
