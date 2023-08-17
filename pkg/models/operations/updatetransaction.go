// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"net/http"
)

type UpdateTransactionSecurity struct {
	XAPIKey string `security:"scheme,type=apiKey,subtype=header,name=X-API-Key"`
}

func (o *UpdateTransactionSecurity) GetXAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XAPIKey
}

// UpdateTransactionRequestBody - Update a Transaction
type UpdateTransactionRequestBody struct {
	// This field corresponds to the merchant's order reference associated with this Bolt transaction.
	DisplayID *string `json:"display_id,omitempty"`
	// Custom metadata associated with this Bolt transaction.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (o *UpdateTransactionRequestBody) GetDisplayID() *string {
	if o == nil {
		return nil
	}
	return o.DisplayID
}

func (o *UpdateTransactionRequestBody) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

type UpdateTransactionRequest struct {
	// A key created by merchants that ensures `POST` and `PATCH` requests are only performed once. [Read more about Idempotent Requests here](/developers/references/idempotency/).
	IdempotencyKey *string `header:"style=simple,explode=false,name=Idempotency-Key"`
	// This is the Bolt transaction reference. (ex. N7Y3-NFKC-VFRF)
	Reference   string                        `pathParam:"style=simple,explode=false,name=REFERENCE"`
	RequestBody *UpdateTransactionRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateTransactionRequest) GetIdempotencyKey() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyKey
}

func (o *UpdateTransactionRequest) GetReference() string {
	if o == nil {
		return ""
	}
	return o.Reference
}

func (o *UpdateTransactionRequest) GetRequestBody() *UpdateTransactionRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

// UpdateTransaction200ApplicationJSONSplitsType - **Nullable** for Transactions Details.
type UpdateTransaction200ApplicationJSONSplitsType string

const (
	UpdateTransaction200ApplicationJSONSplitsTypeNet           UpdateTransaction200ApplicationJSONSplitsType = "net"
	UpdateTransaction200ApplicationJSONSplitsTypeProcessingFee UpdateTransaction200ApplicationJSONSplitsType = "processing_fee"
	UpdateTransaction200ApplicationJSONSplitsTypeBoltFee       UpdateTransaction200ApplicationJSONSplitsType = "bolt_fee"
	UpdateTransaction200ApplicationJSONSplitsTypeAdjustment    UpdateTransaction200ApplicationJSONSplitsType = "adjustment"
	UpdateTransaction200ApplicationJSONSplitsTypeFloat         UpdateTransaction200ApplicationJSONSplitsType = "float"
	UpdateTransaction200ApplicationJSONSplitsTypeReserve       UpdateTransaction200ApplicationJSONSplitsType = "reserve"
)

func (e UpdateTransaction200ApplicationJSONSplitsType) ToPointer() *UpdateTransaction200ApplicationJSONSplitsType {
	return &e
}

func (e *UpdateTransaction200ApplicationJSONSplitsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "net":
		fallthrough
	case "processing_fee":
		fallthrough
	case "bolt_fee":
		fallthrough
	case "adjustment":
		fallthrough
	case "float":
		fallthrough
	case "reserve":
		*e = UpdateTransaction200ApplicationJSONSplitsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateTransaction200ApplicationJSONSplitsType: %v", v)
	}
}

// UpdateTransaction200ApplicationJSONSplits - A split of fees by type and amount.
type UpdateTransaction200ApplicationJSONSplits struct {
	Amount *shared.AmountView `json:"amount,omitempty"`
	// **Nullable** for Transactions Details.
	//
	Type *UpdateTransaction200ApplicationJSONSplitsType `json:"type,omitempty"`
}

func (o *UpdateTransaction200ApplicationJSONSplits) GetAmount() *shared.AmountView {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *UpdateTransaction200ApplicationJSONSplits) GetType() *UpdateTransaction200ApplicationJSONSplitsType {
	if o == nil {
		return nil
	}
	return o.Type
}

