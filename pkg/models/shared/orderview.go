// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OrderView struct {
	// This is the cart object returned in a successful response.
	Cart           *CartView              `json:"cart,omitempty"`
	DynamicContent *IOrderDynamicContent  `json:"dynamic_content,omitempty"`
	ExternalData   *OrderExternalDataView `json:"external_data,omitempty"`
	PlatformUserID *string                `json:"platform_user_id,omitempty"`
	RequiresAction *string                `json:"requires_action,omitempty"`
	Token          *string                `json:"token,omitempty"`
	// Used by shoppers to make extra requests or provide details for gift messages.
	UserNote *string `json:"user_note,omitempty"`
}

func (o *OrderView) GetCart() *CartView {
	if o == nil {
		return nil
	}
	return o.Cart
}

func (o *OrderView) GetDynamicContent() *IOrderDynamicContent {
	if o == nil {
		return nil
	}
	return o.DynamicContent
}

func (o *OrderView) GetExternalData() *OrderExternalDataView {
	if o == nil {
		return nil
	}
	return o.ExternalData
}

func (o *OrderView) GetPlatformUserID() *string {
	if o == nil {
		return nil
	}
	return o.PlatformUserID
}

func (o *OrderView) GetRequiresAction() *string {
	if o == nil {
		return nil
	}
	return o.RequiresAction
}

func (o *OrderView) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OrderView) GetUserNote() *string {
	if o == nil {
		return nil
	}
	return o.UserNote
}
