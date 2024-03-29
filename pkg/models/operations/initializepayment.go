// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/utils"
	"net/http"
)

// SavedPaymentInputInitializePaymentData - Initialize payment for a saved payment method
type SavedPaymentInputInitializePaymentData struct {
	// Payment ID of the saved Bolt Payment method.
	ID *string `json:"id,omitempty"`
}

func (o *SavedPaymentInputInitializePaymentData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SavedPaymentInputInitializeType - The type of the payment attempt
type SavedPaymentInputInitializeType string

const (
	SavedPaymentInputInitializeTypePaypal             SavedPaymentInputInitializeType = "paypal"
	SavedPaymentInputInitializeTypeSavedPaymentMethod SavedPaymentInputInitializeType = "saved_payment_method"
)

func (e SavedPaymentInputInitializeType) ToPointer() *SavedPaymentInputInitializeType {
	return &e
}

func (e *SavedPaymentInputInitializeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paypal":
		fallthrough
	case "saved_payment_method":
		*e = SavedPaymentInputInitializeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedPaymentInputInitializeType: %v", v)
	}
}

type SavedPaymentInputInitializePaymentMethod struct {
	// Initialize payment for a saved payment method
	PaymentData SavedPaymentInputInitializePaymentData `json:"payment_data"`
	// The type of the payment attempt
	Type SavedPaymentInputInitializeType `json:"type"`
}

func (o *SavedPaymentInputInitializePaymentMethod) GetPaymentData() SavedPaymentInputInitializePaymentData {
	if o == nil {
		return SavedPaymentInputInitializePaymentData{}
	}
	return o.PaymentData
}

func (o *SavedPaymentInputInitializePaymentMethod) GetType() SavedPaymentInputInitializeType {
	if o == nil {
		return SavedPaymentInputInitializeType("")
	}
	return o.Type
}

// SavedPaymentInputInitializeShopperIdentity - Identification information for the Shopper. This is only required when creating a new Bolt account.
type SavedPaymentInputInitializeShopperIdentity struct {
	// determines whether to create a bolt account for this shopper
	CreateBoltAccount *bool `json:"create_bolt_account,omitempty"`
	// Email address of the shopper
	Email string `json:"email"`
	// First name of the shopper
	FirstName string `json:"first_name"`
	// Last name of the shopper
	LastName string `json:"last_name"`
	// Phone number of the shopper
	Phone string `json:"phone"`
}

func (o *SavedPaymentInputInitializeShopperIdentity) GetCreateBoltAccount() *bool {
	if o == nil {
		return nil
	}
	return o.CreateBoltAccount
}

func (o *SavedPaymentInputInitializeShopperIdentity) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SavedPaymentInputInitializeShopperIdentity) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *SavedPaymentInputInitializeShopperIdentity) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *SavedPaymentInputInitializeShopperIdentity) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

type SavedPaymentInputInitializeSchemas struct {
	// The details of the cart being purchased with this payment.
	Cart          shared.CartCreate                        `json:"cart"`
	PaymentMethod SavedPaymentInputInitializePaymentMethod `json:"payment_method"`
	// Identification information for the Shopper. This is only required when creating a new Bolt account.
	ShopperIdentity *SavedPaymentInputInitializeShopperIdentity `json:"shopper_identity,omitempty"`
}

func (o *SavedPaymentInputInitializeSchemas) GetCart() shared.CartCreate {
	if o == nil {
		return shared.CartCreate{}
	}
	return o.Cart
}

func (o *SavedPaymentInputInitializeSchemas) GetPaymentMethod() SavedPaymentInputInitializePaymentMethod {
	if o == nil {
		return SavedPaymentInputInitializePaymentMethod{}
	}
	return o.PaymentMethod
}

func (o *SavedPaymentInputInitializeSchemas) GetShopperIdentity() *SavedPaymentInputInitializeShopperIdentity {
	if o == nil {
		return nil
	}
	return o.ShopperIdentity
}

// PaymentData - Initialize payment for a new PayPal order.
type PaymentData struct {
	// Redirect URL for canceled PayPal transaction.
	Cancel *string `json:"cancel,omitempty"`
	// Redirect URL for successful PayPal transaction.
	Success *string `json:"success,omitempty"`
}

func (o *PaymentData) GetCancel() *string {
	if o == nil {
		return nil
	}
	return o.Cancel
}

