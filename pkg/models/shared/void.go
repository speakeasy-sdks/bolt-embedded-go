// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// VoidCause - Specifies why this particular transaction is voided.
type VoidCause string

const (
	VoidCauseMerchantAction          VoidCause = "merchant_action"
	VoidCausePaypalSync              VoidCause = "paypal_sync"
	VoidCauseAmazonPaySync           VoidCause = "amazon_pay_sync"
	VoidCauseIrreversibleReject      VoidCause = "irreversible_reject"
	VoidCauseAuthExpire              VoidCause = "auth_expire"
	VoidCauseAuthVerificationExpired VoidCause = "auth_verification_expired"
	VoidCausePaymentMethodUpdater    VoidCause = "payment_method_updater"
	VoidCauseLessThanNilGreaterThan  VoidCause = "<nil>"
)

func (e VoidCause) ToPointer() *VoidCause {
	return &e
}

func (e *VoidCause) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "merchant_action":
		fallthrough
	case "paypal_sync":
		fallthrough
	case "amazon_pay_sync":
		fallthrough
	case "irreversible_reject":
		fallthrough
	case "auth_expire":
		fallthrough
	case "auth_verification_expired":
		fallthrough
	case "payment_method_updater":
		fallthrough
	case "<nil>":
		*e = VoidCause(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VoidCause: %v", v)
	}
}

// VoidStatus - The status of the void request.
type VoidStatus string

const (
	VoidStatusSucceeded VoidStatus = "succeeded"
	VoidStatusDeclined  VoidStatus = "declined"
	VoidStatusError     VoidStatus = "error"
)

func (e VoidStatus) ToPointer() *VoidStatus {
	return &e
}

func (e *VoidStatus) UnmarshalJSON(data []byte) error {
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
		*e = VoidStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VoidStatus: %v", v)
	}
}

type Void struct {
	// Specifies why this particular transaction is voided.
	Cause *VoidCause `json:"cause,omitempty"`
	// The reference ID associated with a transaction event (auth, capture, refund, void). This is an arbitrary identifier created by the merchant. Bolt does not enforce any uniqueness constraints on this ID. It is up to the merchant to generate identifiers that properly fulfill its needs.
	MerchantEventID *string `json:"merchant_event_id,omitempty"`
	// The status of the void request.
	Status *VoidStatus `json:"status,omitempty"`
	// The void ID returned from the payment processor.
	Void *string `json:"void,omitempty"`
}

func (o *Void) GetCause() *VoidCause {
	if o == nil {
		return nil
	}
	return o.Cause
}

func (o *Void) GetMerchantEventID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantEventID
}

func (o *Void) GetStatus() *VoidStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Void) GetVoid() *string {
	if o == nil {
		return nil
	}
	return o.Void
}
