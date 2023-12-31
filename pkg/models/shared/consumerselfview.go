// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PlatformAccountStatus string

const (
	PlatformAccountStatusNone     PlatformAccountStatus = "none"
	PlatformAccountStatusLinked   PlatformAccountStatus = "linked"
	PlatformAccountStatusUnlinked PlatformAccountStatus = "unlinked"
)

func (e PlatformAccountStatus) ToPointer() *PlatformAccountStatus {
	return &e
}

func (e *PlatformAccountStatus) UnmarshalJSON(data []byte) error {
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
		*e = PlatformAccountStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlatformAccountStatus: %v", v)
	}
}

type ConsumerSelfView struct {
	Authentication        *LoginView             `json:"authentication,omitempty"`
	EmailVerified         *bool                  `json:"email_verified,omitempty"`
	Emails                []EmailView            `json:"emails,omitempty"`
	FirstName             *string                `json:"first_name,omitempty"`
	ID                    *string                `json:"id,omitempty"`
	LastName              *string                `json:"last_name,omitempty"`
	Phones                []PhoneView            `json:"phones,omitempty"`
	PlatformAccountStatus *PlatformAccountStatus `json:"platform_account_status,omitempty"`
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

func (o *ConsumerSelfView) GetPlatformAccountStatus() *PlatformAccountStatus {
	if o == nil {
		return nil
	}
	return o.PlatformAccountStatus
}
