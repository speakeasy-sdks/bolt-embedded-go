# Splits

A split of fees by type and amount.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Amount`                                                         | [*shared.AmountView](../../../pkg/models/shared/amountview.md)   | :heavy_minus_sign:                                               | N/A                                                              |                                                                  |
| `Type`                                                           | [*shared.CaptureType](../../../pkg/models/shared/capturetype.md) | :heavy_minus_sign:                                               | Fee type options. **Nullable** for Transactions Details.<br/>    | processing_fee                                                   |