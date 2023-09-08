# AuthorizeTransactionRequestBody

**Authorize a Transaction**
* • `merchant_credit_card_authorization`: For authorizing with a new, unsaved card. This can be for a guest checkout flow, one-time payment, or an existing Bolt shopper.
* • `merchant_credit_card_authorization_recharge`: For authorizing a card using a shoppers saved payment methods.
* • **Anytime the shopper is paying while logged-in attach their OAuth `access_token` to the request.**



## Supported Types

### MerchantCreditCardAuthorizationRecharge

```go
authorizeTransactionRequestBody := operations.CreateAuthorizeTransactionRequestBodyMerchantCreditCardAuthorizationRecharge(shared.MerchantCreditCardAuthorizationRecharge{/* values here */})
```

### MerchantCreditCardAuthorization

```go
authorizeTransactionRequestBody := operations.CreateAuthorizeTransactionRequestBodyMerchantCreditCardAuthorization(shared.MerchantCreditCardAuthorization{/* values here */})
```

