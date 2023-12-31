// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Unit - The unit for this subscription's frequency.
type Unit string

const (
	UnitDay   Unit = "day"
	UnitWeek  Unit = "week"
	UnitMonth Unit = "month"
	UnitYear  Unit = "year"
)

func (e Unit) ToPointer() *Unit {
	return &e
}

func (e *Unit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "day":
		fallthrough
	case "week":
		fallthrough
	case "month":
		fallthrough
	case "year":
		*e = Unit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Unit: %v", v)
	}
}

// Frequency - Describes how often the subscription recurs.
type Frequency struct {
	// The unit for this subscription's frequency.
	Unit *Unit `json:"unit,omitempty"`
	// The value applied to the unit frequency.
	Value *int64 `json:"value,omitempty"`
}

func (o *Frequency) GetUnit() *Unit {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *Frequency) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}

// Subscription - Describes a product added as a recurring subscription.
type Subscription struct {
	// Describes how often the subscription recurs.
	Frequency *Frequency `json:"frequency,omitempty"`
}

func (o *Subscription) GetFrequency() *Frequency {
	if o == nil {
		return nil
	}
	return o.Frequency
}
