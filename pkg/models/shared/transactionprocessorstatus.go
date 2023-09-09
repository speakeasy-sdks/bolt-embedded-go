// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TransactionProcessorStatus - The processor's status. Only `primary` and `active` processor are displayed.
type TransactionProcessorStatus string

const (
	TransactionProcessorStatusPrimary  TransactionProcessorStatus = "primary"
	TransactionProcessorStatusActive   TransactionProcessorStatus = "active"
	TransactionProcessorStatusInactive TransactionProcessorStatus = "inactive"
)

func (e TransactionProcessorStatus) ToPointer() *TransactionProcessorStatus {
	return &e
}

func (e *TransactionProcessorStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "primary":
		fallthrough
	case "active":
		fallthrough
	case "inactive":
		*e = TransactionProcessorStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionProcessorStatus: %v", v)
	}
}