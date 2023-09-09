// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IFreeShippingDiscountView struct {
	IsFreeShipping     *bool    `json:"is_free_shipping,omitempty"`
	MaximumCostAllowed *float64 `json:"maximum_cost_allowed,omitempty"`
}

func (o *IFreeShippingDiscountView) GetIsFreeShipping() *bool {
	if o == nil {
		return nil
	}
	return o.IsFreeShipping
}

func (o *IFreeShippingDiscountView) GetMaximumCostAllowed() *float64 {
	if o == nil {
		return nil
	}
	return o.MaximumCostAllowed
}