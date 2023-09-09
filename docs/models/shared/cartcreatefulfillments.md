# CartCreateFulfillments

Defines the shipments associated with the cart items.


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `CartItems`                                                                                            | [][CartItem](../../models/shared/cartitem.md)                                                          | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `CartShipment`                                                                                         | [*CartShipment](../../models/shared/cartshipment.md)                                                   | :heavy_minus_sign:                                                                                     | A cart that is being prepared for shipment                                                             |
| `DigitalDelivery`                                                                                      | [*CartCreateFulfillmentsDigitalDelivery](../../models/shared/cartcreatefulfillmentsdigitaldelivery.md) | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `InStoreCartShipment`                                                                                  | [*InStoreCartShipment](../../models/shared/instorecartshipment.md)                                     | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `Type`                                                                                                 | [*CartCreateFulfillmentsType](../../models/shared/cartcreatefulfillmentstype.md)                       | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |