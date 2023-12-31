// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreditCardAuthorizationStatus - The status of the authorization request.
//   - `1` - succeeded
//   - `2` - declined
//   - `3` - error
type CreditCardAuthorizationStatus string

const (
	CreditCardAuthorizationStatusSucceeded CreditCardAuthorizationStatus = "succeeded"
	CreditCardAuthorizationStatusDeclined  CreditCardAuthorizationStatus = "declined"
	CreditCardAuthorizationStatusError     CreditCardAuthorizationStatus = "error"
)

func (e CreditCardAuthorizationStatus) ToPointer() *CreditCardAuthorizationStatus {
	return &e
}

func (e *CreditCardAuthorizationStatus) UnmarshalJSON(data []byte) error {
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
		*e = CreditCardAuthorizationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreditCardAuthorizationStatus: %v", v)
	}
}