type UpdateTransaction200ApplicationJSONTransactionRejectionDetailsAuthRejectionDetails struct {
	ReasonCode        *string `json:"reason_code,omitempty"`
	ReasonDescription *string `json:"reason_description,omitempty"`
}

func (o *UpdateTransaction200ApplicationJSONTransactionRejectionDetailsAuthRejectionDetails) GetReasonCode() *string {
	if o == nil {
		return nil
	}
	return o.ReasonCode
}

func (o *UpdateTransaction200ApplicationJSONTransactionRejectionDetailsAuthRejectionDetails) GetReasonDescription() *string {
	if o == nil {
		return nil
	}
	return o.ReasonDescription
}

type UpdateTransaction200ApplicationJSONTransactionRejectionDetails struct {
	AuthRejectionDetails *UpdateTransaction200ApplicationJSONTransactionRejectionDetailsAuthRejectionDetails `json:"auth_rejection_details,omitempty"`
}

func (o *UpdateTransaction200ApplicationJSONTransactionRejectionDetails) GetAuthRejectionDetails() *UpdateTransaction200ApplicationJSONTransactionRejectionDetailsAuthRejectionDetails {
	if o == nil {
		return nil
	}
	return o.AuthRejectionDetails
}

// UpdateTransaction200ApplicationJSON - Transaction Details Retrieved
type UpdateTransaction200ApplicationJSON struct {
	AddressChangeRequestMetadata *shared.AddressChangeRequestMetadataView `json:"address_change_request_metadata,omitempty"`
	// **Nullable** for Transactions Details.
	//
	AdjustTransactions []shared.TransactionView `json:"adjust_transactions,omitempty"`
	Amount             *shared.AmountView       `json:"amount,omitempty"`
	// Used to track the status of micro-authorizations. **Nullable** for Transactions Details.
	AuthVerificationStatus *shared.AuthorizationVerificationStatus `json:"auth_verification_status,omitempty"`
	Authorization          *shared.CreditCardAuthorizationView     `json:"authorization,omitempty"`
	// The authorization's id.
	AuthorizationID *string `json:"authorization_id,omitempty"`
	// Deprecated. Use `captures`.
	Capture *shared.Capture `json:"capture,omitempty"`
	// All captures associated with the transaction. **Nullable** for Transactions Details.
	Captures           []shared.Capture           `json:"captures,omitempty"`
	ChargebackDetails  *shared.ChargebackDetails  `json:"chargeback_details,omitempty"`
	Credit             *shared.Credit             `json:"credit,omitempty"`
	CustomFields       *shared.CustomFields       `json:"custom_fields,omitempty"`
	CustomerListStatus *shared.CustomerListStatus `json:"customer_list_status,omitempty"`
	// Transaction date. **Nullable** for Transactions Details.
	Date *int64 `json:"date,omitempty"`
	// The credit card user.
	FromConsumer *shared.CreditCardUser `json:"from_consumer,omitempty"`
	// **Nullable** for Transactions Details.
	//
	FromConsumerMembershipUsers *shared.ConsumerUserMembership `json:"from_consumer_membership_users,omitempty"`
	// Contains details about the credit card transaction.
	FromCreditCard *shared.CreditCardView `json:"from_credit_card,omitempty"`
	// The unique ID associated with the transaction. **Nullable** for Transactions Details.
	ID *string `json:"id,omitempty"`
	// Describes whether the transaction is indemnified by Bolt for fraud.
	//
	IndemnificationDecision *shared.TransactionIndemnificationDecision `json:"indemnification_decision,omitempty"`
	// Describes the reason that the transaction is or is not indemnified by Bolt for fraud.
	//
	IndemnificationReason *shared.TransactionIndemnificationReason `json:"indemnification_reason,omitempty"`
	// The card's last 4 digits. **Nullable** for Transactions Details.
	Last4 *string `json:"last4,omitempty"`
	// The last view time as UTC.
	LastViewedUtc    *int64                   `json:"last_viewed_utc,omitempty"`
	ManualDisputes   *shared.ManualDisputes   `json:"manual_disputes,omitempty"`
	Merchant         *shared.Merchant         `json:"merchant,omitempty"`
	MerchantDivision *shared.MerchantDivision `json:"merchant_division,omitempty"`
	// The merchant's internal order number for this transaction.
	MerchantOrderNumber *string           `json:"merchant_order_number,omitempty"`
	Order               *shared.OrderView `json:"order,omitempty"`
	// Decision and score for an order.
	OrderDecision    *shared.OrderDecision `json:"order_decision,omitempty"`
	PlatformMetadata map[string]string     `json:"platform_metadata,omitempty"`
	// The processor used. **Nullable** for Transactions Details.
	Processor *shared.TransactionProcessor `json:"processor,omitempty"`
	// The transaction's 12-digit Bolt reference ID. **Nullable** for Transactions Details.
	Reference *string `json:"reference,omitempty"`
	// **Nullable** for Transactions Details.
	//
	RefundTransactionIds []string `json:"refund_transaction_ids,omitempty"`
	// **Nullable** for Transactions Details.
	//
	RefundTransactions []shared.TransactionView `json:"refund_transactions,omitempty"`
	RefundedAmount     *shared.AmountView       `json:"refunded_amount,omitempty"`
	// Internal use only.
	ReviewTicket *shared.ReviewTicket `json:"review_ticket,omitempty"`
	// Displays fraud decisioning insights based on key factors. This information can either be forwarded via a `risk_insights` transaction webhook type or be polled by sending a `GET` request to Bolt's [transactions endpoint](/api-bolt/#operation/transaction-details).
	//
	RiskInsights *shared.RiskInsightsYml `json:"risk_insights,omitempty"`
	// Describes the current Risk Review status. A transaction could be unreviewed, reviewed, or pending manual review by the Bolt team.
	RiskReviewStatus  *shared.RiskReviewStatus `json:"risk_review_status,omitempty"`
	RiskScore         *int64                   `json:"risk_score,omitempty"`
	SourceTransaction *shared.TransactionView  `json:"source_transaction,omitempty"`
	// A list of splits. **Nullable** for Transactions Details.
	Splits []UpdateTransaction200ApplicationJSONSplits `json:"splits,omitempty"`
	// The transaction's status.
	Status     *shared.TransactionStatus       `json:"status,omitempty"`
	Timeline   *shared.TransactionTimelineView `json:"timeline,omitempty"`
	ToConsumer *shared.ConsumerSelfView        `json:"to_consumer,omitempty"`
	// Contains details about the credit card transaction.
	ToCreditCard                *shared.CreditCardView                                          `json:"to_credit_card,omitempty"`
	TransactionProperties       map[string]string                                               `json:"transaction_properties,omitempty"`
	TransactionRejectionDetails *UpdateTransaction200ApplicationJSONTransactionRejectionDetails `json:"transaction_rejection_details,omitempty"`
	// The type of transaction.
	Type       *shared.TransactionType `json:"type,omitempty"`
	ViewStatus *string                 `json:"view_status,omitempty"`
	Void       *shared.Void            `json:"void,omitempty"`
	VoidCause  *string                 `json:"void_cause,omitempty"`
}

