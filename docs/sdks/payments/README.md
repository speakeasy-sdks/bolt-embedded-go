# Payments
(*Payments*)

## Overview

Create and manage transactions for non credit card payments such as Paypal in your Embedded Accounts experience.


### Available Operations

* [FinalizePayment](#finalizepayment) - Finalize Payment
* [InitializePayment](#initializepayment) - Initialize Payment
* [UpdatePayment](#updatepayment) - Update Payment

## FinalizePayment

Finalize a Bolt Payment. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"log"
)

func main() {
    s := boltembeddedgo.New(
        boltembeddedgo.WithSecurity(shared.Security{
            OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.FinalizePayment(ctx, operations.FinalizePaymentRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.FinalizePaymentRequest](../../pkg/models/operations/finalizepaymentrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.FinalizePaymentResponse](../../pkg/models/operations/finalizepaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## InitializePayment

Initialize a Bolt payment token that will be used to reference this payment to Bolt when it is updated or finalized. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"log"
)

func main() {
    s := boltembeddedgo.New(
        boltembeddedgo.WithSecurity(shared.Security{
            OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.InitializePayment(ctx, operations.InitializePaymentRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.InitializePaymentRequest](../../pkg/models/operations/initializepaymentrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.InitializePaymentResponse](../../pkg/models/operations/initializepaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdatePayment

Update a Bolt payment using the token given after initializing a payment.  Updates will completely replace the original top-level resource (for example, if a cart is sent in with the request it will replace the existing cart).  Any included object should be sent as complete object. NOTE: The authorization header is NOT required for payments associated with users who do not have a Bolt account.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"log"
)

func main() {
    s := boltembeddedgo.New(
        boltembeddedgo.WithSecurity(shared.Security{
            OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Payments.UpdatePayment(ctx, operations.UpdatePaymentRequest{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdatePaymentRequest](../../pkg/models/operations/updatepaymentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdatePaymentResponse](../../pkg/models/operations/updatepaymentresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
