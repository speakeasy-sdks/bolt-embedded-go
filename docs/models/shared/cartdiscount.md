# CartDiscount


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `Amount`                                                                                                     | **float64*                                                                                                   | :heavy_minus_sign:                                                                                           | N/A                                                                                                          | 100                                                                                                          |
| `Code`                                                                                                       | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | N/A                                                                                                          | SUMMER10DISCOUNT                                                                                             |
| `Description`                                                                                                | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | Used to define the discount offering.                                                                        | 10% off our summer collection                                                                                |
| `DetailsURL`                                                                                                 | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | Used to provide a link to additional details, such as a landing page, associated with the discount offering. | https://boltswagstore.com/SUMMERSALE                                                                         |
| `DiscountCategory`                                                                                           | [*CartDiscountDiscountCategory](../../models/shared/cartdiscountdiscountcategory.md)                         | :heavy_minus_sign:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `Reference`                                                                                                  | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | Used to define the reference ID associated with the discount available.                                      | DISC-1234                                                                                                    |
| `Type`                                                                                                       | [*CartDiscountType](../../models/shared/cartdiscounttype.md)                                                 | :heavy_minus_sign:                                                                                           | The type of discount.                                                                                        | percentage                                                                                                   |