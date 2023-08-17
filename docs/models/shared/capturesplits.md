# CaptureSplits

A split of fees by type and amount.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Amount`                                                       | [*AmountView](../../models/shared/amountview.md)               | :heavy_minus_sign:                                             | N/A                                                            |                                                                |
| `Type`                                                         | [*CaptureSplitsType](../../models/shared/capturesplitstype.md) | :heavy_minus_sign:                                             | Fee type options. **Nullable** for Transactions Details.<br/>  | processing_fee                                                 |