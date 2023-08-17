// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConsumerSelfViewPlatformAccountStatus string

const (
	ConsumerSelfViewPlatformAccountStatusNone     ConsumerSelfViewPlatformAccountStatus = "none"
	ConsumerSelfViewPlatformAccountStatusLinked   ConsumerSelfViewPlatformAccountStatus = "linked"
	ConsumerSelfViewPlatformAccountStatusUnlinked ConsumerSelfViewPlatformAccountStatus = "unlinked"
)

func (e ConsumerSelfViewPlatformAccountStatus) ToPointer() *ConsumerSelfViewPlatformAccountStatus {
	return &e
}

func (e *ConsumerSelfViewPlatformAccountStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "linked":
		fallthrough
	case "unlinked":
		*e = ConsumerSelfViewPlatformAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConsumerSelfViewPlatformAccountStatus: %v", v)
	}
}

type ConsumerSelfView struct {
	Authentication        *LoginView                             `json:"authentication,omitempty"`
	EmailVerified         *bool                                  `json:"email_verified,omitempty"`
	Emails                []EmailView                            `json:"emails,omitempty"`
	FirstName             *string                                `json:"first_name,omitempty"`
	ID                    *string                                `json:"id,omitempty"`
	LastName              *string                                `json:"last_name,omitempty"`
	Phones                []PhoneView                            `json:"phones,omitempty"`
	PlatformAccountStatus *ConsumerSelfViewPlatformAccountStatus `json:"platform_account_status,omitempty"`
}

func (o *ConsumerSelfView) GetAuthentication() *LoginView {
	if o == nil {
		return nil
	}
	return o.Authentication
}

func (o *ConsumerSelfView) GetEmailVerified() *bool {
	if o == nil {
		return nil
	}
	return o.EmailVerified
}

func (o *ConsumerSelfView) GetEmails() []EmailView {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *ConsumerSelfView) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *ConsumerSelfView) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConsumerSelfView) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *ConsumerSelfView) GetPhones() []PhoneView {
	if o == nil {
		return nil
	}
	return o.Phones
}

func (o *ConsumerSelfView) GetPlatformAccountStatus() *ConsumerSelfViewPlatformAccountStatus {
	if o == nil {
		return nil
	}
	return o.PlatformAccountStatus
}
