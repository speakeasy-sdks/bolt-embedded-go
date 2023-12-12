# github.com/speakeasy-sdks/bolt-embedded-go

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/bolt-embedded-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
	s := boltembeddedgo.New()

	operationSecurity := operations.AddAddressSecurity{
		OAuth:   "Bearer <YOUR_ACCESS_TOKEN_HERE>",
		XAPIKey: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
		AddressAccount: &shared.AddressAccount{
			Company:        boltembeddedgo.String("Bolt"),
			Country:        boltembeddedgo.String("United States"),
			CountryCode:    "US",
			DoorCode:       boltembeddedgo.String("123456"),
			Email:          "alan.watts@example.com",
			FirstName:      "Alan",
			LastName:       "Watts",
			Locality:       "Brooklyn",
			Metadata:       &shared.Metadata{},
			Name:           boltembeddedgo.String("Alan Watts"),
			Phone:          boltembeddedgo.String("+12125550199"),
			PostalCode:     "10044",
			Region:         "NY",
			RegionCode:     boltembeddedgo.String("NY"),
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
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

### [Transactions](docs/sdks/transactions/README.md)

* [AuthorizeTransaction](docs/sdks/transactions/README.md#authorizetransaction) - Authorize a Card
* [CaptureTransaction](docs/sdks/transactions/README.md#capturetransaction) - Capture a Transaction
* [GetTransactionDetails](docs/sdks/transactions/README.md#gettransactiondetails) - Transaction Details
* [RefundTransaction](docs/sdks/transactions/README.md#refundtransaction) - Refund a Transaction
* [UpdateTransaction](docs/sdks/transactions/README.md#updatetransaction) - Update a Transaction
* [VoidTransaction](docs/sdks/transactions/README.md#voidtransaction) - Void a Transaction

### [OAuth](docs/sdks/oauth/README.md)

* [OAuthToken](docs/sdks/oauth/README.md#oauthtoken) - OAuth Token Endpoint

### [Payments](docs/sdks/payments/README.md)

* [FinalizePayment](docs/sdks/payments/README.md#finalizepayment) - Finalize Payment
* [InitializePayment](docs/sdks/payments/README.md#initializepayment) - Initialize Payment
* [UpdatePayment](docs/sdks/payments/README.md#updatepayment) - Update Payment

### [Testing](docs/sdks/testing/README.md)

* [CreateTestingShopperAccount](docs/sdks/testing/README.md#createtestingshopperaccount) - Create Testing Shopper Account
* [GetTestCreditCardToken](docs/sdks/testing/README.md#gettestcreditcardtoken) - Fetch a Test Credit Card Token
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                    | Status Code                     | Content Type                    |
| ------------------------------- | ------------------------------- | ------------------------------- |
| sdkerrors.ErrorsBoltAPIResponse | 403,404                         | application/json                |
| sdkerrors.SDKError              | 400-600                         | */*                             |

### Example

```go
package main

import (
	"context"
	"errors"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/sdkerrors"
	"log"
)

func main() {
	s := boltembeddedgo.New()

	operationSecurity := operations.DeletePaymentMethodSecurity{
		OAuth:   "Bearer <YOUR_ACCESS_TOKEN_HERE>",
		XAPIKey: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Account.DeletePaymentMethod(ctx, operations.DeletePaymentMethodRequest{
		PaymentMethodID: "string",
	}, operationSecurity)
	if err != nil {

		var e *sdkerrors.ErrorsBoltAPIResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.bolt.com` | None |
| 1 | `https://api-sandbox.bolt.com` | None |
| 2 | `https://api-staging.bolt.com` | None |

#### Example

```go
package main

import (
	"context"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
	s := boltembeddedgo.New(
		boltembeddedgo.WithServerIndex(2),
	)

	operationSecurity := operations.AddAddressSecurity{
		OAuth:   "Bearer <YOUR_ACCESS_TOKEN_HERE>",
		XAPIKey: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
		AddressAccount: &shared.AddressAccount{
			Company:        boltembeddedgo.String("Bolt"),
			Country:        boltembeddedgo.String("United States"),
			CountryCode:    "US",
			DoorCode:       boltembeddedgo.String("123456"),
			Email:          "alan.watts@example.com",
			FirstName:      "Alan",
			LastName:       "Watts",
			Locality:       "Brooklyn",
			Metadata:       &shared.Metadata{},
			Name:           boltembeddedgo.String("Alan Watts"),
			Phone:          boltembeddedgo.String("+12125550199"),
			PostalCode:     "10044",
			Region:         "NY",
			RegionCode:     boltembeddedgo.String("NY"),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
	s := boltembeddedgo.New(
		boltembeddedgo.WithServerURL("https://api.bolt.com"),
	)

	operationSecurity := operations.AddAddressSecurity{
		OAuth:   "Bearer <YOUR_ACCESS_TOKEN_HERE>",
		XAPIKey: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
		AddressAccount: &shared.AddressAccount{
			Company:        boltembeddedgo.String("Bolt"),
			Country:        boltembeddedgo.String("United States"),
			CountryCode:    "US",
			DoorCode:       boltembeddedgo.String("123456"),
			Email:          "alan.watts@example.com",
			FirstName:      "Alan",
			LastName:       "Watts",
			Locality:       "Brooklyn",
			Metadata:       &shared.Metadata{},
			Name:           boltembeddedgo.String("Alan Watts"),
			Phone:          boltembeddedgo.String("+12125550199"),
			PostalCode:     "10044",
			Region:         "NY",
			RegionCode:     boltembeddedgo.String("NY"),
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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `OAuth`      | oauth2       | OAuth2 token |
| `XAPIKey`    | apiKey       | API key      |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
	s := boltembeddedgo.New(
		boltembeddedgo.WithSecurity(shared.Security{
			OAuth: boltembeddedgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"log"
)

func main() {
	s := boltembeddedgo.New()

	operationSecurity := operations.AddAddressSecurity{
		OAuth:   "Bearer <YOUR_ACCESS_TOKEN_HERE>",
		XAPIKey: "<YOUR_API_KEY_HERE>",
	}

	ctx := context.Background()
	res, err := s.Account.AddAddress(ctx, operations.AddAddressRequest{
		AddressAccount: &shared.AddressAccount{
			Company:        boltembeddedgo.String("Bolt"),
			Country:        boltembeddedgo.String("United States"),
			CountryCode:    "US",
			DoorCode:       boltembeddedgo.String("123456"),
			Email:          "alan.watts@example.com",
			FirstName:      "Alan",
			LastName:       "Watts",
			Locality:       "Brooklyn",
			Metadata:       &shared.Metadata{},
			Name:           boltembeddedgo.String("Alan Watts"),
			Phone:          boltembeddedgo.String("+12125550199"),
			PostalCode:     "10044",
			Region:         "NY",
			RegionCode:     boltembeddedgo.String("NY"),
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
