// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CardDisplayNetwork - The card's network. **Nullable** for Transactions Details.
type CardDisplayNetwork string

const (
	CardDisplayNetworkCreditCard      CardDisplayNetwork = "Credit Card"
	CardDisplayNetworkAmericanExpress CardDisplayNetwork = "American Express"
	CardDisplayNetworkDinersClub      CardDisplayNetwork = "Diners Club"
	CardDisplayNetworkDiscover        CardDisplayNetwork = "Discover"
	CardDisplayNetworkJcb             CardDisplayNetwork = "JCB"
	CardDisplayNetworkMasterCard      CardDisplayNetwork = "MasterCard"
	CardDisplayNetworkUnionPay        CardDisplayNetwork = "Union Pay"
	CardDisplayNetworkVisa            CardDisplayNetwork = "Visa"
)

func (e CardDisplayNetwork) ToPointer() *CardDisplayNetwork {
	return &e
}

func (e *CardDisplayNetwork) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Credit Card":
		fallthrough
	case "American Express":
		fallthrough
	case "Diners Club":
		fallthrough
	case "Discover":
		fallthrough
	case "JCB":
		fallthrough
	case "MasterCard":
		fallthrough
	case "Union Pay":
		fallthrough
	case "Visa":
		*e = CardDisplayNetwork(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardDisplayNetwork: %v", v)
	}
}