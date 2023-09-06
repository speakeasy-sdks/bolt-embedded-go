# Transactions

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
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.AuthorizeTransactionSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Transactions.AuthorizeTransaction(ctx, operations.AuthorizeTransactionRequest{
        IdempotencyKey: boltembeddedapi.String("fugit"),
        RequestBody: &operations.AuthorizeTransactionRequestBody{},
        XPublishableKey: boltembeddedapi.String("magni"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.IAuthorizeResultView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AuthorizeTransactionRequest](../../models/operations/authorizetransactionrequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.AuthorizeTransactionSecurity](../../models/operations/authorizetransactionsecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.AuthorizeTransactionResponse](../../models/operations/authorizetransactionresponse.md), error**


## CaptureTransaction

This captures funds for the designated transaction. A capture can be done for any partial amount or for the total authorized amount.

Although the response returns the standard `transaction_view` object, only `captures` and either `id` or `reference` are needed.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.CaptureTransactionSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Transactions.CaptureTransaction(ctx, operations.CaptureTransactionRequest{
        IdempotencyKey: boltembeddedapi.String("odio"),
        CaptureTransactionWithReference: &shared.CaptureTransactionWithReference{
            Amount: 754,
            Currency: "USD",
            MerchantEventID: boltembeddedapi.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            SkipHookNotification: boltembeddedapi.Bool(false),
            TransactionReference: "LBLJ-TWW7-R9VC",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CaptureTransactionRequest](../../models/operations/capturetransactionrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.CaptureTransactionSecurity](../../models/operations/capturetransactionsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.CaptureTransactionResponse](../../models/operations/capturetransactionresponse.md), error**


## GetTransactionDetails

This allows you to pull the full transaction details for a given transaction.

 **Note**: All objects and fields marked `required` in the Transaction Details response are also **nullable**. This includes any sub-components (objects or fields) also marked `required`.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.GetTransactionDetailsSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Transactions.GetTransactionDetails(ctx, operations.GetTransactionDetailsRequest{
        Reference: "sunt",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTransactionDetails200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetTransactionDetailsRequest](../../models/operations/gettransactiondetailsrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.GetTransactionDetailsSecurity](../../models/operations/gettransactiondetailssecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.GetTransactionDetailsResponse](../../models/operations/gettransactiondetailsresponse.md), error**


## RefundTransaction

This refunds a captured transaction. Refunds can be done for any partial amount or for the total authorized amount. These refunds are processed synchronously and return information about the refunded transaction in the standard `transaction_view` object.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.RefundTransactionSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Transactions.RefundTransaction(ctx, operations.RefundTransactionRequest{
        IdempotencyKey: boltembeddedapi.String("ullam"),
        RequestBody: &operations.RefundTransactionRequestBody{
            Amount: 754,
            Currency: "USD",
            MerchantEventID: boltembeddedapi.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            SkipHookNotification: boltembeddedapi.Bool(false),
            TransactionReference: "LBLJ-TWW7-R9VC",
        },
    }, operationSecurity)
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
| `request`                                                                                    | [operations.RefundTransactionRequest](../../models/operations/refundtransactionrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.RefundTransactionSecurity](../../models/operations/refundtransactionsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.RefundTransactionResponse](../../models/operations/refundtransactionresponse.md), error**


## UpdateTransaction

This allows you to update certain transaction properties post-authorization.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.UpdateTransactionSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Transactions.UpdateTransaction(ctx, operations.UpdateTransactionRequest{
        IdempotencyKey: boltembeddedapi.String("nam"),
        Reference: "hic",
        RequestBody: &operations.UpdateTransactionRequestBody{
            DisplayID: boltembeddedapi.String("order-123"),
            Metadata: map[string]string{
                "voluptatem": "cumque",
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateTransaction200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateTransactionRequest](../../models/operations/updatetransactionrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateTransactionSecurity](../../models/operations/updatetransactionsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateTransactionResponse](../../models/operations/updatetransactionresponse.md), error**


## VoidTransaction

This voids the authorization for a given transaction. Voids must be completed before the authorization is captured.
In the request, either `transaction_id` or `transaction_reference` is required.
Although the response returns the standard `transaction_view` object, only `status` and either `id` or `reference` are needed.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.VoidTransactionSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Transactions.VoidTransaction(ctx, operations.VoidTransactionRequest{
        IdempotencyKey: boltembeddedapi.String("soluta"),
        CreditCardVoid: &shared.CreditCardVoid{
            MerchantEventID: boltembeddedapi.String("dbe0cd5d-3261-41d9-ba61-49e5b9d07567"),
            SkipHookNotification: boltembeddedapi.Bool(false),
            TransactionReference: "LBLJ-TWW7-R9VC",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransactionView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.VoidTransactionRequest](../../models/operations/voidtransactionrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.VoidTransactionSecurity](../../models/operations/voidtransactionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.VoidTransactionResponse](../../models/operations/voidtransactionresponse.md), error**

