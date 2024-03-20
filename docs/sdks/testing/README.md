# Testing
(*Testing*)

## Overview

A collection of endpoints that provide useful functionality to assist in testing your Bolt integration.


### Available Operations

* [CreateTestingShopperAccount](#createtestingshopperaccount) - Create Testing Shopper Account
* [GetTestCreditCardToken](#gettestcreditcardtoken) - Fetch a Test Credit Card Token

## CreateTestingShopperAccount

Create a Bolt shopper account for testing purposes. Available for sandbox use only and the created  account will be recycled after a certain time.

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
    res, err := s.Testing.CreateTestingShopperAccount(ctx, operations.CreateTestingShopperAccountRequest{
        RequestBody: &operations.CreateTestingShopperAccountRequestBody{
            DeactivateInDays: boltembeddedgo.Int64(30),
            EmailState: operations.EmailStateVerified.ToPointer(),
            PhoneState: shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailStateVerified.ToPointer(),
        },
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.CreateTestingShopperAccountRequest](../../pkg/models/operations/createtestingshopperaccountrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.CreateTestingShopperAccountResponse](../../pkg/models/operations/createtestingshopperaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTestCreditCardToken

This endpoint fetches a new credit card token for Bolt's universal test credit card number `4111 1111 1111 1004`. This is for testing and is available only in sandbox.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"context"
	"log"
)

func main() {
    s := boltembeddedgo.New(
        boltembeddedgo.WithSecurity(shared.Security{
            OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Testing.GetTestCreditCardToken(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTestCreditCardTokenResponse](../../pkg/models/operations/gettestcreditcardtokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