func (o *PaymentData) GetSuccess() *string {
	if o == nil {
		return nil
	}
	return o.Success
}

// Type - The type of the payment attempt
type Type string

const (
	TypePaypal             Type = "paypal"
	TypeSavedPaymentMethod Type = "saved_payment_method"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paypal":
		fallthrough
	case "saved_payment_method":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type PaymentMethod struct {
	// Initialize payment for a new PayPal order.
	PaymentData PaymentData `json:"payment_data"`
	// The type of the payment attempt
	Type Type `json:"type"`
}

func (o *PaymentMethod) GetPaymentData() PaymentData {
	if o == nil {
		return PaymentData{}
	}
	return o.PaymentData
}

func (o *PaymentMethod) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

// PaypalPaymentInputInitializeShopperIdentity - Identification information for the Shopper. This is only required when creating a new Bolt account.
type PaypalPaymentInputInitializeShopperIdentity struct {
	// determines whether to create a bolt account for this shopper
	CreateBoltAccount *bool `json:"create_bolt_account,omitempty"`
	// Email address of the shopper
	Email string `json:"email"`
	// First name of the shopper
	FirstName string `json:"first_name"`
	// Last name of the shopper
	LastName string `json:"last_name"`
	// Phone number of the shopper
	Phone string `json:"phone"`
}

func (o *PaypalPaymentInputInitializeShopperIdentity) GetCreateBoltAccount() *bool {
	if o == nil {
		return nil
	}
	return o.CreateBoltAccount
}

func (o *PaypalPaymentInputInitializeShopperIdentity) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *PaypalPaymentInputInitializeShopperIdentity) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *PaypalPaymentInputInitializeShopperIdentity) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *PaypalPaymentInputInitializeShopperIdentity) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

type Schemas struct {
	// The details of the cart being purchased with this payment.
	Cart          shared.CartCreate `json:"cart"`
	PaymentMethod PaymentMethod     `json:"payment_method"`
	// Identification information for the Shopper. This is only required when creating a new Bolt account.
	ShopperIdentity *PaypalPaymentInputInitializeShopperIdentity `json:"shopper_identity,omitempty"`
}

func (o *Schemas) GetCart() shared.CartCreate {
	if o == nil {
		return shared.CartCreate{}
	}
	return o.Cart
}

func (o *Schemas) GetPaymentMethod() PaymentMethod {
	if o == nil {
		return PaymentMethod{}
	}
	return o.PaymentMethod
}

func (o *Schemas) GetShopperIdentity() *PaypalPaymentInputInitializeShopperIdentity {
	if o == nil {
		return nil
	}
	return o.ShopperIdentity
}

type InitializePaymentRequestBodyType string

const (
	InitializePaymentRequestBodyTypeSchemas                            InitializePaymentRequestBodyType = "Schemas"
	InitializePaymentRequestBodyTypeSavedPaymentInputInitializeSchemas InitializePaymentRequestBodyType = "saved_payment_input_initialize_Schemas"
)

type InitializePaymentRequestBody struct {
	Schemas                            *Schemas
	SavedPaymentInputInitializeSchemas *SavedPaymentInputInitializeSchemas

	Type InitializePaymentRequestBodyType
}

func CreateInitializePaymentRequestBodySchemas(schemas Schemas) InitializePaymentRequestBody {
	typ := InitializePaymentRequestBodyTypeSchemas

	return InitializePaymentRequestBody{
		Schemas: &schemas,
		Type:    typ,
	}
}

func CreateInitializePaymentRequestBodySavedPaymentInputInitializeSchemas(savedPaymentInputInitializeSchemas SavedPaymentInputInitializeSchemas) InitializePaymentRequestBody {
	typ := InitializePaymentRequestBodyTypeSavedPaymentInputInitializeSchemas

	return InitializePaymentRequestBody{
		SavedPaymentInputInitializeSchemas: &savedPaymentInputInitializeSchemas,
		Type:                               typ,
	}
}

