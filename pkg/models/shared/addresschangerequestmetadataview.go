// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AddressChangeRequestMetadataView struct {
	CanChangeShippingAddress *bool   `json:"can_change_shipping_address,omitempty"`
	ID                       *string `json:"id,omitempty"`
	Status                   *string `json:"status,omitempty"`
	TicketID                 *string `json:"ticket_id,omitempty"`
	TicketStatus             *string `json:"ticket_status,omitempty"`
}

func (o *AddressChangeRequestMetadataView) GetCanChangeShippingAddress() *bool {
	if o == nil {
		return nil
	}
	return o.CanChangeShippingAddress
}

func (o *AddressChangeRequestMetadataView) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddressChangeRequestMetadataView) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AddressChangeRequestMetadataView) GetTicketID() *string {
	if o == nil {
		return nil
	}
	return o.TicketID
}

func (o *AddressChangeRequestMetadataView) GetTicketStatus() *string {
	if o == nil {
		return nil
	}
	return o.TicketStatus
}
