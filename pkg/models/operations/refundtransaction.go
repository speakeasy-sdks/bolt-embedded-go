// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"net/http"
)

type RefundTransactionSecurity struct {
	XAPIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *RefundTransactionSecurity) GetXAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XAPIKey
}

// RefundTransactionRequestBody - Refund a Transaction
type RefundTransactionRequestBody struct {
	// The amount in cents. **Nullable** for Transactions Details.
	Amount int64 `json:"amount"`
	// The 3-digit ISO code for the currency. **Nullable** for Transactions Details.
	Currency string `json:"currency"`
	// The reference ID associated with a transaction event (auth, capture, refund, void). This is an arbitrary identifier created by the merchant. Bolt does not enforce any uniqueness constraints on this ID. It is up to the merchant to generate identifiers that properly fulfill its needs.
	MerchantEventID *string `json:"merchant_event_id,omitempty"`
	// Set to `true` to skip receiving a webhook notification from Bolt that is triggered by this update to the transaction.
	SkipHookNotification *bool `json:"skip_hook_notification,omitempty"`
	// The transaction's 12-digit Bolt reference ID. **Nullable** for Transactions Details.
	TransactionReference string `json:"transaction_reference"`
}

func (o *RefundTransactionRequestBody) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

func (o *RefundTransactionRequestBody) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *RefundTransactionRequestBody) GetMerchantEventID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantEventID
}

func (o *RefundTransactionRequestBody) GetSkipHookNotification() *bool {
	if o == nil {
		return nil
	}
	return o.SkipHookNotification
}

func (o *RefundTransactionRequestBody) GetTransactionReference() string {
	if o == nil {
		return ""
	}
	return o.TransactionReference
}

type RefundTransactionRequest struct {
	// A key created by merchants that ensures `POST` and `PATCH` requests are only performed once. [Read more about Idempotent Requests here](/developers/references/idempotency/).
	IdempotencyKey *string `header:"style=simple,explode=false,name=Idempotency-Key"`
	// Refund a Transaction
	RequestBody *RefundTransactionRequestBody `request:"mediaType=application/json"`
}

func (o *RefundTransactionRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *RefundTransactionRequest) GetRequestBody() *RefundTransactionRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type RefundTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Refund Successful
	TransactionView *shared.TransactionView
}

func (o *RefundTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RefundTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RefundTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RefundTransactionResponse) GetTransactionView() *shared.TransactionView {
	if o == nil {
		return nil
	}
	return o.TransactionView
}