func (u *InitializePaymentRequestBody) UnmarshalJSON(data []byte) error {

	schemas := Schemas{}
	if err := utils.UnmarshalJSON(data, &schemas, "", true, true); err == nil {
		u.Schemas = &schemas
		u.Type = InitializePaymentRequestBodyTypeSchemas
		return nil
	}

	savedPaymentInputInitializeSchemas := SavedPaymentInputInitializeSchemas{}
	if err := utils.UnmarshalJSON(data, &savedPaymentInputInitializeSchemas, "", true, true); err == nil {
		u.SavedPaymentInputInitializeSchemas = &savedPaymentInputInitializeSchemas
		u.Type = InitializePaymentRequestBodyTypeSavedPaymentInputInitializeSchemas
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u InitializePaymentRequestBody) MarshalJSON() ([]byte, error) {
	if u.Schemas != nil {
		return utils.MarshalJSON(u.Schemas, "", true)
	}

	if u.SavedPaymentInputInitializeSchemas != nil {
		return utils.MarshalJSON(u.SavedPaymentInputInitializeSchemas, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type InitializePaymentRequest struct {
	// A key created by merchants that ensures `POST` and `PATCH` requests are only performed once. [Read more about Idempotent Requests here](/developers/references/idempotency/).
	IdempotencyKey *string                       `header:"style=simple,explode=false,name=Idempotency-Key"`
	RequestBody    *InitializePaymentRequestBody `request:"mediaType=application/json"`
	// The publicly viewable identifier used to identify a merchant division. This key is found in the Developer > API section of the Bolt Merchant Dashboard [RECOMMENDED].
	XPublishableKey *string `header:"style=simple,explode=false,name=X-Publishable-Key"`
}

func (o *InitializePaymentRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *InitializePaymentRequest) GetRequestBody() *InitializePaymentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *InitializePaymentRequest) GetXPublishableKey() *string {
	if o == nil {
		return nil
	}
	return o.XPublishableKey
}

// SavedPaymentViewAction - Action after initializing payment
type SavedPaymentViewAction struct {
	// action method
	Method *string `json:"method,omitempty"`
	// action type
	Type *string `json:"type,omitempty"`
	// action URL
	URL *string `json:"url,omitempty"`
}

func (o *SavedPaymentViewAction) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *SavedPaymentViewAction) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SavedPaymentViewAction) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// SavedPaymentViewStatus - The current payment status.
type SavedPaymentViewStatus string

const (
	SavedPaymentViewStatusAwaitingUserConfirmation SavedPaymentViewStatus = "awaiting_user_confirmation"
	SavedPaymentViewStatusPaymentReady             SavedPaymentViewStatus = "payment_ready"
	SavedPaymentViewStatusSuccess                  SavedPaymentViewStatus = "success"
)

func (e SavedPaymentViewStatus) ToPointer() *SavedPaymentViewStatus {
	return &e
}

func (e *SavedPaymentViewStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "awaiting_user_confirmation":
		fallthrough
	case "payment_ready":
		fallthrough
	case "success":
		*e = SavedPaymentViewStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedPaymentViewStatus: %v", v)
	}
}

type SavedPaymentViewSchemas struct {
	// Action after initializing payment
	Action *SavedPaymentViewAction `json:"action,omitempty"`
	// The ID for a Payment Attempt
	ID *string `json:"id,omitempty"`
	// The current payment status.
	Status *SavedPaymentViewStatus `json:"status,omitempty"`
}

func (o *SavedPaymentViewSchemas) GetAction() *SavedPaymentViewAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *SavedPaymentViewSchemas) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SavedPaymentViewSchemas) GetStatus() *SavedPaymentViewStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

// Action after initializing payment
type Action struct {
	// action method
	Method *string `json:"method,omitempty"`
	// action type
	Type *string `json:"type,omitempty"`
	// action URL
	URL *string `json:"url,omitempty"`
}

func (o *Action) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *Action) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Action) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

// Status - The current payment status.
type Status string

const (
	StatusAwaitingUserConfirmation Status = "awaiting_user_confirmation"
	StatusPaymentReady             Status = "payment_ready"
	StatusSuccess                  Status = "success"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "awaiting_user_confirmation":
		fallthrough
	case "payment_ready":
		fallthrough
	case "success":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type PaypalPaymentViewSchemas struct {
	// Action after initializing payment
	Action *Action `json:"action,omitempty"`
	// The ID for a Payment Attempt
	ID *string `json:"id,omitempty"`
	// The current payment status.
	Status *Status `json:"status,omitempty"`
}

func (o *PaypalPaymentViewSchemas) GetAction() *Action {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *PaypalPaymentViewSchemas) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PaypalPaymentViewSchemas) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

type InitializePaymentResponseBodyType string

