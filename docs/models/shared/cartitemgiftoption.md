# CartItemGiftOption

Contains the gift option settings for wrapping and custom messages.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       | Example                                           |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Cost`                                            | **int64*                                          | :heavy_minus_sign:                                | The cost in cents.                                | 770                                               |
| `MerchantProductID`                               | **string*                                         | :heavy_minus_sign:                                | The merchant's unique ID for the product.         | 881                                               |
| `Message`                                         | **string*                                         | :heavy_minus_sign:                                | Includes the gift message written by the shopper. | Happy Anniversary, Smoochy Poo!                   |
| `Wrap`                                            | **bool*                                           | :heavy_minus_sign:                                | Defines whether gift wrapping was requested.      | false                                             |