# Fulfillments

Defines the shipments associated with the cart items.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `CartItems`                                                               | [][shared.CartItem](../../models/shared/cartitem.md)                      | :heavy_minus_sign:                                                        | N/A                                                                       |
| `CartShipment`                                                            | [*shared.CartShipment](../../models/shared/cartshipment.md)               | :heavy_minus_sign:                                                        | A cart that is being prepared for shipment                                |
| `DigitalDelivery`                                                         | [*shared.DigitalDelivery](../../models/shared/digitaldelivery.md)         | :heavy_minus_sign:                                                        | N/A                                                                       |
| `InStoreCartShipment`                                                     | [*shared.InStoreCartShipment](../../models/shared/instorecartshipment.md) | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Type`                                                                    | [*shared.CartCreateType](../../models/shared/cartcreatetype.md)           | :heavy_minus_sign:                                                        | N/A                                                                       |