# Transactions
(*Transactions*)

## Overview

Authorize credit card transactions and perform operations on those transactions with Bolt's transaction API.


### Available Operations

* [AuthorizeTransaction](#authorizetransaction) - Authorize a Card
* [CaptureTransaction](#capturetransaction) - Capture a Transaction
* [GetTransactionDetails](#gettransactiondetails) - Transaction Details
* [RefundTransaction](#refundtransaction) - Refund a Transaction
* [UpdateTransaction](#updatetransaction) - Update a Transaction
* [VoidTransaction](#voidtransaction) - Void a Transaction

## AuthorizeTransaction

This endpoint authorizes card payments and has three main use cases:
* • Authorize a payment using an unsaved payment method for a guest or logged-in shopper.
* • Authorize a payment using a saved payment method for a logged-in shopper.
*  • Re-charge a previous transaction using the `credit_card_id` of the transaction.


### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.AuthorizeTransactionSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.AuthorizeTransaction(ctx, operations.AuthorizeTransactionRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.IAuthorizeResultView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.AuthorizeTransactionRequest](../../pkg/models/operations/authorizetransactionrequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.AuthorizeTransactionSecurity](../../pkg/models/operations/authorizetransactionsecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.AuthorizeTransactionResponse](../../pkg/models/operations/authorizetransactionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CaptureTransaction

This captures funds for the designated transaction. A capture can be done for any partial amount or for the total authorized amount.

Although the response returns the standard `transaction_view` object, only `captures` and either `id` or `reference` are needed.


### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.CaptureTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.CaptureTransaction(ctx, operations.CaptureTransactionRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CaptureTransactionRequest](../../pkg/models/operations/capturetransactionrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.CaptureTransactionSecurity](../../pkg/models/operations/capturetransactionsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.CaptureTransactionResponse](../../pkg/models/operations/capturetransactionresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse          | 403,404                                  | application/json                         |
| sdkerrors.CaptureTransactionResponseBody | 422                                      | application/json                         |
| sdkerrors.SDKError                       | 4xx-5xx                                  | */*                                      |

## GetTransactionDetails

This allows you to pull the full transaction details for a given transaction.

 **Note**: All objects and fields marked `required` in the Transaction Details response are also **nullable**. This includes any sub-components (objects or fields) also marked `required`.


### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.GetTransactionDetailsSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.GetTransactionDetails(ctx, operations.GetTransactionDetailsRequest{
        Reference: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetTransactionDetailsRequest](../../pkg/models/operations/gettransactiondetailsrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.GetTransactionDetailsSecurity](../../pkg/models/operations/gettransactiondetailssecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.GetTransactionDetailsResponse](../../pkg/models/operations/gettransactiondetailsresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,422                         | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## RefundTransaction

This refunds a captured transaction. Refunds can be done for any partial amount or for the total authorized amount. These refunds are processed synchronously and return information about the refunded transaction in the standard `transaction_view` object.

### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.RefundTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.RefundTransaction(ctx, operations.RefundTransactionRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RefundTransactionRequest](../../pkg/models/operations/refundtransactionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.RefundTransactionSecurity](../../pkg/models/operations/refundtransactionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.RefundTransactionResponse](../../pkg/models/operations/refundtransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 422                             | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## UpdateTransaction

This allows you to update certain transaction properties post-authorization.

### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.UpdateTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.UpdateTransaction(ctx, operations.UpdateTransactionRequest{
        Reference: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.UpdateTransactionRequest](../../pkg/models/operations/updatetransactionrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.UpdateTransactionSecurity](../../pkg/models/operations/updatetransactionsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.UpdateTransactionResponse](../../pkg/models/operations/updatetransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,404                         | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## VoidTransaction

This voids the authorization for a given transaction. Voids must be completed before the authorization is captured.
In the request, either `transaction_id` or `transaction_reference` is required.
Although the response returns the standard `transaction_view` object, only `status` and either `id` or `reference` are needed.


### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.VoidTransactionSecurity{
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Transactions.VoidTransaction(ctx, operations.VoidTransactionRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.VoidTransactionRequest](../../pkg/models/operations/voidtransactionrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.VoidTransactionSecurity](../../pkg/models/operations/voidtransactionsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.VoidTransactionResponse](../../pkg/models/operations/voidtransactionresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,404                         | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |
