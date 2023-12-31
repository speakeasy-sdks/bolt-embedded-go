// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CaptureType - Fee type options. **Nullable** for Transactions Details.
type CaptureType string

const (
	CaptureTypeNet            CaptureType = "net"
	CaptureTypeProcessingFee  CaptureType = "processing_fee"
	CaptureTypeRiskFee        CaptureType = "risk_fee"
	CaptureTypeApmFee         CaptureType = "apm_fee"
	CaptureTypeNetworkFee     CaptureType = "network_fee"
	CaptureTypePlatformFee    CaptureType = "platform_fee"
	CaptureTypeBoltAccountFee CaptureType = "bolt_account_fee"
)

func (e CaptureType) ToPointer() *CaptureType {
	return &e
}

func (e *CaptureType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "net":
		fallthrough
	case "processing_fee":
		fallthrough
	case "risk_fee":
		fallthrough
	case "apm_fee":
		fallthrough
	case "network_fee":
		fallthrough
	case "platform_fee":
		fallthrough
	case "bolt_account_fee":
		*e = CaptureType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CaptureType: %v", v)
	}
}

// Splits - A split of fees by type and amount.
type Splits struct {
	Amount *AmountView `json:"amount,omitempty"`
	// Fee type options. **Nullable** for Transactions Details.
	//
	Type *CaptureType `json:"type,omitempty"`
}

func (o *Splits) GetAmount() *AmountView {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *Splits) GetType() *CaptureType {
	if o == nil {
		return nil
	}
	return o.Type
}

// Capture - Deprecated. Use `captures`.
type Capture struct {
	Amount *AmountView `json:"amount,omitempty"`
	// The unique ID for the capture. **Nullable** for Transactions Details.
	ID *string `json:"id,omitempty"`
	// The reference ID associated with a transaction event (auth, capture, refund, void). This is an arbitrary identifier created by the merchant. Bolt does not enforce any uniqueness constraints on this ID. It is up to the merchant to generate identifiers that properly fulfill its needs.
	MerchantEventID *string `json:"merchant_event_id,omitempty"`
	// Additional information about the capture. For example, the processor capture ID. **Nullable** for Transactions Details.
	Metadata map[string]string `json:"metadata,omitempty"`
	// A split of fees by type and amount. **Nullable** for Transactions Details.
	Splits []Splits `json:"splits,omitempty"`
	// The status of the capture. **Nullable** for Transactions Details.
	Status *CaptureStatus `json:"status,omitempty"`
}

func (o *Capture) GetAmount() *AmountView {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *Capture) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Capture) GetMerchantEventID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantEventID
}

func (o *Capture) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Capture) GetSplits() []Splits {
	if o == nil {
		return nil
	}
	return o.Splits
}

func (o *Capture) GetStatus() *CaptureStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
