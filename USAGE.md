<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"log"
)

func main() {
	s := boltembeddedgo.New()

	operationSecurity := operations.AddAddressSecurity{
		OAuth:   "Bearer <YOUR_ACCESS_TOKEN_HERE>",
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
<!-- End SDK Example Usage [usage] -->