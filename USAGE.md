<!-- Start SDK Example Usage [usage] -->
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