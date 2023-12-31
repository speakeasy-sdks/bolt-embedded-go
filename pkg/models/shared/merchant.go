// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Merchant struct {
	// The date the merchant account was created.  **Nullable** for Transactions Details.
	CreatedAt int64 `json:"created_at"`
	// The description of the merchant account. **Nullable** for Transactions Details.
	Description      string                        `json:"description"`
	OnboardingStatus *MerchantOnboardingStatusCode `json:"onboarding_status,omitempty"`
	// **Nullable** for Transactions Details.
	//
	OperationalProcessors []TransactionOperationalProcessor `json:"operational_processors"`
	// The processor used. **Nullable** for Transactions Details.
	Processor TransactionProcessor `json:"processor"`
	// The unique public ID for the merchant's Bolt account. A merchant account contains many merchant divisions.
	PublicID *string `json:"public_id,omitempty"`
	// The merchant's status:
	//   * `1` - Active
	//   * `2` - Inactive
	//   * `3` - Offboarding
	//
	Status *MerchantStatus `json:"status,omitempty"`
	// The timezone of the merchant. **Nullable** for Transactions Details.
	TimeZone string `json:"time_zone"`
}

func (o *Merchant) GetCreatedAt() int64 {
	if o == nil {
		return 0
	}
	return o.CreatedAt
}

func (o *Merchant) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Merchant) GetOnboardingStatus() *MerchantOnboardingStatusCode {
	if o == nil {
		return nil
	}
	return o.OnboardingStatus
}

func (o *Merchant) GetOperationalProcessors() []TransactionOperationalProcessor {
	if o == nil {
		return []TransactionOperationalProcessor{}
	}
	return o.OperationalProcessors
}

func (o *Merchant) GetProcessor() TransactionProcessor {
	if o == nil {
		return TransactionProcessor("")
	}
	return o.Processor
}

func (o *Merchant) GetPublicID() *string {
	if o == nil {
		return nil
	}
	return o.PublicID
}

func (o *Merchant) GetStatus() *MerchantStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Merchant) GetTimeZone() string {
	if o == nil {
		return ""
	}
	return o.TimeZone
}
