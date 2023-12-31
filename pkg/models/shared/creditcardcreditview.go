// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreditCardCreditViewStatus - The status of the refund to a card.
type CreditCardCreditViewStatus string

const (
	CreditCardCreditViewStatusSucceeded  CreditCardCreditViewStatus = "succeeded"
	CreditCardCreditViewStatusDeclined   CreditCardCreditViewStatus = "declined"
	CreditCardCreditViewStatusError      CreditCardCreditViewStatus = "error"
	CreditCardCreditViewStatusPending    CreditCardCreditViewStatus = "pending"
	CreditCardCreditViewStatusInProgress CreditCardCreditViewStatus = "in progress"
)

func (e CreditCardCreditViewStatus) ToPointer() *CreditCardCreditViewStatus {
	return &e
}

func (e *CreditCardCreditViewStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "succeeded":
		fallthrough
	case "declined":
		fallthrough
	case "error":
		fallthrough
	case "pending":
		fallthrough
	case "in progress":
		*e = CreditCardCreditViewStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditCardCreditViewStatus: %v", v)
	}
}

type CreditCardCreditView struct {
	// The reference ID associated with a transaction event (auth, capture, refund, void). This is an arbitrary identifier created by the merchant. Bolt does not enforce any uniqueness constraints on this ID. It is up to the merchant to generate identifiers that properly fulfill its needs.
	MerchantEventID *string `json:"merchant_event_id,omitempty"`
	// The status of the refund to a card.
	Status *CreditCardCreditViewStatus `json:"status,omitempty"`
}

func (o *CreditCardCreditView) GetMerchantEventID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantEventID
}

func (o *CreditCardCreditView) GetStatus() *CreditCardCreditViewStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