func (o *UpdateTransaction200ApplicationJSON) GetAddressChangeRequestMetadata() *shared.AddressChangeRequestMetadataView {
	if o == nil {
		return nil
	}
	return o.AddressChangeRequestMetadata
}

func (o *UpdateTransaction200ApplicationJSON) GetAdjustTransactions() []shared.TransactionView {
	if o == nil {
		return nil
	}
	return o.AdjustTransactions
}

func (o *UpdateTransaction200ApplicationJSON) GetAmount() *shared.AmountView {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *UpdateTransaction200ApplicationJSON) GetAuthVerificationStatus() *shared.AuthorizationVerificationStatus {
	if o == nil {
		return nil
	}
	return o.AuthVerificationStatus
}

func (o *UpdateTransaction200ApplicationJSON) GetAuthorization() *shared.CreditCardAuthorizationView {
	if o == nil {
		return nil
	}
	return o.Authorization
}

func (o *UpdateTransaction200ApplicationJSON) GetAuthorizationID() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationID
}

func (o *UpdateTransaction200ApplicationJSON) GetCapture() *shared.Capture {
	if o == nil {
		return nil
	}
	return o.Capture
}

func (o *UpdateTransaction200ApplicationJSON) GetCaptures() []shared.Capture {
	if o == nil {
		return nil
	}
	return o.Captures
}

