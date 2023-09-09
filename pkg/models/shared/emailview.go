// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EmailView struct {
	Address *string `json:"address,omitempty"`
	ID      *string `json:"id,omitempty"`
	// Describes the card's priority.
	//
	Priority *Priority `json:"priority,omitempty"`
	Status   *string   `json:"status,omitempty"`
}

func (o *EmailView) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *EmailView) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EmailView) GetPriority() *Priority {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *EmailView) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}