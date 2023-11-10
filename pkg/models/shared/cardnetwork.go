// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CardNetwork - The card's network code. **Nullable** for Transactions Details. Note: LEGACY diners_club_us_ca now tagged as mastercard
type CardNetwork string

const (
	CardNetworkUnknown        CardNetwork = "unknown"
	CardNetworkVisa           CardNetwork = "visa"
	CardNetworkMastercard     CardNetwork = "mastercard"
	CardNetworkAmex           CardNetwork = "amex"
	CardNetworkDiscover       CardNetwork = "discover"
	CardNetworkDinersClubUsCa CardNetwork = "diners_club_us_ca"
	CardNetworkJcb            CardNetwork = "jcb"
	CardNetworkUnionpay       CardNetwork = "unionpay"
	CardNetworkAlliancedata   CardNetwork = "alliancedata"
	CardNetworkCitiplcc       CardNetwork = "citiplcc"
)

func (e CardNetwork) ToPointer() *CardNetwork {
	return &e
}

func (e *CardNetwork) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unknown":
		fallthrough
	case "visa":
		fallthrough
	case "mastercard":
		fallthrough
	case "amex":
		fallthrough
	case "discover":
		fallthrough
	case "diners_club_us_ca":
		fallthrough
	case "jcb":
		fallthrough
	case "unionpay":
		fallthrough
	case "alliancedata":
		fallthrough
	case "citiplcc":
		*e = CardNetwork(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardNetwork: %v", v)
	}
}
