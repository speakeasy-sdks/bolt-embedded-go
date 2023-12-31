// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TransactionTimelineViewType string

const (
	TransactionTimelineViewTypeCompleted     TransactionTimelineViewType = "completed"
	TransactionTimelineViewTypeAuthorized    TransactionTimelineViewType = "authorized"
	TransactionTimelineViewTypeReview        TransactionTimelineViewType = "review"
	TransactionTimelineViewTypeNote          TransactionTimelineViewType = "note"
	TransactionTimelineViewTypeVoided        TransactionTimelineViewType = "voided"
	TransactionTimelineViewTypeCaptured      TransactionTimelineViewType = "captured"
	TransactionTimelineViewTypeCredited      TransactionTimelineViewType = "credited"
	TransactionTimelineViewTypeAddressChange TransactionTimelineViewType = "address_change"
)

func (e TransactionTimelineViewType) ToPointer() *TransactionTimelineViewType {
	return &e
}

func (e *TransactionTimelineViewType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "completed":
		fallthrough
	case "authorized":
		fallthrough
	case "review":
		fallthrough
	case "note":
		fallthrough
	case "voided":
		fallthrough
	case "captured":
		fallthrough
	case "credited":
		fallthrough
	case "address_change":
		*e = TransactionTimelineViewType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionTimelineViewType: %v", v)
	}
}

type TransactionTimelineView struct {
	AddressChange *AddressChangeView           `json:"address_change,omitempty"`
	Amount        *AmountView                  `json:"amount,omitempty"`
	Consumer      *ConsumerSummaryView         `json:"consumer,omitempty"`
	Date          *float64                     `json:"date,omitempty"`
	Note          *string                      `json:"note,omitempty"`
	Review        *TransactionReviewView       `json:"review,omitempty"`
	Transaction   *TransactionView             `json:"transaction,omitempty"`
	Type          *TransactionTimelineViewType `json:"type,omitempty"`
	Visibility    *string                      `json:"visibility,omitempty"`
}

func (o *TransactionTimelineView) GetAddressChange() *AddressChangeView {
	if o == nil {
		return nil
	}
	return o.AddressChange
}

func (o *TransactionTimelineView) GetAmount() *AmountView {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *TransactionTimelineView) GetConsumer() *ConsumerSummaryView {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *TransactionTimelineView) GetDate() *float64 {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *TransactionTimelineView) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *TransactionTimelineView) GetReview() *TransactionReviewView {
	if o == nil {
		return nil
	}
	return o.Review
}

func (o *TransactionTimelineView) GetTransaction() *TransactionView {
	if o == nil {
		return nil
	}
	return o.Transaction
}

func (o *TransactionTimelineView) GetType() *TransactionTimelineViewType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *TransactionTimelineView) GetVisibility() *string {
	if o == nil {
		return nil
	}
	return o.Visibility
}
