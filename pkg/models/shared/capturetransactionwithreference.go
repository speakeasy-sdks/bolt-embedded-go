// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CaptureTransactionWithReference - Capture a Transaction
type CaptureTransactionWithReference struct {
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

func (o *CaptureTransactionWithReference) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

func (o *CaptureTransactionWithReference) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *CaptureTransactionWithReference) GetMerchantEventID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantEventID
}

func (o *CaptureTransactionWithReference) GetSkipHookNotification() *bool {
	if o == nil {
		return nil
	}
	return o.SkipHookNotification
}

func (o *CaptureTransactionWithReference) GetTransactionReference() string {
	if o == nil {
		return ""
	}
	return o.TransactionReference
}
