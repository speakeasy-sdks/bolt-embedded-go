// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// V1AccountsView - Has Bolt Account
type V1AccountsView struct {
	HasBoltAccount *bool `json:"has_bolt_account,omitempty"`
}

func (o *V1AccountsView) GetHasBoltAccount() *bool {
	if o == nil {
		return nil
	}
	return o.HasBoltAccount
}
