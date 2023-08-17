// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// EmailPriority - This is the priority of the contact method. This field's contents are not displayed in the transaction details view.
type EmailPriority string

const (
	EmailPriorityPrimary EmailPriority = "primary"
	EmailPriorityListed  EmailPriority = "listed"
)

func (e EmailPriority) ToPointer() *EmailPriority {
	return &e
}

func (e *EmailPriority) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "primary":
		fallthrough
	case "listed":
		*e = EmailPriority(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmailPriority: %v", v)
	}
}
