# CreateAccountInput

The details needed to create a Bolt account.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Addresses`                                                                         | [][shared.AddressAccount](../../../pkg/models/shared/addressaccount.md)             | :heavy_minus_sign:                                                                  | A list of physical shipping addresses associated with this account.                 |
| `PaymentMethods`                                                                    | [][shared.PaymentMethodAccount](../../../pkg/models/shared/paymentmethodaccount.md) | :heavy_minus_sign:                                                                  | A list of payment methods associated with this account.                             |
| `Profile`                                                                           | [shared.Profile](../../../pkg/models/shared/profile.md)                             | :heavy_check_mark:                                                                  | The first name, last name, email address, and phone number of a shopper.            |