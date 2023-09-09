// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PackageWeights struct {
	// The unit of measurement for an item's weight.
	Unit *string `json:"unit,omitempty"`
	// The weight of an item.
	Weight *int64 `json:"weight,omitempty"`
}

func (o *PackageWeights) GetUnit() *string {
	if o == nil {
		return nil
	}
	return o.Unit
}

func (o *PackageWeights) GetWeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Weight
}