func (o *UpdateTransaction200ApplicationJSON) GetChargebackDetails() *shared.ChargebackDetails {
	if o == nil {
		return nil
	}
	return o.ChargebackDetails
}

func (o *UpdateTransaction200ApplicationJSON) GetCredit() *shared.Credit {
	if o == nil {
		return nil
	}
	return o.Credit
}

func (o *UpdateTransaction200ApplicationJSON) GetCustomFields() *shared.CustomFields {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *UpdateTransaction200ApplicationJSON) GetCustomerListStatus() *shared.CustomerListStatus {
	if o == nil {
		return nil
	}
	return o.CustomerListStatus
}

func (o *UpdateTransaction200ApplicationJSON) GetDate() *int64 {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *UpdateTransaction200ApplicationJSON) GetFromConsumer() *shared.CreditCardUser {
	if o == nil {
		return nil
	}
	return o.FromConsumer
}

func (o *UpdateTransaction200ApplicationJSON) GetFromConsumerMembershipUsers() *shared.ConsumerUserMembership {
	if o == nil {
		return nil
	}
	return o.FromConsumerMembershipUsers
}

func (o *UpdateTransaction200ApplicationJSON) GetFromCreditCard() *shared.CreditCardView {
	if o == nil {
		return nil
	}
	return o.FromCreditCard
}

func (o *UpdateTransaction200ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateTransaction200ApplicationJSON) GetIndemnificationDecision() *shared.TransactionIndemnificationDecision {
	if o == nil {
		return nil
	}
	return o.IndemnificationDecision
}

func (o *UpdateTransaction200ApplicationJSON) GetIndemnificationReason() *shared.TransactionIndemnificationReason {
	if o == nil {
		return nil
	}
	return o.IndemnificationReason
}

func (o *UpdateTransaction200ApplicationJSON) GetLast4() *string {
	if o == nil {
		return nil
	}
	return o.Last4
}

func (o *UpdateTransaction200ApplicationJSON) GetLastViewedUtc() *int64 {
	if o == nil {
		return nil
	}
	return o.LastViewedUtc
}

func (o *UpdateTransaction200ApplicationJSON) GetManualDisputes() *shared.ManualDisputes {
	if o == nil {
		return nil
	}
	return o.ManualDisputes
}

func (o *UpdateTransaction200ApplicationJSON) GetMerchant() *shared.Merchant {
	if o == nil {
		return nil
	}
	return o.Merchant
}

func (o *UpdateTransaction200ApplicationJSON) GetMerchantDivision() *shared.MerchantDivision {
	if o == nil {
		return nil
	}
	return o.MerchantDivision
}

func (o *UpdateTransaction200ApplicationJSON) GetMerchantOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.MerchantOrderNumber
}

func (o *UpdateTransaction200ApplicationJSON) GetOrder() *shared.OrderView {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *UpdateTransaction200ApplicationJSON) GetOrderDecision() *shared.OrderDecision {
	if o == nil {
		return nil
	}
	return o.OrderDecision
}

func (o *UpdateTransaction200ApplicationJSON) GetPlatformMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.PlatformMetadata
}

func (o *UpdateTransaction200ApplicationJSON) GetProcessor() *shared.TransactionProcessor {
	if o == nil {
		return nil
	}
	return o.Processor
}

