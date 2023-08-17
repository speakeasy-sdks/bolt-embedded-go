# Amounts

The amount. **Nullable** for Transactions Details.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Amount`                                                                        | *int64*                                                                         | :heavy_check_mark:                                                              | The amount in cents. **Nullable** for Transactions Details.                     | 754                                                                             |
| `Currency`                                                                      | *string*                                                                        | :heavy_check_mark:                                                              | The 3-digit ISO code for the currency. **Nullable** for Transactions Details.   | USD                                                                             |
| `CurrencySymbol`                                                                | *string*                                                                        | :heavy_check_mark:                                                              | The currency symbol used for the amount. **Nullable** for Transactions Details. | $                                                                               |