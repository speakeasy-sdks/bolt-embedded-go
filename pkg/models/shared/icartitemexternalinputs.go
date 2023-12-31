// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ICartItemExternalInputs struct {
	ShopifyLineItemReference       *float64 `json:"shopify_line_item_reference,omitempty"`
	ShopifyProductReference        *float64 `json:"shopify_product_reference,omitempty"`
	ShopifyProductVariantReference *float64 `json:"shopify_product_variant_reference,omitempty"`
}

func (o *ICartItemExternalInputs) GetShopifyLineItemReference() *float64 {
	if o == nil {
		return nil
	}
	return o.ShopifyLineItemReference
}

func (o *ICartItemExternalInputs) GetShopifyProductReference() *float64 {
	if o == nil {
		return nil
	}
	return o.ShopifyProductReference
}

func (o *ICartItemExternalInputs) GetShopifyProductVariantReference() *float64 {
	if o == nil {
		return nil
	}
	return o.ShopifyProductVariantReference
}
