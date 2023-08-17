# github.com/speakeasy-sdks/bolt-embedded-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/bolt-embedded-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


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
    operationSecurity := operations.AddAddressSecurity{
            OAuth: "",
            XAPIKey: "",
        }

    ctx := context.Background()
    res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
        IdempotencyKey: boltembeddedapi.String("corrupti"),
        XPublishableKey: boltembeddedapi.String("provident"),
        AddressAccount: &shared.AddressAccount{
            Company: boltembeddedapi.String("Bolt"),
            Country: boltembeddedapi.String("United States"),
            CountryCode: "US",
            Default: boltembeddedapi.Bool(false),
            DoorCode: boltembeddedapi.String("123456"),
            Email: "alan.watts@example.com",
            FirstName: "Alan",
            LastName: "Watts",
            Locality: "Brooklyn",
            Metadata: &shared.Metadata{
                AdditionalProperties: boltembeddedapi.String("distinctio"),
            },
            Name: boltembeddedapi.String("Alan Watts"),
            Phone: boltembeddedapi.String("+12125550199"),
            PostalCode: "10044",
            Region: "NY",
            RegionCode: boltembeddedapi.String("NY"),
            StreetAddress1: "888 main street",
            StreetAddress2: boltembeddedapi.String("apt 3021"),
            StreetAddress3: boltembeddedapi.String("c/o Alicia Watts"),
            StreetAddress4: boltembeddedapi.String("Bridge Street Apartment Building B"),
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Account](docs/sdks/account/README.md)

* [AddAddress](docs/sdks/account/README.md#addaddress) - Add Address
* [AddPaymentMethod](docs/sdks/account/README.md#addpaymentmethod) - Add Payment Method
* [CreateAccount](docs/sdks/account/README.md#createaccount) - Create Bolt Account
* [DeleteAddress](docs/sdks/account/README.md#deleteaddress) - Delete Address
* [DeletePaymentMethod](docs/sdks/account/README.md#deletepaymentmethod) - Delete Payment Method
* [DetectAccount](docs/sdks/account/README.md#detectaccount) - Detect Account
* [EditAddress](docs/sdks/account/README.md#editaddress) - Edit Address
* [GetAccount](docs/sdks/account/README.md#getaccount) - Get Account Details
* [ReplaceAddress](docs/sdks/account/README.md#replaceaddress) - Replace Address
* [UpdateAccountProfile](docs/sdks/account/README.md#updateaccountprofile) - Update Profile

### [Payments](docs/sdks/payments/README.md)

* [FinalizePayment](docs/sdks/payments/README.md#finalizepayment) - Finalize Payment
* [InitializePayment](docs/sdks/payments/README.md#initializepayment) - Initialize Payment
* [UpdatePayment](docs/sdks/payments/README.md#updatepayment) - Update Payment

### [Testing](docs/sdks/testing/README.md)

* [CreateTestingShopperAccount](docs/sdks/testing/README.md#createtestingshopperaccount) - Create Testing Shopper Account
* [GetTestCreditCardToken](docs/sdks/testing/README.md#gettestcreditcardtoken) - Fetch a Test Credit Card Token

### [Transactions](docs/sdks/transactions/README.md)

* [AuthorizeTransaction](docs/sdks/transactions/README.md#authorizetransaction) - Authorize a Card
* [CaptureTransaction](docs/sdks/transactions/README.md#capturetransaction) - Capture a Transaction
* [GetTransactionDetails](docs/sdks/transactions/README.md#gettransactiondetails) - Transaction Details
* [RefundTransaction](docs/sdks/transactions/README.md#refundtransaction) - Refund a Transaction
* [UpdateTransaction](docs/sdks/transactions/README.md#updatetransaction) - Update a Transaction
* [VoidTransaction](docs/sdks/transactions/README.md#voidtransaction) - Void a Transaction
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
