# TransactionViewVoidCause

Specifies why this particular transaction is voided.


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `TransactionViewVoidCauseMerchantAction`          | merchant_action                                   |
| `TransactionViewVoidCausePaypalSync`              | paypal_sync                                       |
| `TransactionViewVoidCauseAmazonPaySync`           | amazon_pay_sync                                   |
| `TransactionViewVoidCauseIrreversibleReject`      | irreversible_reject                               |
| `TransactionViewVoidCauseAuthExpire`              | auth_expire                                       |
| `TransactionViewVoidCauseAuthVerificationExpired` | auth_verification_expired                         |
| `TransactionViewVoidCausePaymentMethodUpdater`    | payment_method_updater                            |
| `TransactionViewVoidCauseLessThanNilGreaterThan`  | <nil>                                             |