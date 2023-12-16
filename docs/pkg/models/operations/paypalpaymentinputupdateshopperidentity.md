# PaypalPaymentInputUpdateShopperIdentity

Identification information for the Shopper. This is only required when creating a new Bolt account.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `CreateBoltAccount`                                          | **bool*                                                      | :heavy_minus_sign:                                           | determines whether to create a bolt account for this shopper | true                                                         |
| `Email`                                                      | *string*                                                     | :heavy_check_mark:                                           | Email address of the shopper                                 |                                                              |
| `FirstName`                                                  | *string*                                                     | :heavy_check_mark:                                           | First name of the shopper                                    |                                                              |
| `LastName`                                                   | *string*                                                     | :heavy_check_mark:                                           | Last name of the shopper                                     |                                                              |
| `Phone`                                                      | *string*                                                     | :heavy_check_mark:                                           | Phone number of the shopper                                  |                                                              |