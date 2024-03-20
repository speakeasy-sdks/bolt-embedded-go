// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"net/http"
)

type VoidTransactionRequest struct {
	// A key created by merchants that ensures `POST` and `PATCH` requests are only performed once. [Read more about Idempotent Requests here](/developers/references/idempotency/).
	IdempotencyKey *string `header:"style=simple,explode=false,name=Idempotency-Key"`
	// Void a Transaction
	CreditCardVoid *shared.CreditCardVoid `request:"mediaType=application/json"`
}

func (o *VoidTransactionRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *VoidTransactionRequest) GetCreditCardVoid() *shared.CreditCardVoid {
	if o == nil {
		return nil
	}
	return o.CreditCardVoid
}

type VoidTransactionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Void Successful
	TransactionView *shared.TransactionView
}

func (o *VoidTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *VoidTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *VoidTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *VoidTransactionResponse) GetTransactionView() *shared.TransactionView {
	if o == nil {
		return nil
	}
	return o.TransactionView
}