func (o *UpdateTransaction200ApplicationJSON) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *UpdateTransaction200ApplicationJSON) GetRefundTransactionIds() []string {
	if o == nil {
		return nil
	}
	return o.RefundTransactionIds
}

func (o *UpdateTransaction200ApplicationJSON) GetRefundTransactions() []shared.TransactionView {
	if o == nil {
		return nil
	}
	return o.RefundTransactions
}

func (o *UpdateTransaction200ApplicationJSON) GetRefundedAmount() *shared.AmountView {
	if o == nil {
		return nil
	}
	return o.RefundedAmount
}

func (o *UpdateTransaction200ApplicationJSON) GetReviewTicket() *shared.ReviewTicket {
	if o == nil {
		return nil
	}
	return o.ReviewTicket
}

func (o *UpdateTransaction200ApplicationJSON) GetRiskInsights() *shared.RiskInsightsYml {
	if o == nil {
		return nil
	}
	return o.RiskInsights
}

func (o *UpdateTransaction200ApplicationJSON) GetRiskReviewStatus() *shared.RiskReviewStatus {
	if o == nil {
		return nil
	}
	return o.RiskReviewStatus
}

func (o *UpdateTransaction200ApplicationJSON) GetRiskScore() *int64 {
	if o == nil {
		return nil
	}
	return o.RiskScore
}

func (o *UpdateTransaction200ApplicationJSON) GetSourceTransaction() *shared.TransactionView {
	if o == nil {
		return nil
	}
	return o.SourceTransaction
}

func (o *UpdateTransaction200ApplicationJSON) GetSplits() []UpdateTransaction200ApplicationJSONSplits {
	if o == nil {
		return nil
	}
	return o.Splits
}

func (o *UpdateTransaction200ApplicationJSON) GetStatus() *shared.TransactionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UpdateTransaction200ApplicationJSON) GetTimeline() *shared.TransactionTimelineView {
	if o == nil {
		return nil
	}
	return o.Timeline
}

func (o *UpdateTransaction200ApplicationJSON) GetToConsumer() *shared.ConsumerSelfView {
	if o == nil {
		return nil
	}
	return o.ToConsumer
}

func (o *UpdateTransaction200ApplicationJSON) GetToCreditCard() *shared.CreditCardView {
	if o == nil {
		return nil
	}
	return o.ToCreditCard
}

func (o *UpdateTransaction200ApplicationJSON) GetTransactionProperties() map[string]string {
	if o == nil {
		return nil
	}
	return o.TransactionProperties
}

func (o *UpdateTransaction200ApplicationJSON) GetTransactionRejectionDetails() *UpdateTransaction200ApplicationJSONTransactionRejectionDetails {
	if o == nil {
		return nil
	}
	return o.TransactionRejectionDetails
}

func (o *UpdateTransaction200ApplicationJSON) GetType() *shared.TransactionType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateTransaction200ApplicationJSON) GetViewStatus() *string {
	if o == nil {
		return nil
	}
	return o.ViewStatus
}

func (o *UpdateTransaction200ApplicationJSON) GetVoid() *shared.Void {
	if o == nil {
		return nil
	}
	return o.Void
}

func (o *UpdateTransaction200ApplicationJSON) GetVoidCause() *string {
	if o == nil {
		return nil
	}
	return o.VoidCause
}

type UpdateTransactionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Generic Error Schema
	ErrorsBoltAPIResponse *shared.ErrorsBoltAPIResponse
	// Transaction Details Retrieved
	//
	UpdateTransaction200ApplicationJSONObject *UpdateTransaction200ApplicationJSON
}

func (o *UpdateTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTransactionResponse) GetErrorsBoltAPIResponse() *shared.ErrorsBoltAPIResponse {
	if o == nil {
		return nil
	}
	return o.ErrorsBoltAPIResponse
}

func (o *UpdateTransactionResponse) GetUpdateTransaction200ApplicationJSONObject() *UpdateTransaction200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateTransaction200ApplicationJSONObject
}
