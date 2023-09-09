// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IWeight struct {
	Unit   *string  `json:"unit,omitempty"`
	Weight *float64 `json:"weight,omitempty"`
}

func (o *IWeight) GetUnit() *string {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *IWeight) GetWeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Weight
}