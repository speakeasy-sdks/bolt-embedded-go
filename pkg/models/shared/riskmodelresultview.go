// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RiskModelResultView struct {
	Contribution []RiskModelResulContributionView `json:"contribution,omitempty"`
}

func (o *RiskModelResultView) GetContribution() []RiskModelResulContributionView {
	if o == nil {
		return nil
	}
	return o.Contribution
}
