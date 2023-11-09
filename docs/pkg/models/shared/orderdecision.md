# OrderDecision

Decision and score for an order.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `DecisionFactors`                                                                     | [][shared.RiskDecisionFactorYml](../../../pkg/models/shared/riskdecisionfactoryml.md) | :heavy_minus_sign:                                                                    | The top 5 factors of the fraud decision.                                              |                                                                                       |
| `Score`                                                                               | **int64*                                                                              | :heavy_minus_sign:                                                                    | The total fraud risk score of the order.                                              | 680                                                                                   |