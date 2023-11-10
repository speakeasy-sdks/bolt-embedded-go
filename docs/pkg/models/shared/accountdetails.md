# AccountDetails


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Addresses`                                                                          | [][shared.Addresses](../../../pkg/models/shared/addresses.md)                        | :heavy_minus_sign:                                                                   | A list of all addresses associated to the shopper's account.                         |
| `HasBoltAccount`                                                                     | **bool*                                                                              | :heavy_minus_sign:                                                                   | Used to determine whether a Bolt Account exists with this shopper's account details. |
| `PaymentMethods`                                                                     | [][shared.PaymentMethods](../../../pkg/models/shared/paymentmethods.md)              | :heavy_minus_sign:                                                                   | A list of all payment methods associated to the shopper's account.                   |
| `Profile`                                                                            | [*shared.ProfileView](../../../pkg/models/shared/profileview.md)                     | :heavy_minus_sign:                                                                   | The shopper's account profile.                                                       |