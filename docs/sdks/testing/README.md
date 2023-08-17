# Testing

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
	"context"
	"log"
	"github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedapi.New()
    operationSecurity := operations.CreateTestingShopperAccountSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Testing.CreateTestingShopperAccount(ctx, operations.CreateTestingShopperAccountRequest{
        RequestBody: &operations.CreateTestingShopperAccountRequestBody{
            DeactivateInDays: boltembeddedapi.Int64(878601),
            Email: boltembeddedapi.String("Zita.Streich@gmail.com"),
            EmailState: operations.CreateTestingShopperAccountRequestBodyEmailStateVerified.ToPointer(),
            HasAddress: boltembeddedapi.Bool(false),
            Migrated: boltembeddedapi.Bool(false),
            Phone: boltembeddedapi.String("534-551-6044"),
            PhoneState: shared.Onev11testing1shopper1createPostRequestBodyContentApplication1jsonSchemaPropertiesEmailStateVerified.ToPointer(),
        },
        XPublishableKey: boltembeddedapi.String("pariatur"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTestingShopperAccount200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.CreateTestingShopperAccountRequest](../../models/operations/createtestingshopperaccountrequest.md)   | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `security`                                                                                                       | [operations.CreateTestingShopperAccountSecurity](../../models/operations/createtestingshopperaccountsecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |


### Response

**[*operations.CreateTestingShopperAccountResponse](../../models/operations/createtestingshopperaccountresponse.md), error**


## GetTestCreditCardToken

This endpoint fetches a new credit card token for Bolt's universal test credit card number `4111 1111 1111 1004`. This is for testing and is available only in sandbox.

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
    operationSecurity := operations.GetTestCreditCardTokenSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Testing.GetTestCreditCardToken(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTestCreditCardToken200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `security`                                                                                             | [operations.GetTestCreditCardTokenSecurity](../../models/operations/gettestcreditcardtokensecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.GetTestCreditCardTokenResponse](../../models/operations/gettestcreditcardtokenresponse.md), error**