const (
	InitializePaymentResponseBodyTypePaypalPaymentViewSchemas InitializePaymentResponseBodyType = "paypal_payment_view_Schemas"
	InitializePaymentResponseBodyTypeSavedPaymentViewSchemas  InitializePaymentResponseBodyType = "saved_payment_view_Schemas"
)

// InitializePaymentResponseBody - Payment token retrieved.
type InitializePaymentResponseBody struct {
	PaypalPaymentViewSchemas *PaypalPaymentViewSchemas
	SavedPaymentViewSchemas  *SavedPaymentViewSchemas

	Type InitializePaymentResponseBodyType
}

func CreateInitializePaymentResponseBodyPaypalPaymentViewSchemas(paypalPaymentViewSchemas PaypalPaymentViewSchemas) InitializePaymentResponseBody {
	typ := InitializePaymentResponseBodyTypePaypalPaymentViewSchemas

	return InitializePaymentResponseBody{
		PaypalPaymentViewSchemas: &paypalPaymentViewSchemas,
		Type:                     typ,
	}
}

func CreateInitializePaymentResponseBodySavedPaymentViewSchemas(savedPaymentViewSchemas SavedPaymentViewSchemas) InitializePaymentResponseBody {
	typ := InitializePaymentResponseBodyTypeSavedPaymentViewSchemas

	return InitializePaymentResponseBody{
		SavedPaymentViewSchemas: &savedPaymentViewSchemas,
		Type:                    typ,
	}
}

func (u *InitializePaymentResponseBody) UnmarshalJSON(data []byte) error {

	paypalPaymentViewSchemas := PaypalPaymentViewSchemas{}
	if err := utils.UnmarshalJSON(data, &paypalPaymentViewSchemas, "", true, true); err == nil {
		u.PaypalPaymentViewSchemas = &paypalPaymentViewSchemas
		u.Type = InitializePaymentResponseBodyTypePaypalPaymentViewSchemas
		return nil
	}

	savedPaymentViewSchemas := SavedPaymentViewSchemas{}
	if err := utils.UnmarshalJSON(data, &savedPaymentViewSchemas, "", true, true); err == nil {
		u.SavedPaymentViewSchemas = &savedPaymentViewSchemas
		u.Type = InitializePaymentResponseBodyTypeSavedPaymentViewSchemas
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u InitializePaymentResponseBody) MarshalJSON() ([]byte, error) {
	if u.PaypalPaymentViewSchemas != nil {
		return utils.MarshalJSON(u.PaypalPaymentViewSchemas, "", true)
	}

	if u.SavedPaymentViewSchemas != nil {
		return utils.MarshalJSON(u.SavedPaymentViewSchemas, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type InitializePaymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Payment token retrieved.
	OneOf *InitializePaymentResponseBody
}

func (o *InitializePaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InitializePaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InitializePaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *InitializePaymentResponse) GetOneOf() *InitializePaymentResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}

// SavedPaymentInputUpdatePaymentData - Initialize payment for a saved payment method
type SavedPaymentInputUpdatePaymentData struct {
	// Payment ID of the saved Bolt Payment method.
	ID *string `json:"id,omitempty"`
}

func (o *SavedPaymentInputUpdatePaymentData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// SavedPaymentInputUpdateType - The type of the payment attempt
type SavedPaymentInputUpdateType string

const (
	SavedPaymentInputUpdateTypePaypal             SavedPaymentInputUpdateType = "paypal"
	SavedPaymentInputUpdateTypeSavedPaymentMethod SavedPaymentInputUpdateType = "saved_payment_method"
)

func (e SavedPaymentInputUpdateType) ToPointer() *SavedPaymentInputUpdateType {
	return &e
}

func (e *SavedPaymentInputUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paypal":
		fallthrough
	case "saved_payment_method":
		*e = SavedPaymentInputUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedPaymentInputUpdateType: %v", v)
	}
}

type SavedPaymentInputUpdatePaymentMethod struct {
	// Initialize payment for a saved payment method
	PaymentData *SavedPaymentInputUpdatePaymentData `json:"payment_data,omitempty"`
	// The type of the payment attempt
	Type *SavedPaymentInputUpdateType `json:"type,omitempty"`
}

func (o *SavedPaymentInputUpdatePaymentMethod) GetPaymentData() *SavedPaymentInputUpdatePaymentData {
	if o == nil {
		return nil
	}
	return o.PaymentData
}

func (o *SavedPaymentInputUpdatePaymentMethod) GetType() *SavedPaymentInputUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

// SavedPaymentInputUpdateShopperIdentity - Identification information for the Shopper. This is only required when creating a new Bolt account.
type SavedPaymentInputUpdateShopperIdentity struct {
	// determines whether to create a bolt account for this shopper
	CreateBoltAccount *bool `json:"create_bolt_account,omitempty"`
	// Email address of the shopper
	Email string `json:"email"`
	// First name of the shopper
	FirstName string `json:"first_name"`
	// Last name of the shopper
	LastName string `json:"last_name"`
	// Phone number of the shopper
	Phone string `json:"phone"`
}

func (o *SavedPaymentInputUpdateShopperIdentity) GetCreateBoltAccount() *bool {
	if o == nil {
		return nil
	}
	return o.CreateBoltAccount
}

func (o *SavedPaymentInputUpdateShopperIdentity) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SavedPaymentInputUpdateShopperIdentity) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *SavedPaymentInputUpdateShopperIdentity) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *SavedPaymentInputUpdateShopperIdentity) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

