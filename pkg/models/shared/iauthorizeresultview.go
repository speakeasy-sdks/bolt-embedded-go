// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IAuthorizeResultView struct {
	DidCreateBoltAccount *bool            `json:"did_create_bolt_account,omitempty"`
	OrderNumber          *string          `json:"order_number,omitempty"`
	Transaction          *TransactionView `json:"transaction,omitempty"`
}

func (o *IAuthorizeResultView) GetDidCreateBoltAccount() *bool {
	if o == nil {
		return nil
	}
	return o.DidCreateBoltAccount
}

func (o *IAuthorizeResultView) GetOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.OrderNumber
}

func (o *IAuthorizeResultView) GetTransaction() *TransactionView {
	if o == nil {
		return nil
	}
	return o.Transaction
}
