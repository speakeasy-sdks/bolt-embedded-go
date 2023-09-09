// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CartShipmentType - The type corresponding to this shipment, if applicable.
type CartShipmentType string

const (
	CartShipmentTypeDoorDelivery   CartShipmentType = "door_delivery"
	CartShipmentTypeInStorePickup  CartShipmentType = "in_store_pickup"
	CartShipmentTypeShipToStore    CartShipmentType = "ship_to_store"
	CartShipmentTypeShipToHomeOnly CartShipmentType = "ship_to_home_only"
	CartShipmentTypeUnknown        CartShipmentType = "unknown"
)

func (e CartShipmentType) ToPointer() *CartShipmentType {
	return &e
}

func (e *CartShipmentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "door_delivery":
		fallthrough
	case "in_store_pickup":
		fallthrough
	case "ship_to_store":
		fallthrough
	case "ship_to_home_only":
		fallthrough
	case "unknown":
		*e = CartShipmentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CartShipmentType: %v", v)
	}
}