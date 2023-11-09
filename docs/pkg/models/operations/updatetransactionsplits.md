# UpdateTransactionSplits

A split of fees by type and amount.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Amount`                                                                                     | [*shared.AmountView](../../../pkg/models/shared/amountview.md)                               | :heavy_minus_sign:                                                                           | N/A                                                                                          |                                                                                              |
| `Type`                                                                                       | [*operations.UpdateTransactionType](../../../pkg/models/operations/updatetransactiontype.md) | :heavy_minus_sign:                                                                           | **Nullable** for Transactions Details.<br/>                                                  | processing_fee                                                                               |