# Account
(*Account*)

## Overview

Create Embedded Accounts user flows for logged-in and guest experiences by interacting with and updating shopper data.


### Available Operations

* [AddAddress](#addaddress) - Add Address
* [AddPaymentMethod](#addpaymentmethod) - Add Payment Method
* [CreateAccount](#createaccount) - Create Bolt Account
* [DeleteAddress](#deleteaddress) - Delete Address
* [DeletePaymentMethod](#deletepaymentmethod) - Delete Payment Method
* [DetectAccount](#detectaccount) - Detect Account
* [EditAddress](#editaddress) - Edit Address
* [GetAccount](#getaccount) - Get Account Details
* [ReplaceAddress](#replaceaddress) - Replace Address
* [UpdateAccountProfile](#updateaccountprofile) - Update Profile

## AddAddress

Add an address to a shopper's account address book.

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


    operationSecurity := operations.AddAddressSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.AddAddressRequest](../../pkg/models/operations/addaddressrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.AddAddressSecurity](../../pkg/models/operations/addaddresssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.AddAddressResponse](../../pkg/models/operations/addaddressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AddPaymentMethod

Add a payment method to a shopper's Bolt account Wallet. For security purposes, this request must come from your backend because authentication requires the use of your private key.

**Note**: Before using this API, the credit card details must be tokenized using Bolt's JavaScript library function, which is documented in [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).


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


    operationSecurity := operations.AddPaymentMethodSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.AddPaymentMethod(ctx, operations.AddPaymentMethodRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.SavedCreditCardView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AddPaymentMethodRequest](../../pkg/models/operations/addpaymentmethodrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.AddPaymentMethodSecurity](../../pkg/models/operations/addpaymentmethodsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.AddPaymentMethodResponse](../../pkg/models/operations/addpaymentmethodresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateAccount

Create a Bolt shopping account.

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
    res, err := s.Account.CreateAccount(ctx, operations.CreateAccountRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateAccountRequest](../../pkg/models/operations/createaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CreateAccountResponse](../../pkg/models/operations/createaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteAddress

Deletes an existing address in a shopper's address book.


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


    operationSecurity := operations.DeleteAddressSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.DeleteAddress(ctx, operations.DeleteAddressRequest{
        ID: "<id>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteAddressRequest](../../pkg/models/operations/deleteaddressrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.DeleteAddressSecurity](../../pkg/models/operations/deleteaddresssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.DeleteAddressResponse](../../pkg/models/operations/deleteaddressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeletePaymentMethod

Delete a saved payment method from a shopper's Bolt account Wallet.

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


    operationSecurity := operations.DeletePaymentMethodSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.DeletePaymentMethod(ctx, operations.DeletePaymentMethodRequest{
        PaymentMethodID: "<value>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeletePaymentMethodRequest](../../pkg/models/operations/deletepaymentmethodrequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.DeletePaymentMethodSecurity](../../pkg/models/operations/deletepaymentmethodsecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.DeletePaymentMethodResponse](../../pkg/models/operations/deletepaymentmethodresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,404                         | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## DetectAccount

Check whether an account exists using one of `email`, `phone`, or `sha256_email` as the unique identifier.

### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"log"
)

func main() {
    s := boltembeddedgo.New()

    ctx := context.Background()
    res, err := s.Account.DetectAccount(ctx, operations.DetectAccountRequest{
        XPublishableKey: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V1AccountsView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DetectAccountRequest](../../pkg/models/operations/detectaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DetectAccountResponse](../../pkg/models/operations/detectaccountresponse.md), error**
| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 422                             | application/json                |
| sdkerrors.SDKError              | 4xx-5xx                         | */*                             |

## EditAddress

Edit an existing address in a shopper's address book.
This endpoint fully replaces the information for an existing address while retaining the same address ID.


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


    operationSecurity := operations.EditAddressSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.EditAddress(ctx, operations.EditAddressRequest{
        ID: "<id>",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.EditAddressRequest](../../pkg/models/operations/editaddressrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.EditAddressSecurity](../../pkg/models/operations/editaddresssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.EditAddressResponse](../../pkg/models/operations/editaddressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAccount

Fetch a shopper's account details to pre-fill checkout fields. This request must come from your backend for security purposes, as it requires the use of your private key to authenticate. For PCI compliance, only limited information is returned for each credit card available in the shopper’s wallet.

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


    operationSecurity := operations.GetAccountSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.GetAccount(ctx, operations.GetAccountRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.AccountDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetAccountRequest](../../pkg/models/operations/getaccountrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetAccountSecurity](../../pkg/models/operations/getaccountsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetAccountResponse](../../pkg/models/operations/getaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ReplaceAddress

Replace an existing address in a shopper's address book.
These changes delete the existing address and create a new one.


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


    operationSecurity := operations.ReplaceAddressSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.ReplaceAddress(ctx, operations.ReplaceAddressRequest{
        ID: "<id>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ReplaceAddressRequest](../../pkg/models/operations/replaceaddressrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ReplaceAddressSecurity](../../pkg/models/operations/replaceaddresssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ReplaceAddressResponse](../../pkg/models/operations/replaceaddressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateAccountProfile

Update the identifiers for a shopper's profile (first name or last name).

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


    operationSecurity := operations.UpdateAccountProfileSecurity{
            OAuth: "Bearer <YOUR_ACCESS_TOKEN_HERE>",
            XAPIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Account.UpdateAccountProfile(ctx, operations.UpdateAccountProfileRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfileView != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateAccountProfileRequest](../../pkg/models/operations/updateaccountprofilerequest.md)   | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `security`                                                                                             | [operations.UpdateAccountProfileSecurity](../../pkg/models/operations/updateaccountprofilesecurity.md) | :heavy_check_mark:                                                                                     | The security requirements to use for the request.                                                      |


### Response

**[*operations.UpdateAccountProfileResponse](../../pkg/models/operations/updateaccountprofileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
