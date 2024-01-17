# Schemas


## Fields

| Field                                                                                                                                    | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `Cart`                                                                                                                                   | [shared.CartCreate](../../../pkg/models/shared/cartcreate.md)                                                                            | :heavy_check_mark:                                                                                                                       | The details of the cart being purchased with this payment.                                                                               |
| `PaymentMethod`                                                                                                                          | [operations.PaymentMethod](../../../pkg/models/operations/paymentmethod.md)                                                              | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `ShopperIdentity`                                                                                                                        | [*operations.PaypalPaymentInputInitializeShopperIdentity](../../../pkg/models/operations/paypalpaymentinputinitializeshopperidentity.md) | :heavy_minus_sign:                                                                                                                       | Identification information for the Shopper. This is only required when creating a new Bolt account.                                      |