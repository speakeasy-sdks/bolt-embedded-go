// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OrderExternalDataView struct {
	Shopify *string `json:"shopify,omitempty"`
}

func (o *OrderExternalDataView) GetShopify() *string {
	if o == nil {
		return nil
	}
	return o.Shopify
}
