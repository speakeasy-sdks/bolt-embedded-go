// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// InStoreShipment - A cart that is being prepared for shipment
type InStoreShipment struct {
	// The name of the carrier selected.
	Carrier *string `json:"carrier,omitempty"`
	// The amount. **Nullable** for Transactions Details.
	Cost    *Amounts `json:"cost,omitempty"`
	Default *bool    `json:"default,omitempty"`
	// The estimated delivery date.
	EstimatedDeliveryDate *string `json:"estimated_delivery_date,omitempty"`
	// True if shipment is expedited.
	Expedited *bool `json:"expedited,omitempty"`
	// Contains the gift option settings for wrapping and custom messages.
	GiftOptions *GiftOptions `json:"gift_options,omitempty"`
	// ID for billing address
	ID *string `json:"id,omitempty"`
	// Contains the package's width, eight, depth, and unit details.
	PackageDimension *PackageDimension `json:"package_dimension,omitempty"`
	// The type of package.
	PackageType    *string         `json:"package_type,omitempty"`
	PackageWeights *PackageWeights `json:"package_weights,omitempty"`
	// Reference for the object.
	Reference *string `json:"reference,omitempty"`
	// The service name.
	Service         *string                 `json:"service,omitempty"`
	ShippingAddress *ConsumerBillingAddress `json:"shipping_address,omitempty"`
	// The name of the shipping method.
	ShippingMethod *string `json:"shipping_method,omitempty"`
	// The signature.
	Signature *string `json:"signature,omitempty"`
	// The amount. **Nullable** for Transactions Details.
	TaxAmount   *Amounts     `json:"tax_amount,omitempty"`
	TotalWeight *TotalWeight `json:"total_weight,omitempty"`
}

func (o *InStoreShipment) GetCarrier() *string {
	if o == nil {
		return nil
	}
	return o.Carrier
}

func (o *InStoreShipment) GetCost() *Amounts {
	if o == nil {
		return nil
	}
	return o.Cost
}

func (o *InStoreShipment) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *InStoreShipment) GetEstimatedDeliveryDate() *string {
	if o == nil {
		return nil
	}
	return o.EstimatedDeliveryDate
}

func (o *InStoreShipment) GetExpedited() *bool {
	if o == nil {
		return nil
	}
	return o.Expedited
}

func (o *InStoreShipment) GetGiftOptions() *GiftOptions {
	if o == nil {
		return nil
	}
	return o.GiftOptions
}

func (o *InStoreShipment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InStoreShipment) GetPackageDimension() *PackageDimension {
	if o == nil {
		return nil
	}
	return o.PackageDimension
}

func (o *InStoreShipment) GetPackageType() *string {
	if o == nil {
		return nil
	}
	return o.PackageType
}

func (o *InStoreShipment) GetPackageWeights() *PackageWeights {
	if o == nil {
		return nil
	}
	return o.PackageWeights
}

func (o *InStoreShipment) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *InStoreShipment) GetService() *string {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *InStoreShipment) GetShippingAddress() *ConsumerBillingAddress {
	if o == nil {
		return nil
	}
	return o.ShippingAddress
}

func (o *InStoreShipment) GetShippingMethod() *string {
	if o == nil {
		return nil
	}
	return o.ShippingMethod
}

func (o *InStoreShipment) GetSignature() *string {
	if o == nil {
		return nil
	}
	return o.Signature
}

func (o *InStoreShipment) GetTaxAmount() *Amounts {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *InStoreShipment) GetTotalWeight() *TotalWeight {
	if o == nil {
		return nil
	}
	return o.TotalWeight
}
