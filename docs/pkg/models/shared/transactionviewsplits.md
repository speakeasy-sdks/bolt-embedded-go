# TransactionViewSplits

A split of fees by type and amount.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Amount`                                                                         | [*shared.AmountView](../../../pkg/models/shared/amountview.md)                   | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `Type`                                                                           | [*shared.TransactionViewType](../../../pkg/models/shared/transactionviewtype.md) | :heavy_minus_sign:                                                               | Fee type options. **Nullable** for Transactions Details.<br/>                    | processing_fee                                                                   |