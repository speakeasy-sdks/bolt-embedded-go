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
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.AddAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedgo.String("Bolt"),
            Country: boltembeddedgo.String("United States"),
            CountryCode: "US",
            DoorCode: boltembeddedgo.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{},
            Name: boltembeddedgo.String("Alan Watts"),
            Phone: boltembeddedgo.String("+12125550199"),
            PostalCode: "10044",
            Region: "NY",
            RegionCode: boltembeddedgo.String("NY"),
            StreetAddress1: "888 main street",
            StreetAddress2: boltembeddedgo.String("apt 3021"),
            StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
            StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
        },
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.AddAddressRequest](../../pkg/models/operations/addaddressrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.AddAddressSecurity](../../pkg/models/operations/addaddresssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.AddAddressResponse](../../pkg/models/operations/addaddressresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.AddPaymentMethodSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.AddPaymentMethod(ctx, operations.AddPaymentMethodRequest{
        RequestBody: &operations.AddPaymentMethodRequestBody{
            BillingAddress: shared.Address{
                Company: boltembeddedgo.String("Bolt"),
                Country: boltembeddedgo.String("United States"),
                CountryCode: "US",
                Default: boltembeddedgo.Bool(true),
                DoorCode: boltembeddedgo.String("123456"),
                Email: "alan.watts@example.com",
                FirstName: "Alan",
                LastName: "Watts",
                Locality: "Brooklyn",
                Name: boltembeddedgo.String("Alan Watts"),
                Phone: boltembeddedgo.String("+12125550199"),
                PostalCode: "10044",
                Region: "NY",
                RegionCode: boltembeddedgo.String("NY"),
                StreetAddress1: "888 main street",
                StreetAddress2: boltembeddedgo.String("apt 3021"),
                StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
            },
            BillingAddressID: boltembeddedgo.String("null"),
            Bin: boltembeddedgo.String("411111"),
            Currency: boltembeddedgo.String("USD"),
            Expiration: "2025-11",
            Last4: boltembeddedgo.String("1234"),
            Metadata: &shared.Metadata{},
            PostalCode: boltembeddedgo.String("10044"),
            Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
            TokenType: operations.TokenTypeBolt.ToPointer(),
        },
    }, operationSecurity)
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
| sdkerrors.SDKError | 400-600            | */*                |

## CreateAccount

Create a Bolt shopping account.

### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.CreateAccountSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.CreateAccount(ctx, operations.CreateAccountRequest{
        CreateAccountInput: &shared.CreateAccountInput{
            Addresses: []shared.AddressAccount{
                shared.AddressAccount{
                    Company: boltembeddedgo.String("Bolt"),
                    Country: boltembeddedgo.String("United States"),
                    CountryCode: "US",
                    DoorCode: boltembeddedgo.String("123456"),
                    Email: "alan.watts@example.com",
                    FirstName: "Alan",
                    LastName: "Watts",
                    Locality: "Brooklyn",
                    Metadata: &shared.Metadata{},
                    Name: boltembeddedgo.String("Alan Watts"),
                    Phone: boltembeddedgo.String("+12125550199"),
                    PostalCode: "10044",
                    Region: "NY",
                    RegionCode: boltembeddedgo.String("NY"),
                    StreetAddress1: "888 main street",
                    StreetAddress2: boltembeddedgo.String("apt 3021"),
                    StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                    StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                },
            },
            PaymentMethods: []shared.PaymentMethodAccount{
                shared.PaymentMethodAccount{
                    BillingAddress: shared.Address{
                        Company: boltembeddedgo.String("Bolt"),
                        Country: boltembeddedgo.String("United States"),
                        CountryCode: "US",
                        Default: boltembeddedgo.Bool(true),
                        DoorCode: boltembeddedgo.String("123456"),
                        Email: "alan.watts@example.com",
                        FirstName: "Alan",
                        LastName: "Watts",
                        Locality: "Brooklyn",
                        Name: boltembeddedgo.String("Alan Watts"),
                        Phone: boltembeddedgo.String("+12125550199"),
                        PostalCode: "10044",
                        Region: "NY",
                        RegionCode: boltembeddedgo.String("NY"),
                        StreetAddress1: "888 main street",
                        StreetAddress2: boltembeddedgo.String("apt 3021"),
                        StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
                        StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
                    },
                    BillingAddressID: boltembeddedgo.String("null"),
                    Bin: boltembeddedgo.String("411111"),
                    Expiration: "2025-11",
                    Last4: boltembeddedgo.String("1234"),
                    Metadata: &shared.Metadata{},
                    PostalCode: boltembeddedgo.String("10044"),
                    Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                    TokenType: shared.PaymentMethodAccountTokenTypeBolt.ToPointer(),
                },
            },
            Profile: shared.Profile{
                Email: "alan.watts@example.com",
                FirstName: "Alan",
                LastName: "Watts",
                Metadata: &shared.ProfileMetadata{},
                Phone: "+12125550199",
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountDetails != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateAccountRequest](../../pkg/models/operations/createaccountrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.CreateAccountSecurity](../../pkg/models/operations/createaccountsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.CreateAccountResponse](../../pkg/models/operations/createaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
	"net/http"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.DeleteAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.DeleteAddress(ctx, operations.DeleteAddressRequest{
        ID: "<ID>",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError | 400-600            | */*                |

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
	"net/http"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.DeletePaymentMethodSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.DeletePaymentMethod(ctx, operations.DeletePaymentMethodRequest{
        PaymentMethodID: "string",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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
| sdkerrors.SDKError              | 400-600                         | */*                             |

## DetectAccount

Check whether an account exists using one of `email`, `phone`, or `sha256_email` as the unique identifier.

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
            OAuth: boltembeddedgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Account.DetectAccount(ctx, operations.DetectAccountRequest{
        XPublishableKey: "string",
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
| sdkerrors.SDKError              | 400-600                         | */*                             |

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
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.EditAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.EditAddress(ctx, operations.EditAddressRequest{
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedgo.String("Bolt"),
            Country: boltembeddedgo.String("United States"),
            CountryCode: "US",
            DoorCode: boltembeddedgo.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{},
            Name: boltembeddedgo.String("Alan Watts"),
            Phone: boltembeddedgo.String("+12125550199"),
            PostalCode: "10044",
            Region: "NY",
            RegionCode: boltembeddedgo.String("NY"),
            StreetAddress1: "888 main street",
            StreetAddress2: boltembeddedgo.String("apt 3021"),
            StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
            StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
        },
        ID: "<ID>",
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
| sdkerrors.SDKError | 400-600            | */*                |

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
            OAuth: "",
            XAPIKey: "",
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
| sdkerrors.SDKError | 400-600            | */*                |

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
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.ReplaceAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.ReplaceAddress(ctx, operations.ReplaceAddressRequest{
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedgo.String("Bolt"),
            Country: boltembeddedgo.String("United States"),
            CountryCode: "US",
            DoorCode: boltembeddedgo.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{},
            Name: boltembeddedgo.String("Alan Watts"),
            Phone: boltembeddedgo.String("+12125550199"),
            PostalCode: "10044",
            Region: "NY",
            RegionCode: boltembeddedgo.String("NY"),
            StreetAddress1: "888 main street",
            StreetAddress2: boltembeddedgo.String("apt 3021"),
            StreetAddress3: boltembeddedgo.String("c/o Alicia Watts"),
            StreetAddress4: boltembeddedgo.String("Bridge Street Apartment Building B"),
        },
        ID: "<ID>",
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
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateAccountProfile

Update the identifiers for a shopper's profile (first name or last name).

### Example Usage

```go
package main

import(
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"context"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
    s := boltembeddedgo.New()


    operationSecurity := operations.UpdateAccountProfileSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.UpdateAccountProfile(ctx, operations.UpdateAccountProfileRequest{
        RequestBody: &operations.UpdateAccountProfileRequestBody{
            FirstName: boltembeddedgo.String("Alan"),
            LastName: boltembeddedgo.String("Watts"),
            Metadata: &shared.Metadata{},
        },
    }, operationSecurity)
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
| sdkerrors.SDKError | 400-600            | */*                |
