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
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.AddAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
        IdempotencyKey: boltembeddedgo.String("Maserati Bespoke frictionless"),
        XPublishableKey: boltembeddedgo.String("deploy Central"),
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedgo.String("Bolt"),
            Country: boltembeddedgo.String("United States"),
            CountryCode: "US",
            Default: boltembeddedgo.Bool(false),
            DoorCode: boltembeddedgo.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{
                AdditionalProperties: boltembeddedgo.String("Loan Dollar"),
            },
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

    if res.AddAddress200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.AddAddressRequest](../../models/operations/addaddressrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.AddAddressSecurity](../../models/operations/addaddresssecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.AddAddressResponse](../../models/operations/addaddressresponse.md), error**


## AddPaymentMethod

Add a payment method to a shopper's Bolt account Wallet. For security purposes, this request must come from your backend because authentication requires the use of your private key.

**Note**: Before using this API, the credit card details must be tokenized using Bolt's JavaScript library function, which is documented in [Install the Bolt Tokenizer](https://help.bolt.com/developers/references/bolt-tokenizer).


### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.AddPaymentMethodSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.AddPaymentMethod(ctx, operations.AddPaymentMethodRequest{
        IdempotencyKey: boltembeddedgo.String("violet Scandium"),
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
            Cryptogram: boltembeddedgo.String("Gasoline aggregate"),
            Currency: boltembeddedgo.String("USD"),
            Eci: boltembeddedgo.String("District"),
            Expiration: "2025-11",
            Last4: boltembeddedgo.String("1234"),
            Metadata: &shared.Metadata{
                AdditionalProperties: boltembeddedgo.String("male Chips termite"),
            },
            Network: operations.AddPaymentMethodRequestBodyNetworkAmex.ToPointer(),
            Number: boltembeddedgo.String("Fish Transgender ivory"),
            PostalCode: boltembeddedgo.String("10044"),
            Priority: operations.AddPaymentMethodRequestBodyPriorityOne.ToPointer(),
            Save: boltembeddedgo.Bool(false),
            Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
            TokenType: operations.AddPaymentMethodRequestBodyTokenTypeBolt.ToPointer(),
        },
        XPublishableKey: boltembeddedgo.String("katal Buckinghamshire"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.AddPaymentMethodRequest](../../models/operations/addpaymentmethodrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.AddPaymentMethodSecurity](../../models/operations/addpaymentmethodsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.AddPaymentMethodResponse](../../models/operations/addpaymentmethodresponse.md), error**


## CreateAccount

Create a Bolt shopping account.

### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.CreateAccountSecurity{
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.CreateAccount(ctx, operations.CreateAccountRequest{
        IdempotencyKey: boltembeddedgo.String("grey North rare"),
        XPublishableKey: boltembeddedgo.String("content Man"),
        CreateAccountInput: &shared.CreateAccountInput{
            Addresses: []shared.AddressAccount{
                shared.AddressAccount{
                    Company: boltembeddedgo.String("Bolt"),
                    Country: boltembeddedgo.String("United States"),
                    CountryCode: "US",
                    Default: boltembeddedgo.Bool(false),
                    DoorCode: boltembeddedgo.String("123456"),
                    Email: "alan.watts@example.com",
                    FirstName: "Alan",
                    LastName: "Watts",
                    Locality: "Brooklyn",
                    Metadata: &shared.Metadata{
                        AdditionalProperties: boltembeddedgo.String("Handmade"),
                    },
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
                    Cryptogram: boltembeddedgo.String("West Rap"),
                    Default: boltembeddedgo.Bool(false),
                    Eci: boltembeddedgo.String("matrix"),
                    Expiration: "2025-11",
                    Last4: boltembeddedgo.String("1234"),
                    Metadata: &shared.Metadata{
                        AdditionalProperties: boltembeddedgo.String("Bicycle Lauderhill"),
                    },
                    Network: shared.PaymentMethodAccountNetworkAmex.ToPointer(),
                    Number: boltembeddedgo.String("Hybrid frame Alabama"),
                    PostalCode: boltembeddedgo.String("10044"),
                    Priority: shared.PaymentMethodAccountPriorityTwo.ToPointer(),
                    Save: boltembeddedgo.Bool(false),
                    Token: "a1B2c3D4e5F6G7H8i9J0k1L2m3N4o5P6Q7r8S9t0",
                    TokenType: shared.PaymentMethodAccountTokenTypeBolt.ToPointer(),
                },
            },
            Profile: shared.Profile{
                Email: "alan.watts@example.com",
                FirstName: "Alan",
                LastName: "Watts",
                Metadata: &shared.ProfileMetadata{
                    AdditionalProperties: boltembeddedgo.String("payment"),
                },
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateAccountRequest](../../models/operations/createaccountrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.CreateAccountSecurity](../../models/operations/createaccountsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.CreateAccountResponse](../../models/operations/createaccountresponse.md), error**


## DeleteAddress

Deletes an existing address in a shopper's address book.


### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.DeleteAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.DeleteAddress(ctx, operations.DeleteAddressRequest{
        XPublishableKey: boltembeddedgo.String("Grocery Configurable Larissa"),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteAddressRequest](../../models/operations/deleteaddressrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.DeleteAddressSecurity](../../models/operations/deleteaddresssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.DeleteAddressResponse](../../models/operations/deleteaddressresponse.md), error**


## DeletePaymentMethod

Delete a saved payment method from a shopper's Bolt account Wallet.

### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.DeletePaymentMethodSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.DeletePaymentMethod(ctx, operations.DeletePaymentMethodRequest{
        XPublishableKey: boltembeddedgo.String("Rico Security aha"),
        PaymentMethodID: "sed",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeletePaymentMethodRequest](../../models/operations/deletepaymentmethodrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.DeletePaymentMethodSecurity](../../models/operations/deletepaymentmethodsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.DeletePaymentMethodResponse](../../models/operations/deletepaymentmethodresponse.md), error**


## DetectAccount

Check whether an account exists using one of `email`, `phone`, or `sha256_email` as the unique identifier.

### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedgo.New(
        boltembeddedgo.WithSecurity(shared.Security{
            OAuth: boltembeddedgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Account.DetectAccount(ctx, operations.DetectAccountRequest{
        XPublishableKey: "Specialist",
        Email: boltembeddedgo.String("Connor42@gmail.com"),
        Phone: boltembeddedgo.String("935-762-8190 x328"),
        Sha256Email: boltembeddedgo.String("markets Frozen"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DetectAccountRequest](../../models/operations/detectaccountrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DetectAccountResponse](../../models/operations/detectaccountresponse.md), error**


## EditAddress

Edit an existing address in a shopper's address book.
This endpoint fully replaces the information for an existing address while retaining the same address ID.


### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.EditAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.EditAddress(ctx, operations.EditAddressRequest{
        XPublishableKey: boltembeddedgo.String("Latin Lead"),
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedgo.String("Bolt"),
            Country: boltembeddedgo.String("United States"),
            CountryCode: "US",
            Default: boltembeddedgo.Bool(false),
            DoorCode: boltembeddedgo.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{
                AdditionalProperties: boltembeddedgo.String("Nauru Frozen"),
            },
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

    if res.EditAddress200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.EditAddressRequest](../../models/operations/editaddressrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.EditAddressSecurity](../../models/operations/editaddresssecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.EditAddressResponse](../../models/operations/editaddressresponse.md), error**


## GetAccount

Fetch a shopper's account details to pre-fill checkout fields. This request must come from your backend for security purposes, as it requires the use of your private key to authenticate. For PCI compliance, only limited information is returned for each credit card available in the shopper’s wallet.

### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.GetAccountSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.GetAccount(ctx, operations.GetAccountRequest{
        XPublishableKey: boltembeddedgo.String("Market"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetAccountRequest](../../models/operations/getaccountrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetAccountSecurity](../../models/operations/getaccountsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetAccountResponse](../../models/operations/getaccountresponse.md), error**


## ReplaceAddress

Replace an existing address in a shopper's address book.
These changes delete the existing address and create a new one.


### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
)

func main() {
    s := boltembeddedgo.New()
    operationSecurity := operations.ReplaceAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.ReplaceAddress(ctx, operations.ReplaceAddressRequest{
        IdempotencyKey: boltembeddedgo.String("West Northwest logistical"),
        XPublishableKey: boltembeddedgo.String("purple"),
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedgo.String("Bolt"),
            Country: boltembeddedgo.String("United States"),
            CountryCode: "US",
            Default: boltembeddedgo.Bool(false),
            DoorCode: boltembeddedgo.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{
                AdditionalProperties: boltembeddedgo.String("Frozen Zealand Passenger"),
            },
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

    if res.ReplaceAddress200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ReplaceAddressRequest](../../models/operations/replaceaddressrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ReplaceAddressSecurity](../../models/operations/replaceaddresssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.ReplaceAddressResponse](../../models/operations/replaceaddressresponse.md), error**


## UpdateAccountProfile

Update the identifiers for a shopper's profile (first name or last name).

### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
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
            Metadata: &shared.Metadata{
                AdditionalProperties: boltembeddedgo.String("over"),
            },
        },
        XPublishableKey: boltembeddedgo.String("monetize Northwest"),
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateAccountProfileRequest](../../models/operations/updateaccountprofilerequest.md)   | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [operations.UpdateAccountProfileSecurity](../../models/operations/updateaccountprofilesecurity.md) | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[*operations.UpdateAccountProfileResponse](../../models/operations/updateaccountprofileresponse.md), error**

