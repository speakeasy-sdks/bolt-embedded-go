// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ICurrency struct {
	Currency       *string `json:"currency,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty"`
}

func (o *ICurrency) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *ICurrency) GetCurrencySymbol() *string {
	if o == nil {
		return nil
	}
	return o.CurrencySymbol
}