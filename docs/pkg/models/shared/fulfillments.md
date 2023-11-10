# Fulfillments

Defines the shipments associated with the cart items.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `CartItems`                                                                      | [][shared.CartItem](../../../pkg/models/shared/cartitem.md)                      | :heavy_minus_sign:                                                               | N/A                                                                              |
| `CartShipment`                                                                   | [*shared.CartShipment](../../../pkg/models/shared/cartshipment.md)               | :heavy_minus_sign:                                                               | A cart that is being prepared for shipment                                       |
| `DigitalDelivery`                                                                | [*shared.DigitalDelivery](../../../pkg/models/shared/digitaldelivery.md)         | :heavy_minus_sign:                                                               | N/A                                                                              |
| `InStoreCartShipment`                                                            | [*shared.InStoreCartShipment](../../../pkg/models/shared/instorecartshipment.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `Type`                                                                           | [*shared.CartCreateType](../../../pkg/models/shared/cartcreatetype.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |