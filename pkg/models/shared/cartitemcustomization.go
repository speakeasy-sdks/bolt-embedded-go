// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CartItemCustomization struct {
	Attributes map[string]string `json:"attributes,omitempty"`
	Name       *string           `json:"name,omitempty"`
	Price      *AmountView       `json:"price,omitempty"`
}

func (o *CartItemCustomization) GetAttributes() map[string]string {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *CartItemCustomization) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CartItemCustomization) GetPrice() *AmountView {
	if o == nil {
		return nil
	}
	return o.Price
}
