// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CardTokenType - Used to define which payment processor generated the token for this credit card.
type CardTokenType string

const (
	CardTokenTypeVantiv                CardTokenType = "vantiv"
	CardTokenTypeApplepay              CardTokenType = "applepay"
	CardTokenTypeBolt                  CardTokenType = "bolt"
	CardTokenTypeStripe                CardTokenType = "stripe"
	CardTokenTypePlcc                  CardTokenType = "plcc"
	CardTokenTypeApplepayEncryptedBlob CardTokenType = "applepay_encrypted_blob"
)

func (e CardTokenType) ToPointer() *CardTokenType {
	return &e
}

func (e *CardTokenType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vantiv":
		fallthrough
	case "applepay":
		fallthrough
	case "bolt":
		fallthrough
	case "stripe":
		fallthrough
	case "plcc":
		fallthrough
	case "applepay_encrypted_blob":
		*e = CardTokenType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardTokenType: %v", v)
	}
}