type SavedPaymentInputUpdateSchemas struct {
	// The details of the cart being purchased with this payment.
	Cart          *shared.CartCreate                    `json:"cart,omitempty"`
	PaymentMethod *SavedPaymentInputUpdatePaymentMethod `json:"payment_method,omitempty"`
	// Identification information for the Shopper. This is only required when creating a new Bolt account.
	ShopperIdentity *SavedPaymentInputUpdateShopperIdentity `json:"shopper_identity,omitempty"`
}

func (o *SavedPaymentInputUpdateSchemas) GetCart() *shared.CartCreate {
	if o == nil {
		return nil
	}
	return o.Cart
}

func (o *SavedPaymentInputUpdateSchemas) GetPaymentMethod() *SavedPaymentInputUpdatePaymentMethod {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *SavedPaymentInputUpdateSchemas) GetShopperIdentity() *SavedPaymentInputUpdateShopperIdentity {
	if o == nil {
		return nil
	}
	return o.ShopperIdentity
}

// PaypalPaymentInputUpdatePaymentData - Initialize payment for a new PayPal order.
type PaypalPaymentInputUpdatePaymentData struct {
	// Redirect URL for canceled PayPal transaction.
	Cancel *string `json:"cancel,omitempty"`
	// Redirect URL for successful PayPal transaction.
	Success *string `json:"success,omitempty"`
}

func (o *PaypalPaymentInputUpdatePaymentData) GetCancel() *string {
	if o == nil {
		return nil
	}
	return o.Cancel
}

func (o *PaypalPaymentInputUpdatePaymentData) GetSuccess() *string {
	if o == nil {
		return nil
	}
	return o.Success
}

// PaypalPaymentInputUpdateType - The type of the payment attempt
type PaypalPaymentInputUpdateType string

const (
	PaypalPaymentInputUpdateTypePaypal             PaypalPaymentInputUpdateType = "paypal"
	PaypalPaymentInputUpdateTypeSavedPaymentMethod PaypalPaymentInputUpdateType = "saved_payment_method"
)

func (e PaypalPaymentInputUpdateType) ToPointer() *PaypalPaymentInputUpdateType {
	return &e
}

func (e *PaypalPaymentInputUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paypal":
		fallthrough
	case "saved_payment_method":
		*e = PaypalPaymentInputUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaypalPaymentInputUpdateType: %v", v)
	}
}

type PaypalPaymentInputUpdatePaymentMethod struct {
	// Initialize payment for a new PayPal order.
	PaymentData *PaypalPaymentInputUpdatePaymentData `json:"payment_data,omitempty"`
	// The type of the payment attempt
	Type *PaypalPaymentInputUpdateType `json:"type,omitempty"`
}

func (o *PaypalPaymentInputUpdatePaymentMethod) GetPaymentData() *PaypalPaymentInputUpdatePaymentData {
	if o == nil {
		return nil
	}
	return o.PaymentData
}

func (o *PaypalPaymentInputUpdatePaymentMethod) GetType() *PaypalPaymentInputUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

// PaypalPaymentInputUpdateShopperIdentity - Identification information for the Shopper. This is only required when creating a new Bolt account.
type PaypalPaymentInputUpdateShopperIdentity struct {
	// determines whether to create a bolt account for this shopper
	CreateBoltAccount *bool `json:"create_bolt_account,omitempty"`
	// Email address of the shopper
	Email string `json:"email"`
	// First name of the shopper
	FirstName string `json:"first_name"`
	// Last name of the shopper
	LastName string `json:"last_name"`
	// Phone number of the shopper
	Phone string `json:"phone"`
}

