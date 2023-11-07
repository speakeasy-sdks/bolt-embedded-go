# CaptureTransactionResponseBody

Unprocessable Entity


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | Raw HTTP response; suitable for custom response parsing |
| `Errors`                                                | [][sdkerrors.Errors](../../models/errors/errors.md)     | :heavy_minus_sign:                                      | N/A                                                     |
| `Result`                                                | [*sdkerrors.Result](../../models/errors/result.md)      | :heavy_minus_sign:                                      | N/A                                                     |