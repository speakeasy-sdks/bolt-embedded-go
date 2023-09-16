// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CartCreateFees struct {
	// Description of the fee that will appear in the tooltip if the mouse hovers over the fee.
	Description *string `json:"description,omitempty"`
	// Name of the fee that will appear in the order ledger.
	Name     *string `json:"name"`
	Quantity float64 `json:"quantity"`
	// Unique reference used to identify the fee.
	Reference     string  `json:"reference"`
	UnitPrice     float64 `json:"unit_price"`
	UnitTaxAmount float64 `json:"unit_tax_amount"`
}

func (o *CartCreateFees) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CartCreateFees) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CartCreateFees) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *CartCreateFees) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *CartCreateFees) GetUnitPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.UnitPrice
}

func (o *CartCreateFees) GetUnitTaxAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.UnitTaxAmount
}

type CartCreateFulfillmentsDigitalDelivery struct {
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

func (o *CartCreateFulfillmentsDigitalDelivery) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *CartCreateFulfillmentsDigitalDelivery) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

type CartCreateFulfillmentsType string

const (
	CartCreateFulfillmentsTypePhysicalDoorDelivery  CartCreateFulfillmentsType = "physical_door_delivery"
	CartCreateFulfillmentsTypePhysicalShipToStore   CartCreateFulfillmentsType = "physical_ship_to_store"
	CartCreateFulfillmentsTypePhysicalInStorePickup CartCreateFulfillmentsType = "physical_in_store_pickup"
	CartCreateFulfillmentsTypeDigitalDownload       CartCreateFulfillmentsType = "digital_download"
	CartCreateFulfillmentsTypeDigitalNoDelivery     CartCreateFulfillmentsType = "digital_no_delivery"
)

func (e CartCreateFulfillmentsType) ToPointer() *CartCreateFulfillmentsType {
	return &e
}

func (e *CartCreateFulfillmentsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "physical_door_delivery":
		fallthrough
	case "physical_ship_to_store":
		fallthrough
	case "physical_in_store_pickup":
		fallthrough
	case "digital_download":
		fallthrough
	case "digital_no_delivery":
		*e = CartCreateFulfillmentsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CartCreateFulfillmentsType: %v", v)
	}
}

// CartCreateFulfillments - Defines the shipments associated with the cart items.
type CartCreateFulfillments struct {
	CartItems []CartItem `json:"cart_items,omitempty"`
	// A cart that is being prepared for shipment
	CartShipment        *CartShipment                          `json:"cart_shipment,omitempty"`
	DigitalDelivery     *CartCreateFulfillmentsDigitalDelivery `json:"digital_delivery,omitempty"`
	InStoreCartShipment *InStoreCartShipment                   `json:"in_store_cart_shipment,omitempty"`
	Type                *CartCreateFulfillmentsType            `json:"type,omitempty"`
}

func (o *CartCreateFulfillments) GetCartItems() []CartItem {
	if o == nil {
		return nil
	}
	return o.CartItems
}

func (o *CartCreateFulfillments) GetCartShipment() *CartShipment {
	if o == nil {
		return nil
	}
	return o.CartShipment
}

func (o *CartCreateFulfillments) GetDigitalDelivery() *CartCreateFulfillmentsDigitalDelivery {
	if o == nil {
		return nil
	}
	return o.DigitalDelivery
}

func (o *CartCreateFulfillments) GetInStoreCartShipment() *InStoreCartShipment {
	if o == nil {
		return nil
	}
	return o.InStoreCartShipment
}

func (o *CartCreateFulfillments) GetType() *CartCreateFulfillmentsType {
	if o == nil {
		return nil
	}
	return o.Type
}

// CartCreate - The base_cart object contains the core details typically found in most cart objects, including items, discounts, amount totals, shipments, and in-store pickups.
type CartCreate struct {
	AddOns []CartAddOn `json:"add_ons,omitempty"`
	// The Address object is used for billing, shipping, and physical store address use cases.
	BillingAddress *Address `json:"billing_address,omitempty"`
	// Used to provide a link to the cart ID.
	CartURL   *string        `json:"cart_url,omitempty"`
	Currency  string         `json:"currency"`
	Discounts []CartDiscount `json:"discounts,omitempty"`
	// This field, although required, can be an empty string.
	DisplayID            *string                  `json:"display_id,omitempty"`
	Fees                 []CartCreateFees         `json:"fees,omitempty"`
	Fulfillments         []CartCreateFulfillments `json:"fulfillments,omitempty"`
	InStoreCartShipments []InStoreCartShipment    `json:"in_store_cart_shipments,omitempty"`
	// The list of items associated with the cart.
	Items          []CartItem           `json:"items,omitempty"`
	LoyaltyRewards []CartLoyaltyRewards `json:"loyalty_rewards,omitempty"`
	// Optional custom metadata.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Used optionally to pass additional information like order numbers or other IDs as needed.
	OrderDescription *string `json:"order_description,omitempty"`
	// This value is used by Bolt as an external reference to a given order. This reference must be unique per successful transaction.
	OrderReference string         `json:"order_reference"`
	Shipments      []CartShipment `json:"shipments,omitempty"`
	// The total tax amount for all of the items associated with the cart.
	TaxAmount *float64 `json:"tax_amount,omitempty"`
	// The total amount, in cents, of the cart including its items and taxes (if applicable), e.g. $9.00 is 900. This total must match the sum of all other amounts.
	TotalAmount float64 `json:"total_amount"`
}

func (o *CartCreate) GetAddOns() []CartAddOn {
	if o == nil {
		return nil
	}
	return o.AddOns
}

func (o *CartCreate) GetBillingAddress() *Address {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *CartCreate) GetCartURL() *string {
	if o == nil {
		return nil
	}
	return o.CartURL
}

func (o *CartCreate) GetCurrency() string {
	if o == nil {
		return ""
	}
	return o.Currency
}

func (o *CartCreate) GetDiscounts() []CartDiscount {
	if o == nil {
		return nil
	}
	return o.Discounts
}

func (o *CartCreate) GetDisplayID() *string {
	if o == nil {
		return nil
	}
	return o.DisplayID
}

func (o *CartCreate) GetFees() []CartCreateFees {
	if o == nil {
		return nil
	}
	return o.Fees
}

func (o *CartCreate) GetFulfillments() []CartCreateFulfillments {
	if o == nil {
		return nil
	}
	return o.Fulfillments
}

func (o *CartCreate) GetInStoreCartShipments() []InStoreCartShipment {
	if o == nil {
		return nil
	}
	return o.InStoreCartShipments
}

func (o *CartCreate) GetItems() []CartItem {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *CartCreate) GetLoyaltyRewards() []CartLoyaltyRewards {
	if o == nil {
		return nil
	}
	return o.LoyaltyRewards
}

func (o *CartCreate) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CartCreate) GetOrderDescription() *string {
	if o == nil {
		return nil
	}
	return o.OrderDescription
}

func (o *CartCreate) GetOrderReference() string {
	if o == nil {
		return ""
	}
	return o.OrderReference
}

func (o *CartCreate) GetShipments() []CartShipment {
	if o == nil {
		return nil
	}
	return o.Shipments
}

func (o *CartCreate) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *CartCreate) GetTotalAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalAmount
}
