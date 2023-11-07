# CartViewFulfillments


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CartShipment`                                                        | [*shared.ICartShipmentView](../../models/shared/icartshipmentview.md) | :heavy_minus_sign:                                                    | N/A                                                                   |
| `FulfillmentType`                                                     | **string*                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `ID`                                                                  | **string*                                                             | :heavy_minus_sign:                                                    | N/A                                                                   |
| `InStoreCartShipment`                                                 | [*shared.InStoreShipment2](../../models/shared/instoreshipment2.md)   | :heavy_minus_sign:                                                    | A cart that is being prepared for shipment                            |
| `Items`                                                               | [][shared.ICartItemView](../../models/shared/icartitemview.md)        | :heavy_minus_sign:                                                    | N/A                                                                   |