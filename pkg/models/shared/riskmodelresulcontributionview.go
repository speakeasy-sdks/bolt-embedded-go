// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RiskModelResulContributionView struct {
	Category *string `json:"category,omitempty"`
	Weight   *string `json:"weight,omitempty"`
}

func (o *RiskModelResulContributionView) GetCategory() *string {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *RiskModelResulContributionView) GetWeight() *string {
	if o == nil {
		return nil
	}
	return o.Weight
}