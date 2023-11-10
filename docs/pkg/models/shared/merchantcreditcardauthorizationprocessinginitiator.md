# MerchantCreditCardAuthorizationProcessingInitiator

Determines who initiated the transaction (e.g. shopper, merchant) and how they did it (e.g. recurring subscription, on-file card).

* `initial_card_on_file` - The first transaction made for a card. The system then saves this card for future transactions.
* `initial_recurring` - The first time any card is used to pay for a recurring charge. For example, a subscription.
* `stored_cardholder_initiated` - The subsequent (second, third, etc.) transactions a shopper initiates with a stored card. This includes every situation during which a cardholder requests a charge, for example if the cardholder requests a merchant charge their card.
* `stored_merchant_initiated` - The subsequent (second, third, etc.) transactions a merchant initiates with a stored card only when the cardholder does not request the charge. For example, when a customer service representative buys on behalf of a shopper or when a business adds funds to a public transit card.
* `following_recurring` - The subsequent (second, third, etc.) transactions  a card is used to pay for a recurring charge. For example, a subscription.
* `cardholder_initiated` - When a cardholder begins a transaction that isn’t stored in Bolt and won’t be stored in Bolt for future transactions.
* `recurring` - Any time a card is used to pay for a recurring charge (for example, a subscription). Only use this value when you don’t know if it’s the first recurring charge.



## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `MerchantCreditCardAuthorizationProcessingInitiatorInitialCardOnFile`         | initial_card_on_file                                                          |
| `MerchantCreditCardAuthorizationProcessingInitiatorInitialRecurring`          | initial_recurring                                                             |
| `MerchantCreditCardAuthorizationProcessingInitiatorStoredCardholderInitiated` | stored_cardholder_initiated                                                   |
| `MerchantCreditCardAuthorizationProcessingInitiatorStoredMerchantInitiated`   | stored_merchant_initiated                                                     |
| `MerchantCreditCardAuthorizationProcessingInitiatorFollowingRecurring`        | following_recurring                                                           |
| `MerchantCreditCardAuthorizationProcessingInitiatorCardholderInitiated`       | cardholder_initiated                                                          |
| `MerchantCreditCardAuthorizationProcessingInitiatorRecurring`                 | recurring                                                                     |