func (o *PaypalPaymentInputUpdateShopperIdentity) GetCreateBoltAccount() *bool {
	if o == nil {
		return nil
	}
	return o.CreateBoltAccount
}

func (o *PaypalPaymentInputUpdateShopperIdentity) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *PaypalPaymentInputUpdateShopperIdentity) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *PaypalPaymentInputUpdateShopperIdentity) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *PaypalPaymentInputUpdateShopperIdentity) GetPhone() string {
	if o == nil {
		return ""
	}
	return o.Phone
}

type PaypalPaymentInputUpdateSchemas struct {
	// The details of the cart being purchased with this payment.
	Cart          *shared.CartCreate                     `json:"cart,omitempty"`
	PaymentMethod *PaypalPaymentInputUpdatePaymentMethod `json:"payment_method,omitempty"`
	// Identification information for the Shopper. This is only required when creating a new Bolt account.
	ShopperIdentity *PaypalPaymentInputUpdateShopperIdentity `json:"shopper_identity,omitempty"`
}

func (o *PaypalPaymentInputUpdateSchemas) GetCart() *shared.CartCreate {
	if o == nil {
		return nil
	}
	return o.Cart
}

func (o *PaypalPaymentInputUpdateSchemas) GetPaymentMethod() *PaypalPaymentInputUpdatePaymentMethod {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *PaypalPaymentInputUpdateSchemas) GetShopperIdentity() *PaypalPaymentInputUpdateShopperIdentity {
	if o == nil {
		return nil
	}
	return o.ShopperIdentity
}

type UpdatePaymentRequestBodyType string

const (
	UpdatePaymentRequestBodyTypePaypalPaymentInputUpdateSchemas UpdatePaymentRequestBodyType = "paypal_payment_input_update_Schemas"
	UpdatePaymentRequestBodyTypeSavedPaymentInputUpdateSchemas  UpdatePaymentRequestBodyType = "saved_payment_input_update_Schemas"
)

type UpdatePaymentRequestBody struct {
	PaypalPaymentInputUpdateSchemas *PaypalPaymentInputUpdateSchemas
	SavedPaymentInputUpdateSchemas  *SavedPaymentInputUpdateSchemas

	Type UpdatePaymentRequestBodyType
}

func CreateUpdatePaymentRequestBodyPaypalPaymentInputUpdateSchemas(paypalPaymentInputUpdateSchemas PaypalPaymentInputUpdateSchemas) UpdatePaymentRequestBody {
	typ := UpdatePaymentRequestBodyTypePaypalPaymentInputUpdateSchemas

	return UpdatePaymentRequestBody{
		PaypalPaymentInputUpdateSchemas: &paypalPaymentInputUpdateSchemas,
		Type:                            typ,
	}
}

func CreateUpdatePaymentRequestBodySavedPaymentInputUpdateSchemas(savedPaymentInputUpdateSchemas SavedPaymentInputUpdateSchemas) UpdatePaymentRequestBody {
	typ := UpdatePaymentRequestBodyTypeSavedPaymentInputUpdateSchemas

	return UpdatePaymentRequestBody{
		SavedPaymentInputUpdateSchemas: &savedPaymentInputUpdateSchemas,
		Type:                           typ,
	}
}

