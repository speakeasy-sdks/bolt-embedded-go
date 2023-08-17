// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TransactionProcessor - The processor used. **Nullable** for Transactions Details.
type TransactionProcessor string

const (
	TransactionProcessorAdyenGateway    TransactionProcessor = "adyen_gateway"
	TransactionProcessorAdyenPayfac     TransactionProcessor = "adyen_payfac"
	TransactionProcessorAffirm          TransactionProcessor = "affirm"
	TransactionProcessorAfterpay        TransactionProcessor = "afterpay"
	TransactionProcessorAllianceData    TransactionProcessor = "alliance_data"
	TransactionProcessorAmazonPay       TransactionProcessor = "amazon_pay"
	TransactionProcessorAuthorizeNet    TransactionProcessor = "authorize_net"
	TransactionProcessorBraintree       TransactionProcessor = "braintree"
	TransactionProcessorCheckoutCom     TransactionProcessor = "checkout_com"
	TransactionProcessorCybersource     TransactionProcessor = "cybersource"
	TransactionProcessorFirstData       TransactionProcessor = "first_data"
	TransactionProcessorKlarna          TransactionProcessor = "klarna"
	TransactionProcessorNmi             TransactionProcessor = "nmi"
	TransactionProcessorOrbital         TransactionProcessor = "orbital"
	TransactionProcessorPaypal          TransactionProcessor = "paypal"
	TransactionProcessorRadial          TransactionProcessor = "radial"
	TransactionProcessorRadialKlarna    TransactionProcessor = "radial_klarna"
	TransactionProcessorRadialPaypal    TransactionProcessor = "radial_paypal"
	TransactionProcessorRocketgate      TransactionProcessor = "rocketgate"
	TransactionProcessorSezzle          TransactionProcessor = "sezzle"
	TransactionProcessorShopifyPayments TransactionProcessor = "shopify_payments"
	TransactionProcessorStripe          TransactionProcessor = "stripe"
	TransactionProcessorVantiv          TransactionProcessor = "vantiv"
)

func (e TransactionProcessor) ToPointer() *TransactionProcessor {
	return &e
}

func (e *TransactionProcessor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "adyen_gateway":
		fallthrough
	case "adyen_payfac":
		fallthrough
	case "affirm":
		fallthrough
	case "afterpay":
		fallthrough
	case "alliance_data":
		fallthrough
	case "amazon_pay":
		fallthrough
	case "authorize_net":
		fallthrough
	case "braintree":
		fallthrough
	case "checkout_com":
		fallthrough
	case "cybersource":
		fallthrough
	case "first_data":
		fallthrough
	case "klarna":
		fallthrough
	case "nmi":
		fallthrough
	case "orbital":
		fallthrough
	case "paypal":
		fallthrough
	case "radial":
		fallthrough
	case "radial_klarna":
		fallthrough
	case "radial_paypal":
		fallthrough
	case "rocketgate":
		fallthrough
	case "sezzle":
		fallthrough
	case "shopify_payments":
		fallthrough
	case "stripe":
		fallthrough
	case "vantiv":
		*e = TransactionProcessor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransactionProcessor: %v", v)
	}
}
