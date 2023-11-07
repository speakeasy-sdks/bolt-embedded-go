# TransactionViewCause

Specifies why this particular transaction is voided.


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `TransactionViewCauseMerchantAction`          | merchant_action                               |
| `TransactionViewCausePaypalSync`              | paypal_sync                                   |
| `TransactionViewCauseAmazonPaySync`           | amazon_pay_sync                               |
| `TransactionViewCauseIrreversibleReject`      | irreversible_reject                           |
| `TransactionViewCauseAuthExpire`              | auth_expire                                   |
| `TransactionViewCauseAuthVerificationExpired` | auth_verification_expired                     |
| `TransactionViewCausePaymentMethodUpdater`    | payment_method_updater                        |
| `TransactionViewCauseLessThanNilGreaterThan`  | <nil>                                         |