func (u *UpdatePaymentRequestBody) UnmarshalJSON(data []byte) error {

	paypalPaymentInputUpdateSchemas := PaypalPaymentInputUpdateSchemas{}
	if err := utils.UnmarshalJSON(data, &paypalPaymentInputUpdateSchemas, "", true, true); err == nil {
		u.PaypalPaymentInputUpdateSchemas = &paypalPaymentInputUpdateSchemas
		u.Type = UpdatePaymentRequestBodyTypePaypalPaymentInputUpdateSchemas
		return nil
	}

	savedPaymentInputUpdateSchemas := SavedPaymentInputUpdateSchemas{}
	if err := utils.UnmarshalJSON(data, &savedPaymentInputUpdateSchemas, "", true, true); err == nil {
		u.SavedPaymentInputUpdateSchemas = &savedPaymentInputUpdateSchemas
		u.Type = UpdatePaymentRequestBodyTypeSavedPaymentInputUpdateSchemas
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdatePaymentRequestBody) MarshalJSON() ([]byte, error) {
	if u.PaypalPaymentInputUpdateSchemas != nil {
		return utils.MarshalJSON(u.PaypalPaymentInputUpdateSchemas, "", true)
	}

	if u.SavedPaymentInputUpdateSchemas != nil {
		return utils.MarshalJSON(u.SavedPaymentInputUpdateSchemas, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdatePaymentRequest struct {
	// A key created by merchants that ensures `POST` and `PATCH` requests are only performed once. [Read more about Idempotent Requests here](/developers/references/idempotency/).
	IdempotencyKey *string                   `header:"style=simple,explode=false,name=Idempotency-Key"`
	RequestBody    *UpdatePaymentRequestBody `request:"mediaType=application/json"`
	// The publicly viewable identifier used to identify a merchant division. This key is found in the Developer > API section of the Bolt Merchant Dashboard [RECOMMENDED].
	XPublishableKey *string `header:"style=simple,explode=false,name=X-Publishable-Key"`
	// The ID received in the initial v1/payments request.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdatePaymentRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *UpdatePaymentRequest) GetRequestBody() *UpdatePaymentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpdatePaymentRequest) GetXPublishableKey() *string {
	if o == nil {
		return nil
	}
	return o.XPublishableKey
}

func (o *UpdatePaymentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdatePaymentResponseBodyType string

const (
	UpdatePaymentResponseBodyTypePaypalPaymentViewSchemas UpdatePaymentResponseBodyType = "paypal_payment_view_Schemas"
	UpdatePaymentResponseBodyTypeSavedPaymentViewSchemas  UpdatePaymentResponseBodyType = "saved_payment_view_Schemas"
)

// UpdatePaymentResponseBody - Payment updated.
type UpdatePaymentResponseBody struct {
	PaypalPaymentViewSchemas *PaypalPaymentViewSchemas
	SavedPaymentViewSchemas  *SavedPaymentViewSchemas

	Type UpdatePaymentResponseBodyType
}

func CreateUpdatePaymentResponseBodyPaypalPaymentViewSchemas(paypalPaymentViewSchemas PaypalPaymentViewSchemas) UpdatePaymentResponseBody {
	typ := UpdatePaymentResponseBodyTypePaypalPaymentViewSchemas

	return UpdatePaymentResponseBody{
		PaypalPaymentViewSchemas: &paypalPaymentViewSchemas,
		Type:                     typ,
	}
}

func CreateUpdatePaymentResponseBodySavedPaymentViewSchemas(savedPaymentViewSchemas SavedPaymentViewSchemas) UpdatePaymentResponseBody {
	typ := UpdatePaymentResponseBodyTypeSavedPaymentViewSchemas

	return UpdatePaymentResponseBody{
		SavedPaymentViewSchemas: &savedPaymentViewSchemas,
		Type:                    typ,
	}
}

func (u *UpdatePaymentResponseBody) UnmarshalJSON(data []byte) error {

	paypalPaymentViewSchemas := PaypalPaymentViewSchemas{}
	if err := utils.UnmarshalJSON(data, &paypalPaymentViewSchemas, "", true, true); err == nil {
		u.PaypalPaymentViewSchemas = &paypalPaymentViewSchemas
		u.Type = UpdatePaymentResponseBodyTypePaypalPaymentViewSchemas
		return nil
	}

	savedPaymentViewSchemas := SavedPaymentViewSchemas{}
	if err := utils.UnmarshalJSON(data, &savedPaymentViewSchemas, "", true, true); err == nil {
		u.SavedPaymentViewSchemas = &savedPaymentViewSchemas
		u.Type = UpdatePaymentResponseBodyTypeSavedPaymentViewSchemas
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdatePaymentResponseBody) MarshalJSON() ([]byte, error) {
	if u.PaypalPaymentViewSchemas != nil {
		return utils.MarshalJSON(u.PaypalPaymentViewSchemas, "", true)
	}

	if u.SavedPaymentViewSchemas != nil {
		return utils.MarshalJSON(u.SavedPaymentViewSchemas, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdatePaymentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Payment updated.
	OneOf *UpdatePaymentResponseBody
}

func (o *UpdatePaymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePaymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePaymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePaymentResponse) GetOneOf() *UpdatePaymentResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
