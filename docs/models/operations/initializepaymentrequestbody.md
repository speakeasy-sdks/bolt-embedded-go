# InitializePaymentRequestBody


## Fields

| Field                                                                                                                  | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `Cart`                                                                                                                 | [shared.CartCreate](../../models/shared/cartcreate.md)                                                                 | :heavy_check_mark:                                                                                                     | The details of the cart being purchased with this payment.                                                             |
| `ShopperIdentity`                                                                                                      | [*InitializePaymentRequestBodyShopperIdentity](../../models/operations/initializepaymentrequestbodyshopperidentity.md) | :heavy_minus_sign:                                                                                                     | Identification information for the Shopper. This is only required when creating a new Bolt account.                    |