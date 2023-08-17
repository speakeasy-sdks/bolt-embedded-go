// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CardStatus - The card's status. **Nullable** for Transactions Details.
type CardStatus string

const (
	CardStatusActive    CardStatus = "active"
	CardStatusCreated   CardStatus = "created"
	CardStatusInactive  CardStatus = "inactive"
	CardStatusTransient CardStatus = "transient"
)

func (e CardStatus) ToPointer() *CardStatus {
	return &e
}

func (e *CardStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "created":
		fallthrough
	case "inactive":
		fallthrough
	case "transient":
		*e = CardStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardStatus: %v", v)
	}
}
